package anticaptcha

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestAccountService_GetBalance(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "http://anti-captcha.com/res.php?key=F629EBDA-D89A-4A0E-AAA5-069761578237&action=getbalance",
		httpmock.NewStringResponder(200, `1.0`))

	balance, err := client.Account.GetBalance()
	if err != nil {
		t.Errorf("Account.GetBalance returned error: %v", err)
	}
	assert.Equal(t, balance, 1.0)
}
