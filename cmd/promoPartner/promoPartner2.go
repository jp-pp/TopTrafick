package main

import (
	"TopTrafficTest/service/promoPartner"
)


func main() {

	type App interface {
		Start()
	}

	var (
		app 			App
		host			string
		port			uint
		dataFileName	string
	)

	host = "127.0.0.1"
	port = 6002
	dataFileName = "promoPartner2.json"

	app = promoPartner.NewApp(host, port, dataFileName)

	app.Start()
}
