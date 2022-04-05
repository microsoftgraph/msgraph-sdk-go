package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the collection of chat entities.
type ChatMessageType int

const (
    MESSAGE_CHATMESSAGETYPE ChatMessageType = iota
    CHATEVENT_CHATMESSAGETYPE
    TYPING_CHATMESSAGETYPE
    UNKNOWNFUTUREVALUE_CHATMESSAGETYPE
    SYSTEMEVENTMESSAGE_CHATMESSAGETYPE
)

func (i ChatMessageType) String() string {
    return []string{"MESSAGE", "CHATEVENT", "TYPING", "UNKNOWNFUTUREVALUE", "SYSTEMEVENTMESSAGE"}[i]
}
func ParseChatMessageType(v string) (interface{}, error) {
    result := MESSAGE_CHATMESSAGETYPE
    switch strings.ToUpper(v) {
        case "MESSAGE":
            result = MESSAGE_CHATMESSAGETYPE
        case "CHATEVENT":
            result = CHATEVENT_CHATMESSAGETYPE
        case "TYPING":
            result = TYPING_CHATMESSAGETYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_CHATMESSAGETYPE
        case "SYSTEMEVENTMESSAGE":
            result = SYSTEMEVENTMESSAGE_CHATMESSAGETYPE
        default:
            return 0, errors.New("Unknown ChatMessageType value: " + v)
    }
    return &result, nil
}
func SerializeChatMessageType(values []ChatMessageType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
