package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"saerch/workflows"
)

const (
	CHROME        = "google-chrome"
	CHROMIUM      = "chromium"
	FIREFOX       = "firefox"
	FIREFOX_DEV   = "firefox-developer-edition"
	BRAVE_NIGHTLY = "brave-browser-nightly"
	BRAVE_DEV     = "brave-browser-dev"
	BRAVE         = "brave-browser"
)

const HELP = `Usage: go run main.go <browser> <workflow> <search>`

func main() {
	args := os.Args[1:]
	defaultWorkflows := workflows.GetDefaultWorkflows()

	if len(args) == 0 || args[0] == "help" {
		fmt.Println(HELP)
		return
	}

	browser := args[0]
	workflow := args[1]

	userWorkflow := defaultWorkflows.GetWorkflow(workflow)

	search := args[2:]

	switch browser {
	case CHROME:
		runWorkflow(CHROME, workflow, userWorkflow, search)
	case FIREFOX:
		runWorkflow(FIREFOX, workflow, userWorkflow, search)
	case BRAVE:
		runWorkflow(BRAVE, workflow, userWorkflow, search)
	case BRAVE_NIGHTLY:
		runWorkflow(BRAVE_NIGHTLY, workflow, userWorkflow, search)
	case BRAVE_DEV:

	default:
		fmt.Println("Unknown browser")
	}
}

func runWorkflow(browser string, enteredWorkflow string, userWorkflow workflows.Workflow, search []string) {
	if userWorkflow.Name == enteredWorkflow {
		for _, engine := range userWorkflow.Engines {
			openBrowser(browser, engine.SearchQuery+strings.Join(search, "+"))
		}
	}
}

func openBrowser(browser string, url string) {
	exec.Command(browser, url).Start()
}
