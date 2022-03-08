package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the cloudCommunications singleton.
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
    result := INCOMING_CALLSTATE
    switch strings.ToUpper(v) {
        case "INCOMING":
            result = INCOMING_CALLSTATE
        case "ESTABLISHING":
            result = ESTABLISHING_CALLSTATE
        case "ESTABLISHED":
            result = ESTABLISHED_CALLSTATE
        case "HOLD":
            result = HOLD_CALLSTATE
        case "TRANSFERRING":
            result = TRANSFERRING_CALLSTATE
        case "TRANSFERACCEPTED":
            result = TRANSFERACCEPTED_CALLSTATE
        case "REDIRECTING":
            result = REDIRECTING_CALLSTATE
        case "TERMINATING":
            result = TERMINATING_CALLSTATE
        case "TERMINATED":
            result = TERMINATED_CALLSTATE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_CALLSTATE
        default:
            return 0, errors.New("Unknown CallState value: " + v)
    }
    return &result, nil
}
func SerializeCallState(values []CallState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
