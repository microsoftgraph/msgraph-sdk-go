package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_CALENDARROLETYPE, nil
        case "FREEBUSYREAD":
            return FREEBUSYREAD_CALENDARROLETYPE, nil
        case "LIMITEDREAD":
            return LIMITEDREAD_CALENDARROLETYPE, nil
        case "READ":
            return READ_CALENDARROLETYPE, nil
        case "WRITE":
            return WRITE_CALENDARROLETYPE, nil
        case "DELEGATEWITHOUTPRIVATEEVENTACCESS":
            return DELEGATEWITHOUTPRIVATEEVENTACCESS_CALENDARROLETYPE, nil
        case "DELEGATEWITHPRIVATEEVENTACCESS":
            return DELEGATEWITHPRIVATEEVENTACCESS_CALENDARROLETYPE, nil
        case "CUSTOM":
            return CUSTOM_CALENDARROLETYPE, nil
    }
    return 0, errors.New("Unknown CalendarRoleType value: " + v)
}
func SerializeCalendarRoleType(values []CalendarRoleType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
