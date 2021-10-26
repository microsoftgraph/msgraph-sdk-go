package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "SENDER":
            return SENDER_SCHEDULECHANGEREQUESTACTOR, nil
        case "RECIPIENT":
            return RECIPIENT_SCHEDULECHANGEREQUESTACTOR, nil
        case "MANAGER":
            return MANAGER_SCHEDULECHANGEREQUESTACTOR, nil
        case "SYSTEM":
            return SYSTEM_SCHEDULECHANGEREQUESTACTOR, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_SCHEDULECHANGEREQUESTACTOR, nil
    }
    return 0, errors.New("Unknown ScheduleChangeRequestActor value: " + v)
}
func SerializeScheduleChangeRequestActor(values []ScheduleChangeRequestActor) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
