package main

import (
	"app/config"
	"app/injector"
	"app/routers"
)

func main() {
	conf := config.NewConfig()
	cronInterest := injector.InitializedCronInterest(conf.Logger, conf.Database)
	go cronInterest.DeleteInterestAfterOneDay()

	router := routers.NewRouter(conf)
	router.RouterAuth()
	router.RouterPartner()
	router.RouterUser()
	router.Run()
}
