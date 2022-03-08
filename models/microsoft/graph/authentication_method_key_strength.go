package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the collection of drive entities.
type AuthenticationMethodKeyStrength int

const (
    NORMAL_AUTHENTICATIONMETHODKEYSTRENGTH AuthenticationMethodKeyStrength = iota
    WEAK_AUTHENTICATIONMETHODKEYSTRENGTH
    UNKNOWN_AUTHENTICATIONMETHODKEYSTRENGTH
)

func (i AuthenticationMethodKeyStrength) String() string {
    return []string{"NORMAL", "WEAK", "UNKNOWN"}[i]
}
func ParseAuthenticationMethodKeyStrength(v string) (interface{}, error) {
    result := NORMAL_AUTHENTICATIONMETHODKEYSTRENGTH
    switch strings.ToUpper(v) {
        case "NORMAL":
            result = NORMAL_AUTHENTICATIONMETHODKEYSTRENGTH
        case "WEAK":
            result = WEAK_AUTHENTICATIONMETHODKEYSTRENGTH
        case "UNKNOWN":
            result = UNKNOWN_AUTHENTICATIONMETHODKEYSTRENGTH
        default:
            return 0, errors.New("Unknown AuthenticationMethodKeyStrength value: " + v)
    }
    return &result, nil
}
func SerializeAuthenticationMethodKeyStrength(values []AuthenticationMethodKeyStrength) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
