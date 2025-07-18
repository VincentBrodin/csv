package csv_test

import (
	"bytes"
	"testing"

	"github.com/VincentBrodin/csv"
)

func TestEncodeStrings(t *testing.T) {
	buf := bytes.NewBufferString("")
	encoder := csv.NewEncoder(buf)
	str := "Sesame"
	s := &struct {
		Name         string `csv:"name"`
		Email        string `csv:"email"`
		Age          int    `csv:"age"`
		StreetName   *string `csv:"street_name"`
		StreetNumber int    `csv:"street_number"`
		sendEmail    bool   `csv:"send_email"`
		Subscribed   bool   `csv:"subscribed"`
	}{
		Name:         "John Doe",
		Email:        "john.doe@example.com",
		Age:          21,
		StreetName:   &str,
		StreetNumber: 910,
		sendEmail:    true,
		Subscribed:   true,
	}

	if err := encoder.Encode(s); err != nil {
		t.Logf("Failed to encode struct, err:%s\n", err.Error())
		t.Fail()
	} else {
		t.Log(buf.String())
	}
}
