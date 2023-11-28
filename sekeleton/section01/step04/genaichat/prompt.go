package genaichat

import (
	_ "embed"
	"fmt"
	"io"
	"text/template"
)

var (
	//TODO: embedパッケージを使ってprompt_template.txtの値をpromptTmplFileに入れてください
	promptTmplFile string
	promptTmpl     = template.Must(template.New("prompt").Parse(promptTmplFile))
)

type Instruction struct {
	Name string
}

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
	inst := &Instruction{
		Name: p.Name,
	}

	if err := p.Template.Execute(w, inst); err != nil {
		return fmt.Errorf("(*genaichat.Prompt).Write: %w", err)
	}
	return nil
}
