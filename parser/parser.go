package parser

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/borakasmer/motors/Model"
	"github.com/borakasmer/motors/extension"
	"log"
	"net/http"
	"time"
)

func ParseSahibinden(motor string) []Model.Cbr650R {
	var zx6rUrl = "https://www.arabam.com/ikinci-el/motosiklet/" + motor + "?sort=year.desc"
	resultList := make([]Model.Cbr650R, 0)
	client := &http.Client{Timeout: 30 * time.Second}
	res, err := client.Get(zx6rUrl)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	data := doc.Find("#main-listing tbody tr")
	data.Each(func(i int, s *goquery.Selection) {
		if i > 0 {
			resultList = append(resultList,
				Model.Cbr650R{
					Name:       extension.String{s.Find("td:nth-child(2) div").Text()},
					Title:      extension.String{s.Find("td:nth-child(3) div").Text()},
					Year:       extension.String{s.Find("td:nth-child(4) div").Text()},
					Km:         extension.String{s.Find("td:nth-child(5) div").Text()},
					Color:      extension.String{s.Find("td:nth-child(6) div").Text()},
					Price:      extension.String{s.Find("td:nth-child(7) div").Text()},
					CreateDate: extension.String{s.Find("td:nth-child(8) div").Text()},
					Location:   extension.String{s.Find("td:nth-child(9) div a span:nth-child(1)").Text() + " - " + s.Find("td:nth-child(9) div a span:nth-child(2)").Text()},
				})
		}
	})
	return resultList
}
