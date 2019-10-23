// This is an example repository to demonstrate how goCleanse works.
//
// https://github.com/Krlier/goCleanse
//
package main

import (
	"fmt"

	"github.com/Krlier/test/db"
)

func main() {
	databasePassword := "S3CR3TP4SSW0RD"
	AWS_KEY := "A123455690"

	db.Connect(databasePassword)

	fmt.Println("Connected Successfully!")
}
