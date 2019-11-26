package example

import (
	"errors"
	"github.com/fiorix/go-diameter/diam"
	"github.com/fiorix/go-diameter/diam/avp"
	"github.com/fiorix/go-diameter/diam/datatype"
	"github.com/fiorix/go-diameter/diam/dict"
	"github.com/fiorix/go-diameter/diam/sm"
	"github.com/fiorix/go-diameter/diam/sm/smpeer"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"
)

func helloClient() {

	err := dict.Default.LoadFile("hello.xml")
	if err != nil {
		log.Fatal(err)
	}
	//log.Println(dict.Default.String())

	settings := &sm.Settings{
		OriginHost:       "wxdlong.com",
		OriginRealm:      "wxdlong",
		VendorID:         1002,
		ProductName:      "ycat",
		OriginStateID:    0,
		FirmwareRevision: 1,
		HostIPAddresses:  []datatype.Address{datatype.Address(net.ParseIP("127.0.0.1"))},
	}

	sms := sm.New(settings)

	client := sm.Client{
		Dict:               dict.Default,
		Handler:            sms,
		MaxRetransmits:     3,
		RetransmitInterval: time.Second,
		EnableWatchdog:     true,
		WatchdogInterval:   15 * time.Second,
		SupportedVendorID: []*diam.AVP{
			diam.NewAVP(avp.SupportedVendorID, avp.Mbit, 0, datatype.Unsigned32(888)),
		},
		AcctApplicationID: []*diam.AVP{
			diam.NewAVP(avp.AcctApplicationID, avp.Mbit, 0, datatype.Unsigned32(999)),
		},
		AuthApplicationID: []*diam.AVP{
			diam.NewAVP(avp.AuthApplicationID, avp.Mbit, 0, datatype.Unsigned32(4)),
		},
		VendorSpecificApplicationID: []*diam.AVP{
			diam.NewAVP(avp.VendorSpecificApplicationID, avp.Mbit, 0, &diam.GroupedAVP{
				AVP: []*diam.AVP{
					diam.NewAVP(avp.AcctApplicationID, avp.Mbit, 0, datatype.Unsigned32(999)),
					diam.NewAVP(avp.AuthApplicationID, avp.Mbit, 0, datatype.Unsigned32(4)),
				},
			}),
		},
	}
	conn, err := client.Dial("localhost:3868")
	if err != nil {
		log.Fatal("connection ", err)
	}

	defer conn.Close()

	sendHello(conn, settings)

	time.Sleep(20 * time.Second)

}

func sendHello(c diam.Conn, cfg *sm.Settings) error {
	// Get this client's metadata from the connection object,
	// which is set by the state machine after the handshake.
	// It contains the peer's Origin-Host and Realm from the
	// CER/CEA handshake. We use it to populate the AVPs below.
	meta, ok := smpeer.FromContext(c.Context())
	if !ok {
		return errors.New("peer metadata unavailable")
	}
	sid := "session;" + strconv.Itoa(int(rand.Uint32()))
	m := diam.NewRequest(111, 999, c.Dictionary())
	m.NewAVP(avp.SessionID, avp.Mbit, 0, datatype.UTF8String(sid))
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, cfg.OriginHost)
	m.NewAVP(avp.OriginRealm, avp.Mbit, 0, cfg.OriginRealm)
	m.NewAVP(avp.DestinationRealm, avp.Mbit, 0, meta.OriginRealm)
	m.NewAVP(avp.DestinationHost, avp.Mbit, 0, meta.OriginHost)
	m.NewAVP(avp.UserName, avp.Mbit, 0, datatype.UTF8String("foobar"))
	log.Printf("Sending HMR to %s\n%s", c.RemoteAddr(), m)
	_, err := m.WriteTo(c)
	return err
}
