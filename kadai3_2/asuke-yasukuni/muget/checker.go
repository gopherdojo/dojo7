package muget

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"golang.org/x/net/context/ctxhttp"
)

func CheckRanges(ctx context.Context, url string) (int, error) {
	res, err := ctxhttp.Head(ctx, http.DefaultClient, url)
	if err != nil {
		return 0, err
	}

	if res.Header.Get("Accept-Ranges") != "bytes" {
		return 0, fmt.Errorf("not supported range access: %s", url)
	}

	if res.ContentLength <= 0 {
		return 0, errors.New("invalid content length")
	}

	return int(res.ContentLength), nil
}
