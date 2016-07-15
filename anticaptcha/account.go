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

func (s *AccountService) GetBalance() (*Account, error) {
	u := fmt.Sprintf("res.php?key=%s&action=getbalance", s.client.ApiKey)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	data, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	account := new(Account)
	account.Balance, err = strconv.ParseFloat(string(data), 64)
	if err != nil {
		return nil, err
	}
	return account, err
}
