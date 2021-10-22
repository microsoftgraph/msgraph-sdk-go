package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_ACTIONSTATE, nil
        case "PENDING":
            return PENDING_ACTIONSTATE, nil
        case "CANCELED":
            return CANCELED_ACTIONSTATE, nil
        case "ACTIVE":
            return ACTIVE_ACTIONSTATE, nil
        case "DONE":
            return DONE_ACTIONSTATE, nil
        case "FAILED":
            return FAILED_ACTIONSTATE, nil
        case "NOTSUPPORTED":
            return NOTSUPPORTED_ACTIONSTATE, nil
    }
    return 0, errors.New("Unknown ActionState value: " + v)
}
func SerializeActionState(values []ActionState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
