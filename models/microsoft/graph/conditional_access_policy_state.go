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
    switch strings.ToUpper(v) {
        case "ENABLED":
            return ENABLED_CONDITIONALACCESSPOLICYSTATE, nil
        case "DISABLED":
            return DISABLED_CONDITIONALACCESSPOLICYSTATE, nil
        case "ENABLEDFORREPORTINGBUTNOTENFORCED":
            return ENABLEDFORREPORTINGBUTNOTENFORCED_CONDITIONALACCESSPOLICYSTATE, nil
    }
    return 0, errors.New("Unknown ConditionalAccessPolicyState value: " + v)
}
func SerializeConditionalAccessPolicyState(values []ConditionalAccessPolicyState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
