package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Joybaruah/ssl-tracker/pkg/api"
	"github.com/Joybaruah/ssl-tracker/pkg/inits"
)

func init() {
	inits.LoadEnv()
	inits.DBInit()
}

func main() {
	// conn, err := tls.Dial("tcp", "dashboard.eeki.info:443", nil)
	// if err != nil {
	// 	panic("Server doesn't support SSL certificate err: " + err.Error())
	// }

	// err = conn.VerifyHostname("dashboard.eeki.info")
	// if err != nil {
	// 	panic("Hostname doesn't match with certificate: " + err.Error())
	// }
	// expiry := conn.ConnectionState().PeerCertificates[0].NotAfter

	// fmt.Println("Issuer: ", conn.ConnectionState().PeerCertificates[0].Issuer)
	// fmt.Println("Expiry: ", expiry.Format(time.DateOnly))

	// Services
	go api.APIService()

	signalChanel := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(signalChanel, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-signalChanel
		fmt.Println("Server receive signal: ", sig)

		done <- true
	}()

	<-done
	time.Sleep(2 * time.Millisecond)

}
