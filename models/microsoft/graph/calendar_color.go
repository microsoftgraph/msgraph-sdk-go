package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "AUTO":
            return AUTO_CALENDARCOLOR, nil
        case "LIGHTBLUE":
            return LIGHTBLUE_CALENDARCOLOR, nil
        case "LIGHTGREEN":
            return LIGHTGREEN_CALENDARCOLOR, nil
        case "LIGHTORANGE":
            return LIGHTORANGE_CALENDARCOLOR, nil
        case "LIGHTGRAY":
            return LIGHTGRAY_CALENDARCOLOR, nil
        case "LIGHTYELLOW":
            return LIGHTYELLOW_CALENDARCOLOR, nil
        case "LIGHTTEAL":
            return LIGHTTEAL_CALENDARCOLOR, nil
        case "LIGHTPINK":
            return LIGHTPINK_CALENDARCOLOR, nil
        case "LIGHTBROWN":
            return LIGHTBROWN_CALENDARCOLOR, nil
        case "LIGHTRED":
            return LIGHTRED_CALENDARCOLOR, nil
        case "MAXCOLOR":
            return MAXCOLOR_CALENDARCOLOR, nil
    }
    return 0, errors.New("Unknown CalendarColor value: " + v)
}
func SerializeCalendarColor(values []CalendarColor) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
