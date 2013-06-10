package main

import (
	"fmt"
	"github.com/athom/gomixpanel"
)

func main() {
	m := gomixpanel.NewMixpanel()
	//m.SetApiToken("yourApiToken")
	m.SetApiToken("76d89d8aeebb5053719915f18124dbf7")

	//granular api
	e := m.NewEvent()
	e.SetName("Granular API Event")
	e.SetProperties(map[string]interface{}{
		"lib":    "GoMixpanel",
		"author": "Rich Collins"})
	e.SetProperty("version", "20121127")

	if success, err := e.Send(); success {
		fmt.Println("Granular Success")
	} else if err == nil {
		fmt.Println("Granular Failure")
	} else {
		fmt.Println(err)
	}
}
