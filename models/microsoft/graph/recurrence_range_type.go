package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "ENDDATE":
            return ENDDATE_RECURRENCERANGETYPE, nil
        case "NOEND":
            return NOEND_RECURRENCERANGETYPE, nil
        case "NUMBERED":
            return NUMBERED_RECURRENCERANGETYPE, nil
    }
    return 0, errors.New("Unknown RecurrenceRangeType value: " + v)
}
func SerializeRecurrenceRangeType(values []RecurrenceRangeType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
