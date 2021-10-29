package graph
import (
    "strings"
    "errors"
)
// 
type AuthenticationMethodState int

const (
    ENABLED_AUTHENTICATIONMETHODSTATE AuthenticationMethodState = iota
    DISABLED_AUTHENTICATIONMETHODSTATE
)

func (i AuthenticationMethodState) String() string {
    return []string{"ENABLED", "DISABLED"}[i]
}
func ParseAuthenticationMethodState(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "ENABLED":
            return ENABLED_AUTHENTICATIONMETHODSTATE, nil
        case "DISABLED":
            return DISABLED_AUTHENTICATIONMETHODSTATE, nil
    }
    return 0, errors.New("Unknown AuthenticationMethodState value: " + v)
}
func SerializeAuthenticationMethodState(values []AuthenticationMethodState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
