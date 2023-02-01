package app

import (
	"framework/config"

	"github.com/1ets/lets/boot"
)

// Intercept lets initialization
func OnInit() {
	boot.AddInitializer(config.App)
	boot.AddInitializer(config.Database)
}
