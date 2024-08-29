package main

import (
	"github.com/osmanhomek/go2rtc/interior/app"
	"github.com/osmanhomek/go2rtc/interior/rtsp"
	"github.com/osmanhomek/go2rtc/interior/streams"
	"github.com/osmanhomek/go2rtc/pkg/shell"
)

func main() {
	app.Init()
	streams.Init()

	rtsp.Init()

	shell.RunUntilSignal()
}
