package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_PROCESSINTEGRITYLEVEL, nil
        case "UNTRUSTED":
            return UNTRUSTED_PROCESSINTEGRITYLEVEL, nil
        case "LOW":
            return LOW_PROCESSINTEGRITYLEVEL, nil
        case "MEDIUM":
            return MEDIUM_PROCESSINTEGRITYLEVEL, nil
        case "HIGH":
            return HIGH_PROCESSINTEGRITYLEVEL, nil
        case "SYSTEM":
            return SYSTEM_PROCESSINTEGRITYLEVEL, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_PROCESSINTEGRITYLEVEL, nil
    }
    return 0, errors.New("Unknown ProcessIntegrityLevel value: " + v)
}
func SerializeProcessIntegrityLevel(values []ProcessIntegrityLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
