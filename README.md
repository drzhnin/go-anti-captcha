# go-anti-captcha [![Build Status](https://travis-ci.org/andrewdruzhinin/go-anti-captcha.svg?branch=master)](https://travis-ci.org/andrewdruzhinin/go-anti-captcha) [![Coverage Status](https://coveralls.io/repos/github/andrewdruzhinin/go-anti-captcha/badge.svg?branch=master)](https://coveralls.io/github/andrewdruzhinin/go-anti-captcha?branch=master)

Go library for accessing the anti-captcha.com API
## Usage ##

```go
import "github.com/andrewdruzhinin/go-anti-captcha/anticaptcha"
```
Get account balance:
```go
package main

import (
	"fmt"

	"github.com/andrewdruzhinin/go-anti-captcha/anticaptcha"
)

func main() {
	client := anticaptcha.NewClient("apiKey") //Set your apiKey from anti-captcha.com
	balance, err := client.Account.GetBalance()
	if err != nil {
		fmt.Printf("error: %v\n\n", err)
	} else {
		fmt.Printf("Balance: %f\n", balance)
	}
}

```
