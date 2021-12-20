package nationalid

import (
	"testing"
)

func TestNationalID(t *testing.T) {
	tables := []struct {
		value string
		ok    bool
	}{
		{value: "2130396217", ok: false},
		{value: "0000000001", ok: false},
		{value: "abcd1234", ok: false},
		{value: "نمونه", ok: false},
		{value: "۲۱۳۰۳۹۶۲۱۶", ok: false},
		{value: "", ok: false},
		{value: "-123456789", ok: false},
		{value: "-1234567890", ok: false},
		{value: "4608968882", ok: true},
		{value: "1111111111", ok: true},
		{value: "939092001", ok: true},
		{value: "0939092001", ok: true},
	}
	for _, data := range tables {
		ok, err := IsValid(data.value)
		if ok != data.ok {
			t.Errorf("NationalID(%q) = %v, ok %v,err %s", data.value, ok, data.ok, err)
		}
	}
}
