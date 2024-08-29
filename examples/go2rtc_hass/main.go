package main

import (
	"github.com/osmanhomek/go2rtc/interior/api"
	"github.com/osmanhomek/go2rtc/interior/app"
	"github.com/osmanhomek/go2rtc/interior/hass"
	"github.com/osmanhomek/go2rtc/interior/streams"
	"github.com/osmanhomek/go2rtc/pkg/shell"
)

func main() {
	app.Init()
	streams.Init()

	api.Init()

	hass.Init()

	shell.RunUntilSignal()
}
