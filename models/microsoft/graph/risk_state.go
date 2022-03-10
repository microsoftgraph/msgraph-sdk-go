package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the auditLogRoot singleton.
type RiskState int

const (
    NONE_RISKSTATE RiskState = iota
    CONFIRMEDSAFE_RISKSTATE
    REMEDIATED_RISKSTATE
    DISMISSED_RISKSTATE
    ATRISK_RISKSTATE
    CONFIRMEDCOMPROMISED_RISKSTATE
    UNKNOWNFUTUREVALUE_RISKSTATE
)

func (i RiskState) String() string {
    return []string{"NONE", "CONFIRMEDSAFE", "REMEDIATED", "DISMISSED", "ATRISK", "CONFIRMEDCOMPROMISED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseRiskState(v string) (interface{}, error) {
    result := NONE_RISKSTATE
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_RISKSTATE
        case "CONFIRMEDSAFE":
            result = CONFIRMEDSAFE_RISKSTATE
        case "REMEDIATED":
            result = REMEDIATED_RISKSTATE
        case "DISMISSED":
            result = DISMISSED_RISKSTATE
        case "ATRISK":
            result = ATRISK_RISKSTATE
        case "CONFIRMEDCOMPROMISED":
            result = CONFIRMEDCOMPROMISED_RISKSTATE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_RISKSTATE
        default:
            return 0, errors.New("Unknown RiskState value: " + v)
    }
    return &result, nil
}
func SerializeRiskState(values []RiskState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
