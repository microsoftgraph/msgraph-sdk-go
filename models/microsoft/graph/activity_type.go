package graph
import (
    "strings"
    "errors"
)
// 
type ActivityType int

const (
    SIGNIN_ACTIVITYTYPE ActivityType = iota
    USER_ACTIVITYTYPE
    UNKNOWNFUTUREVALUE_ACTIVITYTYPE
)

func (i ActivityType) String() string {
    return []string{"SIGNIN", "USER", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseActivityType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "SIGNIN":
            return SIGNIN_ACTIVITYTYPE, nil
        case "USER":
            return USER_ACTIVITYTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_ACTIVITYTYPE, nil
    }
    return 0, errors.New("Unknown ActivityType value: " + v)
}
func SerializeActivityType(values []ActivityType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
