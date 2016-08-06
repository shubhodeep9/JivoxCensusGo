package api

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/headzoo/surf"
	"log"
	"strings"
	"sync"
)

type Census map[string]Details

type Details struct {
	Population string                       `json:"population"`
	Increase   string                       `json:"increase"`
	Area       string                       `json:"area"`
	Density    string                       `json:"density"`
	Sex        string                       `json:"sex"`
	Literacy   string                       `json:"literacy"`
	Detailed   map[string]map[string]string `json:"details"`
}

func ScrapeDetails(val string) map[string]map[string]string {
	bow := surf.NewBrowser()
	bow.Open("http://www.census2011.co.in/" + val)
	table := bow.Find("#table1").Find("table").Find("tbody")
	detailed := make(map[string]map[string]string)
	table.Find("tr").Each(func(i int, s *goquery.Selection) {
		td := s.Find("td")
		detailed[strings.TrimSpace(td.Eq(0).Text())] = map[string]string{
			"2011": td.Eq(1).Text(),
			"2001": td.Eq(2).Text(),
		}
	})
	return detailed
}

func Scrape() Census {
	bow := surf.NewBrowser()
	err := bow.Open("http://www.census2011.co.in/states.php")
	if err != nil {
		log.Fatal(err)
	}
	var wg sync.WaitGroup
	table := bow.Find(".searchable")
	census := make(Census)
	table.Find("tr").Each(func(i int, s *goquery.Selection) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			detailed := make(map[string]map[string]string)
			if val, exists := s.Find("td").Eq(1).Find("a").Attr("href"); exists {
				detailed = ScrapeDetails(val)
			}
			census[s.Find("td").Eq(1).Text()] = Details{
				Population: s.Find("td").Eq(2).Text(),
				Increase:   s.Find("td").Eq(3).Text(),
				Area:       s.Find("td").Eq(4).Text(),
				Density:    s.Find("td").Eq(5).Text(),
				Sex:        s.Find("td").Eq(6).Text(),
				Literacy:   s.Find("td").Eq(7).Text(),
				Detailed:   detailed,
			}

		}()
	})
	wg.Wait()
	return census
}
