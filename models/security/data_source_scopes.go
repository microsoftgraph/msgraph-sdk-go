package security
import (
    "errors"
    "strings"
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
    var values []string
    for p := DataSourceScopes(1); p <= UNKNOWNFUTUREVALUE_DATASOURCESCOPES; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"none", "allTenantMailboxes", "allTenantSites", "allCaseCustodians", "allCaseNoncustodialDataSources", "unknownFutureValue"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseDataSourceScopes(v string) (any, error) {
    var result DataSourceScopes
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "none":
                result |= NONE_DATASOURCESCOPES
            case "allTenantMailboxes":
                result |= ALLTENANTMAILBOXES_DATASOURCESCOPES
            case "allTenantSites":
                result |= ALLTENANTSITES_DATASOURCESCOPES
            case "allCaseCustodians":
                result |= ALLCASECUSTODIANS_DATASOURCESCOPES
            case "allCaseNoncustodialDataSources":
                result |= ALLCASENONCUSTODIALDATASOURCES_DATASOURCESCOPES
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_DATASOURCESCOPES
            default:
                return 0, errors.New("Unknown DataSourceScopes value: " + v)
        }
    }
    return &result, nil
}
func SerializeDataSourceScopes(values []DataSourceScopes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DataSourceScopes) isMultiValue() bool {
    return true
}
