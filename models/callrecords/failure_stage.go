package callrecords
import (
    "strings"
    "errors"
)
// Provides operations to manage the cloudCommunications singleton.
type FailureStage int

const (
    UNKNOWN_FAILURESTAGE FailureStage = iota
    CALLSETUP_FAILURESTAGE
    MIDCALL_FAILURESTAGE
    UNKNOWNFUTUREVALUE_FAILURESTAGE
)

func (i FailureStage) String() string {
    return []string{"UNKNOWN", "CALLSETUP", "MIDCALL", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseFailureStage(v string) (interface{}, error) {
    result := UNKNOWN_FAILURESTAGE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_FAILURESTAGE
        case "CALLSETUP":
            result = CALLSETUP_FAILURESTAGE
        case "MIDCALL":
            result = MIDCALL_FAILURESTAGE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_FAILURESTAGE
        default:
            return 0, errors.New("Unknown FailureStage value: " + v)
    }
    return &result, nil
}
func SerializeFailureStage(values []FailureStage) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
