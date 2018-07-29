package main

import (
	"fmt"
)

type car struct {
	name     string
	velocity int
}

type ferrari struct {
	car   // campos anonimos
	turbo bool
}

func main() {
	f := ferrari{}
	f.name = "F40"
	f.velocity = 0
	f.turbo = true

	fmt.Printf("A ferari %s esta com turbo ligado? %v\n", f.name, f.turbo)

	fmt.Println(f)

}
