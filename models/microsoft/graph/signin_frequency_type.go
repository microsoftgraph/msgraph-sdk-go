package graph
import (
    "strings"
    "errors"
)
type SigninFrequencyType int

const (
    DAYS_SIGNINFREQUENCYTYPE SigninFrequencyType = iota
    HOURS_SIGNINFREQUENCYTYPE
)

func (i SigninFrequencyType) String() string {
    return []string{"DAYS", "HOURS"}[i]
}
func ParseSigninFrequencyType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "DAYS":
            return DAYS_SIGNINFREQUENCYTYPE, nil
        case "HOURS":
            return HOURS_SIGNINFREQUENCYTYPE, nil
    }
    return 0, errors.New("Unknown SigninFrequencyType value: " + v)
}
func SerializeSigninFrequencyType(values []SigninFrequencyType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
