package signal

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	http.DefaultClient.Timeout = time.Second * 5 //set http client timeout
}

////////////////// signal /////////////////////
func WaitForExit() os.Signal {
	return WaitForSignal(syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
}

func WaitForSignal(sources ...os.Signal) os.Signal {
	var s = make(chan os.Signal, 1)
	defer signal.Stop(s) //the second Ctrl+C will force shutdown

	signal.Notify(s, sources...)
	return <-s //blocked
}
