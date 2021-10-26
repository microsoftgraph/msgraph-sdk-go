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
    switch strings.ToUpper(v) {
        case "FIRST":
            return FIRST_WEEKINDEX, nil
        case "SECOND":
            return SECOND_WEEKINDEX, nil
        case "THIRD":
            return THIRD_WEEKINDEX, nil
        case "FOURTH":
            return FOURTH_WEEKINDEX, nil
        case "LAST":
            return LAST_WEEKINDEX, nil
    }
    return 0, errors.New("Unknown WeekIndex value: " + v)
}
func SerializeWeekIndex(values []WeekIndex) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
