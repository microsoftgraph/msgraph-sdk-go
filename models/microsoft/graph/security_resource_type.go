package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_SECURITYRESOURCETYPE, nil
        case "ATTACKED":
            return ATTACKED_SECURITYRESOURCETYPE, nil
        case "RELATED":
            return RELATED_SECURITYRESOURCETYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_SECURITYRESOURCETYPE, nil
    }
    return 0, errors.New("Unknown SecurityResourceType value: " + v)
}
func SerializeSecurityResourceType(values []SecurityResourceType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
