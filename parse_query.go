package eliotlibs

import (
	"fmt"
	"log"
	"net/url"
	"reflect"
)

func iterator(modelValid interface{}, data string) error {

	v := reflect.ValueOf(modelValid).Elem()

	// Iterar sobre los campos
	values, err := url.ParseQuery(data)
	if err != nil {
		return err
	}
	t := v.Type()
	log.Println(t)
	log.Println(values)

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)

		// Obtener la etiqueta `query`
		queryTag := field.Tag.Get("query")
		value, ok := values[queryTag]
		if ok && fieldValue.CanSet() {
			if len(value) >= 1 {
			}
			fieldValue.SetString(value[0])
		}

		// Imprimir el valor del campo y la etiqueta asociada
		fmt.Printf("Campo: %s, Valor: %v, Query Tag: %s\n", field.Name, fieldValue.Interface(), queryTag)
	}

	return nil
}
