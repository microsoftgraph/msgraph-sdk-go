package graph
import (
    "strings"
    "errors"
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
    return []string{"NONE", "NOTIFYSENDER", "BLOCKACCESS", "BLOCKACCESSEXTERNAL"}[i]
}
func ParseChatMessagePolicyViolationDlpActionTypes(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_CHATMESSAGEPOLICYVIOLATIONDLPACTIONTYPES, nil
        case "NOTIFYSENDER":
            return NOTIFYSENDER_CHATMESSAGEPOLICYVIOLATIONDLPACTIONTYPES, nil
        case "BLOCKACCESS":
            return BLOCKACCESS_CHATMESSAGEPOLICYVIOLATIONDLPACTIONTYPES, nil
        case "BLOCKACCESSEXTERNAL":
            return BLOCKACCESSEXTERNAL_CHATMESSAGEPOLICYVIOLATIONDLPACTIONTYPES, nil
    }
    return 0, errors.New("Unknown ChatMessagePolicyViolationDlpActionTypes value: " + v)
}
func SerializeChatMessagePolicyViolationDlpActionTypes(values []ChatMessagePolicyViolationDlpActionTypes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
