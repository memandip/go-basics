package main

import (
	"fmt"
	"math"
)

func typeSwitch(i interface{}) {
	switch t := i.(type) {
	case int:
		fmt.Printf("Square of %v is %v\n", t, math.Pow(float64(t), 2))
	case string:
		fmt.Printf("Given string value is %v and its length is %v\n", t, len(t))
	case []string:
		fmt.Printf("Given array is of length %v\n", len(t))
	case []interface{}:
		fmt.Printf("Given interface is of length %v\n", len(t))
	default:
		fmt.Printf("Given value %v is of type %T\n", t, t)
	}
}

func main() {
	typeSwitch("Mandip")
	typeSwitch([]interface{}{"Mandip", "Tharu", 9})
	typeSwitch(10)
	typeSwitch(true)
}
