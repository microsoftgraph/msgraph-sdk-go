package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the deviceManagement singleton.
type MobileThreatPartnerTenantState int

const (
    UNAVAILABLE_MOBILETHREATPARTNERTENANTSTATE MobileThreatPartnerTenantState = iota
    AVAILABLE_MOBILETHREATPARTNERTENANTSTATE
    ENABLED_MOBILETHREATPARTNERTENANTSTATE
    UNRESPONSIVE_MOBILETHREATPARTNERTENANTSTATE
)

func (i MobileThreatPartnerTenantState) String() string {
    return []string{"UNAVAILABLE", "AVAILABLE", "ENABLED", "UNRESPONSIVE"}[i]
}
func ParseMobileThreatPartnerTenantState(v string) (interface{}, error) {
    result := UNAVAILABLE_MOBILETHREATPARTNERTENANTSTATE
    switch strings.ToUpper(v) {
        case "UNAVAILABLE":
            result = UNAVAILABLE_MOBILETHREATPARTNERTENANTSTATE
        case "AVAILABLE":
            result = AVAILABLE_MOBILETHREATPARTNERTENANTSTATE
        case "ENABLED":
            result = ENABLED_MOBILETHREATPARTNERTENANTSTATE
        case "UNRESPONSIVE":
            result = UNRESPONSIVE_MOBILETHREATPARTNERTENANTSTATE
        default:
            return 0, errors.New("Unknown MobileThreatPartnerTenantState value: " + v)
    }
    return &result, nil
}
func SerializeMobileThreatPartnerTenantState(values []MobileThreatPartnerTenantState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
