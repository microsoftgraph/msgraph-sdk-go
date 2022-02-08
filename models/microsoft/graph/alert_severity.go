package graph
import (
    "strings"
    "errors"
)
// 
type AlertSeverity int

const (
    UNKNOWN_ALERTSEVERITY AlertSeverity = iota
    INFORMATIONAL_ALERTSEVERITY
    LOW_ALERTSEVERITY
    MEDIUM_ALERTSEVERITY
    HIGH_ALERTSEVERITY
    UNKNOWNFUTUREVALUE_ALERTSEVERITY
)

func (i AlertSeverity) String() string {
    return []string{"UNKNOWN", "INFORMATIONAL", "LOW", "MEDIUM", "HIGH", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseAlertSeverity(v string) (interface{}, error) {
    result := UNKNOWN_ALERTSEVERITY
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_ALERTSEVERITY
        case "INFORMATIONAL":
            result = INFORMATIONAL_ALERTSEVERITY
        case "LOW":
            result = LOW_ALERTSEVERITY
        case "MEDIUM":
            result = MEDIUM_ALERTSEVERITY
        case "HIGH":
            result = HIGH_ALERTSEVERITY
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_ALERTSEVERITY
        default:
            return 0, errors.New("Unknown AlertSeverity value: " + v)
    }
    return &result, nil
}
func SerializeAlertSeverity(values []AlertSeverity) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
