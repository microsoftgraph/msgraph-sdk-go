package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the drive singleton.
type FollowupFlagStatus int

const (
    NOTFLAGGED_FOLLOWUPFLAGSTATUS FollowupFlagStatus = iota
    COMPLETE_FOLLOWUPFLAGSTATUS
    FLAGGED_FOLLOWUPFLAGSTATUS
)

func (i FollowupFlagStatus) String() string {
    return []string{"NOTFLAGGED", "COMPLETE", "FLAGGED"}[i]
}
func ParseFollowupFlagStatus(v string) (interface{}, error) {
    result := NOTFLAGGED_FOLLOWUPFLAGSTATUS
    switch strings.ToUpper(v) {
        case "NOTFLAGGED":
            result = NOTFLAGGED_FOLLOWUPFLAGSTATUS
        case "COMPLETE":
            result = COMPLETE_FOLLOWUPFLAGSTATUS
        case "FLAGGED":
            result = FLAGGED_FOLLOWUPFLAGSTATUS
        default:
            return 0, errors.New("Unknown FollowupFlagStatus value: " + v)
    }
    return &result, nil
}
func SerializeFollowupFlagStatus(values []FollowupFlagStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
