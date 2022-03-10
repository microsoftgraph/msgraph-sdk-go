package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the auditLogRoot singleton.
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
    result := LOW_RISKLEVEL
    switch strings.ToUpper(v) {
        case "LOW":
            result = LOW_RISKLEVEL
        case "MEDIUM":
            result = MEDIUM_RISKLEVEL
        case "HIGH":
            result = HIGH_RISKLEVEL
        case "HIDDEN":
            result = HIDDEN_RISKLEVEL
        case "NONE":
            result = NONE_RISKLEVEL
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_RISKLEVEL
        default:
            return 0, errors.New("Unknown RiskLevel value: " + v)
    }
    return &result, nil
}
func SerializeRiskLevel(values []RiskLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
