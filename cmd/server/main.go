package main

import "github.com/losdmi/timetracker/internal"

// TODO скачать иконки локально
// TODO добавить иконку стопа задачи если у нее нет конца времени

func main() {
	app := internal.BuildApp()

	app.Run()
}
