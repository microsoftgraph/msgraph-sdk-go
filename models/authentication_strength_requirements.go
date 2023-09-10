package models
import (
    "errors"
    "strings"
)
// 
type AuthenticationStrengthRequirements int

const (
    NONE_AUTHENTICATIONSTRENGTHREQUIREMENTS AuthenticationStrengthRequirements = iota
    MFA_AUTHENTICATIONSTRENGTHREQUIREMENTS
    UNKNOWNFUTUREVALUE_AUTHENTICATIONSTRENGTHREQUIREMENTS
)

func (i AuthenticationStrengthRequirements) String() string {
    var values []string
    for p := AuthenticationStrengthRequirements(1); p <= UNKNOWNFUTUREVALUE_AUTHENTICATIONSTRENGTHREQUIREMENTS; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"none", "mfa", "unknownFutureValue"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseAuthenticationStrengthRequirements(v string) (any, error) {
    var result AuthenticationStrengthRequirements
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "none":
                result |= NONE_AUTHENTICATIONSTRENGTHREQUIREMENTS
            case "mfa":
                result |= MFA_AUTHENTICATIONSTRENGTHREQUIREMENTS
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_AUTHENTICATIONSTRENGTHREQUIREMENTS
            default:
                return 0, errors.New("Unknown AuthenticationStrengthRequirements value: " + v)
        }
    }
    return &result, nil
}
func SerializeAuthenticationStrengthRequirements(values []AuthenticationStrengthRequirements) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AuthenticationStrengthRequirements) isMultiValue() bool {
    return true
}
