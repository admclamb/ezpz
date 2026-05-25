package stack

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

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
		{
			ID:           "go-cli",
			Version:      "1.0.0",
			Name:         "Go CLI",
			Description:  "A simple Go CLI starter template",
			TemplateRepo: "github.com/admclamb/templates/go-cli",
			Tags:         []string{"go", "cli"},
		},
	}

	for _, stack := range defaults {
		if err := r.Register(stack); err != nil {
			panic(err)
		}
	}

	return r
}

func (r *Registry) Register(stack Stack) error {
	if r.byKey == nil {
		r.byKey = make(map[string]Stack)
	}
	if r.latestID == nil {
		r.latestID = make(map[string]string)
	}

	stack.ID = normalizeID(stack.ID)
	if stack.ID == "" {
		return fmt.Errorf("stack id is required")
	}

	normalizedVersion, err := normalizeVersion(stack.Version)
	if err != nil {
		return fmt.Errorf("invalid version for stack %q: %w", stack.ID, err)
	}
	stack.Version = normalizedVersion

	if strings.TrimSpace(stack.Name) == "" {
		return fmt.Errorf("stack name is required")
	}

	key := stackKey(stack.ID, stack.Version)
	if _, exists := r.byKey[key]; exists {
		return fmt.Errorf("stack %q version %q is already registered", stack.ID, stack.Version)
	}

	r.byKey[key] = stack

	currentLatestKey, exists := r.latestID[stack.ID]
	if !exists {
		r.latestID[stack.ID] = key
		return nil
	}

	currentLatest := r.byKey[currentLatestKey]
	if compareVersions(stack.Version, currentLatest.Version) > 0 {
		r.latestID[stack.ID] = key
	}

	return nil
}

func (r *Registry) Get(id, version string) (Stack, bool) {
	if r == nil {
		return Stack{}, false
	}

	id = normalizeID(id)
	normalizedVersion, err := normalizeVersion(version)
	if err != nil {
		return Stack{}, false
	}

	stack, ok := r.byKey[stackKey(id, normalizedVersion)]
	return stack, ok
}

func (r *Registry) GetLatest(id string) (Stack, bool) {
	if r == nil {
		return Stack{}, false
	}

	id = normalizeID(id)
	key, ok := r.latestID[id]
	if !ok {
		return Stack{}, false
	}

	stack, ok := r.byKey[key]
	return stack, ok
}

func (r *Registry) List() []Stack {
	if r == nil {
		return nil
	}

	stacks := make([]Stack, 0, len(r.byKey))
	for _, stack := range r.byKey {
		stacks = append(stacks, stack)
	}

	sort.Slice(stacks, func(i, j int) bool {
		if stacks[i].ID == stacks[j].ID {
			return compareVersions(stacks[i].Version, stacks[j].Version) < 0
		}
		return stacks[i].ID < stacks[j].ID
	})

	return stacks
}

func (r *Registry) ListLatest() []Stack {
	if r == nil {
		return nil
	}

	stacks := make([]Stack, 0, len(r.latestID))
	for _, key := range r.latestID {
		stacks = append(stacks, r.byKey[key])
	}

	sort.Slice(stacks, func(i, j int) bool {
		return stacks[i].ID < stacks[j].ID
	})

	return stacks
}

func stackKey(id, version string) string {
	return normalizeID(id) + "@" + version
}

func normalizeID(id string) string {
	return strings.ToLower(strings.TrimSpace(id))
}

func normalizeVersion(version string) (string, error) {
	parts := strings.Split(strings.TrimSpace(version), ".")
	if len(parts) != 3 {
		return "", fmt.Errorf("version must be in major.minor.patch format")
	}

	normalized := make([]string, 3)
	for i, part := range parts {
		if part == "" {
			return "", fmt.Errorf("version contains an empty segment")
		}

		value, err := strconv.Atoi(part)
		if err != nil {
			return "", fmt.Errorf("version segment %q is not numeric", part)
		}
		if value < 0 {
			return "", fmt.Errorf("version segment %q must not be negative", part)
		}

		normalized[i] = strconv.Itoa(value)
	}

	return strings.Join(normalized, "."), nil
}

func compareVersions(a, b string) int {
	aParts := mustParseVersion(a)
	bParts := mustParseVersion(b)

	for i := 0; i < 3; i++ {
		switch {
		case aParts[i] < bParts[i]:
			return -1
		case aParts[i] > bParts[i]:
			return 1
		}
	}

	return 0
}

func mustParseVersion(version string) [3]int {
	normalized, err := normalizeVersion(version)
	if err != nil {
		panic(err)
	}

	parts := strings.Split(normalized, ".")
	var out [3]int
	for i, part := range parts {
		value, err := strconv.Atoi(part)
		if err != nil {
			panic(err)
		}
		out[i] = value
	}

	return out
}
