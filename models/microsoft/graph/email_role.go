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
    result := UNKNOWN_EMAILROLE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_EMAILROLE
        case "SENDER":
            result = SENDER_EMAILROLE
        case "RECIPIENT":
            result = RECIPIENT_EMAILROLE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_EMAILROLE
        default:
            return 0, errors.New("Unknown EmailRole value: " + v)
    }
    return &result, nil
}
func SerializeEmailRole(values []EmailRole) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
