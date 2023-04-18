package models
import (
    "errors"
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
    return []string{"reactionAdded", "reactionRemoved", "actionUndefined", "unknownFutureValue"}[i]
}
func ParseChatMessageActions(v string) (any, error) {
    result := REACTIONADDED_CHATMESSAGEACTIONS
    switch v {
        case "reactionAdded":
            result = REACTIONADDED_CHATMESSAGEACTIONS
        case "reactionRemoved":
            result = REACTIONREMOVED_CHATMESSAGEACTIONS
        case "actionUndefined":
            result = ACTIONUNDEFINED_CHATMESSAGEACTIONS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CHATMESSAGEACTIONS
        default:
            return 0, errors.New("Unknown ChatMessageActions value: " + v)
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
