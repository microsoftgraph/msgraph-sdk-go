package graph
import (
    "strings"
    "errors"
)
// 
type ChatMessageImportance int

const (
    NORMAL_CHATMESSAGEIMPORTANCE ChatMessageImportance = iota
    HIGH_CHATMESSAGEIMPORTANCE
    URGENT_CHATMESSAGEIMPORTANCE
    UNKNOWNFUTUREVALUE_CHATMESSAGEIMPORTANCE
)

func (i ChatMessageImportance) String() string {
    return []string{"NORMAL", "HIGH", "URGENT", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseChatMessageImportance(v string) (interface{}, error) {
    result := NORMAL_CHATMESSAGEIMPORTANCE
    switch strings.ToUpper(v) {
        case "NORMAL":
            result = NORMAL_CHATMESSAGEIMPORTANCE
        case "HIGH":
            result = HIGH_CHATMESSAGEIMPORTANCE
        case "URGENT":
            result = URGENT_CHATMESSAGEIMPORTANCE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_CHATMESSAGEIMPORTANCE
        default:
            return 0, errors.New("Unknown ChatMessageImportance value: " + v)
    }
    return &result, nil
}
func SerializeChatMessageImportance(values []ChatMessageImportance) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
