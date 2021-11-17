package graph
import (
    "strings"
    "errors"
)
// 
type AccessPackageCatalogState int

const (
    UNPUBLISHED_ACCESSPACKAGECATALOGSTATE AccessPackageCatalogState = iota
    PUBLISHED_ACCESSPACKAGECATALOGSTATE
    UNKNOWNFUTUREVALUE_ACCESSPACKAGECATALOGSTATE
)

func (i AccessPackageCatalogState) String() string {
    return []string{"UNPUBLISHED", "PUBLISHED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseAccessPackageCatalogState(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "UNPUBLISHED":
            return UNPUBLISHED_ACCESSPACKAGECATALOGSTATE, nil
        case "PUBLISHED":
            return PUBLISHED_ACCESSPACKAGECATALOGSTATE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_ACCESSPACKAGECATALOGSTATE, nil
    }
    return 0, errors.New("Unknown AccessPackageCatalogState value: " + v)
}
func SerializeAccessPackageCatalogState(values []AccessPackageCatalogState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
