package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the drive singleton.
type ActionState int

const (
    NONE_ACTIONSTATE ActionState = iota
    PENDING_ACTIONSTATE
    CANCELED_ACTIONSTATE
    ACTIVE_ACTIONSTATE
    DONE_ACTIONSTATE
    FAILED_ACTIONSTATE
    NOTSUPPORTED_ACTIONSTATE
)

func (i ActionState) String() string {
    return []string{"NONE", "PENDING", "CANCELED", "ACTIVE", "DONE", "FAILED", "NOTSUPPORTED"}[i]
}
func ParseActionState(v string) (interface{}, error) {
    result := NONE_ACTIONSTATE
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_ACTIONSTATE
        case "PENDING":
            result = PENDING_ACTIONSTATE
        case "CANCELED":
            result = CANCELED_ACTIONSTATE
        case "ACTIVE":
            result = ACTIVE_ACTIONSTATE
        case "DONE":
            result = DONE_ACTIONSTATE
        case "FAILED":
            result = FAILED_ACTIONSTATE
        case "NOTSUPPORTED":
            result = NOTSUPPORTED_ACTIONSTATE
        default:
            return 0, errors.New("Unknown ActionState value: " + v)
    }
    return &result, nil
}
func SerializeActionState(values []ActionState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
