package reciveup

import (
	"fmt"
)

// HelloYourName : recive name and print
func HelloYourName() {
	var name string
	fmt.Printf("Type your name : ")
	fmt.Scanln(&name)

	fmt.Printf("Hello %s. \n", name)
}
