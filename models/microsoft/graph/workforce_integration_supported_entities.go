package graph
import (
    "strings"
    "errors"
)
// 
type WorkforceIntegrationSupportedEntities int

const (
    NONE_WORKFORCEINTEGRATIONSUPPORTEDENTITIES WorkforceIntegrationSupportedEntities = iota
    SHIFT_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
    SWAPREQUEST_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
    USERSHIFTPREFERENCES_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
    OPENSHIFT_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
    OPENSHIFTREQUEST_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
    OFFERSHIFTREQUEST_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
    UNKNOWNFUTUREVALUE_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
)

func (i WorkforceIntegrationSupportedEntities) String() string {
    return []string{"NONE", "SHIFT", "SWAPREQUEST", "USERSHIFTPREFERENCES", "OPENSHIFT", "OPENSHIFTREQUEST", "OFFERSHIFTREQUEST", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseWorkforceIntegrationSupportedEntities(v string) (interface{}, error) {
    result := NONE_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
        case "SHIFT":
            result = SHIFT_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
        case "SWAPREQUEST":
            result = SWAPREQUEST_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
        case "USERSHIFTPREFERENCES":
            result = USERSHIFTPREFERENCES_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
        case "OPENSHIFT":
            result = OPENSHIFT_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
        case "OPENSHIFTREQUEST":
            result = OPENSHIFTREQUEST_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
        case "OFFERSHIFTREQUEST":
            result = OFFERSHIFTREQUEST_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
        default:
            return 0, errors.New("Unknown WorkforceIntegrationSupportedEntities value: " + v)
    }
    return &result, nil
}
func SerializeWorkforceIntegrationSupportedEntities(values []WorkforceIntegrationSupportedEntities) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
