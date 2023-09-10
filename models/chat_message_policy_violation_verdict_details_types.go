package models
import (
    "errors"
    "strings"
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
    var values []string
    for p := ChatMessagePolicyViolationVerdictDetailsTypes(1); p <= ALLOWOVERRIDEWITHJUSTIFICATION_CHATMESSAGEPOLICYVIOLATIONVERDICTDETAILSTYPES; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"none", "allowFalsePositiveOverride", "allowOverrideWithoutJustification", "allowOverrideWithJustification"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseChatMessagePolicyViolationVerdictDetailsTypes(v string) (any, error) {
    var result ChatMessagePolicyViolationVerdictDetailsTypes
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "none":
                result |= NONE_CHATMESSAGEPOLICYVIOLATIONVERDICTDETAILSTYPES
            case "allowFalsePositiveOverride":
                result |= ALLOWFALSEPOSITIVEOVERRIDE_CHATMESSAGEPOLICYVIOLATIONVERDICTDETAILSTYPES
            case "allowOverrideWithoutJustification":
                result |= ALLOWOVERRIDEWITHOUTJUSTIFICATION_CHATMESSAGEPOLICYVIOLATIONVERDICTDETAILSTYPES
            case "allowOverrideWithJustification":
                result |= ALLOWOVERRIDEWITHJUSTIFICATION_CHATMESSAGEPOLICYVIOLATIONVERDICTDETAILSTYPES
            default:
                return 0, errors.New("Unknown ChatMessagePolicyViolationVerdictDetailsTypes value: " + v)
        }
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
func (i ChatMessagePolicyViolationVerdictDetailsTypes) isMultiValue() bool {
    return true
}
