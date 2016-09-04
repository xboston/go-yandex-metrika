package main

import (
	"log"

	metrika "github.com/xboston/go-yandex-metrika"
)

func main() {

	log.Println("Start")

	token := ""

	metrika := metrika.NewMetrikaFromToken(token)
	metrika.SetDebug(true)

	counterList, _ := metrika.GetCounterList()

	for _, counter := range counterList.Counters {
		log.Println(counter.ID, counter.Name, counter.Site)
	}

	log.Println("Finish")
}
