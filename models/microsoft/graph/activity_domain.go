package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_ACTIVITYDOMAIN, nil
        case "WORK":
            return WORK_ACTIVITYDOMAIN, nil
        case "PERSONAL":
            return PERSONAL_ACTIVITYDOMAIN, nil
        case "UNRESTRICTED":
            return UNRESTRICTED_ACTIVITYDOMAIN, nil
    }
    return 0, errors.New("Unknown ActivityDomain value: " + v)
}
func SerializeActivityDomain(values []ActivityDomain) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
