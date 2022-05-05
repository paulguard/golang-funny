package awesomeProject

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"
	"time"
)

func Test(t *testing.T) {

	var stupidAndweak = [5]string{"I", "am", "stupid", "and", "weak"}

	for i := range stupidAndweak {
		if stupidAndweak[i] == "stupid" {
			stupidAndweak[i] = "smart"
		}

		if stupidAndweak[i] == "weak" {
			stupidAndweak[i] = "strong"
		}
	}

	fmt.Println("result:", stupidAndweak)
}

func TestChannel(t *testing.T) {
	ch := make(chan int, 6)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("hello from goroutine")
			ch <- 0
		}
		close(ch)
	}()

	fmt.Println("hello from main")

	for v := range ch {
		fmt.Println("receiving:", v)
	}

	<-ch
}

func TestContext(t *testing.T) {
	baseCtx := context.Background()
	ctx := context.WithValue(baseCtx, "a", "b")
	go func(c context.Context) {
		fmt.Println(c.Value("a"))
	}(ctx)

	timeoutCtx, cancel := context.WithTimeout(baseCtx, time.Second)

	defer cancel()

	go func(ctx context.Context) {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			select {
			case <-ctx.Done():
				fmt.Println("child process interrupt...")
				return
			default:
				fmt.Println("enter default")
			}
		}
	}(timeoutCtx)

	time.Sleep(1 * time.Second)
	select {
	case <-timeoutCtx.Done():
		time.Sleep(1 * time.Second)
		fmt.Println("main process exit!")
	}

}

func TestHttp(t *testing.T) {

	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/healthz1", healthz2)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func healthz2(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "not ok....")
}

func healthz(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "ok!!")
}
