package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string `name tag`
	age  int    `age tag`
}

func main() {
	var float float64 = 5.5
	floatType := reflect.TypeOf(float)
	floatValue := reflect.ValueOf(float)
	fmt.Println(floatType)
	fmt.Println(floatValue.Kind())
	fmt.Printf("%T\n", float)
	fmt.Println(floatValue.Type().String())
	fmt.Println("floatType: ", floatType.String())

	var person1 = Person{"She", 18}
	ty1 := reflect.TypeOf(&person1)
	val1 := reflect.ValueOf(&person1)
	fmt.Println(ty1.Elem().Field(0).Tag)
	fmt.Println(val1.Elem().Field(1).Int())

	var person Person
	ty := reflect.TypeOf(person)
	val := reflect.ValueOf(person)
	fmt.Println(ty.NumField())
	fmt.Println(val.NumField())
	fmt.Println(val.Type())
	fmt.Println(ty.Field(0).Name)
	fmt.Println(val.Type().Field(0))
	fmt.Printf("%+v", val.Type().Field(0))
	fmt.Println(val.Type().Field(0).Tag)
	fmt.Println(val.Type().Field(0).Name)
}
