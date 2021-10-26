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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_USERACCOUNTSECURITYTYPE, nil
        case "STANDARD":
            return STANDARD_USERACCOUNTSECURITYTYPE, nil
        case "POWER":
            return POWER_USERACCOUNTSECURITYTYPE, nil
        case "ADMINISTRATOR":
            return ADMINISTRATOR_USERACCOUNTSECURITYTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_USERACCOUNTSECURITYTYPE, nil
    }
    return 0, errors.New("Unknown UserAccountSecurityType value: " + v)
}
func SerializeUserAccountSecurityType(values []UserAccountSecurityType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
