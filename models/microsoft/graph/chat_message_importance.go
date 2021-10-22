package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "NORMAL":
            return NORMAL_CHATMESSAGEIMPORTANCE, nil
        case "HIGH":
            return HIGH_CHATMESSAGEIMPORTANCE, nil
        case "URGENT":
            return URGENT_CHATMESSAGEIMPORTANCE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_CHATMESSAGEIMPORTANCE, nil
    }
    return 0, errors.New("Unknown ChatMessageImportance value: " + v)
}
func SerializeChatMessageImportance(values []ChatMessageImportance) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
