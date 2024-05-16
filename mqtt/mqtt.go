package mqtt

import (
	"os"
	"strings"

	logger "github.com/qaldak/sysmonmq/internal/logging"
	"github.com/qaldak/sysmonmq/internal/utils"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

/*
Initialize MQTT connection and publish message to MQTT broker.
BrokerURI and topic get from .env file.
*/
func PublishMessage(message string) error {
	brokerURI := os.Getenv("MQTT_BROKER")
	topic := getTopic()
	clientID := utils.GetHostname()

	opts := MQTT.NewClientOptions().AddBroker(brokerURI).SetClientID(strings.ToUpper(clientID))
	logger.Info("opts: ", opts)

	client := MQTT.NewClient(opts)
	logger.Info("client: ", client)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		logger.Error(token.Error())
		return token.Error()
	}
	logger.Info("connected")
	defer client.Disconnect(250)

	token := client.Publish(topic, 0, false, message)
	token.Wait()
	if token.Error() != nil {
		logger.Error("Error publishing message: ", token.Error())
		return token.Error()
	}

	return nil
}

/*
Determine MQTT topic from .env file and replace placeholder '{hostname}' with hostname (uppercase) from local machine.
*/
func getTopic() (topic string) {
	topic = os.Getenv("MQTT_TOPIC")

	// replace placeholder "{hostname}"
	if strings.Contains(topic, "{hostname}") {
		topic = strings.ReplaceAll(topic, "{hostname}", utils.GetHostname(utils.Upper))
	}

	return topic
}
