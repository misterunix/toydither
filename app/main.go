package main

import (
	_ "embed"

	tk "modernc.org/tk9.0"
)

//go:embed embeds/w3.css
var w3embed []byte

//go:embed embeds/slide.css
var slideembed []byte

//go:embed embeds/dropdown.css
var dropdownembed []byte

//go:embed embeds/index.html
var indexembed []byte

//go:embed embeds/htmx.min.js
var htmxembed []byte

//go:embed embeds/formupload.html
var formuploadembed []byte

func main() {
	// Pack the button into the main application window
	tk.Pack(
		// Create a new TButton widget
		tk.TButton(
			// Set the button's text
			tk.Txt("Hello, TK!"),
			// Define the command (function) to run when the button is clicked
			tk.Command(func() {
				tk.Destroy(tk.App) // Close the application
			}),
		),
		// Add padding around the button
		tk.Ipadx(10), tk.Ipady(5), tk.Padx(15), tk.Pady(10),
	)
	// Start the application event loop
	tk.App.Wait()
}
