package utils

import (
	"os"
	"testing"
)

func TestEncodeDecodeTable(t *testing.T) {
	os.Setenv("BASE_62", "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	cases := []struct {
		id int
	}{
		{1}, {62}, {999}, {1234567},
	}

	for _, c := range cases {
		code := Encode(c.id)
		decoded := Decode(code)
		if int(decoded) != c.id {
			t.Errorf("id=%d: expected %d, got %d", c.id, c.id, decoded)
		}
	}
}
