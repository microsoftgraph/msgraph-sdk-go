package graph
import (
    "strings"
    "errors"
)
// 
type ChatMessageType int

const (
    MESSAGE_CHATMESSAGETYPE ChatMessageType = iota
    CHATEVENT_CHATMESSAGETYPE
    TYPING_CHATMESSAGETYPE
    UNKNOWNFUTUREVALUE_CHATMESSAGETYPE
)

func (i ChatMessageType) String() string {
    return []string{"MESSAGE", "CHATEVENT", "TYPING", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseChatMessageType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "MESSAGE":
            return MESSAGE_CHATMESSAGETYPE, nil
        case "CHATEVENT":
            return CHATEVENT_CHATMESSAGETYPE, nil
        case "TYPING":
            return TYPING_CHATMESSAGETYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_CHATMESSAGETYPE, nil
    }
    return 0, errors.New("Unknown ChatMessageType value: " + v)
}
func SerializeChatMessageType(values []ChatMessageType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
