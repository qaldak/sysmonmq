package slack

import (
	"fmt"
	"os"

	logger "github.com/qaldak/sysmonmq/internal/logging"
	"github.com/qaldak/sysmonmq/internal/utils"
	"github.com/slack-go/slack"
)

/*
Publish error in Slack. Input param (error message) used for message body.
*/
func PostSlackMsg(input string) {
	msg := generateSlackMsg(input)
	logger.Debug(msg)

	token := os.Getenv("SLACK_AUTH_TOKEN")
	channelId := os.Getenv("SLACK_CHANNEL_ID")

	bot := slack.New(token)
	_, timestamp, err := bot.PostMessage(channelId, slack.MsgOptionText(msg, false))
	if err != nil {
		logger.Fatal("Failed sending Slack message. Error:", err)
	}

	logger.Info(fmt.Sprintf("Message sent successfully on Channel '%v' at '%v'.", channelId, timestamp))
}

/*
Build error message body for Slack post.
*/
func generateSlackMsg(input string) (msg string) {
	hn := utils.GetHostname()
	msg = fmt.Sprintf(":exclamation: Error! Failed to send system information from '%v' to MQTT broker. Reason: %v.\nSee log for more details", hn, input)
	return
}
