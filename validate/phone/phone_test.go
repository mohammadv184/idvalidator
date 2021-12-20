package phone

import (
	"testing"
)

func TestPhone(t *testing.T) {
	tables := []struct {
		value string
		ok    bool
	}{
		{value: "2130396217", ok: false},
		{value: "+9893909200111", ok: false},
		{value: "0098939092011", ok: false},
		{value: "09390920011", ok: true},
		{value: "+989390920011", ok: true},
		{value: "00989390920011", ok: true},
	}
	for _, data := range tables {
		ok, err := IsValid(data.value)
		if ok != data.ok {
			t.Errorf("PhoneNumber(%q) = %v, ok %v,err %s", data.value, ok, data.ok, err)
		}
	}
}
