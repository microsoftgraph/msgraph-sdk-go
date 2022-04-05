package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the security singleton.
type SecurityResourceType int

const (
    UNKNOWN_SECURITYRESOURCETYPE SecurityResourceType = iota
    ATTACKED_SECURITYRESOURCETYPE
    RELATED_SECURITYRESOURCETYPE
    UNKNOWNFUTUREVALUE_SECURITYRESOURCETYPE
)

func (i SecurityResourceType) String() string {
    return []string{"UNKNOWN", "ATTACKED", "RELATED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseSecurityResourceType(v string) (interface{}, error) {
    result := UNKNOWN_SECURITYRESOURCETYPE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_SECURITYRESOURCETYPE
        case "ATTACKED":
            result = ATTACKED_SECURITYRESOURCETYPE
        case "RELATED":
            result = RELATED_SECURITYRESOURCETYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_SECURITYRESOURCETYPE
        default:
            return 0, errors.New("Unknown SecurityResourceType value: " + v)
    }
    return &result, nil
}
func SerializeSecurityResourceType(values []SecurityResourceType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
