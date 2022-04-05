package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the admin singleton.
type ServiceUpdateSeverity int

const (
    NORMAL_SERVICEUPDATESEVERITY ServiceUpdateSeverity = iota
    HIGH_SERVICEUPDATESEVERITY
    CRITICAL_SERVICEUPDATESEVERITY
    UNKNOWNFUTUREVALUE_SERVICEUPDATESEVERITY
)

func (i ServiceUpdateSeverity) String() string {
    return []string{"NORMAL", "HIGH", "CRITICAL", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseServiceUpdateSeverity(v string) (interface{}, error) {
    result := NORMAL_SERVICEUPDATESEVERITY
    switch strings.ToUpper(v) {
        case "NORMAL":
            result = NORMAL_SERVICEUPDATESEVERITY
        case "HIGH":
            result = HIGH_SERVICEUPDATESEVERITY
        case "CRITICAL":
            result = CRITICAL_SERVICEUPDATESEVERITY
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_SERVICEUPDATESEVERITY
        default:
            return 0, errors.New("Unknown ServiceUpdateSeverity value: " + v)
    }
    return &result, nil
}
func SerializeServiceUpdateSeverity(values []ServiceUpdateSeverity) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
