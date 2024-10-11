package workflows

type Workflow struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Engines     []Engine `json:"engines"`
}

type Engine struct {
	Name        string `json:"name"`
	URL         string `json:"url"`
	SearchQuery string `json:"search_query"`
}

type Workflows struct {
	Workflows []Workflow `json:"workflows"`
}

// Adds a workflow to the list of workflows
func (w *Workflows) AddWorkflow(workflow Workflow) {
	w.Workflows = append(w.Workflows, workflow)
}

// Returns a workflow by name
func (w *Workflows) GetWorkflow(name string) Workflow {
	for _, workflow := range w.Workflows {
		if workflow.Name == name {
			return workflow
		}
	}
	return Workflow{}
}

// Returns a list of default workflows
func GetDefaultWorkflows() *Workflows {
	workflows := &Workflows{
		[]Workflow{
			{
				Name:        "dev",
				Description: "Search for development related topics",
				Engines: []Engine{
					{
						Name:        "Google",
						URL:         "https://www.google.com",
						SearchQuery: "https://www.google.com/search?q=",
					},
					{
						Name:        "Hacker News",
						URL:         "https://hn.algolia.com",
						SearchQuery: "https://hn.algolia.com/?q=",
					},
					{
						Name:        "Stack Overflow",
						URL:         "https://stackoverflow.com",
						SearchQuery: "https://stackoverflow.com/search?q=",
					},
					{
						Name:        "GitHub",
						URL:         "https://github.com",
						SearchQuery: "github.com/search?q=",
					},
					{
						Name:        "Github Issues",
						URL:         "https://github.com",
						SearchQuery: "google.com/search?q=site:github.com+inurl:issues+",
					},
					{
						Name:        "MDN Web Docs",
						URL:         "https://developer.mozilla.org",
						SearchQuery: "https://developer.mozilla.org/en-US/search?q=",
					},
				},
			},
			{
				Name:        "academic",
				Description: "Search for academic related topics",
				Engines: []Engine{
					{
						Name:        "Google Scholar",
						URL:         "https://scholar.google.com",
						SearchQuery: "https://scholar.google.com/scholar?q=",
					},
					{
						Name:        "Semantic Scholar",
						URL:         "https://www.semanticscholar.org",
						SearchQuery: "https://www.semanticscholar.org/search?q=",
					},
				},
			},
		},
	}
	return workflows
}
