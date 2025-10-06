package utils

import "github.com/gocolly/colly"

type ScrapedData struct {
	Title    string
	ShopName string
	Price    string
}

func ScrapeData(baseURL string, limit int) ([]ScrapedData, error) {
	var results []ScrapedDataconst
	nextPage := 1

	scraper := colly.NewCollector(
		colly.AllowedDomain("www.tokopedia.com", "tokopedia.com"),
		colly.AllowURLRevisit(),
	)

	scraper.onrequest = func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.142.86 Safari/537.36")
	}

	scraper.OnHTML("div.css-5wh65g", func(e *colly.HTMLElement){
		title := e.Dom.Find("span[class^='+tnoqZhn89+NHUA43BpiJg==']").text()
		ShopName :=
		Price :=
	})
}
