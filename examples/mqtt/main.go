package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-co-op/gocron"
	"github.com/google/uuid"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	// NOTE: Set log level
	lf := log.Lmsgprefix | log.LstdFlags
	mqtt.ERROR = log.New(os.Stdout, "[ERROR] ", lf)
	mqtt.CRITICAL = log.New(os.Stdout, "[CRITICAL] ", lf)
	mqtt.WARN = log.New(os.Stdout, "[WARN]  ", lf)
	mqtt.DEBUG = log.New(os.Stdout, "[DEBUG] ", lf)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	signal.Notify(c, os.Kill, syscall.SIGTERM)

	// NOTE: MQTT 3.1.1 Client options
	clientOptions := mqtt.NewClientOptions()
	clientOptions.SetClientID("paho" + strings.Replace(uuid.NewString(), "-", "", -1))
	clientOptions.SetCleanSession(true)
	clientOptions.SetKeepAlive(time.Duration(60))
	clientOptions.AddBroker("ws://localhost:9001")
	clientOptions.SetUsername("mambo")
	clientOptions.SetPassword("mambo")
	clientOptions.SetProtocolVersion(4)

	clientOptions.SetConnectionLostHandler(func(client mqtt.Client, err error) {
		log.Println("Connection Lost", err)
	})
	clientOptions.SetReconnectingHandler(func(client mqtt.Client, options *mqtt.ClientOptions) {
		log.Println("Reconnection with ", clientOptions.ClientID)
	})
	clientOptions.SetOnConnectHandler(func(client mqtt.Client) {
		log.Println("Connected")
		subF := func(client mqtt.Client, message mqtt.Message) {
			log.Println("Subscribed: ", string(message.Payload()))
		}

		topic := "$SYS/broker/version"
		token := client.Subscribe(topic, 0, subF)
		token.Wait()
	})

	// NOTE: connect broker
	client := mqtt.NewClient(clientOptions)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// NOTE: Publish scheduling tasks
	s := gocron.NewScheduler(time.UTC)
	_, err := s.Every(10).Seconds().Do(func() {
		if client.IsConnected() {
			text := fmt.Sprintf("Message from %s", time.Now().Format(time.RFC3339))
			token := client.Publish("test/message", 0, false, text)
			token.Wait()
		}
	})
	if err != nil {
		return
	}
	s.StartAsync()

	<-c
}
