package graph
import (
    "strings"
    "errors"
)
// 
type UserAccountSecurityType int

const (
    UNKNOWN_USERACCOUNTSECURITYTYPE UserAccountSecurityType = iota
    STANDARD_USERACCOUNTSECURITYTYPE
    POWER_USERACCOUNTSECURITYTYPE
    ADMINISTRATOR_USERACCOUNTSECURITYTYPE
    UNKNOWNFUTUREVALUE_USERACCOUNTSECURITYTYPE
)

func (i UserAccountSecurityType) String() string {
    return []string{"UNKNOWN", "STANDARD", "POWER", "ADMINISTRATOR", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseUserAccountSecurityType(v string) (interface{}, error) {
    result := UNKNOWN_USERACCOUNTSECURITYTYPE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_USERACCOUNTSECURITYTYPE
        case "STANDARD":
            result = STANDARD_USERACCOUNTSECURITYTYPE
        case "POWER":
            result = POWER_USERACCOUNTSECURITYTYPE
        case "ADMINISTRATOR":
            result = ADMINISTRATOR_USERACCOUNTSECURITYTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_USERACCOUNTSECURITYTYPE
        default:
            return 0, errors.New("Unknown UserAccountSecurityType value: " + v)
    }
    return &result, nil
}
func SerializeUserAccountSecurityType(values []UserAccountSecurityType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
