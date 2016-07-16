package anticaptcha

import (
	"fmt"
	"strconv"
)

type AccountService struct {
	client *Client
}

type Account struct {
	Balance float64
}

func (s *AccountService) GetBalance() (float64, error) {
	u := fmt.Sprintf("res.php?key=%s&action=getbalance", s.client.ApiKey)

	req, err := s.client.NewRequest("GET", u, nil)
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
