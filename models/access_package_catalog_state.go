package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the identityGovernance singleton.
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
    result := UNPUBLISHED_ACCESSPACKAGECATALOGSTATE
    switch strings.ToUpper(v) {
        case "UNPUBLISHED":
            result = UNPUBLISHED_ACCESSPACKAGECATALOGSTATE
        case "PUBLISHED":
            result = PUBLISHED_ACCESSPACKAGECATALOGSTATE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_ACCESSPACKAGECATALOGSTATE
        default:
            return 0, errors.New("Unknown AccessPackageCatalogState value: " + v)
    }
    return &result, nil
}
func SerializeAccessPackageCatalogState(values []AccessPackageCatalogState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
