package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the authenticationMethodsPolicy singleton.
type AuthenticationMethodState int

const (
    ENABLED_AUTHENTICATIONMETHODSTATE AuthenticationMethodState = iota
    DISABLED_AUTHENTICATIONMETHODSTATE
)

func (i AuthenticationMethodState) String() string {
    return []string{"ENABLED", "DISABLED"}[i]
}
func ParseAuthenticationMethodState(v string) (interface{}, error) {
    result := ENABLED_AUTHENTICATIONMETHODSTATE
    switch strings.ToUpper(v) {
        case "ENABLED":
            result = ENABLED_AUTHENTICATIONMETHODSTATE
        case "DISABLED":
            result = DISABLED_AUTHENTICATIONMETHODSTATE
        default:
            return 0, errors.New("Unknown AuthenticationMethodState value: " + v)
    }
    return &result, nil
}
func SerializeAuthenticationMethodState(values []AuthenticationMethodState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
