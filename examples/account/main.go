package main

import (
	"fmt"

	"github.com/andrewdruzhinin/go-anti-captcha/anticaptcha"
)

func main() {
	client := anticaptcha.NewClient("your_api_key")
	account, err := client.Account.GetBalance()
	if err != nil {
		fmt.Printf("error: %v\n\n", err)
	} else {
		fmt.Printf("Balance: %f\n", account.Balance)
	}
}
