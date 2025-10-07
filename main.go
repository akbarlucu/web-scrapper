package main

import (
	"fmt"
	"log"

	"github.com/akbarlucu/web-scrapper/utils"
)

func main() {
	url := "https://www.tokopedia.com/search?q=gundam"

	data, err := utils.ScrapeData(url, 10)
	if err != nil {
		log.Fatalf("Scrapping failed: %v", err)
	}

	for i, item := range data {
		fmt.Printf("Product %d:\n", i+1)
		fmt.Printf("Title %s:\n", item.Title)
		fmt.Printf("Shop %s:\n", item.ShopName)
		fmt.Printf("Price %s:\n", item.Price)
		fmt.Print("")
	}
}
