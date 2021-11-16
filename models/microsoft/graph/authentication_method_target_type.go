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
    switch strings.ToUpper(v) {
        case "USER":
            return USER_AUTHENTICATIONMETHODTARGETTYPE, nil
        case "GROUP":
            return GROUP_AUTHENTICATIONMETHODTARGETTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_AUTHENTICATIONMETHODTARGETTYPE, nil
    }
    return 0, errors.New("Unknown AuthenticationMethodTargetType value: " + v)
}
func SerializeAuthenticationMethodTargetType(values []AuthenticationMethodTargetType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
