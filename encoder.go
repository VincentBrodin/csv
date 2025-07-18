package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"reflect"
)

type Encoder struct {
	writer  *csv.Writer
	headers []string
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{
		writer: csv.NewWriter(w),
	}
}

func (enc *Encoder) Encode(in any) error {
	if enc.headers == nil {
		enc.headers = getHeaders(in)
		enc.writer.Write(enc.headers)
	}

	return enc.writer.Write(marshalCsvRecord(in, enc.headers))
}

func (enc *Encoder) Flush() error {
	enc.writer.Flush()
	return enc.writer.Error()
}

func getHeaders(in any) []string {
	val := reflect.ValueOf(in).Elem()
	typ := val.Type()

	numField := typ.NumField()

	headers := make([]string, 0, numField)

	for i := range typ.NumField() {
		field := typ.Field(i)
		if !field.IsExported() {
			continue
		}
		tag := field.Tag.Get("csv")
		headers = append(headers, tag)

	}
	return headers
}

func marshalCsvRecord(in any, headers []string) []string {
	val := reflect.ValueOf(in).Elem()
	typ := val.Type()

	values := make(map[string]string)

	for i := range typ.NumField() {
		field := typ.Field(i)
		if !field.IsExported() {
			continue
		}
		tag := field.Tag.Get("csv")
		values[tag] = getField(val.Field(i))
	}

	out := make([]string, len(headers))

	for i, head := range headers {
		value, ok := values[head]
		if ok {
			out[i] = value
		}
	}

	return out
}

func getField(fieldVal reflect.Value) string {
	switch fieldVal.Kind() {
	case reflect.Pointer:
		return getField(fieldVal.Elem())
	case reflect.String:
		return fmt.Sprintf("%s", fieldVal.Interface())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%d", fieldVal.Interface())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprintf("%d", fieldVal.Interface())
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%f", fieldVal.Interface())
	case reflect.Bool:
		return fmt.Sprintf("%t", fieldVal.Interface())
	}
	return ""
}
