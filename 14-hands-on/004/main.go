package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

type restaurant struct {
	Name string
	Menu menu
}

type menu []meal

type meal struct {
	Meal string
	Item []item
}

type item struct {
	Name, Descrip string
	Price         float64
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	R1 := restaurant{
		Name: "Federicos",
		Menu: menu{
			meal{
				Meal: "Breakfast",
				Item: []item{
					item{
						Name:    "Oatmeal",
						Descrip: "yum yum",
						Price:   4.95,
					},
					item{
						Name:    "Cheerios",
						Descrip: "American eating food traditional now",
						Price:   3.95,
					},
					item{
						Name:    "Juice Orange",
						Descrip: "Delicious drinking in throat squeezed fresh",
						Price:   2.95,
					},
				},
			},
			meal{
				Meal: "Lunch",
				Item: []item{
					item{
						Name:    "Hamburger",
						Descrip: "Delicous good eating for you",
						Price:   6.95,
					},
					item{
						Name:    "Cheese Melted Sandwich",
						Descrip: "Make cheese bread melt grease hot",
						Price:   3.95,
					},
					item{
						Name:    "French Fries",
						Descrip: "French eat potatoe fingers",
						Price:   2.95,
					},
				},
			},
			meal{
				Meal: "Dinner",
				Item: []item{
					item{
						Name:    "Pasta Bolognese",
						Descrip: "From Italy delicious eating",
						Price:   7.95,
					},
					item{
						Name:    "Steak",
						Descrip: "Dead cow grilled bloody",
						Price:   13.95,
					},
					item{
						Name:    "Bistro Potatoe",
						Descrip: "Bistro bar wood American bacon",
						Price:   6.95,
					},
				},
			},
		},
	}

	R2 := restaurant{
		Name: "Domenicos",
		Menu: menu{
			meal{
				Meal: "Breakfast",
				Item: []item{
					item{
						Name:    "Oatmeal",
						Descrip: "yum yum",
						Price:   4.95,
					},
					item{
						Name:    "Cheerios",
						Descrip: "American eating food traditional now",
						Price:   3.95,
					},
					item{
						Name:    "Juice Orange",
						Descrip: "Delicious drinking in throat squeezed fresh",
						Price:   2.95,
					},
				},
			},
			meal{
				Meal: "Lunch",
				Item: []item{
					item{
						Name:    "Hamburger",
						Descrip: "Delicous good eating for you",
						Price:   6.95,
					},
					item{
						Name:    "Cheese Melted Sandwich",
						Descrip: "Make cheese bread melt grease hot",
						Price:   3.95,
					},
					item{
						Name:    "French Fries",
						Descrip: "French eat potatoe fingers",
						Price:   2.95,
					},
				},
			},
			meal{
				Meal: "Dinner",
				Item: []item{
					item{
						Name:    "Pasta Bolognese",
						Descrip: "From Italy delicious eating",
						Price:   7.95,
					},
					item{
						Name:    "Steak",
						Descrip: "Dead cow grilled bloody",
						Price:   13.95,
					},
					item{
						Name:    "Bistro Potatoe",
						Descrip: "Bistro bar wood American bacon",
						Price:   6.95,
					},
				},
			},
		},
	}
	Restaurants := []restaurant{R1, R2}

	for _, res := range Restaurants {
		fmt.Println(res.Name)
		fmt.Println("___________")
		for _, men := range res.Menu {
			fmt.Println("\t\t", men.Meal)
			for ix, itm := range men.Item {
				fmt.Println("\t\t\t\t Item:", ix)
				fmt.Println("\t\t\t\t\t\t\t\t Name", "-", itm.Name)
				fmt.Println("\t\t\t\t\t\t\t\t Description", "-", itm.Descrip)
				fmt.Println("\t\t\t\t\t\t\t\t Price", "-", itm.Price)
			}
		}
		fmt.Println("************************************************************", "Menu Ends", "*****************************************************************************")
	}

	err := tpl.Execute(os.Stdout, Restaurants)

	if err != nil {
		log.Fatalln(err)
	}

}
