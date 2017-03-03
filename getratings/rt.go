package getratings

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func RtScraper(mname string, year string) string {
	movieName := strings.Replace(mname, " ", "_", 9)
	urlis := "https://www.rottentomatoes.com/m/" + movieName
	if len(year) == 4 {
		urlis = urlis + "_" + year
	}
	doc, err := goquery.NewDocument(urlis)
	if err != nil {
		fmt.Println("error occured")
	}
	rating := doc.Find(".meter-value.superPageFontColor span").Text()
	if len(rating) == 0 {
		return "-1"
	} else if len(rating) > 4 {
		return rating[:3]
	} else {
		return rating[:2]
	}
}