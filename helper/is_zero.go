package helper

import "reflect"

func IsZero(source any) bool {
	return reflect.ValueOf(source).IsZero()
}
