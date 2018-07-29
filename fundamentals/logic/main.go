package main

import "fmt"

func buy(work1 bool, work2 bool) (bool, bool, bool) {
	tv50 := work1 && work2
	tv32 := work1 != work2 // ou exclusivo,
	iceCream := work1 || work2

	return tv50, tv32, iceCream
}

func main() {

	tv50, tv32, iceCream := buy(true, true)
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t",
		tv50, tv32, iceCream, !iceCream)
}
