package graph
import (
    "strings"
    "errors"
)
// 
type TaskStatus int

const (
    NOTSTARTED_TASKSTATUS TaskStatus = iota
    INPROGRESS_TASKSTATUS
    COMPLETED_TASKSTATUS
    WAITINGONOTHERS_TASKSTATUS
    DEFERRED_TASKSTATUS
)

func (i TaskStatus) String() string {
    return []string{"NOTSTARTED", "INPROGRESS", "COMPLETED", "WAITINGONOTHERS", "DEFERRED"}[i]
}
func ParseTaskStatus(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "NOTSTARTED":
            return NOTSTARTED_TASKSTATUS, nil
        case "INPROGRESS":
            return INPROGRESS_TASKSTATUS, nil
        case "COMPLETED":
            return COMPLETED_TASKSTATUS, nil
        case "WAITINGONOTHERS":
            return WAITINGONOTHERS_TASKSTATUS, nil
        case "DEFERRED":
            return DEFERRED_TASKSTATUS, nil
    }
    return 0, errors.New("Unknown TaskStatus value: " + v)
}
func SerializeTaskStatus(values []TaskStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
