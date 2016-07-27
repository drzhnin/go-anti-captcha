package anticaptcha

import (
	"log"
	"reflect"
	"strconv"
)

//AdditionalParams handles communication with the captcha additional params
type AdditionalParams service

//CaptchaParams contains additional parametrs for captcha
type CaptchaParams struct {
	Phrase      int    `t:"phrase"`
	Regsense    int    `t:"regsense"`
	Calc        int    `t:"calc"`
	IsRussian   int    `t:"is_russian"`
	HeaderAcao  int    `t:"header_acao"`
	AllowEmpty  int    `t:"allow_empty"`
	Numeric     int    `t:"numeric"`
	MinLength   int    `t:"min_len"`
	MaxLength   int    `t:"max_len"`
	SoftID      int    `t:"soft_id"`
	CaptchaType string `t:"type" default:""`
	Comment     string `t:"comment" default:""`
}

var params CaptchaParams

//ResetToDefault represents reset params to default values
func (ap *AdditionalParams) ResetToDefault() {
	params.Phrase = 0
	params.Regsense = 0
	params.Calc = 0
	params.IsRussian = 0
	params.HeaderAcao = 0
	params.AllowEmpty = 0
	params.Numeric = 0
	params.MinLength = 0
	params.MaxLength = 0
	params.SoftID = 0
	params.CaptchaType = ""
	params.Comment = ""
}

//LoadParams represents return all params
func (ap *AdditionalParams) LoadParams() map[string]string {
	m := params.toMap("t")
	return m
}

//EnablePhrate represents captcha has 2-3 words
func (ap *AdditionalParams) EnablePhrate() {
	params.Phrase = 1
}

//DisablePhrate represents disable option Phrate
func (ap *AdditionalParams) DisablePhrate() {
	params.Phrase = 0
}

//EnableRegsense represents captcha is case sensitive
func (ap *AdditionalParams) EnableRegsense() {
	params.Regsense = 1
}

//DisableRegsense represents disable option Regsense
func (ap *AdditionalParams) DisableRegsense() {
	params.Regsense = 0
}

//EnableCalc represents arithmetical operation must be performed
func (ap *AdditionalParams) EnableCalc() {
	params.Calc = 1
}

//DisableCalc represents disable option Calc
func (ap *AdditionalParams) DisableCalc() {
	params.Calc = 0
}

//EnableIsRussian represents captcha goes to Russian Queue
func (ap *AdditionalParams) EnableIsRussian() {
	params.IsRussian = 1
}

//DisableIsRussian represents disable option IsRussian
func (ap *AdditionalParams) DisableIsRussian() {
	params.IsRussian = 0
}

//EnableHeaderAcao represents API sends Access-Control-Allow-Origin: * parameter in response header. (Required for cross-domain AJAX requests from client-side applications).
func (ap *AdditionalParams) EnableHeaderAcao() {
	params.HeaderAcao = 1
}

//DisableHeaderAcao represents disable option HeaderAcao
func (ap *AdditionalParams) DisableHeaderAcao() {
	params.HeaderAcao = 0
}

//EnableAllowEmpty represents Allow empty response for Google Recaptcha.
//This is useful if you want to allow our workers to mark captcha as unsolvable because no matching objects were found on your captcha.
//If you send value "1", workers will have a button "no objects found" with each recaptcha marked this way..
//If worker pushed this button, API will return text "EMPTY_ANSWER" as captcha result.
func (ap *AdditionalParams) EnableAllowEmpty() {
	params.AllowEmpty = 1
}

//DisableAllowEmpty represents disable AllowEmpty option
func (ap *AdditionalParams) DisableAllowEmpty() {
	params.AllowEmpty = 0
}

//CaptchaContainsNumericOnly represents captcha consists of digits only
func (ap *AdditionalParams) CaptchaContainsNumericOnly() {
	params.Numeric = 1
}

//CaptchaNotContainAnyDigits represents captcha does not contain any digits
func (ap *AdditionalParams) CaptchaNotContainAnyDigits() {
	params.Numeric = 2
}

//CaptchaContainsAnyCharacters represents set default value for Numeric option
func (ap *AdditionalParams) CaptchaContainsAnyCharacters() {
	params.Numeric = 0
}

//SetMinCaptchaLength represents set minimum length of captcha text required to input
func (ap *AdditionalParams) SetMinCaptchaLength(val int) {
	if val < 1 || val > 20 {
		params.MinLength = 0
	} else {
		params.MinLength = val
	}

}

//ResetMinCaptchaLengthToDefault represents reset option MinLength to default value
func (ap *AdditionalParams) ResetMinCaptchaLengthToDefault() {
	params.MinLength = 0
}

//SetMaxCaptchaLength represents set maximum length of captcha text required to input
func (ap *AdditionalParams) SetMaxCaptchaLength(val int) {
	if val < 1 || val > 20 {
		params.MaxLength = 0
	} else {
		params.MaxLength = val
	}

}

//ResetMaxCaptchaLengthToDefault represents reset option MsxLength to default value
func (ap *AdditionalParams) ResetMaxCaptchaLengthToDefault() {
	params.MaxLength = 0
}

//SetSoftID represents set AppCenter Application ID used for comission earnings
func (ap *AdditionalParams) SetSoftID(val int) {
	params.SoftID = val
}

//SetCaptchaType represents empty = default value
//recaptcha2 = Use this value for Recaptcha2 images. Image must have size ratio 1x1, minimum height 200 pixels and come along with "comment" parameter.
//This is where you should specify English name of the object which worker must choose (cat, road sign, burger, etc.).
//recaptcha2_44 = Same thing as "recaptcha2" but we will put this captcha in a mask of 16 squares (4 x 4).
//recaptcha2_24 = Recaptcha2 in a mask of 8 squares (2 x 4).
//It is very important that you send only captcha image without blue heading or any comments embeded in image.
//audio = Use this value to send audio captchas in MP3 format.
func (ap *AdditionalParams) SetCaptchaType(val string) {
	params.CaptchaType = val
}

//AddComment represents empty = default value
//Option 1. Send along with any captcha to make it more clear for workers.
//Option 2. Send without captcha if you just want to ask a question (example: "What color is the sky?")
func (ap *AdditionalParams) AddComment(val string) {
	if len(val) > 100 {
		params.Comment = ""
	} else {
		params.Comment = val
	}

}

//DebugParams represents show all captcha additional params
func (ap *AdditionalParams) DebugParams() {
	log.Printf("%#v\n", params)
}

//toMap uses tags on struct fields to decide which fields to add to the
//returned map.
func (s *CaptchaParams) toMap(tag string) map[string]string {
	out := make(map[string]string)

	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// gets us a StructField
		fi := typ.Field(i)
		if tagv := fi.Tag.Get(tag); tagv != "" {
			data := v.Field(i).Interface()
			switch data.(type) {
			case int:
				if data != 0 {
					out[tagv] = strconv.Itoa(data.(int))
				}
			case string:
				if data != "" {
					out[tagv] = data.(string)
				}
			}
		}
	}
	return out
}
