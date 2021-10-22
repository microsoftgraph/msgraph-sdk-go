package graph
import (
    "strings"
    "errors"
)
type VppTokenAccountType int

const (
    BUSINESS_VPPTOKENACCOUNTTYPE VppTokenAccountType = iota
    EDUCATION_VPPTOKENACCOUNTTYPE
)

func (i VppTokenAccountType) String() string {
    return []string{"BUSINESS", "EDUCATION"}[i]
}
func ParseVppTokenAccountType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "BUSINESS":
            return BUSINESS_VPPTOKENACCOUNTTYPE, nil
        case "EDUCATION":
            return EDUCATION_VPPTOKENACCOUNTTYPE, nil
    }
    return 0, errors.New("Unknown VppTokenAccountType value: " + v)
}
func SerializeVppTokenAccountType(values []VppTokenAccountType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
