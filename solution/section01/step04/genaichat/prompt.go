package genaichat

import (
	_ "embed"
	"fmt"
	"io"
	"text/template"
)

var (
	//go:embed prompt_template.txt
	promptTmplFile string
	promptTmpl     = template.Must(template.New("prompt").Parse(promptTmplFile))
)

type Prompt struct {
	Name     string
	Template *template.Template
}

func NewPrompt(name string) *Prompt {
	return &Prompt{
		Name:     name,
		Template: promptTmpl,
	}
}

func (p *Prompt) Write(w io.Writer) error {

	err := p.Template.Execute(w, p.Name)
	if err != nil {
		return fmt.Errorf("(*genaichat.Prompt).Write: %w", err)
	}

	return nil
}
