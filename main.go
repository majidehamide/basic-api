package main

import "fmt"

type App struct {
}

func (app *App) Run() error {
	fmt.Println("Method Run")
	return nil
}

func main() {
	fmt.Println("Test rest api course")

	app := App{}

	if err := app.Run(); err != nil {
		fmt.Println("error")
		fmt.Println(err)
	}
}
