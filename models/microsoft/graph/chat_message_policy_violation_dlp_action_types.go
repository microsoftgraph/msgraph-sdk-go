package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the collection of chat entities.
type ChatMessagePolicyViolationDlpActionTypes int

const (
    NONE_CHATMESSAGEPOLICYVIOLATIONDLPACTIONTYPES ChatMessagePolicyViolationDlpActionTypes = iota
    NOTIFYSENDER_CHATMESSAGEPOLICYVIOLATIONDLPACTIONTYPES
    BLOCKACCESS_CHATMESSAGEPOLICYVIOLATIONDLPACTIONTYPES
    BLOCKACCESSEXTERNAL_CHATMESSAGEPOLICYVIOLATIONDLPACTIONTYPES
)

func (i ChatMessagePolicyViolationDlpActionTypes) String() string {
    return []string{"NONE", "NOTIFYSENDER", "BLOCKACCESS", "BLOCKACCESSEXTERNAL"}[i]
}
func ParseChatMessagePolicyViolationDlpActionTypes(v string) (interface{}, error) {
    result := NONE_CHATMESSAGEPOLICYVIOLATIONDLPACTIONTYPES
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_CHATMESSAGEPOLICYVIOLATIONDLPACTIONTYPES
        case "NOTIFYSENDER":
            result = NOTIFYSENDER_CHATMESSAGEPOLICYVIOLATIONDLPACTIONTYPES
        case "BLOCKACCESS":
            result = BLOCKACCESS_CHATMESSAGEPOLICYVIOLATIONDLPACTIONTYPES
        case "BLOCKACCESSEXTERNAL":
            result = BLOCKACCESSEXTERNAL_CHATMESSAGEPOLICYVIOLATIONDLPACTIONTYPES
        default:
            return 0, errors.New("Unknown ChatMessagePolicyViolationDlpActionTypes value: " + v)
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
