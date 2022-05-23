package hello

import "fmt"

func Print() {
	x, err := fmt.Println("こんにちは、世界")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(x)
}
