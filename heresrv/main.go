package main

import (
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/hibooboo2/systemutils/activity"
)

var currentComp, prevComp string
var mu sync.Mutex

const (
	localHost   = "localHost"
	none        = "none"
	localHostIP = "127.0.0.1"
)

func main() {
	w, err := activity.NewActivityMonitor(time.Millisecond*500, time.Millisecond*10)
	if err != nil {
		log.Fatal(err)
	}
	here := w.UserActiveChan()
	gone := w.UserInactiveChan()
	go func() {
		for {
			select {
			case _ = <-here:
				setCurrentComp(localHost)
			case _ = <-gone:
				setCurrentComp(none)
			}
		}
	}()
	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			mu.Lock()
			if currentComp == prevComp {
				mu.Unlock()
				continue
			}
			log.Println(currentComp)
			mu.Unlock()
		}
	}()
	http.HandleFunc("/here", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		setCurrentComp(r.RemoteAddr)
	})
	http.HandleFunc("/gone", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		setCurrentComp(none)
	})
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func setCurrentComp(comp string) {
	if strings.Contains(comp, localHostIP) {
		comp = localHost
	}
	mu.Lock()
	prevComp = currentComp
	currentComp = comp
	mu.Unlock()
}
