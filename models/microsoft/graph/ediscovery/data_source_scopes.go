package ediscovery
import (
    "strings"
    "errors"
)
// 
type DataSourceScopes int

const (
    NONE_DATASOURCESCOPES DataSourceScopes = iota
    ALLTENANTMAILBOXES_DATASOURCESCOPES
    ALLTENANTSITES_DATASOURCESCOPES
    ALLCASECUSTODIANS_DATASOURCESCOPES
    ALLCASENONCUSTODIALDATASOURCES_DATASOURCESCOPES
    UNKNOWNFUTUREVALUE_DATASOURCESCOPES
)

func (i DataSourceScopes) String() string {
    return []string{"NONE", "ALLTENANTMAILBOXES", "ALLTENANTSITES", "ALLCASECUSTODIANS", "ALLCASENONCUSTODIALDATASOURCES", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseDataSourceScopes(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_DATASOURCESCOPES, nil
        case "ALLTENANTMAILBOXES":
            return ALLTENANTMAILBOXES_DATASOURCESCOPES, nil
        case "ALLTENANTSITES":
            return ALLTENANTSITES_DATASOURCESCOPES, nil
        case "ALLCASECUSTODIANS":
            return ALLCASECUSTODIANS_DATASOURCESCOPES, nil
        case "ALLCASENONCUSTODIALDATASOURCES":
            return ALLCASENONCUSTODIALDATASOURCES_DATASOURCESCOPES, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_DATASOURCESCOPES, nil
    }
    return 0, errors.New("Unknown DataSourceScopes value: " + v)
}
func SerializeDataSourceScopes(values []DataSourceScopes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
