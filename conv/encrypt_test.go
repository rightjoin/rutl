package conv

import "testing"

func TestEncryption(t *testing.T) {
	data := []struct {
		input string
		key   string
	}{
		{"Foo", "Boo"},
		{"Bar", "Car"},
		{"Bar", ""},
		{"", "Car"},
		{"Long input with more than 16 characters", "Car"},
	}
	for _, d := range data {
		enc, err := Encrypt(d.input, d.key, true)
		if err != nil {
			t.Errorf("Unable to encrypt '%v' with key '%v': %v", d.input, d.key, err)
			continue
		}
		dec, err := Decrypt(enc, d.key)
		if err != nil {
			t.Errorf("Unable to decrypt '%v' with key '%v': %v", enc, d.key, err)
			continue
		}
		if dec != d.input {
			t.Errorf("Decrypt Key %v\n  Input: %v\n  Expect: %v\n  Actual: %v", d.key, enc, d.input, enc)
		}
	}
}
