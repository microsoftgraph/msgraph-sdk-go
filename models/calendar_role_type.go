package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the drive singleton.
type CalendarRoleType int

const (
    NONE_CALENDARROLETYPE CalendarRoleType = iota
    FREEBUSYREAD_CALENDARROLETYPE
    LIMITEDREAD_CALENDARROLETYPE
    READ_CALENDARROLETYPE
    WRITE_CALENDARROLETYPE
    DELEGATEWITHOUTPRIVATEEVENTACCESS_CALENDARROLETYPE
    DELEGATEWITHPRIVATEEVENTACCESS_CALENDARROLETYPE
    CUSTOM_CALENDARROLETYPE
)

func (i CalendarRoleType) String() string {
    return []string{"NONE", "FREEBUSYREAD", "LIMITEDREAD", "READ", "WRITE", "DELEGATEWITHOUTPRIVATEEVENTACCESS", "DELEGATEWITHPRIVATEEVENTACCESS", "CUSTOM"}[i]
}
func ParseCalendarRoleType(v string) (interface{}, error) {
    result := NONE_CALENDARROLETYPE
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_CALENDARROLETYPE
        case "FREEBUSYREAD":
            result = FREEBUSYREAD_CALENDARROLETYPE
        case "LIMITEDREAD":
            result = LIMITEDREAD_CALENDARROLETYPE
        case "READ":
            result = READ_CALENDARROLETYPE
        case "WRITE":
            result = WRITE_CALENDARROLETYPE
        case "DELEGATEWITHOUTPRIVATEEVENTACCESS":
            result = DELEGATEWITHOUTPRIVATEEVENTACCESS_CALENDARROLETYPE
        case "DELEGATEWITHPRIVATEEVENTACCESS":
            result = DELEGATEWITHPRIVATEEVENTACCESS_CALENDARROLETYPE
        case "CUSTOM":
            result = CUSTOM_CALENDARROLETYPE
        default:
            return 0, errors.New("Unknown CalendarRoleType value: " + v)
    }
    return &result, nil
}
func SerializeCalendarRoleType(values []CalendarRoleType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
