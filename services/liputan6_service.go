package services

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Article struct {
	Headline string
	URL      string
	Category string
	Summary  string
}

func CrawlLiputan6() (articles []Article) {
	res, err := http.Get("https://liputan6.com")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	fmt.Println("status code: ", res.StatusCode)
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Extract article data
	doc.Find(".articles--iridescent-list .articles--iridescent-list--text-item").Each(func(i int, s *goquery.Selection) {
		title := strings.TrimSpace(s.Find("h4").Text())
		category := strings.TrimSpace(s.Find(".articles--iridescent-list--text-item__category").Text())
		summary := strings.TrimSpace(s.Find(".articles--iridescent-list--text-item__summary.articles--iridescent-list--text-item__summary-seo").Text())
		url, _ := s.Find("a").Attr("href")

		article := Article{
			Headline: title,
			Category: category,
			Summary:  summary,
			URL:      url,
		}

		articles = append(articles, article)
	})

	return articles
}
