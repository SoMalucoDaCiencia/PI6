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

var AppleRestCalls = 0

func MainRoutine(async bool) error {
	start := time.Now()

	// Abre o arquivo dos distritos de São Paulo.
	// ######################################################
	file, err := os.OpenFile("misc/distritos.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Print("err", err.Error())
		panic(err)
	}
	defer file.Close()

	// Faz o parse dos objetos.
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
			if async {
				wg.Add(1)
			}
			if i+10 < len(vec) {
				callWorker(async, &wg, atg, ad0, vec[i:i+10])
				continue
			}
			callWorker(async, &wg, atg, ad0, vec[i:])
		}
	}

	// Espera todas as threads e finaliza a rotina.
	// ######################################################
	if async {
		wg.Wait()
	}
	log.WriteLog(log.LogOk, fmt.Sprintf("routine completed in %s with %d calls", time.Since(start).String(), AppleRestCalls), "")
	AppleRestCalls = 0
	return nil
}

var Pendencies = map[string][][]*models.Address{}

func callWorker(async bool, wg *sync.WaitGroup, atg models.AppleTokenGetter, ad0 *models.Address, to []*models.Address) {
	if !async {
		err := worker(nil, atg, ad0, to)
		if err != nil {
			log.WriteLog(log.LogErr, err.Error(), "")
			Pendencies[ad0.GetUuid()] = append(Pendencies[ad0.GetUuid()], to)
		}
		return
	}
	go func() {
		err := worker(wg, atg, ad0, to)
		if err != nil {
			log.WriteLog(log.LogErr, err.Error(), "")
			Pendencies[ad0.GetUuid()] = append(Pendencies[ad0.GetUuid()], to)
		}
	}()
}

func worker(wg *sync.WaitGroup, atg models.AppleTokenGetter, from *models.Address, to []*models.Address) error {
	if wg != nil {
		defer wg.Done()
	}

	// Abre uma conexão com o banco.
	// ######################################################
	db, err := database.GetConn()
	if err != nil {
		return err
	}

	result, err := entity.ExtractRegister(atg.AccessToken, from, to)
	if err != nil {
		err = fmt.Errorf("an error occurred while extracting register: %v", err.Error())
		log.WriteLog(log.LogErr, err.Error(), "apple")
		panic(err.Error())
	}
	AppleRestCalls += 3

	// Faz a inserção no banco.
	// ######################################################
	err = db.Create(result).Error
	if err != nil {
		return err
	}

	// Fecha a conexão.
	// ######################################################
	return database.CloseConn(db)
}
