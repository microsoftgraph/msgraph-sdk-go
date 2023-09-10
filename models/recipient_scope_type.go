package models
import (
    "errors"
    "strings"
)
// 
type RecipientScopeType int

const (
    NONE_RECIPIENTSCOPETYPE RecipientScopeType = iota
    INTERNAL_RECIPIENTSCOPETYPE
    EXTERNAL_RECIPIENTSCOPETYPE
    EXTERNALPARTNER_RECIPIENTSCOPETYPE
    EXTERNALNONPARTNER_RECIPIENTSCOPETYPE
)

func (i RecipientScopeType) String() string {
    var values []string
    for p := RecipientScopeType(1); p <= EXTERNALNONPARTNER_RECIPIENTSCOPETYPE; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"none", "internal", "external", "externalPartner", "externalNonPartner"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseRecipientScopeType(v string) (any, error) {
    var result RecipientScopeType
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "none":
                result |= NONE_RECIPIENTSCOPETYPE
            case "internal":
                result |= INTERNAL_RECIPIENTSCOPETYPE
            case "external":
                result |= EXTERNAL_RECIPIENTSCOPETYPE
            case "externalPartner":
                result |= EXTERNALPARTNER_RECIPIENTSCOPETYPE
            case "externalNonPartner":
                result |= EXTERNALNONPARTNER_RECIPIENTSCOPETYPE
            default:
                return 0, errors.New("Unknown RecipientScopeType value: " + v)
        }
    }
    return &result, nil
}
func SerializeRecipientScopeType(values []RecipientScopeType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RecipientScopeType) isMultiValue() bool {
    return true
}
