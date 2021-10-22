package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "WHITE":
            return WHITE_SCHEDULEENTITYTHEME, nil
        case "BLUE":
            return BLUE_SCHEDULEENTITYTHEME, nil
        case "GREEN":
            return GREEN_SCHEDULEENTITYTHEME, nil
        case "PURPLE":
            return PURPLE_SCHEDULEENTITYTHEME, nil
        case "PINK":
            return PINK_SCHEDULEENTITYTHEME, nil
        case "YELLOW":
            return YELLOW_SCHEDULEENTITYTHEME, nil
        case "GRAY":
            return GRAY_SCHEDULEENTITYTHEME, nil
        case "DARKBLUE":
            return DARKBLUE_SCHEDULEENTITYTHEME, nil
        case "DARKGREEN":
            return DARKGREEN_SCHEDULEENTITYTHEME, nil
        case "DARKPURPLE":
            return DARKPURPLE_SCHEDULEENTITYTHEME, nil
        case "DARKPINK":
            return DARKPINK_SCHEDULEENTITYTHEME, nil
        case "DARKYELLOW":
            return DARKYELLOW_SCHEDULEENTITYTHEME, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_SCHEDULEENTITYTHEME, nil
    }
    return 0, errors.New("Unknown ScheduleEntityTheme value: " + v)
}
func SerializeScheduleEntityTheme(values []ScheduleEntityTheme) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
