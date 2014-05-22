
package main

import (
	"log"
	"reflect"
	"encoding/json"
	"fmt"
)
type Foo struct {
	FirstName string `tag_name:"tag 1"`
	LastName  string `tag_name:"tag 2"`
	Age       int    `tag_name:"tag 3"`
}

func (f *Foo) reflect() {
	val := reflect.ValueOf(f).Elem()
 
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		tag := typeField.Tag
 
		fmt.Printf("Field Name: %s,\t Field Value: %v,\t Tag Value: %s\n", typeField.Name, valueField.Interface(), tag.Get("tag_name"))
	}
}

func main() {

	log.Println("hello")
	//common_args = make([]reflect.Type)
	args := make(map[string]int)
	//key := reflect.TypeOf( "" )
	//args["zhong"] =  reflect.TypeOf( 12 )

	values := []byte("{\"zhong\":12}")

	json.Unmarshal(values, args)
	log.Println(args["zhong"])

	word := "中国"
	log.Println(len(word))

	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())

	type T struct {
		A int
		B string
	}

	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}

	f := &Foo{
		FirstName: "Drew",
		LastName:  "Olson",
		Age:       30,
	}
 
	f.reflect()
}