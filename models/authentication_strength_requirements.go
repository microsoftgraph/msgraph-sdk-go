package models
import (
    "errors"
)
// 
type AuthenticationStrengthRequirements int

const (
    NONE_AUTHENTICATIONSTRENGTHREQUIREMENTS AuthenticationStrengthRequirements = iota
    MFA_AUTHENTICATIONSTRENGTHREQUIREMENTS
    UNKNOWNFUTUREVALUE_AUTHENTICATIONSTRENGTHREQUIREMENTS
)

func (i AuthenticationStrengthRequirements) String() string {
    return []string{"none", "mfa", "unknownFutureValue"}[i]
}
func ParseAuthenticationStrengthRequirements(v string) (any, error) {
    result := NONE_AUTHENTICATIONSTRENGTHREQUIREMENTS
    switch v {
        case "none":
            result = NONE_AUTHENTICATIONSTRENGTHREQUIREMENTS
        case "mfa":
            result = MFA_AUTHENTICATIONSTRENGTHREQUIREMENTS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_AUTHENTICATIONSTRENGTHREQUIREMENTS
        default:
            return 0, errors.New("Unknown AuthenticationStrengthRequirements value: " + v)
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
