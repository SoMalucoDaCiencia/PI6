package main

import (
	"PI6/database"
	"PI6/models/entity"
	"PI6/share/log"
	"fmt"
	mgu "github.com/artking28/myGoUtils"
	"github.com/gocolly/colly"
	"math"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	t := time.Now()

	//file, err := os.ReadFile("../misc/proxies.json")
	//if err != nil {
	//    panic(err)
	//}
	//var proxies []models.ProxyObj
	//_ = json.Unmarshal(file, &proxies)

	allLinks := findAllLinks(true)
	var allProds []entity.Chemical

	var count float64
	tcl := mgu.NewThreadControl(50)
	for _, link := range allLinks {
		tcl.Begin()
		go func() {
			//c, err := entity.NewChemical("https://www.drogariasaopaulo.com.br/janumet-xr-100mg-1000mg-merck-30-comprimidos/p")
			c, err := entity.NewChemical(link, http.DefaultClient)
			//c, err := entity.NewChemical(link, share.GetRandomProxy(proxies))
			if err != nil {
				//panic(err)
				log.WriteLog(log.LogErr, err.Error(), "")
			}
			if c == nil {
				return
			}
			tcl.Lock()
			count++
			time.Sleep(800 * time.Millisecond)
			relative := int(math.Floor(count / 10_000))
			hashtags := strings.Repeat("#", relative)
			dashes := strings.Repeat("-", 100-relative)
			fmt.Printf("\rLoading %s%s %.4f%%", hashtags, dashes, count/10_000)
			allProds = append(allProds, *c)
			tcl.Unlock()
			tcl.Done()
		}()
	}
	tcl.Wait()

	db, err := database.GetConn()
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(allProds); i++ {
		prod := allProds[i]
		var count int64
		db = db.Where("sku = ? or ean = ?", prod.ExternalId, prod.Ean).Count(&count)
		if db.Error != nil {
			log.WriteLog(log.LogErr, db.Error.Error(), "")
		}
		if count > 0 {
			allProds = append(allProds[:i], allProds[i+1:]...)
		}
	}

	db = db.Create(allProds)
	if db.Error != nil {
		panic(err)
	}

	println("Completed. Extracted %d elements in", len(allLinks), time.Since(t).Milliseconds())
}

func findAllLinks(saveToFile bool) []string {
	tcl := mgu.NewThreadControl(300)

	var generalLinkList []string
	for _, brandName := range entity.GetChemicalBrand() {

		local := mgu.NewThreadControl(100)
		counter := 50
		page := 1

		tcl.Begin()
		go func() {

			var localLinkList []string
			for ; counter >= 50; page++ {
				counter = 0
				c := colly.NewCollector()
				c.OnHTML("div.main a.collection-link", func(e *colly.HTMLElement) {
					local.Begin()
					go func() {
						local.Lock()
						counter++
						localLinkList = append(localLinkList, e.Attr("href"))
						local.Unlock()
						local.Done()
					}()
				})

				err := c.Visit(fmt.Sprintf("https://www.drogariasaopaulo.com.br/medicamentos/%s?PS=50&PageNumber=%d", brandName, page))
				if err != nil {
					log.WriteLog(log.LogErr, fmt.Sprintf("A error occurred at %s in page number %d [%s]", brandName, page, err.Error()), "")
				}
				local.Wait()
			}

			local.Lock()
			generalLinkList = append(generalLinkList, localLinkList...)
			local.Unlock()
			tcl.Done()
		}()
	}
	tcl.Wait()

	if saveToFile && len(generalLinkList) > 0 {
		content := []byte(strings.Join(generalLinkList, "\n"))
		name := fmt.Sprintf("../misc/productsLinks_%d.txt", time.Now().Weekday())
		err := os.WriteFile(name, content, 666)
		if err != nil {
			log.WriteLog(log.LogErr, err.Error(), "")
		}
	}
	return generalLinkList
}
