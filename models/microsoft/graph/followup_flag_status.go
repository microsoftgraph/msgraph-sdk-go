package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "NOTFLAGGED":
            return NOTFLAGGED_FOLLOWUPFLAGSTATUS, nil
        case "COMPLETE":
            return COMPLETE_FOLLOWUPFLAGSTATUS, nil
        case "FLAGGED":
            return FLAGGED_FOLLOWUPFLAGSTATUS, nil
    }
    return 0, errors.New("Unknown FollowupFlagStatus value: " + v)
}
func SerializeFollowupFlagStatus(values []FollowupFlagStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
