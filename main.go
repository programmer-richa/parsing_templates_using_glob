package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

// init() it initialises tpl variable and executed while program initialisation.
// This practise makes our program more efficient.
func init() {
	// Passes all files available in the template directory with extension .gohtml
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}
func main() {
	// Passes all files available in the template directory with extension .gohtml
	// tpl, err := template.ParseGlob("templates/*.gohtml")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// Executes template first.gohtml
	err := tpl.ExecuteTemplate(os.Stdout, "first.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}

	// Executes template second.gohtml
	err = tpl.ExecuteTemplate(os.Stdout, "second.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}
}
