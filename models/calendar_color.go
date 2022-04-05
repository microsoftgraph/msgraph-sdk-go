package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the collection of drive entities.
type CalendarColor int

const (
    AUTO_CALENDARCOLOR CalendarColor = iota
    LIGHTBLUE_CALENDARCOLOR
    LIGHTGREEN_CALENDARCOLOR
    LIGHTORANGE_CALENDARCOLOR
    LIGHTGRAY_CALENDARCOLOR
    LIGHTYELLOW_CALENDARCOLOR
    LIGHTTEAL_CALENDARCOLOR
    LIGHTPINK_CALENDARCOLOR
    LIGHTBROWN_CALENDARCOLOR
    LIGHTRED_CALENDARCOLOR
    MAXCOLOR_CALENDARCOLOR
)

func (i CalendarColor) String() string {
    return []string{"AUTO", "LIGHTBLUE", "LIGHTGREEN", "LIGHTORANGE", "LIGHTGRAY", "LIGHTYELLOW", "LIGHTTEAL", "LIGHTPINK", "LIGHTBROWN", "LIGHTRED", "MAXCOLOR"}[i]
}
func ParseCalendarColor(v string) (interface{}, error) {
    result := AUTO_CALENDARCOLOR
    switch strings.ToUpper(v) {
        case "AUTO":
            result = AUTO_CALENDARCOLOR
        case "LIGHTBLUE":
            result = LIGHTBLUE_CALENDARCOLOR
        case "LIGHTGREEN":
            result = LIGHTGREEN_CALENDARCOLOR
        case "LIGHTORANGE":
            result = LIGHTORANGE_CALENDARCOLOR
        case "LIGHTGRAY":
            result = LIGHTGRAY_CALENDARCOLOR
        case "LIGHTYELLOW":
            result = LIGHTYELLOW_CALENDARCOLOR
        case "LIGHTTEAL":
            result = LIGHTTEAL_CALENDARCOLOR
        case "LIGHTPINK":
            result = LIGHTPINK_CALENDARCOLOR
        case "LIGHTBROWN":
            result = LIGHTBROWN_CALENDARCOLOR
        case "LIGHTRED":
            result = LIGHTRED_CALENDARCOLOR
        case "MAXCOLOR":
            result = MAXCOLOR_CALENDARCOLOR
        default:
            return 0, errors.New("Unknown CalendarColor value: " + v)
    }
    return &result, nil
}
func SerializeCalendarColor(values []CalendarColor) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
