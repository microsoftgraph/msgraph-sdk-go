package models
import (
    "errors"
)
// 
type ScopeOperatorMultiValuedComparisonType int

const (
    ALL_SCOPEOPERATORMULTIVALUEDCOMPARISONTYPE ScopeOperatorMultiValuedComparisonType = iota
    ANY_SCOPEOPERATORMULTIVALUEDCOMPARISONTYPE
)

func (i ScopeOperatorMultiValuedComparisonType) String() string {
    return []string{"All", "Any"}[i]
}
func ParseScopeOperatorMultiValuedComparisonType(v string) (any, error) {
    result := ALL_SCOPEOPERATORMULTIVALUEDCOMPARISONTYPE
    switch v {
        case "All":
            result = ALL_SCOPEOPERATORMULTIVALUEDCOMPARISONTYPE
        case "Any":
            result = ANY_SCOPEOPERATORMULTIVALUEDCOMPARISONTYPE
        default:
            return 0, errors.New("Unknown ScopeOperatorMultiValuedComparisonType value: " + v)
    }
    return &result, nil
}
func SerializeScopeOperatorMultiValuedComparisonType(values []ScopeOperatorMultiValuedComparisonType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
