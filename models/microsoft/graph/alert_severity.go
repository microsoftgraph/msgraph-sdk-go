package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_ALERTSEVERITY, nil
        case "INFORMATIONAL":
            return INFORMATIONAL_ALERTSEVERITY, nil
        case "LOW":
            return LOW_ALERTSEVERITY, nil
        case "MEDIUM":
            return MEDIUM_ALERTSEVERITY, nil
        case "HIGH":
            return HIGH_ALERTSEVERITY, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_ALERTSEVERITY, nil
    }
    return 0, errors.New("Unknown AlertSeverity value: " + v)
}
func SerializeAlertSeverity(values []AlertSeverity) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
