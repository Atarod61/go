package main

import (
	"fmt"
	"reflect"
)

func main() {
	var number = 23.5
	fmt.Println(number)
	fmt.Println(reflect.TypeOf(number))
}
