package main

import (
	"encoding/json"
	"fmt"
)

type PostConfig struct {
	ID      int               `json:"id,omitempty"`
	Type    string            `json:"type,omitempty"`
	Tags    []string          `json:"tags,omitempty"`
	Filters map[string]string `json:"filters,omitempty"`
}

type BlogConfig struct {
	Name    string     `json:"name,omitempty"`
	Authors []string   `json:"authors,omitempty"`
	Configs PostConfig `json:"configs,omitzero"`
}

func (p *PostConfig) IsZero() bool {
	if p.ID == 0 && p.Type == "Post" && len(p.Tags) == 0 && len(p.Filters) == 0 {
		return true
	}
	return false
}

func main() {
	blogConfig := BlogConfig{
		Configs: PostConfig{
			Type: "Post",
		},
	}
	b, _ := json.Marshal(blogConfig)
	fmt.Println(string(b))
}
