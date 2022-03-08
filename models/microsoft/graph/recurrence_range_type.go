package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the identityGovernance singleton.
type RecurrenceRangeType int

const (
    ENDDATE_RECURRENCERANGETYPE RecurrenceRangeType = iota
    NOEND_RECURRENCERANGETYPE
    NUMBERED_RECURRENCERANGETYPE
)

func (i RecurrenceRangeType) String() string {
    return []string{"ENDDATE", "NOEND", "NUMBERED"}[i]
}
func ParseRecurrenceRangeType(v string) (interface{}, error) {
    result := ENDDATE_RECURRENCERANGETYPE
    switch strings.ToUpper(v) {
        case "ENDDATE":
            result = ENDDATE_RECURRENCERANGETYPE
        case "NOEND":
            result = NOEND_RECURRENCERANGETYPE
        case "NUMBERED":
            result = NUMBERED_RECURRENCERANGETYPE
        default:
            return 0, errors.New("Unknown RecurrenceRangeType value: " + v)
    }
    return &result, nil
}
func SerializeRecurrenceRangeType(values []RecurrenceRangeType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
