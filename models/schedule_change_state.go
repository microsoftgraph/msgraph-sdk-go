package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the drive singleton.
type ScheduleChangeState int

const (
    PENDING_SCHEDULECHANGESTATE ScheduleChangeState = iota
    APPROVED_SCHEDULECHANGESTATE
    DECLINED_SCHEDULECHANGESTATE
    UNKNOWNFUTUREVALUE_SCHEDULECHANGESTATE
)

func (i ScheduleChangeState) String() string {
    return []string{"PENDING", "APPROVED", "DECLINED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseScheduleChangeState(v string) (interface{}, error) {
    result := PENDING_SCHEDULECHANGESTATE
    switch strings.ToUpper(v) {
        case "PENDING":
            result = PENDING_SCHEDULECHANGESTATE
        case "APPROVED":
            result = APPROVED_SCHEDULECHANGESTATE
        case "DECLINED":
            result = DECLINED_SCHEDULECHANGESTATE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_SCHEDULECHANGESTATE
        default:
            return 0, errors.New("Unknown ScheduleChangeState value: " + v)
    }
    return &result, nil
}
func SerializeScheduleChangeState(values []ScheduleChangeState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
