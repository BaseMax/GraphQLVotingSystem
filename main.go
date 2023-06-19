package main

import (
	"fmt"
)

func main() {
	r := setupRouter()
	r.Run(":8080")
	fmt.Println("hi projeact")
}
