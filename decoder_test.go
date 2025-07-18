package csv_test

import (
	"io"
	"strings"
	"testing"

	"github.com/VincentBrodin/csv"
)

func TestDecodeStrings(t *testing.T) {

	csvStr := "a,b,c,d,e\n1,2,3,4,5\napple,bannana,pear,orange,mango"
	reader := strings.NewReader(csvStr)

	decoder := csv.NewDecoder(reader)

	for {
		s := &struct {
			A string `csv:"a"`
			B string `csv:"b"`
			C string `csv:"c"`
			D string `csv:"d"`
			E string `csv:"e"`
		}{}
		if err := decoder.Decode(s); err == io.EOF {
			break
		} else if err != nil {
			t.Logf("Failed to decode struct, err:%s\n", err.Error())
			t.Fail()
		} else if s.A == "" || s.B == "" || s.C == "" || s.D == "" || s.E == "" {
			t.Log("Values are nil")
			t.Fail()
		} else {
			t.Log(*s)
		}
	}

}

func TestDecodeInts(t *testing.T) {
	csvStr := "a,b,c,d,e\n1,2,3,4,5\n41,42,43,44,45"
	reader := strings.NewReader(csvStr)

	decoder := csv.NewDecoder(reader)

	for {
		s := &struct {
			A int   `csv:"a"`
			B int8  `csv:"b"`
			C int16 `csv:"c"`
			D int32 `csv:"d"`
			E int64 `csv:"e"`
		}{}
		if err := decoder.Decode(s); err == io.EOF {
			break
		} else if err != nil {
			t.Logf("Failed to decode struct, err:%s\n", err.Error())
			t.Fail()
		} else if s.A == 0 || s.B == 0 || s.C == 0 || s.D == 0 || s.E == 0 {
			t.Log("Values are nil")
			t.Fail()
		} else {
			t.Log(*s)
		}
	}

}

func TestDecodeUints(t *testing.T) {
	csvStr := "a,b,c,d,e\n1,2,3,4,5\n41,42,43,44,45"
	reader := strings.NewReader(csvStr)

	decoder := csv.NewDecoder(reader)

	for {
		s := &struct {
			A uint   `csv:"a"`
			B uint8  `csv:"b"`
			C uint16 `csv:"c"`
			D uint32 `csv:"d"`
			E uint64 `csv:"e"`
		}{}
		if err := decoder.Decode(s); err == io.EOF {
			break
		} else if err != nil {
			t.Logf("Failed to decode struct, err:%s\n", err.Error())
			t.Fail()
		} else if s.A == 0 || s.B == 0 || s.C == 0 || s.D == 0 || s.E == 0 {
			t.Log("Values are nil")
			t.Fail()
		} else {
			t.Log(*s)
		}
	}

}

func TestDecodeFloats(t *testing.T) {
	csvStr := "a,b\n1,2\n3.123,4.456"
	reader := strings.NewReader(csvStr)

	decoder := csv.NewDecoder(reader)

	for {
		s := &struct {
			A float32 `csv:"a"`
			B float64 `csv:"b"`
		}{}
		if err := decoder.Decode(s); err == io.EOF {
			break
		} else if err != nil {
			t.Logf("Failed to decode struct, err:%s\n", err.Error())
			t.Fail()
		} else if s.A == 0 || s.B == 0 {
			t.Log("Values are nil")
			t.Fail()
		} else {
			t.Log(*s)
		}
	}

}

func TestDecodeBools(t *testing.T) {
	csvStr := "a\ntrue"
	reader := strings.NewReader(csvStr)

	decoder := csv.NewDecoder(reader)

	s := &struct {
		A bool `csv:"a"`
	}{}

	if err := decoder.Decode(s); err != nil {
		t.Logf("Failed to decode struct, err:%s\n", err.Error())
		t.Fail()
	} else if s.A == false {
		t.Log("Value is nil")
		t.Fail()
	} else {
		t.Log(*s)
	}

}
