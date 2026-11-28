package main

import (
	"fmt"

	"charm.land/bubbles/v2/table"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var baseStyle = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("#641882"))

type appState int

const (
	stateClass appState = iota
	stateProblem
)

type model struct {
	state             appState
	table             table.Model
	classeSelecionada string
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "esc":
			switch m.state {
			case stateProblem:
				m.state = stateClass
				m.table = tableClass()

			case stateClass:
				return m, tea.Quit
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			switch m.state {
			case stateClass:
				m.classeSelecionada = m.table.SelectedRow()[0]
				m.state = stateProblem
				m.table = tableProblem(m.classeSelecionada)
				return m, nil
			}
		case "space":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}
func (m model) View() tea.View {
	return tea.NewView(baseStyle.Render(m.table.View()) + "v0.1.0" + "\n " + m.table.HelpView() + "\n ")
}
func tableClass() table.Model {
	columns := []table.Column{
		{Title: "VortexDrive", Width: 50},
	}
	rows := []table.Row{
		{"🌐 Internet"},
		{"🐢 Lentidão"},
		{"🖨️ Impressora"},
	}
	menuC := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(10),
		table.WithWidth(50),
	)
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#641882")).
		BorderBottom(true).
		Bold(false)
	s.Cell = s.Cell.
		Foreground(lipgloss.Color("#ab0fe9"))
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("#418222")).
		Background(lipgloss.Color("#222321")).
		BorderBottom(true).
		Bold(false)
	menuC.SetStyles(s)

	return menuC
}
func tableProblem(classeSelecionada string) table.Model {
	titulo := fmt.Sprintf("VortexDrive: %s", classeSelecionada)
	columns := []table.Column{
		{Title: titulo, Width: 50},
	}
	var rows []table.Row
	for _, problem := range Catalog[classeSelecionada] {
		rows = append(rows, table.Row{problem.Options})
	}
	menuP := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(5),
		table.WithWidth(42),
	)
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#641882")).
		Foreground(lipgloss.Color("#ab0fe9")).
		BorderBottom(true).
		Bold(false)
	s.Cell = s.Cell.
		Foreground(lipgloss.Color("#ab0fe9"))
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("#418822")).
		Background(lipgloss.Color("#222321")).
		BorderBottom(true).
		Bold(false)
	menuP.SetStyles(s)

	return menuP

}
func main() {
	m := model{table: tableClass()}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Erro ao iniciar o programa", err)
	}
}
