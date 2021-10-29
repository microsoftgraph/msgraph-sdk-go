package graph
import (
    "strings"
    "errors"
)
// 
type CallState int

const (
    INCOMING_CALLSTATE CallState = iota
    ESTABLISHING_CALLSTATE
    ESTABLISHED_CALLSTATE
    HOLD_CALLSTATE
    TRANSFERRING_CALLSTATE
    TRANSFERACCEPTED_CALLSTATE
    REDIRECTING_CALLSTATE
    TERMINATING_CALLSTATE
    TERMINATED_CALLSTATE
    UNKNOWNFUTUREVALUE_CALLSTATE
)

func (i CallState) String() string {
    return []string{"INCOMING", "ESTABLISHING", "ESTABLISHED", "HOLD", "TRANSFERRING", "TRANSFERACCEPTED", "REDIRECTING", "TERMINATING", "TERMINATED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseCallState(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "INCOMING":
            return INCOMING_CALLSTATE, nil
        case "ESTABLISHING":
            return ESTABLISHING_CALLSTATE, nil
        case "ESTABLISHED":
            return ESTABLISHED_CALLSTATE, nil
        case "HOLD":
            return HOLD_CALLSTATE, nil
        case "TRANSFERRING":
            return TRANSFERRING_CALLSTATE, nil
        case "TRANSFERACCEPTED":
            return TRANSFERACCEPTED_CALLSTATE, nil
        case "REDIRECTING":
            return REDIRECTING_CALLSTATE, nil
        case "TERMINATING":
            return TERMINATING_CALLSTATE, nil
        case "TERMINATED":
            return TERMINATED_CALLSTATE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_CALLSTATE, nil
    }
    return 0, errors.New("Unknown CallState value: " + v)
}
func SerializeCallState(values []CallState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
