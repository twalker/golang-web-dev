package main

import (
	"html/template"
	"log"
	"os"
)

type Dish struct {
	Meal string
	Name string
	Cost float32
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	dishes := []Dish{
		{
			Meal: "Dinner",
			Name: "Pizza",
			Cost: 12.50,
		},
		{
			Meal: "Breakfast",
			Name: "Pancakes",
			Cost: 8.99,
		},
	}

	err := tpl.Execute(os.Stdout, dishes)
	if err != nil {
		log.Fatalln("template can't be executed")
	}
}
