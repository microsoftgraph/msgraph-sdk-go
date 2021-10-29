package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "SUNDAY":
            return SUNDAY_DAYOFWEEK, nil
        case "MONDAY":
            return MONDAY_DAYOFWEEK, nil
        case "TUESDAY":
            return TUESDAY_DAYOFWEEK, nil
        case "WEDNESDAY":
            return WEDNESDAY_DAYOFWEEK, nil
        case "THURSDAY":
            return THURSDAY_DAYOFWEEK, nil
        case "FRIDAY":
            return FRIDAY_DAYOFWEEK, nil
        case "SATURDAY":
            return SATURDAY_DAYOFWEEK, nil
    }
    return 0, errors.New("Unknown DayOfWeek value: " + v)
}
func SerializeDayOfWeek(values []DayOfWeek) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
