package externalconnectors
import (
    "strings"
    "errors"
)
// 
type AccessType int

const (
    GRANT_ACCESSTYPE AccessType = iota
    DENY_ACCESSTYPE
    UNKNOWNFUTUREVALUE_ACCESSTYPE
)

func (i AccessType) String() string {
    return []string{"GRANT", "DENY", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseAccessType(v string) (interface{}, error) {
    result := GRANT_ACCESSTYPE
    switch strings.ToUpper(v) {
        case "GRANT":
            result = GRANT_ACCESSTYPE
        case "DENY":
            result = DENY_ACCESSTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_ACCESSTYPE
        default:
            return 0, errors.New("Unknown AccessType value: " + v)
    }
    return &result, nil
}
func SerializeAccessType(values []AccessType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
