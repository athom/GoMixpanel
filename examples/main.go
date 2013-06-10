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

	//terse api
	success, err := m.SendEvent("Terse API Event", map[string]interface{}{
		"lib":     "GoMixpanel",
		"author":  "Rich Collins",
		"version": "20121127"})

	if success {
		fmt.Println("Terse Success")
	} else if err == nil {
		fmt.Println("Terse Failure")
	} else {
		fmt.Println(err)
	}

	//Singleton usage
	//gomixpanel.SetApiToken("anotherToken")
	gomixpanel.SetApiToken("76d89d8aeebb5053719915f18124dbf7")
	gomixpanel.Track("Simple Track Demo", map[string]interface{}{
		"lib":     "GoMixpanel",
		"authors": []string{"Rich Collins", "athom"},
		"age":     12,
		"version": 20130610,
	})

	gomixpanel.TrackWithCallback("Singleton API Event", map[string]interface{}{
		"lib":     "GoMixpanel",
		"authors": []string{"Rich Collins", "athom"},
		"age":     12,
		"version": 20130610,
	}, func() {
		fmt.Println("Put Singleton Event Success")
	})
}
