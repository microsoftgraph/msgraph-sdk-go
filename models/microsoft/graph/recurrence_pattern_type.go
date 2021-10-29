package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "DAILY":
            return DAILY_RECURRENCEPATTERNTYPE, nil
        case "WEEKLY":
            return WEEKLY_RECURRENCEPATTERNTYPE, nil
        case "ABSOLUTEMONTHLY":
            return ABSOLUTEMONTHLY_RECURRENCEPATTERNTYPE, nil
        case "RELATIVEMONTHLY":
            return RELATIVEMONTHLY_RECURRENCEPATTERNTYPE, nil
        case "ABSOLUTEYEARLY":
            return ABSOLUTEYEARLY_RECURRENCEPATTERNTYPE, nil
        case "RELATIVEYEARLY":
            return RELATIVEYEARLY_RECURRENCEPATTERNTYPE, nil
    }
    return 0, errors.New("Unknown RecurrencePatternType value: " + v)
}
func SerializeRecurrencePatternType(values []RecurrencePatternType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
