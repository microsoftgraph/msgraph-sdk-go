package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the security singleton.
type AlertStatus int

const (
    UNKNOWN_ALERTSTATUS AlertStatus = iota
    NEWALERT_ALERTSTATUS
    INPROGRESS_ALERTSTATUS
    RESOLVED_ALERTSTATUS
    DISMISSED_ALERTSTATUS
    UNKNOWNFUTUREVALUE_ALERTSTATUS
)

func (i AlertStatus) String() string {
    return []string{"UNKNOWN", "NEWALERT", "INPROGRESS", "RESOLVED", "DISMISSED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseAlertStatus(v string) (interface{}, error) {
    result := UNKNOWN_ALERTSTATUS
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_ALERTSTATUS
        case "NEWALERT":
            result = NEWALERT_ALERTSTATUS
        case "INPROGRESS":
            result = INPROGRESS_ALERTSTATUS
        case "RESOLVED":
            result = RESOLVED_ALERTSTATUS
        case "DISMISSED":
            result = DISMISSED_ALERTSTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_ALERTSTATUS
        default:
            return 0, errors.New("Unknown AlertStatus value: " + v)
    }
    return &result, nil
}
func SerializeAlertStatus(values []AlertStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
