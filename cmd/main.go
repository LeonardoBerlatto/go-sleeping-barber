package main

import (
	"go-sleeping-barber/internal/shop"
	logger "go-sleeping-barber/pkg/log"
)

var log = logger.GetLogger()

func main() {
	log.Info("Starting the sleeping barbershop example")
	shop.Run()
}
