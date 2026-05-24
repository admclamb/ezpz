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
