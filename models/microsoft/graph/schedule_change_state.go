package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "PENDING":
            return PENDING_SCHEDULECHANGESTATE, nil
        case "APPROVED":
            return APPROVED_SCHEDULECHANGESTATE, nil
        case "DECLINED":
            return DECLINED_SCHEDULECHANGESTATE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_SCHEDULECHANGESTATE, nil
    }
    return 0, errors.New("Unknown ScheduleChangeState value: " + v)
}
func SerializeScheduleChangeState(values []ScheduleChangeState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
