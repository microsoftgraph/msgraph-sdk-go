package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the identityGovernance singleton.
type AccessPackageSubjectType int

const (
    NOTSPECIFIED_ACCESSPACKAGESUBJECTTYPE AccessPackageSubjectType = iota
    USER_ACCESSPACKAGESUBJECTTYPE
    SERVICEPRINCIPAL_ACCESSPACKAGESUBJECTTYPE
    UNKNOWNFUTUREVALUE_ACCESSPACKAGESUBJECTTYPE
)

func (i AccessPackageSubjectType) String() string {
    return []string{"NOTSPECIFIED", "USER", "SERVICEPRINCIPAL", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseAccessPackageSubjectType(v string) (interface{}, error) {
    result := NOTSPECIFIED_ACCESSPACKAGESUBJECTTYPE
    switch strings.ToUpper(v) {
        case "NOTSPECIFIED":
            result = NOTSPECIFIED_ACCESSPACKAGESUBJECTTYPE
        case "USER":
            result = USER_ACCESSPACKAGESUBJECTTYPE
        case "SERVICEPRINCIPAL":
            result = SERVICEPRINCIPAL_ACCESSPACKAGESUBJECTTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_ACCESSPACKAGESUBJECTTYPE
        default:
            return 0, errors.New("Unknown AccessPackageSubjectType value: " + v)
    }
    return &result, nil
}
func SerializeAccessPackageSubjectType(values []AccessPackageSubjectType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
