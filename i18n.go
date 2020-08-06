package bleach

import (
	"encoding/json"
	"io/ioutil"
)

const (
	i18nCheckLengthTooShort     = "check_length_too_short"
	i18nCheckLengthTooLong      = "check_length_too_long"
	i18nCheckLengthOutOfRange   = "check_length_out_of_range"
	i18nCheckTypeNotConvertible = "check_type_not_convertible"
)

func init() {
	LoadTranslations("translations/en_us.json")
}

func LoadTranslations(path string) error {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	dst := make(map[string]string)

	if err := json.Unmarshal(data, &dst); err != nil {
		return err
	}

	i18n.load(dst)

	return nil
}

type translations struct {
	messages map[string]string
}

var i18n = &translations{
	messages: make(map[string]string),
}

func (t *translations) load(src map[string]string) {
	for k, v := range src {
		t.messages[k] = v
	}
}

func (t *translations) get(key string) string {
	if msg, exists := t.messages[key]; exists {
		return msg
	}

	return ""
}
