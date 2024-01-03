package main

import (
	"fmt"
	"github.com/pterm/pterm"
	"os"
)

func graphHours(s string, i int) {
	weeklyGraph("fn", 3)
}

func weeklyGraph(fn string, i int) {
	// init drawing area
	area, err := pterm.DefaultArea.WithFullscreen().WithCenter().Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// destroy drawing area when done
	defer area.Stop()
	barData := []pterm.Bar{
		{Label: "Jan", Value: 10},
		{Label: "Feb", Value: 20},
		{Label: "Mar", Value: 30},
		{Label: "Apr", Value: 40},
		{Label: "May", Value: 45},
		{Label: "Jun", Value: 45},
		{Label: "Jul", Value: 35},
		{Label: "Aug", Value: 20},
		{Label: "Sep", Value: 15},
		{Label: "Oct", Value: 35},
		{Label: "Nov", Value: 22},
		{Label: "Dec", Value: 26},
	}
	// workhours
	area.Update()
	pterm.DefaultBasicText.Println("Monthly Workhours:")
	pterm.DefaultBarChart.WithBars(barData).WithShowValue().Render()

	result, err := pterm.DefaultInteractiveConfirm.WithDefaultText("Show overtime?").Show()
	if err != nil {
		fmt.Println(err)
	}
	pterm.Println()
	if result {
		// overtime
		area.Update()
		pterm.DefaultBasicText.Println("Monthly ovetime:")
		pterm.DefaultBarChart.WithHorizontal().WithBars(barData).WithShowValue().Render()
	}
	if !result {
		pterm.Println("Exiting ...")
		os.Exit(0)
	}
}
