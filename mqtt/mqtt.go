package mqtt

import (
	"fmt"
	"log"
	"os"

	logger "github.com/qaldak/SysMonMQ/logging"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func PublishMessage(client MQTT.Client, topic, message string) {
    token := client.Publish(topic, 0, false, message)
    token.Wait()
    if token.Error() != nil {
        fmt.Printf("Error publishing message: %v\n", token.Error())
    }
}

func InitMQTT() {
    brokerURI := os.Getenv("MQTT_BROKER")
    topic := os.Getenv("MQTT_TOPIC")
    clientID := os.Getenv("HOSTNAME")

    opts := MQTT.NewClientOptions().AddBroker(brokerURI).SetClientID(clientID)
    client := MQTT.NewClient(opts)

    if token := client.Connect(); token.Wait() && token.Error() != nil {
        logger.Fatal(token.Error())
    }
    defer client.Disconnect(250)

    if token := client.Subscribe(topic, 0, onMessageReceived); token.Wait() && token.Error() != nil {
        logger.Fatal(token.Error())
        log.Fatal(token.Error())
    }
}

func onMessageReceived(client MQTT.Client, message MQTT.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", message.Payload(), message.Topic())
}
