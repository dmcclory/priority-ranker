/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/charmbracelet/lipgloss"
)

func warningStyle() lipgloss.Style {
  style := lipgloss.NewStyle().
    Bold(true).
    PaddingTop(1).
    PaddingLeft(4).
    Foreground(lipgloss.Color("5"))

	return style
  // return style.Render("There are no options in the list, use `ranker option add` to add some!")
}

