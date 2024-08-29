package expr

import (
	"errors"

	"github.com/osmanhomek/go2rtc/interior/app"
	"github.com/osmanhomek/go2rtc/interior/streams"
	"github.com/osmanhomek/go2rtc/pkg/expr"
)

func Init() {
	log := app.GetLogger("expr")

	streams.RedirectFunc("expr", func(url string) (string, error) {
		v, err := expr.Run(url[5:])
		if err != nil {
			return "", err
		}

		log.Debug().Msgf("[expr] url=%s", url)

		if url = v.(string); url == "" {
			return "", errors.New("expr: result is empty")
		}

		return url, nil
	})
}
