package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the security singleton.
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
    result := UNKNOWN_ALERTFEEDBACK
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_ALERTFEEDBACK
        case "TRUEPOSITIVE":
            result = TRUEPOSITIVE_ALERTFEEDBACK
        case "FALSEPOSITIVE":
            result = FALSEPOSITIVE_ALERTFEEDBACK
        case "BENIGNPOSITIVE":
            result = BENIGNPOSITIVE_ALERTFEEDBACK
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_ALERTFEEDBACK
        default:
            return 0, errors.New("Unknown AlertFeedback value: " + v)
    }
    return &result, nil
}
func SerializeAlertFeedback(values []AlertFeedback) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
