package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	go_arg "github.com/alexflint/go-arg"
)

const (
	DayFormat string = "2006-01-02 15:04:05 MST"
)

var args struct {
	Url    string `arg:"required,positional"`
	Follow bool   `arg:"-f,--follow"`
	Quiet  bool   `arg:"-q,--quiet"`
}

func main() {
	var output string
	var now time.Time
	var lastStatus string
	go_arg.MustParse(&args)

	for {
		resp, err := http.Get(args.Url)
		if err != nil {
			fmt.Printf("Error returned in GET request: %v\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		now = time.Now()
		currentStatus := resp.Status
		output = now.Format(DayFormat) + " " + args.Url + ": " + currentStatus

		if !args.Quiet || lastStatus != currentStatus {
			fmt.Println(output)
		}

		lastStatus = currentStatus

		if !args.Follow {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
}
