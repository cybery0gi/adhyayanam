package main

import (
	"github.com/rivo/tview"
	"github.com/gdamore/tcell/v2"
)

// Define string constants for the splash screen
const (
	WelcomeMessage     = "[::b][yellow]Welcome to [green]Adhyayanam[::] - [blue]Press [::b][magenta]Enter[::] to continue"
	LandingScreenTitle = "Vedonitamadheeyataam"
)

func displaySplashPage(app *tview.Application, switchToMainPage func()) {
	// Create a new text view to display the landing message
	landingText := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetText(WelcomeMessage).
		SetDynamicColors(true)

	// Set border and title for the splash screen
	landingText.SetBorder(true).SetTitle("[cyan]" + LandingScreenTitle)

	// Handle input for when the user presses "Enter"
	landingText.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEnter {
			// When Enter is pressed, proceed to the main page
			switchToMainPage()
		}
		return event
	})

	// Spacer to center content vertically
	spacer := tview.NewTextView().SetText("\n\n\n\n").SetTextAlign(tview.AlignCenter)

	// Layout for splash page
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(spacer, 0, 1, false).
		AddItem(landingText, 0, 3, true)

	// Set the root for the application
	app.SetRoot(flex, true)
}

func displayMainPage(app *tview.Application) {
	// Create the left pane for the main page
	leftPane := tview.NewTextView().SetText("Left Pane").SetBorder(true).SetTitle("Left Pane")

	// Create the middle and right panes (which will be stacked vertically)
	middlePane := tview.NewTextView().SetText("Middle Pane").SetBorder(true).SetTitle("Middle Pane")
	rightPane := tview.NewTextView().SetText("Right Pane").SetBorder(true).SetTitle("Right Pane")

	// Create a vertical Flex layout to stack middle and right panes
	rightFlex := tview.NewFlex().
		SetDirection(tview.FlexRow). // Vertical stacking (rows)
		AddItem(middlePane, 0, 1, false). // Middle pane (takes half the vertical space)
		AddItem(rightPane, 0, 3, false)   // Right pane (takes the other half of vertical space)

	// Create the main layout: left pane (horizontal) and rightFlex (vertical stack of middle + right)
	mainFlex := tview.NewFlex().
		AddItem(leftPane, 0, 1, false). // Left pane (one part of the width)
		AddItem(rightFlex, 0, 2, false) // Right pane (stacked middle + right, two parts of the width)

	// Set the root for the main page
	app.SetRoot(mainFlex, true)
}

func main() {
	// Create the application
	app := tview.NewApplication()

	// Define a function to switch to the main page
	switchToMainPage := func() {
		displayMainPage(app)
	}

	// Display the splash page initially
	displaySplashPage(app, switchToMainPage)

	// Run the application
	if err := app.Run(); err != nil {
		panic(err)
	}
}
