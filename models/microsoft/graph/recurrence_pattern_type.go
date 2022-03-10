package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the collection of drive entities.
type RecurrencePatternType int

const (
    DAILY_RECURRENCEPATTERNTYPE RecurrencePatternType = iota
    WEEKLY_RECURRENCEPATTERNTYPE
    ABSOLUTEMONTHLY_RECURRENCEPATTERNTYPE
    RELATIVEMONTHLY_RECURRENCEPATTERNTYPE
    ABSOLUTEYEARLY_RECURRENCEPATTERNTYPE
    RELATIVEYEARLY_RECURRENCEPATTERNTYPE
)

func (i RecurrencePatternType) String() string {
    return []string{"DAILY", "WEEKLY", "ABSOLUTEMONTHLY", "RELATIVEMONTHLY", "ABSOLUTEYEARLY", "RELATIVEYEARLY"}[i]
}
func ParseRecurrencePatternType(v string) (interface{}, error) {
    result := DAILY_RECURRENCEPATTERNTYPE
    switch strings.ToUpper(v) {
        case "DAILY":
            result = DAILY_RECURRENCEPATTERNTYPE
        case "WEEKLY":
            result = WEEKLY_RECURRENCEPATTERNTYPE
        case "ABSOLUTEMONTHLY":
            result = ABSOLUTEMONTHLY_RECURRENCEPATTERNTYPE
        case "RELATIVEMONTHLY":
            result = RELATIVEMONTHLY_RECURRENCEPATTERNTYPE
        case "ABSOLUTEYEARLY":
            result = ABSOLUTEYEARLY_RECURRENCEPATTERNTYPE
        case "RELATIVEYEARLY":
            result = RELATIVEYEARLY_RECURRENCEPATTERNTYPE
        default:
            return 0, errors.New("Unknown RecurrencePatternType value: " + v)
    }
    return &result, nil
}
func SerializeRecurrencePatternType(values []RecurrencePatternType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
