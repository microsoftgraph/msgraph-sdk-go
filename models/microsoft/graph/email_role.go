package graph
import (
    "strings"
    "errors"
)
// 
type EmailRole int

const (
    UNKNOWN_EMAILROLE EmailRole = iota
    SENDER_EMAILROLE
    RECIPIENT_EMAILROLE
    UNKNOWNFUTUREVALUE_EMAILROLE
)

func (i EmailRole) String() string {
    return []string{"UNKNOWN", "SENDER", "RECIPIENT", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseEmailRole(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_EMAILROLE, nil
        case "SENDER":
            return SENDER_EMAILROLE, nil
        case "RECIPIENT":
            return RECIPIENT_EMAILROLE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_EMAILROLE, nil
    }
    return 0, errors.New("Unknown EmailRole value: " + v)
}
func SerializeEmailRole(values []EmailRole) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
