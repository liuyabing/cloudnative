package main

import "fmt"

func main() {
	strArray := [5]string{"I", "am", "stupid", "and", "weak"}
	fmt.Println("before learning ", strArray)

	for i, v := range strArray {
		if v == "stupid" {
			strArray[i] = "smart"
		} else if v == "weak" {
			strArray[i] = "strong"
		}
	}

	fmt.Println("after learning", strArray)

}
