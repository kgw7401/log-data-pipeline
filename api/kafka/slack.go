package kafka

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/slack-go/slack"
)

func SendSlack(token string, channel string, e string, b []byte) {
	api := slack.New(token)
	var out bytes.Buffer
	json.Indent(&out, b, "", "\t")
	attachment1 := slack.Attachment{
		Title:   "스키마 오류",
		Pretext: "에러 메시지",
		Text:    e,
	}

	attachment2 := slack.Attachment{
		Title:      "스키마 오류",
		Fallback:   "선택 오류",
		CallbackID: "schema_error",
		Pretext:    "페이로드",
		Text:       "```" + out.String() + "```",
		// Actions: []slack.AttachmentAction{
		// 	{
		// 		Name: "Resolve",
		// 		Text: "Resolve",
		// 		Type: slack.ActionType("button"),
		// 		Confirm: &slack.ConfirmationField{
		// 			Title:       "Resolve Issue",
		// 			Text:        "해결 처리 하시겠습니까?",
		// 			OkText:      "Yes",
		// 			DismissText: "No",
		// 		},
		// 		Value: "resolve",
		// 	},
		// },
	}
	message := slack.MsgOptionAttachments(attachment1, attachment2)
	channelID, timestamp, err := api.PostMessage(channel, slack.MsgOptionText("", false), message)
	if err != nil {
		fmt.Printf("Could not send message: %v", err)
	}
	fmt.Printf("Message with buttons sucessfully sent to channel %s at %s", channelID, timestamp)
}
