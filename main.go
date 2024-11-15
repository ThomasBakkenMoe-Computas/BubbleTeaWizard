package main

// TODO: Follow the rest of the tutorial
// https://youtu.be/Gl31diSVP8M?t=782

import (
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	BorderColor lipgloss.Color
	InputField 	lipgloss.Style
}

func DefaultStyles () *Styles {
	s := new(Styles)
	s.BorderColor = lipgloss.Color("202")
	s.InputField = lipgloss.NewStyle().BorderForeground(s.BorderColor).BorderStyle(lipgloss.NormalBorder()).Padding(1).Width(80)

	return s
}

type model struct{
	questions 	[]string
	width 		int
	height 		int
	index 		int
	answerField textinput.Model
	styles 		*Styles
}

func New (questions []string) *model {
	styles := DefaultStyles()
	answerField := textinput.New()
	answerField.Placeholder = "Your answer here"
	return &model{
		questions:		questions,
		answerField: 	answerField,
		styles: 		styles,
	}
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
	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Center,
			m.questions[m.index],
			m.styles.InputField.Render(m.answerField.View()),
		),
	)
}

func main() {
	questions := []string{
		"What is your name?",
		"What is your favoirite editor?",
		"What is you favorite quote?",
	}
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