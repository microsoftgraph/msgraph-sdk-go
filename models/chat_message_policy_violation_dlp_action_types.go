package models
import (
    "errors"
    "strings"
)
// 
type ChatMessagePolicyViolationDlpActionTypes int

const (
    NONE_CHATMESSAGEPOLICYVIOLATIONDLPACTIONTYPES ChatMessagePolicyViolationDlpActionTypes = iota
    NOTIFYSENDER_CHATMESSAGEPOLICYVIOLATIONDLPACTIONTYPES
    BLOCKACCESS_CHATMESSAGEPOLICYVIOLATIONDLPACTIONTYPES
    BLOCKACCESSEXTERNAL_CHATMESSAGEPOLICYVIOLATIONDLPACTIONTYPES
)

func (i ChatMessagePolicyViolationDlpActionTypes) String() string {
    var values []string
    for p := ChatMessagePolicyViolationDlpActionTypes(1); p <= BLOCKACCESSEXTERNAL_CHATMESSAGEPOLICYVIOLATIONDLPACTIONTYPES; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"none", "notifySender", "blockAccess", "blockAccessExternal"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseChatMessagePolicyViolationDlpActionTypes(v string) (any, error) {
    var result ChatMessagePolicyViolationDlpActionTypes
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "none":
                result |= NONE_CHATMESSAGEPOLICYVIOLATIONDLPACTIONTYPES
            case "notifySender":
                result |= NOTIFYSENDER_CHATMESSAGEPOLICYVIOLATIONDLPACTIONTYPES
            case "blockAccess":
                result |= BLOCKACCESS_CHATMESSAGEPOLICYVIOLATIONDLPACTIONTYPES
            case "blockAccessExternal":
                result |= BLOCKACCESSEXTERNAL_CHATMESSAGEPOLICYVIOLATIONDLPACTIONTYPES
            default:
                return 0, errors.New("Unknown ChatMessagePolicyViolationDlpActionTypes value: " + v)
        }
    }
    return &result, nil
}
func SerializeChatMessagePolicyViolationDlpActionTypes(values []ChatMessagePolicyViolationDlpActionTypes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ChatMessagePolicyViolationDlpActionTypes) isMultiValue() bool {
    return true
}
