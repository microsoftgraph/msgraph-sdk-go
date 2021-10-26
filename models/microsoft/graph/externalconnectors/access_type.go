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
    switch strings.ToUpper(v) {
        case "GRANT":
            return GRANT_ACCESSTYPE, nil
        case "DENY":
            return DENY_ACCESSTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_ACCESSTYPE, nil
    }
    return 0, errors.New("Unknown AccessType value: " + v)
}
func SerializeAccessType(values []AccessType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
