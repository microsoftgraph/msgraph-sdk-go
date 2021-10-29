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
    switch strings.ToUpper(v) {
        case "NOTSTARTED":
            return NOTSTARTED_CALLTRANSCRIPTIONSTATE, nil
        case "ACTIVE":
            return ACTIVE_CALLTRANSCRIPTIONSTATE, nil
        case "INACTIVE":
            return INACTIVE_CALLTRANSCRIPTIONSTATE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_CALLTRANSCRIPTIONSTATE, nil
    }
    return 0, errors.New("Unknown CallTranscriptionState value: " + v)
}
func SerializeCallTranscriptionState(values []CallTranscriptionState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
