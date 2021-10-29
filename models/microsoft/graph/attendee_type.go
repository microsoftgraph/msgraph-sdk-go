package graph
import (
    "strings"
    "errors"
)
// 
type AttendeeType int

const (
    REQUIRED_ATTENDEETYPE AttendeeType = iota
    OPTIONAL_ATTENDEETYPE
    RESOURCE_ATTENDEETYPE
)

func (i AttendeeType) String() string {
    return []string{"REQUIRED", "OPTIONAL", "RESOURCE"}[i]
}
func ParseAttendeeType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "REQUIRED":
            return REQUIRED_ATTENDEETYPE, nil
        case "OPTIONAL":
            return OPTIONAL_ATTENDEETYPE, nil
        case "RESOURCE":
            return RESOURCE_ATTENDEETYPE, nil
    }
    return 0, errors.New("Unknown AttendeeType value: " + v)
}
func SerializeAttendeeType(values []AttendeeType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
