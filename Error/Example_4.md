package main

import (
	"errors"
	"fmt"
)

func Divide (a  int , b int) (int , error) {
	if b == 0 {
		return 0 , errors.New("The Value Of b Does Not be a Zero ")
	}

	return a / b , nil

}
func main () {
	result , err := Divide(10 , 8) 

	if err != nil  {
		fmt.Println(err)
	}else {
		fmt.Println(result)
	}

	esult , err := Divide(10 , 0) 
	
	if err != nil  {
		fmt.Println(err)
	}else {
		fmt.Println(esult)
	}

}