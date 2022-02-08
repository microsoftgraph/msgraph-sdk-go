package graph
import (
    "strings"
    "errors"
)
// 
type ChatMessagePolicyViolationVerdictDetailsTypes int

const (
    NONE_CHATMESSAGEPOLICYVIOLATIONVERDICTDETAILSTYPES ChatMessagePolicyViolationVerdictDetailsTypes = iota
    ALLOWFALSEPOSITIVEOVERRIDE_CHATMESSAGEPOLICYVIOLATIONVERDICTDETAILSTYPES
    ALLOWOVERRIDEWITHOUTJUSTIFICATION_CHATMESSAGEPOLICYVIOLATIONVERDICTDETAILSTYPES
    ALLOWOVERRIDEWITHJUSTIFICATION_CHATMESSAGEPOLICYVIOLATIONVERDICTDETAILSTYPES
)

func (i ChatMessagePolicyViolationVerdictDetailsTypes) String() string {
    return []string{"NONE", "ALLOWFALSEPOSITIVEOVERRIDE", "ALLOWOVERRIDEWITHOUTJUSTIFICATION", "ALLOWOVERRIDEWITHJUSTIFICATION"}[i]
}
func ParseChatMessagePolicyViolationVerdictDetailsTypes(v string) (interface{}, error) {
    result := NONE_CHATMESSAGEPOLICYVIOLATIONVERDICTDETAILSTYPES
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_CHATMESSAGEPOLICYVIOLATIONVERDICTDETAILSTYPES
        case "ALLOWFALSEPOSITIVEOVERRIDE":
            result = ALLOWFALSEPOSITIVEOVERRIDE_CHATMESSAGEPOLICYVIOLATIONVERDICTDETAILSTYPES
        case "ALLOWOVERRIDEWITHOUTJUSTIFICATION":
            result = ALLOWOVERRIDEWITHOUTJUSTIFICATION_CHATMESSAGEPOLICYVIOLATIONVERDICTDETAILSTYPES
        case "ALLOWOVERRIDEWITHJUSTIFICATION":
            result = ALLOWOVERRIDEWITHJUSTIFICATION_CHATMESSAGEPOLICYVIOLATIONVERDICTDETAILSTYPES
        default:
            return 0, errors.New("Unknown ChatMessagePolicyViolationVerdictDetailsTypes value: " + v)
    }
    return &result, nil
}
func SerializeChatMessagePolicyViolationVerdictDetailsTypes(values []ChatMessagePolicyViolationVerdictDetailsTypes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
