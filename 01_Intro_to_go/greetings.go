package main

import (
"fmt"
"text/template"
"os"
)

func main(){
	tpl, err := template.ParseFiles("letter.phpasp")

	if err != nil {
		fmt.Println("There was an error parsing file", err)
	}

	friends := []string{"Anshuman", "Anusha", "Abhijay"}

	err = tpl.Execute(os.Stdout, friends)
	if err != nil {
		fmt.Println("There was an error executing template", err)
	}
}