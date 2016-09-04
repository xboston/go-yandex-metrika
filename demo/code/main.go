package main

import (
	"log"

	metrika "github.com/xboston/go-yandex-metrika"
)

func main() {

	log.Println("Start")

	code := ""
	clientID := ""
	clientSecret := ""

	metrika := metrika.NewMetrikaFromCode(code, clientID, clientSecret)
	metrika.SetDebug(true)

	if err := metrika.Authorize(); err != nil {
		log.Fatal(err.Error())
	}

	counterList, _ := metrika.GetCounterList()

	for _, counter := range counterList.Counters {
		log.Println(counter.ID, counter.Name, counter.Site)
	}

	log.Println("Finish")
}
