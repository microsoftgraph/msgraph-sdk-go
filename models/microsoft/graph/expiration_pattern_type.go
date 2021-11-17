package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "NOTSPECIFIED":
            return NOTSPECIFIED_EXPIRATIONPATTERNTYPE, nil
        case "NOEXPIRATION":
            return NOEXPIRATION_EXPIRATIONPATTERNTYPE, nil
        case "AFTERDATETIME":
            return AFTERDATETIME_EXPIRATIONPATTERNTYPE, nil
        case "AFTERDURATION":
            return AFTERDURATION_EXPIRATIONPATTERNTYPE, nil
    }
    return 0, errors.New("Unknown ExpirationPatternType value: " + v)
}
func SerializeExpirationPatternType(values []ExpirationPatternType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
