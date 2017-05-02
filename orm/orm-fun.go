package main

import "github.com/jinzhu/gorm"
import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

type Product struct {
	gorm.Model
	Code       string `gorm:"size 5"`
	Price      uint
	IgnoreMe   string `gorm:"-"` // Ignore this field
	Vertical   Vertical `gorm:"ForeignKey:VerticalID"`
	VerticalID uint
}

type Vertical struct {
	gorm.Model
	Name string `gorm:"unique"`
}

//Default Name would be products
func (p *Product) TableName() string {
	return "MeraProduct"
}

func main() {
	db, err := gorm.Open("mysql", "root@/aman?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	migrate(db)

	playProduct(db)

	//db.Create(&Product{Code: "LongCode", Price: 4})

	//schemaAlterPlay(db)
}

func schemaAlterPlay(db *gorm.DB) {
	db.Model(&Product{}).DropColumn("code")
}

func playProduct(db *gorm.DB) {
	createVertical(db)

	// Create
	db.Create(&Product{Code: "L1212", Price: 1000, VerticalID: 1})
	// Read
	var product Product
	db.First(&product, 1)
	// find product with id 1
	db.First(&product, "code = ?", "L1212").Preload("Vertical")
	fmt.Println("Vertical ID:", product.VerticalID, "Vertical Name:", product.Vertical.Name)
	// find product with code l1212
	// Update - update product's price to 2000
	db.Model(&product).Update("Price", 2000)
	// Delete - delete product
	db.Delete(&product)
}
func migrate(db *gorm.DB) {
	/** Clear Old Tables */
	db.DropTable(&Product{}, &Vertical{})
	// Migrate the schema
	db.AutoMigrate(&Product{}, &Vertical{})
}

func createVertical(db *gorm.DB) {
	vertical := &Vertical{}
	db.First(vertical, "name=?", "Shirts")
	db.Create(&Vertical{Name: "Shirts"})
}
