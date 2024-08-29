package bubble

import (
	"github.com/osmanhomek/go2rtc/interior/streams"
	"github.com/osmanhomek/go2rtc/pkg/bubble"
	"github.com/osmanhomek/go2rtc/pkg/core"
)

func Init() {
	streams.HandleFunc("bubble", func(source string) (core.Producer, error) {
		return bubble.Dial(source)
	})
}
