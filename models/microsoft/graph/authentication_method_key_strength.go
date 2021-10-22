package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "NORMAL":
            return NORMAL_AUTHENTICATIONMETHODKEYSTRENGTH, nil
        case "WEAK":
            return WEAK_AUTHENTICATIONMETHODKEYSTRENGTH, nil
        case "UNKNOWN":
            return UNKNOWN_AUTHENTICATIONMETHODKEYSTRENGTH, nil
    }
    return 0, errors.New("Unknown AuthenticationMethodKeyStrength value: " + v)
}
func SerializeAuthenticationMethodKeyStrength(values []AuthenticationMethodKeyStrength) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
