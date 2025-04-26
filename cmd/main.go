/**
   Copyright 2025 Vassili Dzuba

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.	
**/

package main

import (
	"strconv"

	"github.com/vassilidzuba/yacictfe/internals"
	
	"github.com/rivo/tview"
	tcell "github.com/gdamore/tcell/v2"
)

var ProjectList *tview.List
var BranchList *tview.List
var BuildList *tview.List
var StepList *tview.List

var ProjectLabel *tview.TextView
var BranchLabel *tview.TextView
var BuildLabel *tview.TextView
var CommandLabel *tview.TextView
var LogView *tview.TextView

var App *tview.Application

var CurrentFocus string

func main() {
	yacicclient.InitProjectList()
	
	App = tview.NewApplication()
	App.SetInputCapture(inputCaptureHandler)
	
	ProjectList = tview.NewList()
	ProjectList.SetSelectedFocusOnly(true)
	ProjectList.SetSelectedFunc(projectListSelectHandler);
		
	for _, e := range yacicclient.Projects {
		ProjectList.AddItem(e.ProjectId, e.Repo, '-', nil)
	}

	BranchList = tview.NewList()
	BranchList.SetSelectedFocusOnly(true)
	BranchList.ShowSecondaryText(false)
	BranchList.SetSelectedFunc(branchListSelectHandler);

	BuildList = tview.NewList()
	BuildList.SetSelectedFocusOnly(true)
	BuildList.ShowSecondaryText(false)
	BuildList.SetSelectedFunc(buildListSelectHandler);
	
	BuildList.AddItem("aaaa", "", 0, nil)
	BuildList.AddItem("bbb", "", 0, nil)
	BuildList.AddItem("ccc", "", 0, nil)
	BuildList.AddItem("ddd", "", 0, nil)

	StepList = tview.NewList()
	StepList.SetSelectedFocusOnly(true)
	StepList.ShowSecondaryText(false)

	StepList.AddItem("xxx", "", 0, nil)
	StepList.AddItem("yyy", "", 0, nil)
	StepList.AddItem("zzz", "", 0, nil)
	
	ProjectLabel = tview.NewTextView()
	BranchLabel = tview.NewTextView()
	BuildLabel = tview.NewTextView()
	CommandLabel = tview.NewTextView().SetText("p:Project b:Branch u:build s:step n:next l:log r:run")

	flex3 := tview.NewFlex()
	flex3.SetDirection(tview.FlexRow)
	flex3.AddItem(ProjectLabel, 1, 1, false)
	flex3.AddItem(BranchList, 0, 1, false)
	
	flex4 := tview.NewFlex()
	flex4.SetDirection(tview.FlexRow)
	flex4.AddItem(BranchLabel, 1, 1, false)
	flex4.AddItem(BuildList, 0, 1, false)

	flex5 := tview.NewFlex()
	flex5.SetDirection(tview.FlexRow)
	flex5.AddItem(BuildLabel, 1, 1, false)
	flex5.AddItem(StepList, 0, 1, false)

	flex2 := tview.NewFlex()
	flex2.AddItem(ProjectList, 0, 2, true)
	flex2.AddItem(flex3, 0, 1, false)
	flex2.AddItem(flex4, 0, 2, false)
	flex2.AddItem(flex5, 0, 1, false)

	LogView = tview.NewTextView()
	
	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow)
	flex.AddItem(flex2, 0, 1, true)
	flex.AddItem(LogView, 0, 3, false)
	flex.AddItem(CommandLabel, 1, 1, false)

	CurrentFocus = "p"
	
	if err := App.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}

func projectListSelectHandler(index int, mainText string, secondaryText string, shortcut rune) {
	BranchList.Clear()
	
	ProjectLabel.SetText(mainText)
	
	branches := yacicclient.GetBranches(mainText)
	for _, b := range branches {
		BranchList.AddItem(b.Branch, "", '-', nil)
	}
}


func branchListSelectHandler(index int, mainText string, secondaryText string, shortcut rune) {
	BuildList.Clear()
	
	BranchLabel.SetText(mainText)
	
	yacicclient.InitBuildList(ProjectLabel.GetText(true), mainText)
	
	for _, b := range yacicclient.Builds {
		text := b.Timestamp + " " + strconv.Itoa(b.Duration / 1000) + "ms " + b.Status
		BuildList.AddItem(text, "", '-', nil)
	}
}



func buildListSelectHandler(index int, mainText string, secondaryText string, shortcut rune) {
	StepList.Clear()
	
	BuildLabel.SetText(mainText)
	
	yacicclient.InitStepList(ProjectLabel.GetText(true), BranchLabel.GetText(true), extractTimestamp(mainText))
	
	for _, s := range yacicclient.Steps {
		text := strconv.Itoa(s.Seq)+ " " + s.StepId + " " + strconv.Itoa(s.Duration / 1000) + "s " + s.Status
		StepList.AddItem(text, "", '-', nil)
	}
}

func displayLog() {
	LogView.SetText("lorem ipsum")
}


func runPipeline() {
	LogView.SetText("shmuld run pipeline")
}

func inputCaptureHandler(event *tcell.EventKey) *tcell.EventKey {
	if event.Key() == tcell.KeyRune {
		ch := event.Rune()
		
		if ch == 'p' {
			App.SetFocus(ProjectList)
			return nil;
		}
		if ch == 'b' {
			App.SetFocus(BranchList)
			return nil;
		}
		if ch == 'u' {
			App.SetFocus(BuildList)
			return nil;
		}
		if ch == 's' {
			App.SetFocus(StepList)
			return nil;
		}
		if ch == 'n' {
			if CurrentFocus == "p" {
				App.SetFocus(BranchList)
				CurrentFocus = "b"
			} else if CurrentFocus == "b" {
				App.SetFocus(BuildList)
				CurrentFocus = "u"
			} else if CurrentFocus == "u" {
				App.SetFocus(StepList)
				CurrentFocus = "s"
			} else if CurrentFocus == "b" {
				App.SetFocus(ProjectList)
				CurrentFocus = "p"
			}
			return nil;
		}
		if ch == 'l' {
			displayLog()
			return nil
		}
		if ch == 'r' {
			runPipeline()
			return nil
		}
	}
	
	return event;
}

func extractTimestamp(ts string) string {
	// timestamp is ascii, no need to convert into runes
	if ts == "" {
		return ""	
	} else {
		return ts[0:14]
	}
}
