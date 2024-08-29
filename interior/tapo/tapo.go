package tapo

import (
	"github.com/osmanhomek/go2rtc/interior/streams"
	"github.com/osmanhomek/go2rtc/pkg/core"
	"github.com/osmanhomek/go2rtc/pkg/kasa"
	"github.com/osmanhomek/go2rtc/pkg/tapo"
)

func Init() {
	streams.HandleFunc("kasa", func(source string) (core.Producer, error) {
		return kasa.Dial(source)
	})

	streams.HandleFunc("tapo", func(source string) (core.Producer, error) {
		return tapo.Dial(source)
	})
}
