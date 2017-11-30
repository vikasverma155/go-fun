package crawler

import (
	"github.com/amanhigh/go-fun/util"
	"sync"
	"context"
	"fmt"
	. "github.com/amanhigh/go-fun/models/crawler"
	"io/ioutil"
	"strings"
)

const (
	GOOD_URL_FILE = "/tmp/good.url"
	BAD_URL_FILE  = "/tmp/bad.url"
	BUFFER_SIZE   = 512
)

type Crawler interface {
	GetBaseUrl() string
	GatherLinks(page *util.Page, ch chan CrawlInfo)
	NextPageLink(page *util.Page) (string, bool)
	PrintSet(good []CrawlInfo, bad []CrawlInfo) bool
}

type CrawlerManager struct {
	Crawler    Crawler
	ctx        context.Context
	cancelFunc context.CancelFunc

	/* Counts to track collected & required */
	//collectCount  int32
	//RequiredCount int32

	infoChannel chan CrawlInfo
	goodInfo    []CrawlInfo
	badInfo     []CrawlInfo
}

func NewCrawlerManager(crawler Crawler, requiredCount int) *CrawlerManager {
	return &CrawlerManager{
		Crawler: crawler,
		//RequiredCount: int32(requiredCount),
		infoChannel: make(chan CrawlInfo, BUFFER_SIZE),
	}
}

func (self *CrawlerManager) Crawl() {
	topPage := util.NewPage(self.Crawler.GetBaseUrl())

	/* Fire First Crawler */
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(1)
	go self.crawlRecursive(topPage, waitGroup)

	/* Collect & Organise Crawled Links */
	go self.BuildSet()

	/* Wait for all Crawlers to Return */
	waitGroup.Wait()
	close(self.infoChannel)

	/* Print Organised Links */
	self.PrintSet(self.goodInfo, self.badInfo)
}

func (self *CrawlerManager) BuildSet() {
	/* Fire Parallel Consumer to Separate Movies */
	for info := range self.infoChannel {
		if info.GoodBad() {
			self.goodInfo = append(self.goodInfo, info)
		} else {
			self.badInfo = append(self.badInfo, info)
		}
	}
}

func (self *CrawlerManager) PrintSet(good []CrawlInfo, bad []CrawlInfo) {
	/* Check if Crawler want us to print or already has printed required info */
	if ok := self.Crawler.PrintSet(good, bad); ok {
		/* Output Good/Bad Info in Separate Sections */
		util.PrintGreen(fmt.Sprintf("Passed Info: %v", len(good)))
		printWriteCrawledInfo(good, GOOD_URL_FILE)

		util.PrintYellow(fmt.Sprintf("Failed Info: %v", len(bad)))
		printWriteCrawledInfo(bad, BAD_URL_FILE)
	}
}

/**
	Print Info using interface and write extracted links to
	GOOD/BAD Files for Chrome Processing
 */
func printWriteCrawledInfo(good []CrawlInfo, filePath string) {
	var urls []string
	for _, info := range good {
		info.Print()
		urls = append(urls, info.ToUrl())
	}
	ioutil.WriteFile(filePath, []byte(strings.Join(urls, "\n")), util.DEFAULT_PERM)
}

/**
	Recursively Crawl Given Page moving to next if next link is available.
	Write all Movies of current page onto channel
 */
func (self *CrawlerManager) crawlRecursive(page *util.Page, waitGroup *sync.WaitGroup) {
	util.PrintYellow(fmt.Sprintf("Processing: %v Collected: %v", page.Document.Url.String(), 1))

	/* If Next Link is Present Crawl It */
	if link, ok := self.Crawler.NextPageLink(page); ok {
		waitGroup.Add(1)
		go self.crawlRecursive(util.NewPage(link), waitGroup)
	}
	/* Find Links for this Page */
	self.Crawler.GatherLinks(page, self.infoChannel)
	waitGroup.Done()
}
