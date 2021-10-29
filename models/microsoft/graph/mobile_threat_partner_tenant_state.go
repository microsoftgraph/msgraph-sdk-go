package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "UNAVAILABLE":
            return UNAVAILABLE_MOBILETHREATPARTNERTENANTSTATE, nil
        case "AVAILABLE":
            return AVAILABLE_MOBILETHREATPARTNERTENANTSTATE, nil
        case "ENABLED":
            return ENABLED_MOBILETHREATPARTNERTENANTSTATE, nil
        case "UNRESPONSIVE":
            return UNRESPONSIVE_MOBILETHREATPARTNERTENANTSTATE, nil
    }
    return 0, errors.New("Unknown MobileThreatPartnerTenantState value: " + v)
}
func SerializeMobileThreatPartnerTenantState(values []MobileThreatPartnerTenantState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
