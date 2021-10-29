package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_ALERTSTATUS, nil
        case "NEWALERT":
            return NEWALERT_ALERTSTATUS, nil
        case "INPROGRESS":
            return INPROGRESS_ALERTSTATUS, nil
        case "RESOLVED":
            return RESOLVED_ALERTSTATUS, nil
        case "DISMISSED":
            return DISMISSED_ALERTSTATUS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_ALERTSTATUS, nil
    }
    return 0, errors.New("Unknown AlertStatus value: " + v)
}
func SerializeAlertStatus(values []AlertStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
