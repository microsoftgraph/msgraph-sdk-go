package graph
import (
    "strings"
    "errors"
)
// 
type AlertFeedback int

const (
    UNKNOWN_ALERTFEEDBACK AlertFeedback = iota
    TRUEPOSITIVE_ALERTFEEDBACK
    FALSEPOSITIVE_ALERTFEEDBACK
    BENIGNPOSITIVE_ALERTFEEDBACK
    UNKNOWNFUTUREVALUE_ALERTFEEDBACK
)

func (i AlertFeedback) String() string {
    return []string{"UNKNOWN", "TRUEPOSITIVE", "FALSEPOSITIVE", "BENIGNPOSITIVE", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseAlertFeedback(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_ALERTFEEDBACK, nil
        case "TRUEPOSITIVE":
            return TRUEPOSITIVE_ALERTFEEDBACK, nil
        case "FALSEPOSITIVE":
            return FALSEPOSITIVE_ALERTFEEDBACK, nil
        case "BENIGNPOSITIVE":
            return BENIGNPOSITIVE_ALERTFEEDBACK, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_ALERTFEEDBACK, nil
    }
    return 0, errors.New("Unknown AlertFeedback value: " + v)
}
func SerializeAlertFeedback(values []AlertFeedback) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
