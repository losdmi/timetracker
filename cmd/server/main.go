package main

import "github.com/losdmi/timetracker/internal"

func main() {
	app := internal.BuildApp()

	app.Run()
}
