package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/bobbae/q"
	"github.com/markcheno/go-quote"
	"gopkg.in/toast.v1"
)

//  go run main.go -ql=all -qo=stderr aapl amzn
func main() {
	flag.Parse()
	q.Q(flag.Args())
	quotes, err := quote.NewQuotesFromYahooSyms(flag.Args(), "", "", quote.Daily, false)
	if err != nil {
		log.Fatalln(err)
	}
	q.Q(quotes)
	actions := []toast.Action{}
	qmsg := ""
	for _, q := range quotes {
		qmsg += fmt.Sprintf("%s:%v\n",
			q.Symbol, q.Close,
		)
	}
	notification := &toast.Notification{
		AppID: "ToastQuote",
		//Title:   "Toast Quote",
		Message: qmsg,
		Actions: actions,
		//ActivationType:      "protocol",
		//ActivationArguments: "https://google.com",
		Audio: "ms-winsoundevent:Notification.Looping.Call6",
		//Loop:                false,
		Duration: "long",
	}

	if err := notification.Push(); err != nil {
		log.Fatalln(err)
	}
	q.Q(notification)
}
