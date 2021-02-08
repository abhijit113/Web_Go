package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

// create a FuncMap to register functions.
/*
Funcs adds the elements of the argument map to the template's function map.
It must be called before the template is parsed. It panics if a value in the map is not a function with appropriate return type
or if the name cannot be used syntactically as a function in a template. It is legal to overwrite elements of the map.
 The return value is the template, so calls can be chained.
*/
// "uc" is what the func will be called in the template
// "uc" is the ToUpper func from package strings
// "ft" is a func I declared
// "ft" slices a string, returning the first three characters
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
	//you need to have funcs before you parse it
	//Out of 4 choices, the last 2 choices need to go as the are parsing it.
	//But template pointer needs to be returned.
	//So go with 1st two choices, so that before you parse the function
	//you need to pass the func map
	//that's why new doesn't have any nane and so template can not parse immediately

	/*
		func Must(t *Template, err error) *Template
		func New(name string) *Template
		func ParseFiles(filenames ...string) (*Template, error)
		func ParseGlob(pattern string) (*Template, error)
	*/
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
}

func main() {

	b := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	g := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	m := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed.",
	}

	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	c := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	sages := []sage{b, g, m}
	cars := []car{f, c}

	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		sages,
		cars,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
