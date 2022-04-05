package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the collection of drive entities.
type ScheduleChangeRequestActor int

const (
    SENDER_SCHEDULECHANGEREQUESTACTOR ScheduleChangeRequestActor = iota
    RECIPIENT_SCHEDULECHANGEREQUESTACTOR
    MANAGER_SCHEDULECHANGEREQUESTACTOR
    SYSTEM_SCHEDULECHANGEREQUESTACTOR
    UNKNOWNFUTUREVALUE_SCHEDULECHANGEREQUESTACTOR
)

func (i ScheduleChangeRequestActor) String() string {
    return []string{"SENDER", "RECIPIENT", "MANAGER", "SYSTEM", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseScheduleChangeRequestActor(v string) (interface{}, error) {
    result := SENDER_SCHEDULECHANGEREQUESTACTOR
    switch strings.ToUpper(v) {
        case "SENDER":
            result = SENDER_SCHEDULECHANGEREQUESTACTOR
        case "RECIPIENT":
            result = RECIPIENT_SCHEDULECHANGEREQUESTACTOR
        case "MANAGER":
            result = MANAGER_SCHEDULECHANGEREQUESTACTOR
        case "SYSTEM":
            result = SYSTEM_SCHEDULECHANGEREQUESTACTOR
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_SCHEDULECHANGEREQUESTACTOR
        default:
            return 0, errors.New("Unknown ScheduleChangeRequestActor value: " + v)
    }
    return &result, nil
}
func SerializeScheduleChangeRequestActor(values []ScheduleChangeRequestActor) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
