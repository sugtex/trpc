package test

import (
	"reflect"
)

var StructGroup map[int]reflect.Type

//协议号（一对多）
const (
	PGreet =iota
)
func RegisterRoutes(){
	StructGroup=make(map[int]reflect.Type)
	StructGroup[PGreet]=reflect.TypeOf(Greet{})
}