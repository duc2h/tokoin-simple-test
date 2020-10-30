package models

import "reflect"

type Search struct {
	FieldName string
	FieldType reflect.Kind
	Value     string
}
