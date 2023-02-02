package config

import (
	"encoding/base64"
	"fmt"
	"reflect"

	"github.com/mitchellh/mapstructure"
)

func byteArrayFromBase64StringHook() mapstructure.DecodeHookFuncType {
	return func(
		f reflect.Type,
		t reflect.Type,
		data any,
	) (any, error) {
		// Check that the target type is a byte array
		if !(t.Kind() == reflect.Array && t.Elem().Kind() == reflect.Uint8) {
			return data, nil
		}

		// Check that the data is string
		if f.Kind() != reflect.String {
			return data, fmt.Errorf("byte array type only supports base64 encoded strings as input. Provided type is %s", f.Kind())
		}

		// Parse content as base64
		value, err := base64.StdEncoding.DecodeString(data.(string))
		if err != nil {
			return nil, fmt.Errorf("failed to parse data as base64: %w", err)
		}

		// Validate data length
		if len(value) != t.Len() {
			return nil, fmt.Errorf("provided base64 has length %d, but length %d is expected", len(value), t.Len())
		}

		// Convert slice to array
		result := reflect.New(t).Elem()
		reflect.Copy(result, reflect.ValueOf(value))

		// Parse successful
		return result.Interface(), nil
	}
}
