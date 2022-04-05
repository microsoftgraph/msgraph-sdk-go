package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the identityGovernance singleton.
type AccessPackageCatalogType int

const (
    USERMANAGED_ACCESSPACKAGECATALOGTYPE AccessPackageCatalogType = iota
    SERVICEDEFAULT_ACCESSPACKAGECATALOGTYPE
    SERVICEMANAGED_ACCESSPACKAGECATALOGTYPE
    UNKNOWNFUTUREVALUE_ACCESSPACKAGECATALOGTYPE
)

func (i AccessPackageCatalogType) String() string {
    return []string{"USERMANAGED", "SERVICEDEFAULT", "SERVICEMANAGED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseAccessPackageCatalogType(v string) (interface{}, error) {
    result := USERMANAGED_ACCESSPACKAGECATALOGTYPE
    switch strings.ToUpper(v) {
        case "USERMANAGED":
            result = USERMANAGED_ACCESSPACKAGECATALOGTYPE
        case "SERVICEDEFAULT":
            result = SERVICEDEFAULT_ACCESSPACKAGECATALOGTYPE
        case "SERVICEMANAGED":
            result = SERVICEMANAGED_ACCESSPACKAGECATALOGTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_ACCESSPACKAGECATALOGTYPE
        default:
            return 0, errors.New("Unknown AccessPackageCatalogType value: " + v)
    }
    return &result, nil
}
func SerializeAccessPackageCatalogType(values []AccessPackageCatalogType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
