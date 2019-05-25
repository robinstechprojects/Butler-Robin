package main

import (
	"github.com/robinstechprojects/Butler-Robin/core" //Import Core to access handler Class
	"fmt" //fmt for printing & scanning
)

func main(){
	//variables
	var dckey string //Variable for User-Input
	var tokeninput string //Variable for use as Parameter to call Connection Method
	//User-Input
	fmt.Println("Bitte Token eingeben: ")
	fmt.Scan(&dckey)

	tokeninput = "Bot " + dckey //Add Bot to the Key to make it work
	core.MakeConnection(tokeninput) //Call ConnectionHandler from Class handler in Core
}