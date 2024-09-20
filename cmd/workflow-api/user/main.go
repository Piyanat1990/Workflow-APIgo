package main

import (
	"fmt"

	"github.com/Piyanat1990/workflow/internal/user"
)

func main() {
	password, err := user.HashPassword("secret")
	if err!=nil{
		fmt.Println("err",err.Error())
	}
	fmt.Println("password",password)

}