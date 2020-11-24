package template

import (
	"html/template"
	"log"
	"os"
	"testing"
)

func TestHello(t *testing.T) {
	tml ,err := template.ParseFiles("layout.html")
	if err!=nil {
		log.Fatal(err)
	}
	log.Println("Template Name: ",tml.Name())
	tml.Execute(os.Stdout, "Hello template!")
}
