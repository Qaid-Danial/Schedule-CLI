package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/rivo/tview"
)

type Subject struct {
	Name     string  `json:"name"`
	Code     string  `json:"code"`
	Section  string  `json:"Section"`
	Time     float32 `json:"time"`
	Location string  `json:"location"`
}

var (
	classes   = []Subject{}
	subjectDB = "SubjectDB.json"
)

// Function - Reading the JSON files for all the subjects
func readSubjects() {
	if _, err := os.Stat(subjectDB); err == nil {
		data, err := os.ReadFile(subjectDB)
		if err != nil {
			log.Fatal("Error while reading file: ", err)
		}
		json.Unmarshal(data, &classes)
	}
}

// Function - Saving new subjects onto the JSON file
func addSubject() {
	data, err := json.MarshalIndent(classes, "", " ")
	if err != nil {
		log.Fatal("Error while saving file: ", err)
	}
	os.WriteFile(subjectDB, data, 0644)
}

func main() {
	app := tview.NewApplication()

	readSubjects()

	subjectList := tview.NewTextView()
	subjectList.SetBorder(true).
		SetTitle(" Subject List ")

	refreshSubject := func() {
		subjectList.Clear()
		if len(classes) == 0 {
			fmt.Fprintln(subjectList, "No Classes currently! YIPEE!!")
		} else {
			for i, subject := range classes {
				fmt.Fprintf(subjectList, "%v - %v\n", i+1, subject.Name)
			}
		}
	}

	classForm := tview.NewForm()
	classForm.SetBorder(true).
		SetTitle(" Class Form ")

	classNameInput := tview.NewInputField().SetLabel("Class Name: ")
	classCodeInput := tview.NewInputField().SetLabel("Class Code: ")
	// subjectSectionInput
	// subjectTimeInput
	// subjectLocationInput

	classForm.AddFormItem(classNameInput).
		AddFormItem(classCodeInput).
		AddButton("Register Class", func() {
			name := classNameInput.GetText()
			code := classCodeInput.GetText()

			if name != "" && code != "" {
				classes = append(classes, Subject{Name: name, Code: code, Section: "", Time: 0, Location: ""})
				addSubject()
				refreshSubject()
				classNameInput.SetText("")
				classCodeInput.SetText("")
				app.SetFocus(classForm)
			}
		})

	flex := tview.NewFlex().
		AddItem(classForm, 0, 1, true).
		AddItem(subjectList, 0, 1, false)

	refreshSubject()

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}
