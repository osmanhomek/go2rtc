package gopro

import (
	"net/http"

	"github.com/osmanhomek/go2rtc/interior/api"
	"github.com/osmanhomek/go2rtc/interior/streams"
	"github.com/osmanhomek/go2rtc/pkg/core"
	"github.com/osmanhomek/go2rtc/pkg/gopro"
)

func Init() {
	streams.HandleFunc("gopro", func(source string) (core.Producer, error) {
		return gopro.Dial(source)
	})

	api.HandleFunc("api/gopro", apiGoPro)
}

func apiGoPro(w http.ResponseWriter, r *http.Request) {
	var items []*api.Source

	for _, host := range gopro.Discovery() {
		items = append(items, &api.Source{Name: host, URL: "gopro://" + host})
	}

	api.ResponseSources(w, items)
}
