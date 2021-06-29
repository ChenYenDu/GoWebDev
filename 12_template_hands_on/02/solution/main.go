package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type hotel struct {
	Name, Address, City, Zip, Region string
}

type region struct {
	Region string
	Hotels []hotel
}

type Region []region

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	hs := Region{
		{
			Region: "Southern",
			Hotels: []hotel{
				{
					Name:    "Hotel California",
					Address: "42 Sunset Boulevard",
					City:    "Los Angeles",
					Zip:     "95612",
					Region:  "southern",
				},
				{
					Name:    "Hotel California",
					Address: "42 Sunset Boulevard",
					City:    "Los Angeles",
					Zip:     "95612",
					Region:  "southern",
				},
			},
		},
		{
			Region: "Northern",
			Hotels: []hotel{
				{
					Name:    "Hotel California",
					Address: "42 Sunset Boulevard",
					City:    "Los Angeles",
					Zip:     "95612",
					Region:  "northern",
				},
				{
					Name:    "Hotel California",
					Address: "42 Sunset Boulevard",
					City:    "Los Angeles",
					Zip:     "95612",
					Region:  "northern",
				},
			},
		},
		{
			Region: "Central",
			Hotels: []hotel{
				{
					Name:    "Hotel California",
					Address: "42 Sunset Boulevard",
					City:    "Los Angeles",
					Zip:     "95612",
					Region:  "central",
				},
				{
					Name:    "Hotel California",
					Address: "42 Sunset Boulevard",
					City:    "Los Angeles",
					Zip:     "95612",
					Region:  "central",
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, hs)

	if err != nil {
		log.Fatalln(err)
	}
}
