package routine

import (
	"PI6/database"
	"PI6/models/entity"
	"PI6/share"
	"PI6/share/log"
	"encoding/json"
	"fmt"
	mgu "github.com/artking28/myGoUtils"
	"math"
	"net/http"
	"strings"
	"time"
)

func MainRoutine() error {
	tcl := mgu.NewThreadControl(30)
	start := time.Now()

	// Pega todos os SKU's divididos em pacotes de 50 elementos
	// ########################################################################
	allLinks := share.FindSkuPacks()

	// Sets para mapear produtos ja existentes
	// ########################################################################
	setSku := map[string]*uint64{}
	setEan := map[string]*uint64{}

	// Busca os produtos e os salva nos sets
	// ########################################################################
	tcl.Begin()
	go func() {
		db, err := database.GetConn()
		if err != nil {
			panic(err)
		}

		var vector []entity.Chemical
		db = db.Model(entity.Chemical{}).Find(&vector)
		if db.Error != nil {
			log.WriteLog(log.LogErr, db.Error.Error(), "database")
		}

		for _, chemical := range vector {
			tcl.Lock()
			setSku[chemical.ExternalId] = chemical.ID
			setEan[chemical.Ean] = chemical.ID
			tcl.Unlock()
		}
		tcl.Done()
	}()

	// Produtos q serão salvos e contagem
	// ########################################################################
	var inserts []entity.Chemical
	var insertsPrices []entity.PriceUnity
	var count float64

	for _, group := range allLinks[:1] {

		// All group SKU's
		base := "https://www.drogariasaopaulo.com.br/api/catalog_system/pub/products/search?_from=0&_to=49"
		for _, v := range group {
			add := fmt.Sprintf("&fq=skuId:%s", v)
			if add != "&fq=skuId:" {
				base += add
			}
		}

		tcl.Begin()
		go func() {

			var res []byte
			var allProds []entity.ChemicalJson
			var err error

			eCode, err := share.Rest(http.MethodGet, base, &res, nil, nil, nil)
			if eCode == 500 {
				time.Sleep(time.Second)
				eCode, err = share.Rest(http.MethodGet, base, &res, nil, nil, nil)
			}
			if err != nil {
				log.WriteLog(log.LogErr, err.Error(), "drogariasaopaulo")
				tcl.Done()
				return
			}

			if err != nil {
				println(base)
				log.WriteLog(log.LogErr, err.Error(), "drogariasaopaulo")
				tcl.Done()
				return
			}

			err = json.Unmarshal(res, &allProds)
			if err != nil {
				panic(err)
			}

			for _, v := range allProds {
				inserts = append(inserts, v.Adapt())
			}

			tcl.Lock()
			count++
			time.Sleep(750 * time.Millisecond)
			n := count / float64(len(allLinks)) * 100
			relative := int(math.Floor(n))
			hashtags := strings.Repeat("#", relative)
			dashes := strings.Repeat("-", 100-relative)
			fmt.Printf("\rLoading %s%s %.4f%%", hashtags, dashes, n)
			tcl.Unlock()
			tcl.Done()
		}()
	}

	// Espera todas as threads e finaliza a rotina.
	// ########################################################################
	tcl.Wait()
	println()

	// Abre uma conexão com o banco.
	// ########################################################################
	db, err := database.GetConn()
	if err != nil {
		return err
	}

	// Faz um for-loop por todos os remedios e os insere se n existirem
	// ########################################################################
	for i := 0; i < len(inserts); i++ {
		novo := inserts[i]
		if setSku[novo.ExternalId] != nil || setEan[novo.Ean] != nil {
			inserts = append(inserts[:i], inserts[i+1:]...)
			i--

			price := novo.Prices[0]
			price.ChemicalID = setSku[novo.ExternalId]
			insertsPrices = append(insertsPrices, price)
		}
	}

	if len(inserts) > 0 {
		db = db.Create(inserts)
		if db.Error != nil {
			panic(db.Error)
		}
	}

	if len(insertsPrices) > 0 {
		db, err = database.GetConn()
		if err != nil {
			return err
		}

		db = db.Model(entity.PriceUnity{}).Omit("Chemical").Save(insertsPrices)
		if db.Error != nil {
			panic(err)
		}
	}

	// Finaliza o programa
	// ########################################################################
	timeSince := time.Since(start).String()
	size := len(allLinks)
	str := fmt.Sprintf("Completed. %d elements has been extracted in %s", size, timeSince)
	log.WriteLog(log.LogOk, str, "")
	return nil
}
