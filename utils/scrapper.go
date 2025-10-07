package utils

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

type ScrapedData struct {
	Title    string
	ShopName string
	Price    string
}

func ScrapeData(baseURL string, limit int) ([]ScrapedData, error) {
	var results []ScrapedData
	nextPage := 1

	scraper := colly.NewCollector(
		colly.AllowedDomains("www.tokopedia.com", "tokopedia.com"),
		colly.AllowURLRevisit(),
	)

	scraper.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.142.86 Safari/537.36")
	})

	scraper.OnHTML("div.css-5wh65g", func(e *colly.HTMLElement) {
		title := e.DOM.Find("span[class^='+tnoqZhn89+NHUA43BpiJg==']").Text()
		ShopName := e.DOM.Find("span[class^='si3CNdiG8AR0EaXvf6bFbQ==']").Text()
		Price := e.DOM.Find("span[class^='urMOIDHH7I0Iy1Dv2oFaNw==']").Text()

		if title != "" {
			data := ScrapedData{
				Title:    title,
				ShopName: ShopName,
				Price:    Price,
			}
			results = append(results, data)
		}
	})

	scraper.OnError(func(r *colly.Response, err error) {
		log.Printf("Request failed with status %d: %s\n", r.StatusCode, err)
	})
	for nextPage <= limit {
		pageURL := fmt.Sprintf("%spage=%d", baseURL, nextPage)
		err := scraper.Visit(pageURL)
		if err != nil {
			return nil, fmt.Errorf("failed to visit URL: %W", err)
		}
		nextPage++
	}

	scraper.Wait()

	return results, nil
}
