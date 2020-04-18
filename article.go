package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type article struct {
	title string
	time  time.Time
	url   string
}

func parseArticle(s *goquery.Selection) article {
	title := strings.TrimSpace(
		s.Find("h1").Text())

	url, ok := s.Find("a").Attr("href")
	if !ok {
		panic(fmt.Errorf("url is missing: %q", title))
	}

	timeVal, ok := s.Find("time").Attr("datetime")
	if !ok {
		panic(fmt.Errorf("time is missing: %q", title))
	}

	t, err := time.Parse(time.RFC3339, timeVal)
	if err != nil {
		panic(fmt.Errorf("cannot parse datetime: %s", timeVal))
	}

	return article{
		title,
		t,
		url,
	}
}

func scrapeArticles() ([]article, error) {
	res, err := http.Get("https://yle.fi/uutiset/osasto/news/")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		err := fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	articles := []article{}

	doc.Find("#oikea_palsta").Find("article").Each(func(i int, s *goquery.Selection) {
		a := parseArticle(s)

		articles = append(articles, a)
	})

	return articles, nil
}
