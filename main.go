package main

import (
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct{
	questions []string
	width int
	height int
	answerField textinput.Model

}

func New (questions []string) *model {
	return &model{questions: questions}

}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
		case tea.WindowSizeMsg:
			m.width = msg.Width
			m.height = msg.Height

		case tea.KeyMsg:
			switch msg.String() {
				case "ctrl+c":
					return m, tea.Quit

			}
	}
	return m, nil
}

func (m model) View() string {
	if m.width == 0 {
		return "Loading..."
	}
	return "Loaded!"
}

func main() {
	questions := []string{"What is your name?", "What is your favoirite editor?", "What is you favorite quote?"}
	m := New(questions)


	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}


}