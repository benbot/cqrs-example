package events

import (
	"errors"
	"reflect"
)

type DeconstructedEvent struct {
	Type  string
	Event map[string]interface{}
}

func (d *DeconstructedEvent) ConvertTo(obj interface{}) (interface{}, error) {
	for s, v := range d.Event {
		if err := setField(obj, s, v); err != nil {
			return nil, err
		}
	}

	return obj, nil
}

func setField(obj interface{}, name string, value interface{}) error {
	structType := reflect.TypeOf(obj)
	structField, found := structType.FieldByNameFunc(func(n string) bool {
		f, found := structType.FieldByName(n)
		if !found {
			return false
		}

		t := f.Tag.Get("bson")

		if t != name {
			return false
		}

		return true
	})

	if !found {
		return errors.New("Field not found")
	}

	v := reflect.ValueOf(obj).Elem()
	f := v.FieldByIndex(structField.Index)

	if !f.IsValid() {
		return errors.New("struct field not valid")
	}

	if !f.CanSet() {
		return errors.New("can't set field")
	}

	val := reflect.ValueOf(value)

	if val.Type() != f.Type() {
		return errors.New("fields are of different types")
	}

	v.FieldByIndex(structField.Index).Set(val)

	return nil
}
