package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the identityGovernance singleton.
type ExpirationPatternType int

const (
    NOTSPECIFIED_EXPIRATIONPATTERNTYPE ExpirationPatternType = iota
    NOEXPIRATION_EXPIRATIONPATTERNTYPE
    AFTERDATETIME_EXPIRATIONPATTERNTYPE
    AFTERDURATION_EXPIRATIONPATTERNTYPE
)

func (i ExpirationPatternType) String() string {
    return []string{"NOTSPECIFIED", "NOEXPIRATION", "AFTERDATETIME", "AFTERDURATION"}[i]
}
func ParseExpirationPatternType(v string) (interface{}, error) {
    result := NOTSPECIFIED_EXPIRATIONPATTERNTYPE
    switch strings.ToUpper(v) {
        case "NOTSPECIFIED":
            result = NOTSPECIFIED_EXPIRATIONPATTERNTYPE
        case "NOEXPIRATION":
            result = NOEXPIRATION_EXPIRATIONPATTERNTYPE
        case "AFTERDATETIME":
            result = AFTERDATETIME_EXPIRATIONPATTERNTYPE
        case "AFTERDURATION":
            result = AFTERDURATION_EXPIRATIONPATTERNTYPE
        default:
            return 0, errors.New("Unknown ExpirationPatternType value: " + v)
    }
    return &result, nil
}
func SerializeExpirationPatternType(values []ExpirationPatternType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
