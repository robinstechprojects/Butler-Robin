package main

import (
	"github.com/robinstechprojects/Butler-Robin/core" //Import Core to access handler Class
	"fmt" //fmt for printing & scanning
)

func main(){

		setupLang()
		setupClient()
}

func setupLang() bool{
	var templang string
    	
	//User-Input for Language
	fmt.Println("Enter Language / Sprache Eingeben : (EN/DE)")
	fmt.Scan(&templang)
	core.SetLang(templang)
	fmt.Println("Language set to : " + templang)
	return true
}

func setupClient()bool{

	//variables
	var dckey string //Variable for User-Input
	var tokeninput string //Variable for use as Parameter to call Connection Method

	//User-Input for Token
	fmt.Println("Enter Token: ")
	fmt.Scan(&dckey)

	tokeninput = "Bot " + dckey //Add Bot to the Key to make it work
	core.MakeConnection(tokeninput) //Call ConnectionHandler from Class handler in Core
	return true
}