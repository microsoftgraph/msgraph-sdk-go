package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the identityGovernance singleton.
type AccessPackageRequestType int

const (
    NOTSPECIFIED_ACCESSPACKAGEREQUESTTYPE AccessPackageRequestType = iota
    USERADD_ACCESSPACKAGEREQUESTTYPE
    USERUPDATE_ACCESSPACKAGEREQUESTTYPE
    USERREMOVE_ACCESSPACKAGEREQUESTTYPE
    ADMINADD_ACCESSPACKAGEREQUESTTYPE
    ADMINUPDATE_ACCESSPACKAGEREQUESTTYPE
    ADMINREMOVE_ACCESSPACKAGEREQUESTTYPE
    SYSTEMADD_ACCESSPACKAGEREQUESTTYPE
    SYSTEMUPDATE_ACCESSPACKAGEREQUESTTYPE
    SYSTEMREMOVE_ACCESSPACKAGEREQUESTTYPE
    ONBEHALFADD_ACCESSPACKAGEREQUESTTYPE
    UNKNOWNFUTUREVALUE_ACCESSPACKAGEREQUESTTYPE
)

func (i AccessPackageRequestType) String() string {
    return []string{"NOTSPECIFIED", "USERADD", "USERUPDATE", "USERREMOVE", "ADMINADD", "ADMINUPDATE", "ADMINREMOVE", "SYSTEMADD", "SYSTEMUPDATE", "SYSTEMREMOVE", "ONBEHALFADD", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseAccessPackageRequestType(v string) (interface{}, error) {
    result := NOTSPECIFIED_ACCESSPACKAGEREQUESTTYPE
    switch strings.ToUpper(v) {
        case "NOTSPECIFIED":
            result = NOTSPECIFIED_ACCESSPACKAGEREQUESTTYPE
        case "USERADD":
            result = USERADD_ACCESSPACKAGEREQUESTTYPE
        case "USERUPDATE":
            result = USERUPDATE_ACCESSPACKAGEREQUESTTYPE
        case "USERREMOVE":
            result = USERREMOVE_ACCESSPACKAGEREQUESTTYPE
        case "ADMINADD":
            result = ADMINADD_ACCESSPACKAGEREQUESTTYPE
        case "ADMINUPDATE":
            result = ADMINUPDATE_ACCESSPACKAGEREQUESTTYPE
        case "ADMINREMOVE":
            result = ADMINREMOVE_ACCESSPACKAGEREQUESTTYPE
        case "SYSTEMADD":
            result = SYSTEMADD_ACCESSPACKAGEREQUESTTYPE
        case "SYSTEMUPDATE":
            result = SYSTEMUPDATE_ACCESSPACKAGEREQUESTTYPE
        case "SYSTEMREMOVE":
            result = SYSTEMREMOVE_ACCESSPACKAGEREQUESTTYPE
        case "ONBEHALFADD":
            result = ONBEHALFADD_ACCESSPACKAGEREQUESTTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_ACCESSPACKAGEREQUESTTYPE
        default:
            return 0, errors.New("Unknown AccessPackageRequestType value: " + v)
    }
    return &result, nil
}
func SerializeAccessPackageRequestType(values []AccessPackageRequestType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
