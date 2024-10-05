package routine

import (
	"PI6/database"
	"PI6/share/log"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var AppleRestCalls = 0

func MainRoutine(async bool) error {
	start := time.Now()

	// Abre uma conexão com o banco.
	// ######################################################
	_, err := database.GetConn()
	if err != nil {
		return err
	}

	// Faz um for-loop por todos os remedios.
	// ######################################################
	var wg sync.WaitGroup
	for {
		if rand.Int()%2 == 0 {
			break
		}
		wg.Add(1)
		go func() {
			time.Sleep(5 * time.Second)
		}()
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

//func callWorker(async bool, wg *sync.WaitGroup, atg models.AppleTokenGetter, ad0 *models.Address, to []*models.Address) {
//	if !async {
//		err := worker(nil, atg, ad0, to)
//		if err != nil {
//			log.WriteLog(log.LogErr, err.Error(), "")
//			// Pendencies[ad0.GetUuid()] = append(Pendencies[ad0.GetUuid()], to)
//		}
//		return
//	}
//	go func() {
//		err := worker(wg, atg, ad0, to)
//		if err != nil {
//			log.WriteLog(log.LogErr, err.Error(), "")
//			// Pendencies[ad0.GetUuid()] = append(Pendencies[ad0.GetUuid()], to)
//		}
//	}()
//}
//
//func worker(wg *sync.WaitGroup, atg models.AppleTokenGetter, from *models.Address, to []*models.Address) error {
//	if wg != nil {
//		defer wg.Done()
//	}
//
//	// Abre uma conexão com o banco.
//	// ######################################################
//	db, err := database.GetConn()
//	if err != nil {
//		return err
//	}
//
//	result, err := entity.ExtractRegister(atg.AccessToken, from, to)
//	if err != nil {
//		err = fmt.Errorf("an error occurred while extracting register: %v", err.Error())
//		log.WriteLog(log.LogErr, err.Error(), "apple")
//		panic(err.Error())
//	}
//	AppleRestCalls += 3
//
//	// Faz a inserção no banco.
//	// ######################################################
//	err = db.Create(result).Error
//	if err != nil {
//		return err
//	}
//
//	// Fecha a conexão.
//	// ######################################################
//	return database.CloseConn(db)
//}
