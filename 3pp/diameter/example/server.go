package example

import (
	"github.com/fiorix/go-diameter/diam"
	"github.com/fiorix/go-diameter/diam/dict"
	"github.com/fiorix/go-diameter/diam/sm"
	"log"
)

func helloServer() {

	err := dict.Default.LoadFile("hello.xml")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(dict.Default.String())

	settings := &sm.Settings{
		OriginHost:       "wxdlong.com",
		OriginRealm:      "wxdlong",
		VendorID:         1002,
		ProductName:      "ycat",
		OriginStateID:    0,
		FirmwareRevision: 0,
		HostIPAddresses:  nil,
	}

	sms := sm.New(settings)

	sms.HandleFunc("ALL", handleALL) // Catch all.)
	err = diam.ListenAndServe("localhost:3868", sms, dict.Default)
	if err != nil {
		log.Fatal("ListenServe: ", err)
	}

}

func handleALL(c diam.Conn, m *diam.Message) {
	log.Printf("Received unexpected message from %s:\n%s", c.RemoteAddr(), m)
}
