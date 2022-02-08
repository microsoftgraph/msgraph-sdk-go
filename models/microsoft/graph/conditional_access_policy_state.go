package graph
import (
    "strings"
    "errors"
)
// 
type ConditionalAccessPolicyState int

const (
    ENABLED_CONDITIONALACCESSPOLICYSTATE ConditionalAccessPolicyState = iota
    DISABLED_CONDITIONALACCESSPOLICYSTATE
    ENABLEDFORREPORTINGBUTNOTENFORCED_CONDITIONALACCESSPOLICYSTATE
)

func (i ConditionalAccessPolicyState) String() string {
    return []string{"ENABLED", "DISABLED", "ENABLEDFORREPORTINGBUTNOTENFORCED"}[i]
}
func ParseConditionalAccessPolicyState(v string) (interface{}, error) {
    result := ENABLED_CONDITIONALACCESSPOLICYSTATE
    switch strings.ToUpper(v) {
        case "ENABLED":
            result = ENABLED_CONDITIONALACCESSPOLICYSTATE
        case "DISABLED":
            result = DISABLED_CONDITIONALACCESSPOLICYSTATE
        case "ENABLEDFORREPORTINGBUTNOTENFORCED":
            result = ENABLEDFORREPORTINGBUTNOTENFORCED_CONDITIONALACCESSPOLICYSTATE
        default:
            return 0, errors.New("Unknown ConditionalAccessPolicyState value: " + v)
    }
    return &result, nil
}
func SerializeConditionalAccessPolicyState(values []ConditionalAccessPolicyState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
