package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

var choices = []string{"*", "+", "-", "/"}

type model struct {
	cursor int
	choice string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		case "enter":
			m.choice = choices[m.cursor]
			return m, tea.Quit

		case "down", "j":
			m.cursor++
			if m.cursor >= len(choices) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(choices) - 1
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	s := strings.Builder{}
	s.WriteString("Choose an operator:\n\n")

	for i := 0; i < len(choices); i++ {
		if m.cursor == i {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}
		s.WriteString(choices[i])
		s.WriteString("\n")
	}
	s.WriteString("\n(press q to quit)\n")

	return s.String()
}

func main() {
	var number1 float64
	var number2 float64

	fmt.Print("Enter a number: ")
	fmt.Scan(&number1)
	fmt.Print("Enter a number: ")
	fmt.Scan(&number2)

	p := tea.NewProgram(model{})

	m, err := p.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if m, ok := m.(model); ok && m.choice != "" {
		switch m.choice {
		case "*":
			total := number1 * number2
			fmt.Printf("The total is: %.2f\n", total)
		case "+":
			total := number1 + number2
			fmt.Printf("The total is: %.2f\n", total)
		case "-":
			total := number1 - number2
			fmt.Printf("The total is: %.2f\n", total)
		case "/":
			total := number1 / number2
			fmt.Printf("The total is: %.2f\n", total)
		}
	}
}
