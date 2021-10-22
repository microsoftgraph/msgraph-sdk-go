package callrecords
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_CALLTYPE, nil
        case "GROUPCALL":
            return GROUPCALL_CALLTYPE, nil
        case "PEERTOPEER":
            return PEERTOPEER_CALLTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_CALLTYPE, nil
    }
    return 0, errors.New("Unknown CallType value: " + v)
}
func SerializeCallType(values []CallType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
