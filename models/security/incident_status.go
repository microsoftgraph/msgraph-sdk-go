package security
import (
    "errors"
)
// Provides operations to manage the collection of agreement entities.
type IncidentStatus int

const (
    ACTIVE_INCIDENTSTATUS IncidentStatus = iota
    RESOLVED_INCIDENTSTATUS
    REDIRECTED_INCIDENTSTATUS
    UNKNOWNFUTUREVALUE_INCIDENTSTATUS
)

func (i IncidentStatus) String() string {
    return []string{"active", "resolved", "redirected", "unknownFutureValue"}[i]
}
func ParseIncidentStatus(v string) (interface{}, error) {
    result := ACTIVE_INCIDENTSTATUS
    switch v {
        case "active":
            result = ACTIVE_INCIDENTSTATUS
        case "resolved":
            result = RESOLVED_INCIDENTSTATUS
        case "redirected":
            result = REDIRECTED_INCIDENTSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_INCIDENTSTATUS
        default:
            return 0, errors.New("Unknown IncidentStatus value: " + v)
    }
    return &result, nil
}
func SerializeIncidentStatus(values []IncidentStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
