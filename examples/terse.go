package main

import (
	"fmt"
	"github.com/athom/gomixpanel"
)

func main() {
	m := gomixpanel.NewMixpanel()
	m.SetApiToken("yourApiToken")

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
}
