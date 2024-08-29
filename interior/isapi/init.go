package isapi

import (
	"github.com/osmanhomek/go2rtc/interior/streams"
	"github.com/osmanhomek/go2rtc/pkg/core"
	"github.com/osmanhomek/go2rtc/pkg/isapi"
)

func Init() {
	streams.HandleFunc("isapi", func(source string) (core.Producer, error) {
		return isapi.Dial(source)
	})
}
