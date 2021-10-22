package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_CHATMESSAGEPOLICYVIOLATIONUSERACTIONTYPES, nil
        case "OVERRIDE":
            return OVERRIDE_CHATMESSAGEPOLICYVIOLATIONUSERACTIONTYPES, nil
        case "REPORTFALSEPOSITIVE":
            return REPORTFALSEPOSITIVE_CHATMESSAGEPOLICYVIOLATIONUSERACTIONTYPES, nil
    }
    return 0, errors.New("Unknown ChatMessagePolicyViolationUserActionTypes value: " + v)
}
func SerializeChatMessagePolicyViolationUserActionTypes(values []ChatMessagePolicyViolationUserActionTypes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
