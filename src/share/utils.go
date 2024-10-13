package share

import (
	"PI6/models"
	"PI6/share/log"
	"fmt"
	mgu "github.com/artking28/myGoUtils"
	"github.com/gocolly/colly"
	"os"
	"strings"
	"time"
)

func FindSkuPacks() [][50]string {
	tcl := mgu.NewThreadControl(10000)

	var generalLinkList [][50]string
	var lastPack [50]string
	var lastInsertion int
	for _, brandName := range models.GetChemicalBrand() {

		local := mgu.NewThreadControl(100)
		counter := 50
		page := 1

		tcl.Begin()
		go func() {

			for ; counter >= 50; page++ {
				counter = 0
				c := colly.NewCollector()
				c.OnHTML("div.main div[data-trustvox-product-code]", func(e *colly.HTMLElement) {
					local.Begin()
					go func() {
						tcl.Lock()
						counter++
						if lastInsertion >= 50 {
							generalLinkList = append(generalLinkList, lastPack)
							lastPack = [50]string{}
							lastInsertion = 0
						}
						if len(fmt.Sprintf("%s", e.Attr("data-trustvox-product-code"))) > 0 {
							lastPack[lastInsertion] = e.Attr("data-trustvox-product-code")
							lastInsertion++
						}
						tcl.Unlock()
						local.Done()
					}()
				})

				err := c.Visit(fmt.Sprintf("https://www.drogariasaopaulo.com.br/medicamentos/%s?PS=50&PageNumber=%d", brandName, page))
				if err != nil {
					log.WriteLog(log.LogErr, fmt.Sprintf("A error occurred at %s in page number %d [%s]", brandName, page, err.Error()), "")
				}
				local.Wait()
			}
			tcl.Done()
		}()
	}
	tcl.Wait()

	return generalLinkList
}

func FindAllLinks(saveToFile bool) []string {
	tcl := mgu.NewThreadControl(300)

	var generalLinkList []string
	for _, brandName := range models.GetChemicalBrand() {

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
