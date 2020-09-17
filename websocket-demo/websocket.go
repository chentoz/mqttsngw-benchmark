package main

import (
	"fmt"
	"log"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/golang/protobuf/proto"
	BLELocationpb "github.com/krylovsk/mqtt-benchmark/BLELocation"
)

func main() {
	donePub := make(chan bool)
	doneRecv := make(chan bool)
	looping := true
	ready := 0
	go pub(donePub)
	go get(doneRecv)

	for looping {
		select {
		case <-donePub:
			ready++
		case <-doneRecv:
			ready++
		default:
			if ready == 2 {
				fmt.Println("breaked")
				looping = false
			}
			fmt.Println("Looping")
		}
	}
}

func pub(donePub chan bool) {
	onConnected := func(client mqtt.Client) {
		count := 0
		for true { // continue publishing
			//			location := &Locationpb.Location{
			//				Id: []byte("9100000"),
			//			}

			bleLocation := &BLELocationpb.BLELocation{
				Id: []byte("91100000"),
			}

			//data, err := proto.Marshal(location)
			data, err := proto.Marshal(bleLocation)
			if err != nil {
				log.Fatal("marshaling error: ", err)
			}
			token := client.Publish("test", 0, false, data)
			token.Wait()

			if token.Error() != nil {
				log.Printf("CLIENT %v Error sending message: %v\n", "golang-client-pub", token.Error())
			} else {
				log.Println("published")
			}
			count++
			if count == 1000000000 {
				donePub <- true
			}
		}
	}

	opts := mqtt.NewClientOptions().
		AddBroker("udp://www.spidersens.cn:31337").
		SetClientID("golang-client-pub").
		SetCleanSession(true).
		SetAutoReconnect(true).
		SetOnConnectHandler(onConnected).
		SetConnectionLostHandler(func(client mqtt.Client, reason error) {
			log.Printf("CLIENT %v lost connection to the broker: %v. Will reconnect...\n", "golang-client", reason.Error())
		})

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic("CLIENT %v had error connecting to the broker: %v\n")
	}
}

func get(doneRecv chan bool) {
	count := 1000
	var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		bleLocation := &BLELocationpb.BLELocation{}
		proto.Unmarshal(msg.Payload(), bleLocation)
		length := len(bleLocation.Id)
		fmt.Printf("Get : %v \n", string(bleLocation.Id[:length]))
		count++
		if count >= 1000 {
			doneRecv <- true
		}
	}

	onConnected := func(client mqtt.Client) {
		if token := client.Subscribe("test", 0, f); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
			os.Exit(1)
		} else {
			fmt.Print("Subscribe topic " + "test" + " success\n")
		}
	}

	opts := mqtt.NewClientOptions().
		AddBroker("ws://www.spidersens.cn:3000").
		SetClientID("golang-client-get").
		//		SetDefaultPublishHandler(f).
		SetCleanSession(true).
		SetAutoReconnect(true).
		SetOnConnectHandler(onConnected).
		SetConnectionLostHandler(func(client mqtt.Client, reason error) {
			log.Printf("CLIENT %v lost connection to the broker: %v. Will reconnect...\n", "golang-client", reason.Error())
		})

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic("CLIENT %v had error connecting to the broker: %v\n")
	}
}
