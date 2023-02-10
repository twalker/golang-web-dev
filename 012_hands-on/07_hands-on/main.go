package main

import (
	"log"
	"os"
	"text/template"
)

type restaurant struct {
	Name, Location string
	Menu           items
}

type item struct {
	Name, Descrip, Meal string
	Price               float64
}

type items []item

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	r1 := restaurant{
		Name:     "4 Spoons",
		Location: "Greenwood",
		Menu: items{
			{
				Name:    "Oatmeal",
				Descrip: "yum yum",
				Meal:    "Breakfast",
				Price:   4.95,
			},
			{
				Name:    "Hamburger",
				Descrip: "Delicous good eating for you",
				Meal:    "Lunch",
				Price:   6.95,
			},
		},
	}
	r2 := restaurant{
		Name:     "Paseo's",
		Location: "Shilshole",
		Menu: items{
			{
				Name:    "Pasta Bolognese",
				Descrip: "From Italy delicious eating",
				Meal:    "Dinner",
				Price:   7.95,
			},
		},
	}

	err := tpl.Execute(os.Stdout, []restaurant{r1, r2})
	if err != nil {
		log.Fatalln(err)
	}
}
