// https://blog.golang.org/laws-of-reflection

package reflect

import (
	"fmt"
	"reflect"
)

func TypeOf(values ...interface{}) {
	for _, v := range values {
		fmt.Printf("val: %v, type: %v\n", v, reflect.TypeOf(v))
	}
}

func ValueOf(values ...interface{}) {
	for _, value := range values {
		reflectVal := reflect.ValueOf(value)
		fmt.Printf("val: %v, type: %v, kind: %v\n", reflectVal, reflectVal.Type(), reflectVal.Kind())
	}
}

func Interface(values ...interface{}) {
	for _, value := range values {
		reflectVal := reflect.ValueOf(value)
		inter := reflectVal.Interface()
		valType := reflectVal.Type()
		fmt.Printf("val: %v, interface: %v\n", reflectVal, inter)

		// inter is interface{}, can use type switch
		switch inter.(type) {
		case string:
			fmt.Println(value, " is string")
		case int, int32, int64, uint, uint8, uint32, uint64:
			fmt.Println(value, " is integer")
		case float64, float32:
			fmt.Println(value, " is float")
		case bool:
			fmt.Println(value, " is boolean")
		default:
			fmt.Println(value, " is not basic data type, type: ", valType)
		}
	}
}
