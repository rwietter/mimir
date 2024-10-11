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
	FIREFOX_DEV   = "firefox-developer-edition"
	BRAVE_NIGHTLY = "brave-browser-nightly"
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
	worflow := args[1]

	userWorkflow := defaultWorkflows.GetWorkflow(worflow)

	search := args[2:]

	switch browser {
	case "chrome":
		// exec.Command("google-chrome", google_url).Start()
	case "firefox":
		// exec.Command("firefox", google_url).Start()
	case "safari":
		// exec.Command("open", google_url).Start()
	case BRAVE_NIGHTLY:
		if userWorkflow.Name == worflow {
			for _, engine := range userWorkflow.Engines {
				exec.Command(BRAVE_NIGHTLY, engine.SearchQuery+strings.Join(search, "+")).Start()
			}
		}
	default:
		fmt.Println("Unknown browser")
	}
}
