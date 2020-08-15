package main

import (
	"fmt"
	"w8clr/dev/fonts"
)

func main() {

	fmt.Println("w8clr: Cleanup windows")
	fmt.Println("NB: please run this under Administrator account for get results…")

	// names, err := fonts.Read()
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// 	return
	// }
	//
	// fmt.Println("Fonts")
	// for _, name := range names {
	// 	fmt.Println(" ", name)
	// }

	fmt.Println("1. Fonts")
	err := fonts.DeleteWin81()
	if err != nil {
	 	fmt.Println("Error: ", err)
	 	return
	 }

}
