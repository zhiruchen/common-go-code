package context

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// ctx done, error: context deadline exceeded
func TestCrawlWebTimeout(t *testing.T) {
	timeOut := 500 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	res, err := CrawlWeb(ctx, "https://blog.golang.org/context")
	assert.NotNil(t, err)
	assert.Equal(t, context.DeadlineExceeded, err)
	assert.Nil(t, res)
}

func TestCrawlWebWithoutTimeout(t *testing.T) {
	res, err := CrawlWeb(context.Background(), "https://blog.golang.org/context")
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

// ctx done, error: context deadline exceeded
// res:  <nil>
// err:  context deadline exceeded
func TestCrawlWebWithDeadline(t *testing.T) {
	timeOut := 500 * time.Millisecond
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(timeOut))
	defer cancel()

	res, err := CrawlWeb(ctx, "https://blog.golang.org/context")
	log.Println("res: ", res)
	log.Println("err: ", err)
	assert.NotNil(t, err)
	assert.Equal(t, context.DeadlineExceeded, err)
	assert.Nil(t, res)
}

func TestCrawlWebNotTimeout(t *testing.T) {
	timeOut := 3 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	res, err := CrawlWeb(ctx, "https://blog.golang.org/context")
	assert.Nil(t, err)
	assert.NotNil(t, res)
	if res != nil {
		fmt.Println("res: ", *res)
	}
}

// 500ms reached, cancel this request
// ctx done, error: context canceled
// res:  <nil> , error:  context canceled
func TestCrawlWebCancel(t *testing.T) {
	type crawlResult struct {
		res *result
		err error
	}
	var c = make(chan crawlResult, 1)
	// if the result not returned in 1 second will cancel this.
	timer := time.NewTimer(500 * time.Millisecond)
	defer timer.Stop()

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		res, err := CrawlWeb(ctx, "https://blog.golang.org/context")
		log.Println("res: ", res, ", error: ", err)
		assert.Equal(t, context.Canceled, err)
		c <- crawlResult{
			res: res,
			err: err,
		}
	}()
	select {
	case <-timer.C: // reached 500ms, but not get result. cancel this request.
		log.Println("500ms reached, cancel this request")
		cancel()
	case res := <-c:
		log.Println("get result in 500ms, great!")
		assert.Nil(t, res.err)
		assert.NotNil(t, res.res)
	}
}

// get value from context({}): Hello,, true
func TestPassValueThroughContext(t *testing.T) {
	ctx := context.WithValue(context.Background(), ctxKey, "Hello,")
	res, err := CrawlWeb(ctx, "https://blog.golang.org/context")
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
