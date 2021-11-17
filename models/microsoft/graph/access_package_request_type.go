package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "NOTSPECIFIED":
            return NOTSPECIFIED_ACCESSPACKAGEREQUESTTYPE, nil
        case "USERADD":
            return USERADD_ACCESSPACKAGEREQUESTTYPE, nil
        case "USERUPDATE":
            return USERUPDATE_ACCESSPACKAGEREQUESTTYPE, nil
        case "USERREMOVE":
            return USERREMOVE_ACCESSPACKAGEREQUESTTYPE, nil
        case "ADMINADD":
            return ADMINADD_ACCESSPACKAGEREQUESTTYPE, nil
        case "ADMINUPDATE":
            return ADMINUPDATE_ACCESSPACKAGEREQUESTTYPE, nil
        case "ADMINREMOVE":
            return ADMINREMOVE_ACCESSPACKAGEREQUESTTYPE, nil
        case "SYSTEMADD":
            return SYSTEMADD_ACCESSPACKAGEREQUESTTYPE, nil
        case "SYSTEMUPDATE":
            return SYSTEMUPDATE_ACCESSPACKAGEREQUESTTYPE, nil
        case "SYSTEMREMOVE":
            return SYSTEMREMOVE_ACCESSPACKAGEREQUESTTYPE, nil
        case "ONBEHALFADD":
            return ONBEHALFADD_ACCESSPACKAGEREQUESTTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_ACCESSPACKAGEREQUESTTYPE, nil
    }
    return 0, errors.New("Unknown AccessPackageRequestType value: " + v)
}
func SerializeAccessPackageRequestType(values []AccessPackageRequestType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
