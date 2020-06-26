// https://blog.golang.org/laws-of-reflection

package reflect

import (
	"fmt"
	"reflect"
)

type Programmer struct {
	Name     string   `tag:"name"`
	Age      uint     `tag:"age"`
	LangList []string `tag:"lang_list"`
	Salary   float32  `tag:"salary"`
}

func InspectStructFields(s interface{}) {
	fmt.Printf("inspecting strcut: %+v\n", s)
	t := reflect.TypeOf(s)

	if t.Kind() != reflect.Struct {
		fmt.Println(s, " is not a struct")
		return
	}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("field: %s, type: %v\n", f.Name, f.Type.Kind())
	}
}

func GetStructTag(s interface{}) {
	structValue := reflect.ValueOf(s)
	structType := reflect.TypeOf(s)
	for i := 0; i < structValue.NumField(); i++ {
		f := structValue.Field(i)
		value := f.Interface()
		tp := structType.Field(i)
		fmt.Printf("field: %v, type: %v, tagName: %s\n", value, tp.Type.Kind(), tp.Tag.Get("tag"))
	}
}

func TypeOf(values ...interface{}) {
	for _, v := range values {
		typeOfValue := reflect.TypeOf(v)
		fmt.Printf("val: %v, type: %v\n", v, typeOfValue.Kind())
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
