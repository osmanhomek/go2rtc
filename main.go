package main

import (
	"github.com/AlexxIT/go2rtc/interior/api"
	"github.com/AlexxIT/go2rtc/interior/api/ws"
	"github.com/AlexxIT/go2rtc/interior/app"
	"github.com/AlexxIT/go2rtc/interior/bubble"
	"github.com/AlexxIT/go2rtc/interior/debug"
	"github.com/AlexxIT/go2rtc/interior/dvrip"
	"github.com/AlexxIT/go2rtc/interior/echo"
	"github.com/AlexxIT/go2rtc/interior/exec"
	"github.com/AlexxIT/go2rtc/interior/expr"
	"github.com/AlexxIT/go2rtc/interior/ffmpeg"
	"github.com/AlexxIT/go2rtc/interior/gopro"
	"github.com/AlexxIT/go2rtc/interior/hass"
	"github.com/AlexxIT/go2rtc/interior/hls"
	"github.com/AlexxIT/go2rtc/interior/homekit"
	"github.com/AlexxIT/go2rtc/interior/http"
	"github.com/AlexxIT/go2rtc/interior/isapi"
	"github.com/AlexxIT/go2rtc/interior/ivideon"
	"github.com/AlexxIT/go2rtc/interior/mjpeg"
	"github.com/AlexxIT/go2rtc/interior/mp4"
	"github.com/AlexxIT/go2rtc/interior/mpegts"
	"github.com/AlexxIT/go2rtc/interior/nest"
	"github.com/AlexxIT/go2rtc/interior/ngrok"
	"github.com/AlexxIT/go2rtc/interior/onvif"
	"github.com/AlexxIT/go2rtc/interior/roborock"
	"github.com/AlexxIT/go2rtc/interior/rtmp"
	"github.com/AlexxIT/go2rtc/interior/rtsp"
	"github.com/AlexxIT/go2rtc/interior/srtp"
	"github.com/AlexxIT/go2rtc/interior/streams"
	"github.com/AlexxIT/go2rtc/interior/tapo"
	"github.com/AlexxIT/go2rtc/interior/webrtc"
	"github.com/AlexxIT/go2rtc/interior/webtorrent"
	"github.com/AlexxIT/go2rtc/pkg/shell"
)

func main() {
	app.Version = "1.9.4"

	// 1. Core modules: app, api/ws, streams

	app.Init() // init config and logs

	api.Init() // init API before all others
	ws.Init()  // init WS API endpoint

	streams.Init() // streams module

	// 2. Main sources and servers

	rtsp.Init()   // rtsp source, RTSP server
	webrtc.Init() // webrtc source, WebRTC server

	// 3. Main API

	mp4.Init()   // MP4 API
	hls.Init()   // HLS API
	mjpeg.Init() // MJPEG API

	// 4. Other sources and servers

	hass.Init()       // hass source, Hass API server
	onvif.Init()      // onvif source, ONVIF API server
	webtorrent.Init() // webtorrent source, WebTorrent module

	// 5. Other sources

	rtmp.Init()     // rtmp source
	exec.Init()     // exec source
	ffmpeg.Init()   // ffmpeg source
	echo.Init()     // echo source
	ivideon.Init()  // ivideon source
	http.Init()     // http/tcp source
	dvrip.Init()    // dvrip source
	tapo.Init()     // tapo source
	isapi.Init()    // isapi source
	mpegts.Init()   // mpegts passive source
	roborock.Init() // roborock source
	homekit.Init()  // homekit source
	nest.Init()     // nest source
	bubble.Init()   // bubble source
	expr.Init()     // expr source
	gopro.Init()    // gopro source

	// 6. Helper modules

	ngrok.Init() // ngrok module
	srtp.Init()  // SRTP server
	debug.Init() // debug API

	// 7. Go

	shell.RunUntilSignal()
}
