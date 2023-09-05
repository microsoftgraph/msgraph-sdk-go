package models
import (
    "errors"
    "strings"
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
    var values []string
    for p := WorkforceIntegrationSupportedEntities(1); p <= UNKNOWNFUTUREVALUE_WORKFORCEINTEGRATIONSUPPORTEDENTITIES; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"none", "shift", "swapRequest", "userShiftPreferences", "openShift", "openShiftRequest", "offerShiftRequest", "unknownFutureValue"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseWorkforceIntegrationSupportedEntities(v string) (any, error) {
    var result WorkforceIntegrationSupportedEntities
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "none":
                result |= NONE_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
            case "shift":
                result |= SHIFT_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
            case "swapRequest":
                result |= SWAPREQUEST_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
            case "userShiftPreferences":
                result |= USERSHIFTPREFERENCES_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
            case "openShift":
                result |= OPENSHIFT_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
            case "openShiftRequest":
                result |= OPENSHIFTREQUEST_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
            case "offerShiftRequest":
                result |= OFFERSHIFTREQUEST_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_WORKFORCEINTEGRATIONSUPPORTEDENTITIES
            default:
                return 0, errors.New("Unknown WorkforceIntegrationSupportedEntities value: " + v)
        }
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
func (i WorkforceIntegrationSupportedEntities) isMultiValue() bool {
    return true
}
