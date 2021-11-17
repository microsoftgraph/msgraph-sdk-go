package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "USERMANAGED":
            return USERMANAGED_ACCESSPACKAGECATALOGTYPE, nil
        case "SERVICEDEFAULT":
            return SERVICEDEFAULT_ACCESSPACKAGECATALOGTYPE, nil
        case "SERVICEMANAGED":
            return SERVICEMANAGED_ACCESSPACKAGECATALOGTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_ACCESSPACKAGECATALOGTYPE, nil
    }
    return 0, errors.New("Unknown AccessPackageCatalogType value: " + v)
}
func SerializeAccessPackageCatalogType(values []AccessPackageCatalogType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
