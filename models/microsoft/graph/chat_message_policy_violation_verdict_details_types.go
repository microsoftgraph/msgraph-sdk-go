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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_CHATMESSAGEPOLICYVIOLATIONVERDICTDETAILSTYPES, nil
        case "ALLOWFALSEPOSITIVEOVERRIDE":
            return ALLOWFALSEPOSITIVEOVERRIDE_CHATMESSAGEPOLICYVIOLATIONVERDICTDETAILSTYPES, nil
        case "ALLOWOVERRIDEWITHOUTJUSTIFICATION":
            return ALLOWOVERRIDEWITHOUTJUSTIFICATION_CHATMESSAGEPOLICYVIOLATIONVERDICTDETAILSTYPES, nil
        case "ALLOWOVERRIDEWITHJUSTIFICATION":
            return ALLOWOVERRIDEWITHJUSTIFICATION_CHATMESSAGEPOLICYVIOLATIONVERDICTDETAILSTYPES, nil
    }
    return 0, errors.New("Unknown ChatMessagePolicyViolationVerdictDetailsTypes value: " + v)
}
func SerializeChatMessagePolicyViolationVerdictDetailsTypes(values []ChatMessagePolicyViolationVerdictDetailsTypes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
