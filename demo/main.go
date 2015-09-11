package main

import (
	"fmt"
	"log"

	Metrika "github.com/xboston/yametrikago"
)

func main() {

	log.Println("Start")

	var (
		clientId, clientSecret, username, password, token, code string
		err                                                     error
	)

	//
	clientId = ""
	clientSecret = ""
	username = ""
	password = ""

	//
	code = ""

	// при указании токена остальные параметры не обязательны
	token = ""

	metrika := Metrika.NewMetrika(clientId, clientSecret, username, password, token, code)
	metrika.SetDebug(true)

	err = metrika.Authorize()
	PanicIfErr(err)

	counterList, err := metrika.GetCounterList()
	PanicIfErr(err)

	Debug(counterList)

	counterId := counterList.Counters[0].ID

	counterData, err := metrika.GetCounter(counterId)
	PanicIfErr(err)

	Debug(counterData)

	statTrafficSummary, err := metrika.GetStatTrafficSummary(counterId)
	PanicIfErr(err)

	Debug(statTrafficSummary.Data[6])

	log.Println("Finish")
}

func PanicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Debug(data interface{}) {
	fmt.Printf("%#v", data)
}
