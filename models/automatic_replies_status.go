package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the collection of drive entities.
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
    result := DISABLED_AUTOMATICREPLIESSTATUS
    switch strings.ToUpper(v) {
        case "DISABLED":
            result = DISABLED_AUTOMATICREPLIESSTATUS
        case "ALWAYSENABLED":
            result = ALWAYSENABLED_AUTOMATICREPLIESSTATUS
        case "SCHEDULED":
            result = SCHEDULED_AUTOMATICREPLIESSTATUS
        default:
            return 0, errors.New("Unknown AutomaticRepliesStatus value: " + v)
    }
    return &result, nil
}
func SerializeAutomaticRepliesStatus(values []AutomaticRepliesStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
