package main

import (
	"log"

	"github.com/andrewdruzhinin/go-anti-captcha/anticaptcha"
)

func main() {
	client := anticaptcha.NewClient("apiKey") //Set your apiKey
	balance, err := client.Account.GetBalance()
	if err != nil {
		log.Fatalf("error: %v\n", err)
	} else {
		log.Printf("Balance: %f\n", balance)
	}
	sysStat, err := client.GetSystemStat()
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	log.Printf("Amount of workers waiting for a captcha from English Queue: %v", sysStat.Waiting)
	log.Printf("Amount of workers waiting for a captcha from Russian Queue: %v", sysStat.WaitingRU)
	log.Printf("Represents current demand/supply ratio: %v", sysStat.Load)
	log.Printf("Minimum bid required to enter English Queue. Account's maximum bid must be higher than this value.: %v", sysStat.Minbid)
	log.Printf("Minimum bid required to enter Russian Queue: %v", sysStat.MinbidRU)
	log.Printf("Average captcha recognition time in English Queue: %v", sysStat.AverageRecognitionTime)
	log.Printf("Average captcha recognition time in Russian Queue: %v", sysStat.AverageRecognitionTimeRU)
}
