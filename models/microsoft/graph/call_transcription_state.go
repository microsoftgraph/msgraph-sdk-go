package graph
import (
    "strings"
    "errors"
)
// 
type CallTranscriptionState int

const (
    NOTSTARTED_CALLTRANSCRIPTIONSTATE CallTranscriptionState = iota
    ACTIVE_CALLTRANSCRIPTIONSTATE
    INACTIVE_CALLTRANSCRIPTIONSTATE
    UNKNOWNFUTUREVALUE_CALLTRANSCRIPTIONSTATE
)

func (i CallTranscriptionState) String() string {
    return []string{"NOTSTARTED", "ACTIVE", "INACTIVE", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseCallTranscriptionState(v string) (interface{}, error) {
    result := NOTSTARTED_CALLTRANSCRIPTIONSTATE
    switch strings.ToUpper(v) {
        case "NOTSTARTED":
            result = NOTSTARTED_CALLTRANSCRIPTIONSTATE
        case "ACTIVE":
            result = ACTIVE_CALLTRANSCRIPTIONSTATE
        case "INACTIVE":
            result = INACTIVE_CALLTRANSCRIPTIONSTATE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_CALLTRANSCRIPTIONSTATE
        default:
            return 0, errors.New("Unknown CallTranscriptionState value: " + v)
    }
    return &result, nil
}
func SerializeCallTranscriptionState(values []CallTranscriptionState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
