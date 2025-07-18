package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"reflect"
	"strconv"
)

type Decoder struct {
	reader  *csv.Reader
	headers []string
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		reader: csv.NewReader(r),
	}
}

func (dec *Decoder) Decode(out any) error {
	// If we don't have any headers we set them
	if dec.headers == nil {
		var err error
		dec.headers, err = dec.reader.Read()
		if err != nil {
			return err
		}
	}

	rec, err := dec.reader.Read()
	if err != nil {
		return err
	}

	return unmarshalCsvRecord(rec, dec.headers, out)
}

func unmarshalCsvRecord(record, headers []string, out any) error {
	val := reflect.ValueOf(out).Elem()
	typ := val.Type()

	for i := range typ.NumField() {
		field := typ.Field(i)
		tag := field.Tag.Get("csv")

		// This matches the json's ignore sytstem
		if tag == "" || tag == "-" {
			continue
		}

		for idx, head := range headers {
			if head == tag {
				if err := setField(val.Field(i), record[idx]); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func setField(fieldVal reflect.Value, val string) error {
	if !fieldVal.CanSet() {
		return fmt.Errorf("Can't set the value of %s", fieldVal.Type().Name())
	}
	switch fieldVal.Kind() {
	case reflect.String:
		fieldVal.SetString(val)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		var size int
		switch fieldVal.Kind() {
		case reflect.Int:
			size = strconv.IntSize
		case reflect.Int8:
			size = 8
		case reflect.Int16:
			size = 16
		case reflect.Int32:
			size = 32
		case reflect.Int64:
			size = 64
		}
		i, err := strconv.ParseInt(val, 10, size)
		if err != nil {
			return err
		}
		fieldVal.SetInt(i)
case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		var size int
		switch fieldVal.Kind() {
		case reflect.Uint:
			size = strconv.IntSize
		case reflect.Uint8:
			size = 8
		case reflect.Uint16:
			size = 16
		case reflect.Uint32:
			size = 32
		case reflect.Uint64:
			size = 64
		}
		i, err := strconv.ParseUint(val, 10, size)
		if err != nil {
			return err
		}
		fieldVal.SetUint(i)

	case reflect.Float32, reflect.Float64:
		var size int
		switch fieldVal.Kind() {
		case reflect.Float32:
			size = 32
		case reflect.Float64:
			size = 64
		}
		f, err := strconv.ParseFloat(val, size)
		if err != nil {
			return err
		}
		fieldVal.SetFloat(f)

	case reflect.Bool:
		b, err := strconv.ParseBool(val)
		if err != nil {
			return err
		}
		fieldVal.SetBool(b)
	}
	return nil
}
