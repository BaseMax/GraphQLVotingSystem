package main

import (
	"fmt"
)

func main() {
	r := setupApiRouter()
	r.Run(":8080")
	fmt.Println("hi projeact")
}
