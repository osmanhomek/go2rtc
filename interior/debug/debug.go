package debug

import (
	"github.com/osmanhomek/go2rtc/interior/api"
)

func Init() {
	api.HandleFunc("api/stack", stackHandler)
}
