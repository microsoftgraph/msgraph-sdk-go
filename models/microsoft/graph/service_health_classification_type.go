package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the admin singleton.
type ServiceHealthClassificationType int

const (
    ADVISORY_SERVICEHEALTHCLASSIFICATIONTYPE ServiceHealthClassificationType = iota
    INCIDENT_SERVICEHEALTHCLASSIFICATIONTYPE
    UNKNOWNFUTUREVALUE_SERVICEHEALTHCLASSIFICATIONTYPE
)

func (i ServiceHealthClassificationType) String() string {
    return []string{"ADVISORY", "INCIDENT", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseServiceHealthClassificationType(v string) (interface{}, error) {
    result := ADVISORY_SERVICEHEALTHCLASSIFICATIONTYPE
    switch strings.ToUpper(v) {
        case "ADVISORY":
            result = ADVISORY_SERVICEHEALTHCLASSIFICATIONTYPE
        case "INCIDENT":
            result = INCIDENT_SERVICEHEALTHCLASSIFICATIONTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_SERVICEHEALTHCLASSIFICATIONTYPE
        default:
            return 0, errors.New("Unknown ServiceHealthClassificationType value: " + v)
    }
    return &result, nil
}
func SerializeServiceHealthClassificationType(values []ServiceHealthClassificationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
