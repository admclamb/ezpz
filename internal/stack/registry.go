package stack

type Stack struct {
	ID           string
	Version      string
	Name         string
	Description  string
	TemplateRepo string
	Variables    map[string]string
	Tags         []string
}

type Registry struct {
	byKey    map[string]Stack
	latestID map[string]string
}

func NewDefaultRegistry() *Registry {
	r := &Registry{
		byKey:    make(map[string]Stack),
		latestID: make(map[string]string),
	}

	defaults := []Stack{
		{
			ID:           "nextjs-simple",
			Version:      "1.0.0",
			Name:         "Next.js Simple",
			Description:  "A simple Next.js starter template",
			TemplateRepo: "github.com/admclamb/templates/nextjs-simple",
			Tags:         []string{"nextjs", "react", "web"},
		},
	}

	for _, stack := range defaults {
		if err := r.Register(stack); err != nil {
			panic(err)
		}
	}

	return r
}
