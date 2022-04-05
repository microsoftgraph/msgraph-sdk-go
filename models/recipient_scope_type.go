package models
import (
    "strings"
    "errors"
)
// Provides operations to call the getMailTips method.
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
    result := NONE_RECIPIENTSCOPETYPE
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_RECIPIENTSCOPETYPE
        case "INTERNAL":
            result = INTERNAL_RECIPIENTSCOPETYPE
        case "EXTERNAL":
            result = EXTERNAL_RECIPIENTSCOPETYPE
        case "EXTERNALPARTNER":
            result = EXTERNALPARTNER_RECIPIENTSCOPETYPE
        case "EXTERNALNONPARTNER":
            result = EXTERNALNONPARTNER_RECIPIENTSCOPETYPE
        default:
            return 0, errors.New("Unknown RecipientScopeType value: " + v)
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
