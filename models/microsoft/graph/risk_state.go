package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_RISKSTATE, nil
        case "CONFIRMEDSAFE":
            return CONFIRMEDSAFE_RISKSTATE, nil
        case "REMEDIATED":
            return REMEDIATED_RISKSTATE, nil
        case "DISMISSED":
            return DISMISSED_RISKSTATE, nil
        case "ATRISK":
            return ATRISK_RISKSTATE, nil
        case "CONFIRMEDCOMPROMISED":
            return CONFIRMEDCOMPROMISED_RISKSTATE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_RISKSTATE, nil
    }
    return 0, errors.New("Unknown RiskState value: " + v)
}
func SerializeRiskState(values []RiskState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
