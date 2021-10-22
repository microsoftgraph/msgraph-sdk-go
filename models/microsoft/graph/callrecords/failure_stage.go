package callrecords
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_FAILURESTAGE, nil
        case "CALLSETUP":
            return CALLSETUP_FAILURESTAGE, nil
        case "MIDCALL":
            return MIDCALL_FAILURESTAGE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_FAILURESTAGE, nil
    }
    return 0, errors.New("Unknown FailureStage value: " + v)
}
func SerializeFailureStage(values []FailureStage) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
