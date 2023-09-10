package models
import (
    "errors"
    "strings"
)
// 
type ChatMessageActions int

const (
    REACTIONADDED_CHATMESSAGEACTIONS ChatMessageActions = iota
    REACTIONREMOVED_CHATMESSAGEACTIONS
    ACTIONUNDEFINED_CHATMESSAGEACTIONS
    UNKNOWNFUTUREVALUE_CHATMESSAGEACTIONS
)

func (i ChatMessageActions) String() string {
    var values []string
    for p := ChatMessageActions(1); p <= UNKNOWNFUTUREVALUE_CHATMESSAGEACTIONS; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"reactionAdded", "reactionRemoved", "actionUndefined", "unknownFutureValue"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseChatMessageActions(v string) (any, error) {
    var result ChatMessageActions
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "reactionAdded":
                result |= REACTIONADDED_CHATMESSAGEACTIONS
            case "reactionRemoved":
                result |= REACTIONREMOVED_CHATMESSAGEACTIONS
            case "actionUndefined":
                result |= ACTIONUNDEFINED_CHATMESSAGEACTIONS
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_CHATMESSAGEACTIONS
            default:
                return 0, errors.New("Unknown ChatMessageActions value: " + v)
        }
    }
    return &result, nil
}
func SerializeChatMessageActions(values []ChatMessageActions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ChatMessageActions) isMultiValue() bool {
    return true
}
