package graph
import (
    "strings"
    "errors"
)
type RecipientScopeType int

const (
    NONE_RECIPIENTSCOPETYPE RecipientScopeType = iota
    INTERNAL_RECIPIENTSCOPETYPE
    EXTERNAL_RECIPIENTSCOPETYPE
    EXTERNALPARTNER_RECIPIENTSCOPETYPE
    EXTERNALNONPARTNER_RECIPIENTSCOPETYPE
)

func (i RecipientScopeType) String() string {
    return []string{"NONE", "INTERNAL", "EXTERNAL", "EXTERNALPARTNER", "EXTERNALNONPARTNER"}[i]
}
func ParseRecipientScopeType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_RECIPIENTSCOPETYPE, nil
        case "INTERNAL":
            return INTERNAL_RECIPIENTSCOPETYPE, nil
        case "EXTERNAL":
            return EXTERNAL_RECIPIENTSCOPETYPE, nil
        case "EXTERNALPARTNER":
            return EXTERNALPARTNER_RECIPIENTSCOPETYPE, nil
        case "EXTERNALNONPARTNER":
            return EXTERNALNONPARTNER_RECIPIENTSCOPETYPE, nil
    }
    return 0, errors.New("Unknown RecipientScopeType value: " + v)
}
func SerializeRecipientScopeType(values []RecipientScopeType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
