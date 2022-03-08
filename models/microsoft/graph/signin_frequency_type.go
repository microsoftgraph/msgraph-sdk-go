package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the identityContainer singleton.
type SigninFrequencyType int

const (
    DAYS_SIGNINFREQUENCYTYPE SigninFrequencyType = iota
    HOURS_SIGNINFREQUENCYTYPE
)

func (i SigninFrequencyType) String() string {
    return []string{"DAYS", "HOURS"}[i]
}
func ParseSigninFrequencyType(v string) (interface{}, error) {
    result := DAYS_SIGNINFREQUENCYTYPE
    switch strings.ToUpper(v) {
        case "DAYS":
            result = DAYS_SIGNINFREQUENCYTYPE
        case "HOURS":
            result = HOURS_SIGNINFREQUENCYTYPE
        default:
            return 0, errors.New("Unknown SigninFrequencyType value: " + v)
    }
    return &result, nil
}
func SerializeSigninFrequencyType(values []SigninFrequencyType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
