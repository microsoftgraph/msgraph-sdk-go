package graph
import (
    "strings"
    "errors"
)
// 
type AuthenticationMethodTargetType int

const (
    USER_AUTHENTICATIONMETHODTARGETTYPE AuthenticationMethodTargetType = iota
    GROUP_AUTHENTICATIONMETHODTARGETTYPE
    UNKNOWNFUTUREVALUE_AUTHENTICATIONMETHODTARGETTYPE
)

func (i AuthenticationMethodTargetType) String() string {
    return []string{"USER", "GROUP", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseAuthenticationMethodTargetType(v string) (interface{}, error) {
    result := USER_AUTHENTICATIONMETHODTARGETTYPE
    switch strings.ToUpper(v) {
        case "USER":
            result = USER_AUTHENTICATIONMETHODTARGETTYPE
        case "GROUP":
            result = GROUP_AUTHENTICATIONMETHODTARGETTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_AUTHENTICATIONMETHODTARGETTYPE
        default:
            return 0, errors.New("Unknown AuthenticationMethodTargetType value: " + v)
    }
    return &result, nil
}
func SerializeAuthenticationMethodTargetType(values []AuthenticationMethodTargetType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
