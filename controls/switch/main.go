package main

import (
	"fmt"
	"time"
)

func getType(i interface{}) string {
	switch i.(type) {
	case int:
		return "integer"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "functions"
	default:
		return "not result"
	}
}

func getSwitchCondition(i float64) string {
	switch {
	case i <= 1:
		return "bebe"
	case i < 11:
		return "criança"
	case i < 14:
		return "pré adolecente"
	case i < 17:
		return "adolecente"
	default:
		return "adulto"
	}
}

func getSwitchSimple(i float64) string {
	switch i {
	case 1:
		return "bebe"
	case 11:
		return "criança"
	case 14:
		return "pré adolecente"
	case 17:
		return "adolecente"
	default:
		return "adulto"
	}
}

func main() {
	fmt.Println(getSwitchCondition(1))
	fmt.Println(getSwitchSimple(11))
	fmt.Println(getType(2.34))
	fmt.Println(getType(2))
	fmt.Println(getType("OI"))
	fmt.Println(getType(func() {}))
	fmt.Println(getType(time.Now()))
}
