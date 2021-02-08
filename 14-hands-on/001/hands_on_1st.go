package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

type course struct {
	Number string
	Name   string
	Units  string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	AcaYear string
	Sem     semester
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	year1 := year{
		AcaYear: "2020-2021",
		Sem: semester{
			Term: "Fall",
			Courses: []course{
				course{"CSCI-40", "Introduction to Programming in Go", "4"},
				course{"CSCI-130", "Introduction to Web Programming with Go", "4"},
				course{"CSCI-140", "Mobile Apps Using Go", "4"},
			},
		},
	}

	year2 := year{
		AcaYear: "2020-2021",
		Sem: semester{
			Term: "Spring",
			Courses: []course{
				course{"CSCI-40", "Introduction to Programming in Go", "5"},
				course{"CSCI-130", "Introduction to Web Programming with Go", "5"},
				course{"CSCI-140", "Mobile Apps Using Go", "5"},
			},
		},
	}

	year3 := year{
		AcaYear: "2020-2021",
		Sem: semester{
			Term: "Fall",
			Courses: []course{
				course{"CSCI-40", "Introduction to Programming in Go", "5"},
				course{"CSCI-130", "Introduction to Web Programming with Go", "5"},
				course{"CSCI-140", "Mobile Apps Using Go", "5"},
			},
		},
	}

	year4 := year{
		AcaYear: "2020-2021",
		Sem: semester{
			Term: "Fall",
			Courses: []course{
				course{"CSCI-40", "Introduction to Programming in Go", "5"},
				course{"CSCI-130", "Introduction to Web Programming with Go", "5"},
				course{"CSCI-140", "Mobile Apps Using Go", "5"},
			},
		},
	}

	Years := []year{year1, year2, year3, year4}

	for _, v := range Years {
		fmt.Println(v.AcaYear, v.Sem.Term)
		for _, val := range v.Sem.Courses {
			fmt.Println(val.Name, val.Number, val.Units)
		}
		fmt.Println("****************")
	}

	err := tpl.Execute(os.Stdout, Years)

	if err != nil {
		log.Fatalln(err)
	}

}
