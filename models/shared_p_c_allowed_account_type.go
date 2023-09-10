package models
import (
    "errors"
    "strings"
)
// Type of accounts that are allowed to share the PC.
type SharedPCAllowedAccountType int

const (
    // Only guest accounts.
    GUEST_SHAREDPCALLOWEDACCOUNTTYPE SharedPCAllowedAccountType = iota
    // Only domain-joined accounts.
    DOMAIN_SHAREDPCALLOWEDACCOUNTTYPE
)

func (i SharedPCAllowedAccountType) String() string {
    var values []string
    for p := SharedPCAllowedAccountType(1); p <= DOMAIN_SHAREDPCALLOWEDACCOUNTTYPE; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"guest", "domain"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseSharedPCAllowedAccountType(v string) (any, error) {
    var result SharedPCAllowedAccountType
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "guest":
                result |= GUEST_SHAREDPCALLOWEDACCOUNTTYPE
            case "domain":
                result |= DOMAIN_SHAREDPCALLOWEDACCOUNTTYPE
            default:
                return 0, errors.New("Unknown SharedPCAllowedAccountType value: " + v)
        }
    }
    return &result, nil
}
func SerializeSharedPCAllowedAccountType(values []SharedPCAllowedAccountType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SharedPCAllowedAccountType) isMultiValue() bool {
    return true
}
