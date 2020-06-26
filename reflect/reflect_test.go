package reflect

import "testing"

type TestStruct struct {
}

type MyInteger int32

/*
val: 1, type: int
val: 6, type: int32
val: 9, type: uint
val: 100, type: uint32
val: 1000, type: int64
val: ABC, type: string
val: true, type: bool
val: false, type: bool
val: &{}, type: ptr
*/
func TestTypeOf(t *testing.T) {
	TypeOf(
		1, int32(6), uint(9), uint32(100), int64(1000),
		"ABC", true, false, &TestStruct{},
	)
}

/*
val: 1, type: int, kind: int
val: 6, type: int32, kind: int32
val: 9, type: uint, kind: uint
val: 100, type: uint32, kind: uint32
val: 1000, type: int64, kind: int64
val: ABC, type: string, kind: string
val: true, type: bool, kind: bool
val: false, type: bool, kind: bool
val: &{}, type: *reflect.TestStruct, kind: ptr
val: 1000, type: reflect.MyInteger, kind: int32
*/
func TestValueOf(t *testing.T) {
	ValueOf(
		1, int32(6), uint(9), uint32(100), int64(1000),
		"ABC", true, false, &TestStruct{},
		MyInteger(1000),
	)
}

/*
val: 1, interface: 1
1  is integer
val: 6, interface: 6
6  is integer
val: 9, interface: 9
9  is integer
val: 100, interface: 100
100  is integer
val: 1000, interface: 1000
1000  is integer
val: ABC, interface: ABC
ABC  is string
val: true, interface: true
true  is boolean
val: false, interface: false
false  is boolean
val: &{}, interface: &{}
&{}  is not basic data type, type:  *reflect.TestStruct
val: 1000, interface: 1000
1000  is not basic data type, type:  reflect.MyInteger
*/
func TestInterface(t *testing.T) {
	Interface(
		1, int32(6), uint(9), uint32(100), int64(1000),
		"ABC", true, false, &TestStruct{},
		MyInteger(1000),
	)
}

func TestInspectStructFields(t *testing.T) {
	InspectStructFields(Programmer{
		Name:     "Bob",
		Age:      26,
		LangList: []string{"GO", "Java", "Python"},
		Salary:   1000000,
	})
}

func TestGetStrcutTag(t *testing.T) {
	GetStructTag(Programmer{
		Name:     "Bob",
		Age:      26,
		LangList: []string{"GO", "Java", "Python"},
		Salary:   1000000,
	})
}
