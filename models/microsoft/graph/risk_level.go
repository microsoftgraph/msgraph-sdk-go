package graph
import (
    "strings"
    "errors"
)
// 
type RiskLevel int

const (
    LOW_RISKLEVEL RiskLevel = iota
    MEDIUM_RISKLEVEL
    HIGH_RISKLEVEL
    HIDDEN_RISKLEVEL
    NONE_RISKLEVEL
    UNKNOWNFUTUREVALUE_RISKLEVEL
)

func (i RiskLevel) String() string {
    return []string{"LOW", "MEDIUM", "HIGH", "HIDDEN", "NONE", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseRiskLevel(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "LOW":
            return LOW_RISKLEVEL, nil
        case "MEDIUM":
            return MEDIUM_RISKLEVEL, nil
        case "HIGH":
            return HIGH_RISKLEVEL, nil
        case "HIDDEN":
            return HIDDEN_RISKLEVEL, nil
        case "NONE":
            return NONE_RISKLEVEL, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_RISKLEVEL, nil
    }
    return 0, errors.New("Unknown RiskLevel value: " + v)
}
func SerializeRiskLevel(values []RiskLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
