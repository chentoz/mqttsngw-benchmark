package mqttsn

import (
	"encoding/binary"
	"log"
	"net"
	"time"

	common "github.com/chentoz/mqttsngw-benchmark/common"

	BLE "github.com/chentoz/mqttsngw-benchmark/BLELocation"

	"github.com/GaryBoone/GoStats/stats"
	"github.com/golang/protobuf/proto"
)

// Message describes a message
// UDP Message
type UDPMessage struct {
	ShortTopic string
	QoS        byte
	Payload    *[]byte
	Sent       time.Time
	Delivered  time.Time
	Error      bool
	MsgType    uint8
	MsgId      uint16
	Flags      uint8
}

type UDPClient struct {
	ID         int
	BrokerURL  string
	BrokerUser string
	BrokerPass string
	MsgTopic   string
	MsgSize    int
	MsgCount   int
	MsgQoS     byte
	Quiet      bool
}

func (c *UDPClient) Run(res chan *common.RunResults) {
	newMsgs := make(chan *UDPMessage)
	pubMsgs := make(chan *UDPMessage)
	doneGen := make(chan bool)
	donePub := make(chan bool)
	doneGet := make(chan bool)

	runResults := new(common.RunResults)

	started := time.Now()
	// start generator
	go c.genMessages(newMsgs, doneGen)
	// start publisher
	go c.pubMessages(newMsgs, pubMsgs, doneGen, donePub)

	runResults.ID = c.ID
	times := []float64{}
	for {
		select {
		case m := <-pubMsgs:
			if m.Error {
				log.Printf("CLIENT %v ERROR publishing message: %v: at %v\n", c.ID, c.MsgTopic, m.Sent.Unix())
				runResults.Failures++
			} else {
				log.Printf("Message published: %v: sent: %v delivered: %v flight time: %v\n", c.MsgTopic, m.Sent, m.Delivered, m.Delivered.Sub(m.Sent))
				runResults.Successes++
				times = append(times, m.Delivered.Sub(m.Sent).Seconds()*1000) // in milliseconds
			}
		case <-doneGet:
			// calculate results
			duration := time.Now().Sub(started)
			runResults.MsgTimeMin = stats.StatsMin(times)
			runResults.MsgTimeMax = stats.StatsMax(times)
			runResults.MsgTimeMean = stats.StatsMean(times)
			runResults.RunTime = duration.Seconds()
			runResults.MsgsPerSec = float64(runResults.Successes) / duration.Seconds()
			// calculate std if sample is > 1, otherwise leave as 0 (convention)
			if c.MsgCount > 1 {
				runResults.MsgTimeStd = stats.StatsSampleStandardDeviation(times)
			}

			// report results and exit
			res <- runResults
			return
		}
	}
}

func (c *UDPClient) genMessages(ch chan *UDPMessage, done chan bool) {
	for i := 0; i < c.MsgCount; i++ {
		msg := &UDPMessage{
			MsgType:    0x0c,
			MsgId:      0x0b00,
			Flags:      0x62, // 0110 0010 : qos -3
			ShortTopic: c.MsgTopic[0:2],
			QoS:        c.MsgQoS,
		}
		// ble location object
		bleLocation := &BLE.BLELocation{
			Id:              []byte("123456789"),
			Timestamp:       0,
			AccuracyArray:   []byte{0},
			MajorMinorArray: []byte{0},
			Barometer:       0,
			Status:          []byte{0},
		}
		payload, err := proto.Marshal(bleLocation)
		if err != nil {
			log.Fatal("marshaling error: ", err)
		}
		msg.Payload = &payload
		ch <- msg
	}
	done <- true
	return
}

func (c *UDPClient) pubMessages(in, out chan *UDPMessage, doneGen, donePub chan bool) {
	if !c.Quiet {
		log.Printf("CLIENT %v is connected to the broker %v\n", c.ID, c.BrokerURL)
	}
	ctr := 0
	for {
		select {
		case m := <-in:
			m.Sent = time.Now()

			data := c.genQosM1Packet(m)

			err := c.sendUDP(c.BrokerURL, &data)

			if err != nil {
				log.Printf("CLIENT %v Error sending message: %v\n", c.ID, err)
				m.Error = true
			} else {
				m.Delivered = time.Now()
				m.Error = false
			}
			out <- m

			if ctr > 0 && ctr%100 == 0 {
				if !c.Quiet {
					log.Printf("CLIENT %v published %v messages and keeps publishing...\n", c.ID, ctr)
				}
			}
			ctr++
		case <-doneGen:
			donePub <- true
			if !c.Quiet {
				log.Printf("CLIENT %v is done publishing\n", c.ID)
			}
			return
		}
	}

	/* 	opts := mqtt.NewClientOptions().
	   		AddBroker(c.BrokerURL).
	   		SetClientID(fmt.Sprintf("mqtt-benchmark-%v-%v", time.Now().Format(time.RFC3339Nano), c.ID)).
	   		SetCleanSession(true).
	   		SetAutoReconnect(true).
	   		SetOnConnectHandler(onConnected).
	   		SetConnectionLostHandler(func(client mqtt.Client, reason error) {
	   			log.Printf("CLIENT %v lost connection to the broker: %v. Will reconnect...\n", c.ID, reason.Error())
	   		})
	   	if c.BrokerUser != "" && c.BrokerPass != "" {
	   		opts.SetUsername(c.BrokerUser)
	   		opts.SetPassword(c.BrokerPass)
	   	} */
	//client := mqtt.NewClient(opts)
	//token := client.Connect()
	//token.Wait()

	//if token.Error() != nil {
	//log.Printf("CLIENT %v had error connecting to the broker: %v\n", c.ID, token.Error())
	//}
}

func (c *UDPClient) genQosM1Packet(m *UDPMessage) []byte {
	msgTypeByte := byte(m.MsgType)
	flagByte := byte(m.Flags)
	tidBytes := make([]byte, 2)
	tidBytes[0] = byte(m.ShortTopic[0])
	tidBytes[1] = byte(m.ShortTopic[1])
	midBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(midBytes, m.MsgId)
	lenByte := byte(1 + 1 + 1 + 2 + 2 + binary.Size(m.Payload))

	packet := make([]byte, lenByte)
	(packet)[0] = lenByte
	(packet)[1] = msgTypeByte
	(packet)[2] = flagByte
	copy((packet)[3:5], tidBytes)
	copy((packet)[5:7], midBytes)
	copy((packet)[7:], *m.Payload)

	//fmt.Println(packet)
	return packet
}

func (c *UDPClient) sendUDP(brokerURL string, data *[]byte) error {
	remoteAddr, err := net.ResolveUDPAddr("udp4", brokerURL)
	if err != nil {
		return err
	}
	conn, err := net.DialUDP("udp", nil, remoteAddr)
	if err != nil {
		return err
	}
	defer conn.Close()
	_, err = conn.Write(*data)
	if err != nil {
		return err
	}
	return nil
}
