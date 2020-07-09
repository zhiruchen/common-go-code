package context

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/parnurzeal/gorequest"
)

type result struct {
	content string
	err     error
}

func CrawlWeb(ctx context.Context, url string) (*result, error) {
	var c = make(chan result, 1)
	go func() {
		now := time.Now()
		request := gorequest.New()
		_, body, errs := request.Get(url).End()
		if len(errs) > 0 {
			var errStr []string
			for _, e := range errs {
				errStr = append(errStr, e.Error())
			}

			log.Printf("get error: %v, at: %v\n", errStr, time.Since(now))
			c <- result{err: fmt.Errorf(strings.Join(errStr, ","))}
			return
		}
		log.Printf("get result: %s, at: %v\n", body, time.Since(now))
		c <- result{content: body}
	}()

	select {
	case <-ctx.Done():
		ctxErr := ctx.Err()
		log.Printf("ctx done, error: %v\n", ctxErr)
		return nil, ctxErr
	case res := <-c:
		return &res, nil
	}
}
