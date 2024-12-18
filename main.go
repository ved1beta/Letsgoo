package main

import (
	"fmt"
	"frmBasics/imports"
)

func main() {
	newTicket := imports.Ticket{
		ID:    123,
		Event: "FEM course",
	}
	fmt.Println(newTicket)
}
