package anticaptcha

import (
	"fmt"
	"strconv"
)

//AccountService handles communication with the account actions of the
//anti-captcha API
type AccountService service

//Account represents account struct
type Account struct {
	Balance float64
}

//GetBalance represents balance from account
func (s *AccountService) GetBalance() (float64, error) {
	reqURL := fmt.Sprintf("res.php?key=%s&action=getbalance", s.client.APIKey)

	req, err := s.client.NewRequest("GET", reqURL, nil)
	if err != nil {
		return 0.0, err
	}

	data, err := s.client.Do(req)
	if err != nil {
		return 0.0, err
	}

	account := new(Account)
	account.Balance, err = strconv.ParseFloat(string(data), 64)
	if err != nil {
		return 0.0, err
	}
	return account.Balance, err
}

// //GetStats represents statistics from account
// func (s *AccountService) GetStats() {
//
// }
