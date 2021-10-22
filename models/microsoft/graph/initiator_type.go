package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "USER":
            return USER_INITIATORTYPE, nil
        case "APPLICATION":
            return APPLICATION_INITIATORTYPE, nil
        case "SYSTEM":
            return SYSTEM_INITIATORTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_INITIATORTYPE, nil
    }
    return 0, errors.New("Unknown InitiatorType value: " + v)
}
func SerializeInitiatorType(values []InitiatorType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
