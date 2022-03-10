package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the deviceAppManagement singleton.
type VppTokenAccountType int

const (
    BUSINESS_VPPTOKENACCOUNTTYPE VppTokenAccountType = iota
    EDUCATION_VPPTOKENACCOUNTTYPE
)

func (i VppTokenAccountType) String() string {
    return []string{"BUSINESS", "EDUCATION"}[i]
}
func ParseVppTokenAccountType(v string) (interface{}, error) {
    result := BUSINESS_VPPTOKENACCOUNTTYPE
    switch strings.ToUpper(v) {
        case "BUSINESS":
            result = BUSINESS_VPPTOKENACCOUNTTYPE
        case "EDUCATION":
            result = EDUCATION_VPPTOKENACCOUNTTYPE
        default:
            return 0, errors.New("Unknown VppTokenAccountType value: " + v)
    }
    return &result, nil
}
func SerializeVppTokenAccountType(values []VppTokenAccountType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
