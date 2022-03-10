package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the collection of chat entities.
type ChatMessagePolicyViolationUserActionTypes int

const (
    NONE_CHATMESSAGEPOLICYVIOLATIONUSERACTIONTYPES ChatMessagePolicyViolationUserActionTypes = iota
    OVERRIDE_CHATMESSAGEPOLICYVIOLATIONUSERACTIONTYPES
    REPORTFALSEPOSITIVE_CHATMESSAGEPOLICYVIOLATIONUSERACTIONTYPES
)

func (i ChatMessagePolicyViolationUserActionTypes) String() string {
    return []string{"NONE", "OVERRIDE", "REPORTFALSEPOSITIVE"}[i]
}
func ParseChatMessagePolicyViolationUserActionTypes(v string) (interface{}, error) {
    result := NONE_CHATMESSAGEPOLICYVIOLATIONUSERACTIONTYPES
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_CHATMESSAGEPOLICYVIOLATIONUSERACTIONTYPES
        case "OVERRIDE":
            result = OVERRIDE_CHATMESSAGEPOLICYVIOLATIONUSERACTIONTYPES
        case "REPORTFALSEPOSITIVE":
            result = REPORTFALSEPOSITIVE_CHATMESSAGEPOLICYVIOLATIONUSERACTIONTYPES
        default:
            return 0, errors.New("Unknown ChatMessagePolicyViolationUserActionTypes value: " + v)
    }
    return &result, nil
}
func SerializeChatMessagePolicyViolationUserActionTypes(values []ChatMessagePolicyViolationUserActionTypes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
