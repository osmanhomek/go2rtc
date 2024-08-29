package main

import (
	"github.com/AlexxIT/go2rtc/interior/api"
	"github.com/AlexxIT/go2rtc/interior/app"
	"github.com/AlexxIT/go2rtc/interior/hass"
	"github.com/AlexxIT/go2rtc/interior/streams"
	"github.com/AlexxIT/go2rtc/pkg/shell"
)

func main() {
	app.Init()
	streams.Init()

	api.Init()

	hass.Init()

	shell.RunUntilSignal()
}
