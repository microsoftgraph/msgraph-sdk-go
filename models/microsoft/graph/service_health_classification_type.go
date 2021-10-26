package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "ADVISORY":
            return ADVISORY_SERVICEHEALTHCLASSIFICATIONTYPE, nil
        case "INCIDENT":
            return INCIDENT_SERVICEHEALTHCLASSIFICATIONTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_SERVICEHEALTHCLASSIFICATIONTYPE, nil
    }
    return 0, errors.New("Unknown ServiceHealthClassificationType value: " + v)
}
func SerializeServiceHealthClassificationType(values []ServiceHealthClassificationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
