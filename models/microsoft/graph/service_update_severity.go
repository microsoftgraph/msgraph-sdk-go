package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "NORMAL":
            return NORMAL_SERVICEUPDATESEVERITY, nil
        case "HIGH":
            return HIGH_SERVICEUPDATESEVERITY, nil
        case "CRITICAL":
            return CRITICAL_SERVICEUPDATESEVERITY, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_SERVICEUPDATESEVERITY, nil
    }
    return 0, errors.New("Unknown ServiceUpdateSeverity value: " + v)
}
func SerializeServiceUpdateSeverity(values []ServiceUpdateSeverity) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
