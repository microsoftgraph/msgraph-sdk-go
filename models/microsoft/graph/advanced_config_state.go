package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "DEFAULT":
            return DEFAULT_ADVANCEDCONFIGSTATE, nil
        case "ENABLED":
            return ENABLED_ADVANCEDCONFIGSTATE, nil
        case "DISABLED":
            return DISABLED_ADVANCEDCONFIGSTATE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_ADVANCEDCONFIGSTATE, nil
    }
    return 0, errors.New("Unknown AdvancedConfigState value: " + v)
}
func SerializeAdvancedConfigState(values []AdvancedConfigState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
