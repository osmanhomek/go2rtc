package ivideon

import (
	"github.com/osmanhomek/go2rtc/interior/streams"
	"github.com/osmanhomek/go2rtc/pkg/core"
	"github.com/osmanhomek/go2rtc/pkg/ivideon"
)

func Init() {
	streams.HandleFunc("ivideon", func(source string) (core.Producer, error) {
		return ivideon.Dial(source)
	})
}
