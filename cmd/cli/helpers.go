package main

func (app *application) printSuccess(message string) {
	app.cobra.Println("✅ ", message)
}

func (app *application) printError(message string) {
	app.cobra.Println("🟥 ", message)
}
