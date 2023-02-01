package main

import (
	"framework/app"

	"github.com/1ets/lets/boot"
)

// Initiate lets engine
func init() {
	app.OnInit()
	boot.OnInit()
}

// Bootstrap applications
func main() {
	boot.OnMain()
}
