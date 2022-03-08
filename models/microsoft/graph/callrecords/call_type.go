package callrecords
import (
    "strings"
    "errors"
)
// Provides operations to manage the cloudCommunications singleton.
type CallType int

const (
    UNKNOWN_CALLTYPE CallType = iota
    GROUPCALL_CALLTYPE
    PEERTOPEER_CALLTYPE
    UNKNOWNFUTUREVALUE_CALLTYPE
)

func (i CallType) String() string {
    return []string{"UNKNOWN", "GROUPCALL", "PEERTOPEER", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseCallType(v string) (interface{}, error) {
    result := UNKNOWN_CALLTYPE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_CALLTYPE
        case "GROUPCALL":
            result = GROUPCALL_CALLTYPE
        case "PEERTOPEER":
            result = PEERTOPEER_CALLTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_CALLTYPE
        default:
            return 0, errors.New("Unknown CallType value: " + v)
    }
    return &result, nil
}
func SerializeCallType(values []CallType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
