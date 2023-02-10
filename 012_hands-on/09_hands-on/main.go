package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type stuff struct {
	Date, Open, High, Low, Close, Volume, AdjClose string
}

func main() {
	// bs, err := os.ReadFile("table.csv")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// s := strings.Split(string(bs), "/n")
	// fmt.Println("first in byteslice %S", len(s))
	f, err := os.Open("table.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	var s []stuff
	for _, line := range lines {
		item := stuff{
			Date: line[0],
			Open: line[1],
		}
		// item := stuff{line[0], line[1]}
		s = append(s, item)
	}

	tpl.Execute(os.Stdout, s)

}
