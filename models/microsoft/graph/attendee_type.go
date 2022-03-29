package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the educationRoot singleton.
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
    result := REQUIRED_ATTENDEETYPE
    switch strings.ToUpper(v) {
        case "REQUIRED":
            result = REQUIRED_ATTENDEETYPE
        case "OPTIONAL":
            result = OPTIONAL_ATTENDEETYPE
        case "RESOURCE":
            result = RESOURCE_ATTENDEETYPE
        default:
            return 0, errors.New("Unknown AttendeeType value: " + v)
    }
    return &result, nil
}
func SerializeAttendeeType(values []AttendeeType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
