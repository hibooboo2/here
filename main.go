package main

import (
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/hibooboo2/systemutils/activity"
)

func main() {
	w, err := activity.NewActivityMonitor(time.Millisecond*500, time.Millisecond*10)
	if err != nil {
		log.Fatal(err)
	}
	here := w.UserActiveChan()
	gone := w.UserInactiveChan()
	for {
		select {
		case _ = <-here:
			log.Println("Active")
			http.Get("http://localhost:8080/here")
		case _ = <-gone:
			log.Println("Gone")
			http.Get("http://localhost:8080/gone")
		}
	}

	// devs, err := keylogger.NewDevices()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	//
	// for _, val := range devs {
	// 	go func(val *keylogger.InputDevice) {
	// 		switch {
	// 		case strings.Contains(strings.ToLower(val.Name), "mouse"):
	// 		case strings.Contains(strings.ToLower(val.Name), "keyboard"):
	// 		case strings.Contains(strings.ToLower(val.Name), "pad"):
	// 		case strings.Contains(strings.ToLower(val.Name), "controller"):
	// 		case strings.Contains(strings.ToLower(val.Name), "key"):
	// 		default:
	// 			// return
	// 		}
	// 		//keyboard device file, on your system it will be diffrent!
	// 		rd := keylogger.NewKeyLogger(devs[val.Id])
	//
	// 		in, err := rd.Read()
	// 		if err != nil {
	// 			fmt.Println(err)
	// 			return
	// 		}
	// 		fmt.Println("Id->", val.Id, "Device->", val.Name)
	// 		for i := range in {
	// 			//we only need keypress
	// 			if i.Type == keylogger.EV_KEY {
	// 				fmt.Println(val.Name, i.KeyString())
	//
	// 			}
	// 		}
	//
	// 	}(val)
	// }
	runtime.Goexit()
}
