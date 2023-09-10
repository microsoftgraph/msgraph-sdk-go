package models
import (
    "errors"
    "strings"
)
// 
type ChatMessagePolicyViolationUserActionTypes int

const (
    NONE_CHATMESSAGEPOLICYVIOLATIONUSERACTIONTYPES ChatMessagePolicyViolationUserActionTypes = iota
    OVERRIDE_CHATMESSAGEPOLICYVIOLATIONUSERACTIONTYPES
    REPORTFALSEPOSITIVE_CHATMESSAGEPOLICYVIOLATIONUSERACTIONTYPES
)

func (i ChatMessagePolicyViolationUserActionTypes) String() string {
    var values []string
    for p := ChatMessagePolicyViolationUserActionTypes(1); p <= REPORTFALSEPOSITIVE_CHATMESSAGEPOLICYVIOLATIONUSERACTIONTYPES; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"none", "override", "reportFalsePositive"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseChatMessagePolicyViolationUserActionTypes(v string) (any, error) {
    var result ChatMessagePolicyViolationUserActionTypes
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "none":
                result |= NONE_CHATMESSAGEPOLICYVIOLATIONUSERACTIONTYPES
            case "override":
                result |= OVERRIDE_CHATMESSAGEPOLICYVIOLATIONUSERACTIONTYPES
            case "reportFalsePositive":
                result |= REPORTFALSEPOSITIVE_CHATMESSAGEPOLICYVIOLATIONUSERACTIONTYPES
            default:
                return 0, errors.New("Unknown ChatMessagePolicyViolationUserActionTypes value: " + v)
        }
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
func (i ChatMessagePolicyViolationUserActionTypes) isMultiValue() bool {
    return true
}
