package main

import (
	"log"

	metrika "github.com/xboston/go-yandex-metrika"
)

func main() {

	log.Println("Start")

	var (
		clientID, clientSecret, username, password, token, code string
	)

	token = "398a167c1a864b06badf9f5e40bab675"

	metrika := metrika.NewMetrika(clientID, clientSecret, username, password, token, code)
	metrika.SetDebug(true)

	metrika.Authorize()
	counterList, _ := metrika.GetCounterList()

	for _, counter := range counterList.Counters {
		log.Println(counter.ID, counter.Name, counter.Site)
	}

	log.Println("Finish")
}
