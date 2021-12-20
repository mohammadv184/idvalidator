package card

import (
	"testing"
)

func TestCard(t *testing.T) {
	tables := []struct {
		value string
		ok    bool
	}{
		{value: "abcdefghi1234555", ok: false},
		{value: "622106104944", ok: false},
		{value: "3221061049447982", ok: false},
		{value: "6221061049447983", ok: false},
		{value: "۶۲۲۱۰۶۱۰۴۹۴۴۷۹۸۲", ok: false},
		{value: "6221061049447982", ok: true},
		{value: "6395991166685826", ok: true},
	}
	for _, data := range tables {
		ok, err := IsValid(data.value)
		if ok != data.ok {
			t.Errorf("CardNumber(%q) = %v, ok %v,err %s", data.value, ok, data.ok, err)
		}
	}
}
