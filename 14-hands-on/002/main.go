package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

type hotel struct {
	Name, Address, City, Zip, Region string
}

type region struct {
	Region string
	Hotels []hotel
}

type Regions []region

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	h := Regions{

		region{
			Region: "Southern",
			Hotels: []hotel{
				hotel{
					Name:    "Hotel California",
					Address: "42 Sunset Boulevard",
					City:    "Los Angeles",
					Zip:     "95612",
					Region:  "southern",
				},
				hotel{
					Name:    "H",
					Address: "4",
					City:    "L",
					Zip:     "95612",
					Region:  "southern",
				},
			},
		},

		region{
			Region: "Northern",
			Hotels: []hotel{
				hotel{
					Name:    "Hotel California",
					Address: "42 Sunset Boulevard",
					City:    "Los Angeles",
					Zip:     "95612",
					Region:  "southern",
				},
				hotel{
					Name:    "H",
					Address: "4",
					City:    "L",
					Zip:     "95612",
					Region:  "southern",
				},
			},
		},

		region{
			Region: "Central",
			Hotels: []hotel{
				hotel{
					Name:    "Hotel California",
					Address: "42 Sunset Boulevard",
					City:    "Los Angeles",
					Zip:     "95612",
					Region:  "southern",
				},
				hotel{
					Name:    "H",
					Address: "4",
					City:    "L",
					Zip:     "95612",
					Region:  "southern",
				},
			},
		},
	}

	for _, v := range h {
		fmt.Println(v.Region)
		for _, val := range v.Hotels {
			fmt.Println(val.Address)
			fmt.Println(val.City)
			fmt.Println(val.Name)
			fmt.Println(val.Region)
			fmt.Println(val.Zip)
		}
		fmt.Println("******************")
	}

	err := tpl.Execute(os.Stdout, h)

	if err != nil {
		log.Fatalln(err)
	}

}
