package routine

import (
	"PI6/database"
	"PI6/models"
	"PI6/models/entity"
	log "PI6/share/log"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/gocarina/gocsv"
)

func MainRoutine() error {
	start := time.Now()

	// Abre o arquivo dos distritos de São Paulo.
	// ######################################################
	file, err := os.OpenFile("../misc/distritos.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Print("err", err.Error())
		panic(err)
	}
	defer file.Close()

	// Faz o parse para os objetos.
	// ######################################################
	var all []*models.Address
	err = gocsv.UnmarshalFile(file, &all)
	if err != nil {
		return err
	}

	// Abre uma conexão com o banco.
	// ######################################################
	db, err := database.GetConn()
	if err != nil {
		return err
	}

	// Seleciona um token válido no banco.
	// ######################################################
	var tokens []entity.Token
	db = db.Model(&entity.Token{}).Where("createdAt >= DATEADD(week, -1, GETDATE())").First(&tokens)
	if db.Error != nil {
		return db.Error
	}

	// Usa o token como seed para extrair um token válido no servidor da apple.
	// ######################################################################################
	atg, err := entity.ExtractToken(tokens[0].Value)
	if db.Error != nil {
		return db.Error
	}

	// Faz um for-loop por todos os endereços.
	// ######################################################
	var wg sync.WaitGroup
	for k, ad0 := range all {

		// Manipulação de arrays para evitar confusão com ponteiros
		// ##########################################################
		vec := make([]*models.Address, len(all))
		copy(vec[:], all[:])
		vec = append(vec[:k], vec[k+1:]...)

		// Se o token não estiver mais valido, extraia outro.
		// ######################################################
		if !atg.IsValid() {
			err := atg.Renew()
			if err != nil {
				return err
			}
		}

		// Separa os destinos em blocos de dez unidades .
		// ######################################################
		for i := 0; i < len(vec); i += 10 {
			wg.Add(1)
			if i+9 < len(vec) {
				callWorker(&wg, atg, ad0, vec[i:i+9])
				continue
			}
			callWorker(&wg, atg, ad0, vec[i:])
		}
	}

	wg.Wait()
	log.WriteLog(log.LogOk, "routine completed in "+time.Since(start).String(), "")
	return nil
}

var Pendencies = map[string][][]*models.Address{}

func callWorker(wg *sync.WaitGroup, atg models.AppleTokenGetter, ad0 *models.Address, to []*models.Address) {
	go func() {
		err := worker(wg, atg, ad0, to)
		if err != nil {
			Pendencies["ds"] = append(Pendencies["dsa"], to)
		}
	}()
}

func worker(wg *sync.WaitGroup, atg models.AppleTokenGetter, from *models.Address, to []*models.Address) error {
	defer wg.Done()

	db, err := database.GetConn()
	if err != nil {
		return err
	}

	origin := models.NewGeoPos(from.Lat, from.Long)
	var destinies []models.GeoPos
	for _, add := range to {
		destinies = append(destinies, models.NewGeoPos(add.Lat, add.Long))
	}

	result, err := entity.ExtractRegister(atg.AccessToken, origin, destinies)
	if err != nil {
		err = fmt.Errorf("an error occurred while extracting register: %v", err.Error())
		log.WriteLog(log.LogErr, err.Error(), "apple")
		panic(err.Error())
	}

	err = db.Create(result).Error
	if err != nil {
		return err
	}

	return database.CloseConn(db)
}
