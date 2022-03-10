package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the security singleton.
type ProcessIntegrityLevel int

const (
    UNKNOWN_PROCESSINTEGRITYLEVEL ProcessIntegrityLevel = iota
    UNTRUSTED_PROCESSINTEGRITYLEVEL
    LOW_PROCESSINTEGRITYLEVEL
    MEDIUM_PROCESSINTEGRITYLEVEL
    HIGH_PROCESSINTEGRITYLEVEL
    SYSTEM_PROCESSINTEGRITYLEVEL
    UNKNOWNFUTUREVALUE_PROCESSINTEGRITYLEVEL
)

func (i ProcessIntegrityLevel) String() string {
    return []string{"UNKNOWN", "UNTRUSTED", "LOW", "MEDIUM", "HIGH", "SYSTEM", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseProcessIntegrityLevel(v string) (interface{}, error) {
    result := UNKNOWN_PROCESSINTEGRITYLEVEL
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_PROCESSINTEGRITYLEVEL
        case "UNTRUSTED":
            result = UNTRUSTED_PROCESSINTEGRITYLEVEL
        case "LOW":
            result = LOW_PROCESSINTEGRITYLEVEL
        case "MEDIUM":
            result = MEDIUM_PROCESSINTEGRITYLEVEL
        case "HIGH":
            result = HIGH_PROCESSINTEGRITYLEVEL
        case "SYSTEM":
            result = SYSTEM_PROCESSINTEGRITYLEVEL
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_PROCESSINTEGRITYLEVEL
        default:
            return 0, errors.New("Unknown ProcessIntegrityLevel value: " + v)
    }
    return &result, nil
}
func SerializeProcessIntegrityLevel(values []ProcessIntegrityLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
