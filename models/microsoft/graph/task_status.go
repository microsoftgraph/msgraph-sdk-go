package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the collection of drive entities.
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
    result := NOTSTARTED_TASKSTATUS
    switch strings.ToUpper(v) {
        case "NOTSTARTED":
            result = NOTSTARTED_TASKSTATUS
        case "INPROGRESS":
            result = INPROGRESS_TASKSTATUS
        case "COMPLETED":
            result = COMPLETED_TASKSTATUS
        case "WAITINGONOTHERS":
            result = WAITINGONOTHERS_TASKSTATUS
        case "DEFERRED":
            result = DEFERRED_TASKSTATUS
        default:
            return 0, errors.New("Unknown TaskStatus value: " + v)
    }
    return &result, nil
}
func SerializeTaskStatus(values []TaskStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
