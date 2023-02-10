package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

/*
* contains information about California hotels including Name, Address, City, Zip, Region
* region can be: Southern, Central, Northern
* can hold an unlimited number of hotels
 */
// type hotel struct {
// 	Name    string
// 	Address string
// 	City    string
// 	Zip     string
// }

type hotel struct {
	Name, Address, City, Zip string
}

type region struct {
	Name   string
	Hotels []hotel
}

func main() {
	northern := region{
		Name: "Northern",
		Hotels: []hotel{
			{
				Name:    "N hotel 1",
				Address: "123 Main St",
				City:    "Oakland",
				Zip:     "12345",
			},
		},
	}

	central := region{
		Name: "Central",
		Hotels: []hotel{
			{
				Name:    "C hotel 1",
				Address: "123 Main St",
				City:    "Oakland",
				Zip:     "12345",
			},
			{
				Name:    "C hotel 2",
				Address: "123 Main St",
				City:    "Oakland",
				Zip:     "12345",
			},
		},
	}
	southern := region{
		Name: "Southern",
		Hotels: []hotel{
			{
				Name:    "S hotel 1",
				Address: "123 Main St",
				City:    "San Francisco",
				Zip:     "12345",
			},
		},
	}

	rh := []region{northern, central, southern}
	err := tpl.Execute(os.Stdout, rh)
	if err != nil {
		log.Fatal("Can't execute template")
	}
}
