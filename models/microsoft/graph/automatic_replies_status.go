package graph
import (
    "strings"
    "errors"
)
// 
type AutomaticRepliesStatus int

const (
    DISABLED_AUTOMATICREPLIESSTATUS AutomaticRepliesStatus = iota
    ALWAYSENABLED_AUTOMATICREPLIESSTATUS
    SCHEDULED_AUTOMATICREPLIESSTATUS
)

func (i AutomaticRepliesStatus) String() string {
    return []string{"DISABLED", "ALWAYSENABLED", "SCHEDULED"}[i]
}
func ParseAutomaticRepliesStatus(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "DISABLED":
            return DISABLED_AUTOMATICREPLIESSTATUS, nil
        case "ALWAYSENABLED":
            return ALWAYSENABLED_AUTOMATICREPLIESSTATUS, nil
        case "SCHEDULED":
            return SCHEDULED_AUTOMATICREPLIESSTATUS, nil
    }
    return 0, errors.New("Unknown AutomaticRepliesStatus value: " + v)
}
func SerializeAutomaticRepliesStatus(values []AutomaticRepliesStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
