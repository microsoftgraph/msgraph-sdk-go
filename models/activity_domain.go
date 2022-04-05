package models
import (
    "strings"
    "errors"
)
// Provides operations to call the findMeetingTimes method.
type ActivityDomain int

const (
    UNKNOWN_ACTIVITYDOMAIN ActivityDomain = iota
    WORK_ACTIVITYDOMAIN
    PERSONAL_ACTIVITYDOMAIN
    UNRESTRICTED_ACTIVITYDOMAIN
)

func (i ActivityDomain) String() string {
    return []string{"UNKNOWN", "WORK", "PERSONAL", "UNRESTRICTED"}[i]
}
func ParseActivityDomain(v string) (interface{}, error) {
    result := UNKNOWN_ACTIVITYDOMAIN
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_ACTIVITYDOMAIN
        case "WORK":
            result = WORK_ACTIVITYDOMAIN
        case "PERSONAL":
            result = PERSONAL_ACTIVITYDOMAIN
        case "UNRESTRICTED":
            result = UNRESTRICTED_ACTIVITYDOMAIN
        default:
            return 0, errors.New("Unknown ActivityDomain value: " + v)
    }
    return &result, nil
}
func SerializeActivityDomain(values []ActivityDomain) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
