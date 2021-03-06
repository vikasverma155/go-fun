package components

import (
	"fmt"
	"time"

	"github.com/vikasverma155/go-fun/kohan/commander/tools"
	"github.com/vikasverma155/go-fun/util"
)

func GetVersion(pkgName string, host string, versionType string, comment string) {
	switch versionType {
	case "dpkg":
		GetDpkgVersion(pkgName, host)
	case "latest":
		GetLatestVersion(pkgName, host, comment)
	}
}

func GetDpkgVersion(pkgName string, host string) {
	util.PrintBlue(fmt.Sprintf("Fetching Config for %v from %v", pkgName, host))
	cmd := fmt.Sprintf(`ssh %v dpkg -l | grep "%v" | tail -1 | awk '{print $3}'`, host, pkgName)
	dpkgVersion := tools.RunCommandPrintError(cmd)
	versionString := fmt.Sprintf("\n%v - %v - HostVersion: %v", pkgName, dpkgVersion, util.FormatTime(time.Now(), util.PRINT_LAYOUT))
	util.PrintYellow(versionString)
	util.AppendFile(util.RELEASE_FILE, versionString)
}

func GetLatestVersion(pkgName string, host string, comment string) {
	util.PrintBlue(fmt.Sprintf("Fetching LatestVersion for %v from %v", pkgName, host))
	cmd := fmt.Sprintf(`ssh %v "sudo apt-get update > /dev/null; apt-cache madison %v | head -1" | awk '{print $3}'`, host, pkgName)
	latestVersion := tools.RunCommandPrintError(cmd)
	versionString := fmt.Sprintf("\n%v - %v - LatestVersion [ %v ]", pkgName, latestVersion, comment)
	util.PrintYellow(versionString)
	util.AppendFile(util.RELEASE_FILE, versionString)
}
