package api

import (
	"JivoxCensusGo/models"
	"github.com/PuerkitoBio/goquery"
	"github.com/headzoo/surf"
	"log"
	"strings"
	"sync"
)

/*
Function to scrape the details from individual state pages
@params value of url string to be parsed
@return a map having string keys of map having string keys of string
*/
func ScrapeDetails(val string) map[string]map[string]string {
	bow := surf.NewBrowser()
	bow.Open("http://www.census2011.co.in/" + val)
	table := bow.Find("#table1").Find("table").Find("tbody")
	detailed := make(map[string]map[string]string)
	table.Find("tr").Each(func(i int, s *goquery.Selection) {
		td := s.Find("td")
		detailed[strings.Replace(td.Eq(0).Text(), " ", "", -1)] = map[string]string{
			"2011": td.Eq(1).Text(),
			"2001": td.Eq(2).Text(),
		}
	})
	return detailed
}

/*
Function to scrape the Census data
@return Census (A struct having all the required data types)
@params null
*/
func Scrape() models.Census {
	bow := surf.NewBrowser()
	err := bow.Open("http://www.census2011.co.in/states.php")
	if err != nil {
		log.Fatal(err)
	}
	//Mutex wait group to wait till all the coroutines are completed
	var wg sync.WaitGroup
	table := bow.Find(".searchable")
	census := make(models.Census)
	table.Find("tr").Each(func(i int, s *goquery.Selection) {

		wg.Add(1)
		/*
			The go routine used to concurrently all the pages in a single time
		*/
		go func() {
			defer wg.Done()
			detailed := make(map[string]map[string]string)
			if val, exists := s.Find("td").Eq(1).Find("a").Attr("href"); exists {
				detailed = ScrapeDetails(val)
			}
			census[s.Find("td").Eq(1).Text()] = models.Details{
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
