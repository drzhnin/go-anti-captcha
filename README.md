# go-anti-captcha [![Build Status](https://travis-ci.org/drzhnin/go-anti-captcha.svg?branch=master)](https://travis-ci.org/andrewdruzhinin/go-anti-captcha) [![Coverage Status](https://coveralls.io/repos/github/andrewdruzhinin/go-anti-captcha/badge.svg?branch=master)](https://coveralls.io/github/andrewdruzhinin/go-anti-captcha?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/andrewdruzhinin/go-anti-captcha)](https://goreportcard.com/report/github.com/andrewdruzhinin/go-anti-captcha)

Go library for accessing the anti-captcha.com API
## Usage ##
See
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

Upload captcha from url and get text:
```go
package main

import (
	"fmt"

	"github.com/andrewdruzhinin/go-anti-captcha/anticaptcha"
)

func main() {
	client := anticaptcha.NewClient("apiKey") //Set your apiKey
	ID, err := client.Captcha.UploadCaptchaFromURL("https://s3-us-west-2.amazonaws.com/captcha-test/1045.png")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Printf("Captcha ID: %d\n", ID)
		res, err := client.Captcha.GetText(ID)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
		fmt.Printf("Captcha Text: %s\n", res)
	}

}
```
Take a look at ./examples/ to know more how to use anti-captcha api.

## Additional captcha parameters ##
You can use optional captcha parameters:
```go
package main

import (
	"fmt"

	"github.com/andrewdruzhinin/go-anti-captcha/anticaptcha"
)

func main() {
	client := anticaptcha.NewClient("APIKey")
	client.CaptchaAdditionalParams.EnablePhrate()
	client.CaptchaAdditionalParams.EnableAllowEmpty()
	client.CaptchaAdditionalParams.AddComment("What color is the sky?")
	fmt.Println(client.CaptchaAdditionalParams.LoadParams())
	ID, err := client.Captcha.UploadCaptchaFromURL("https://s3-us-west-2.amazonaws.com/captcha-test/1045.png")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Printf("Captcha ID: %d\n", ID)
		res, err := client.Captcha.GetText(ID)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
		fmt.Printf("Captcha Text: %s\n", res)
}
```
Parameter |	Type |	Possible values |	Description
------------ | ------------- | -------------|  -------------
phrase |	integer | 0, 1 | 0 = default value, 1 = captcha has 2-3 words
regsense |	integer | 0, 1 | 0 = default value, 1 = captcha is case sensitive
numeric |	integer | 0, 1, 2 | 0 = default value, 1 = captcha consists of digits only, 2 = captcha does not contain any digits
calc |	integer | 0, 1 | 0 = default value, 1 = arithmetical operation must be performed
min_len |	integer | 0..20 | 0 = default value, 1..20 = minimum length of captcha text required to input
max_len |	integer | 0..20 | 0 = default value, 1..20 = maximum length of captcha text required to input
is_russian | integer | 0, 1 | 0 = default value, 1 = captcha goes to Russian Queue
soft_id	| integer | | AppCenter Application ID used for comission earnings
header_acao |	integer | 0, 1 | 0 = default value, 1 = API sends Access-Control-Allow-Origin: * parameter in response header. (Required for cross-domain AJAX requests from client-side applications).
type | string | recaptcha2, recaptcha2_44, recaptcha2_24, audio | empty = default value, recaptcha2 = Use this value for Recaptcha2 images. Image must have size ratio 1x1, minimum height 200 pixels and come along with "comment" parameter. This is where you should specify English name of the object which worker must choose (cat, road sign, burger, etc.). See workers interface screenshot. recaptcha2_44 = Same thing as "recaptcha2" but we will put this captcha in a mask of 16 squares (4 x 4). recaptcha2_24 = Recaptcha2 in a mask of 8 squares (2 x 4). It is very important that you send only captcha image without blue heading or any comments embeded in image. Examples of how some of our clients misunderstand this you can see here and here. audio = Use this value to send audio captchas in MP3 format.
comment	| string max (100 bytes) | | empty = default value, Option 1. Send along with any captcha to make it more clear for workers. Option 2. Send without captcha if you just want to ask a question (example: "What color is the sky?")
allow_empty | integer | 0, 1 | 0 = default value, 1 = Allow empty response for Google Recaptcha. This is useful if you want to allow our workers to mark captcha as unsolvable because no matching objects were found on your captcha. If you send value "1", workers will have a button "no objects found" with each recaptcha marked this way.. If worker pushed this button, API will return text "EMPTY_ANSWER" as captcha result.
