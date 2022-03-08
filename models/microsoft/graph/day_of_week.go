package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the drive singleton.
type DayOfWeek int

const (
    SUNDAY_DAYOFWEEK DayOfWeek = iota
    MONDAY_DAYOFWEEK
    TUESDAY_DAYOFWEEK
    WEDNESDAY_DAYOFWEEK
    THURSDAY_DAYOFWEEK
    FRIDAY_DAYOFWEEK
    SATURDAY_DAYOFWEEK
)

func (i DayOfWeek) String() string {
    return []string{"SUNDAY", "MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY"}[i]
}
func ParseDayOfWeek(v string) (interface{}, error) {
    result := SUNDAY_DAYOFWEEK
    switch strings.ToUpper(v) {
        case "SUNDAY":
            result = SUNDAY_DAYOFWEEK
        case "MONDAY":
            result = MONDAY_DAYOFWEEK
        case "TUESDAY":
            result = TUESDAY_DAYOFWEEK
        case "WEDNESDAY":
            result = WEDNESDAY_DAYOFWEEK
        case "THURSDAY":
            result = THURSDAY_DAYOFWEEK
        case "FRIDAY":
            result = FRIDAY_DAYOFWEEK
        case "SATURDAY":
            result = SATURDAY_DAYOFWEEK
        default:
            return 0, errors.New("Unknown DayOfWeek value: " + v)
    }
    return &result, nil
}
func SerializeDayOfWeek(values []DayOfWeek) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
