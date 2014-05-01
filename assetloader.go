package v1go

import "fmt"
import "reflect"
import "errors"

func QueryFor(v interface{}) error {
	rv := reflect.ValueOf(v)
	
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
	  return errors.New("Not a valid thingie!")
	}
	
	ofType := rv.Type().Elem().Elem()
	
	setter := rv.Elem()
	
	fmt.Println(setter.CanSet())
	
	newItem := reflect.New(ofType).Elem()
	newItem.Field(0).SetString("Hello")
	setter.Set(reflect.Append(setter, newItem ))
	
	
	return nil
}
