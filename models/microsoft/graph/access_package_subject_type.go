package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "NOTSPECIFIED":
            return NOTSPECIFIED_ACCESSPACKAGESUBJECTTYPE, nil
        case "USER":
            return USER_ACCESSPACKAGESUBJECTTYPE, nil
        case "SERVICEPRINCIPAL":
            return SERVICEPRINCIPAL_ACCESSPACKAGESUBJECTTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_ACCESSPACKAGESUBJECTTYPE, nil
    }
    return 0, errors.New("Unknown AccessPackageSubjectType value: " + v)
}
func SerializeAccessPackageSubjectType(values []AccessPackageSubjectType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
