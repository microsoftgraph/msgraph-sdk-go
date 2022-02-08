package graph
import (
    "strings"
    "errors"
)
// 
type WeekIndex int

const (
    FIRST_WEEKINDEX WeekIndex = iota
    SECOND_WEEKINDEX
    THIRD_WEEKINDEX
    FOURTH_WEEKINDEX
    LAST_WEEKINDEX
)

func (i WeekIndex) String() string {
    return []string{"FIRST", "SECOND", "THIRD", "FOURTH", "LAST"}[i]
}
func ParseWeekIndex(v string) (interface{}, error) {
    result := FIRST_WEEKINDEX
    switch strings.ToUpper(v) {
        case "FIRST":
            result = FIRST_WEEKINDEX
        case "SECOND":
            result = SECOND_WEEKINDEX
        case "THIRD":
            result = THIRD_WEEKINDEX
        case "FOURTH":
            result = FOURTH_WEEKINDEX
        case "LAST":
            result = LAST_WEEKINDEX
        default:
            return 0, errors.New("Unknown WeekIndex value: " + v)
    }
    return &result, nil
}
func SerializeWeekIndex(values []WeekIndex) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
