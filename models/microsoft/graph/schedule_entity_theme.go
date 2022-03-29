package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the educationRoot singleton.
type ScheduleEntityTheme int

const (
    WHITE_SCHEDULEENTITYTHEME ScheduleEntityTheme = iota
    BLUE_SCHEDULEENTITYTHEME
    GREEN_SCHEDULEENTITYTHEME
    PURPLE_SCHEDULEENTITYTHEME
    PINK_SCHEDULEENTITYTHEME
    YELLOW_SCHEDULEENTITYTHEME
    GRAY_SCHEDULEENTITYTHEME
    DARKBLUE_SCHEDULEENTITYTHEME
    DARKGREEN_SCHEDULEENTITYTHEME
    DARKPURPLE_SCHEDULEENTITYTHEME
    DARKPINK_SCHEDULEENTITYTHEME
    DARKYELLOW_SCHEDULEENTITYTHEME
    UNKNOWNFUTUREVALUE_SCHEDULEENTITYTHEME
)

func (i ScheduleEntityTheme) String() string {
    return []string{"WHITE", "BLUE", "GREEN", "PURPLE", "PINK", "YELLOW", "GRAY", "DARKBLUE", "DARKGREEN", "DARKPURPLE", "DARKPINK", "DARKYELLOW", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseScheduleEntityTheme(v string) (interface{}, error) {
    result := WHITE_SCHEDULEENTITYTHEME
    switch strings.ToUpper(v) {
        case "WHITE":
            result = WHITE_SCHEDULEENTITYTHEME
        case "BLUE":
            result = BLUE_SCHEDULEENTITYTHEME
        case "GREEN":
            result = GREEN_SCHEDULEENTITYTHEME
        case "PURPLE":
            result = PURPLE_SCHEDULEENTITYTHEME
        case "PINK":
            result = PINK_SCHEDULEENTITYTHEME
        case "YELLOW":
            result = YELLOW_SCHEDULEENTITYTHEME
        case "GRAY":
            result = GRAY_SCHEDULEENTITYTHEME
        case "DARKBLUE":
            result = DARKBLUE_SCHEDULEENTITYTHEME
        case "DARKGREEN":
            result = DARKGREEN_SCHEDULEENTITYTHEME
        case "DARKPURPLE":
            result = DARKPURPLE_SCHEDULEENTITYTHEME
        case "DARKPINK":
            result = DARKPINK_SCHEDULEENTITYTHEME
        case "DARKYELLOW":
            result = DARKYELLOW_SCHEDULEENTITYTHEME
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_SCHEDULEENTITYTHEME
        default:
            return 0, errors.New("Unknown ScheduleEntityTheme value: " + v)
    }
    return &result, nil
}
func SerializeScheduleEntityTheme(values []ScheduleEntityTheme) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
