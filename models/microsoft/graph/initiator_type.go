package graph
import (
    "strings"
    "errors"
)
// 
type InitiatorType int

const (
    USER_INITIATORTYPE InitiatorType = iota
    APPLICATION_INITIATORTYPE
    SYSTEM_INITIATORTYPE
    UNKNOWNFUTUREVALUE_INITIATORTYPE
)

func (i InitiatorType) String() string {
    return []string{"USER", "APPLICATION", "SYSTEM", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseInitiatorType(v string) (interface{}, error) {
    result := USER_INITIATORTYPE
    switch strings.ToUpper(v) {
        case "USER":
            result = USER_INITIATORTYPE
        case "APPLICATION":
            result = APPLICATION_INITIATORTYPE
        case "SYSTEM":
            result = SYSTEM_INITIATORTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_INITIATORTYPE
        default:
            return 0, errors.New("Unknown InitiatorType value: " + v)
    }
    return &result, nil
}
func SerializeInitiatorType(values []InitiatorType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
