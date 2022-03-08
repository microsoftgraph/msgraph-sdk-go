package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the authenticationMethodsPolicy singleton.
type AdvancedConfigState int

const (
    DEFAULT_ADVANCEDCONFIGSTATE AdvancedConfigState = iota
    ENABLED_ADVANCEDCONFIGSTATE
    DISABLED_ADVANCEDCONFIGSTATE
    UNKNOWNFUTUREVALUE_ADVANCEDCONFIGSTATE
)

func (i AdvancedConfigState) String() string {
    return []string{"DEFAULT", "ENABLED", "DISABLED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseAdvancedConfigState(v string) (interface{}, error) {
    result := DEFAULT_ADVANCEDCONFIGSTATE
    switch strings.ToUpper(v) {
        case "DEFAULT":
            result = DEFAULT_ADVANCEDCONFIGSTATE
        case "ENABLED":
            result = ENABLED_ADVANCEDCONFIGSTATE
        case "DISABLED":
            result = DISABLED_ADVANCEDCONFIGSTATE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_ADVANCEDCONFIGSTATE
        default:
            return 0, errors.New("Unknown AdvancedConfigState value: " + v)
    }
    return &result, nil
}
func SerializeAdvancedConfigState(values []AdvancedConfigState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
