package main

import (
	"github.com/AlexxIT/go2rtc/interior/app"
	"github.com/AlexxIT/go2rtc/interior/rtsp"
	"github.com/AlexxIT/go2rtc/interior/streams"
	"github.com/AlexxIT/go2rtc/pkg/shell"
)

func main() {
	app.Init()
	streams.Init()

	rtsp.Init()

	shell.RunUntilSignal()
}
