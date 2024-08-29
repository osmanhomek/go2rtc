package debug

import (
	"github.com/AlexxIT/go2rtc/interior/api"
)

func Init() {
	api.HandleFunc("api/stack", stackHandler)
}
