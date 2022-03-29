package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the deviceManagement singleton.
type ComplianceStatus int

const (
    UNKNOWN_COMPLIANCESTATUS ComplianceStatus = iota
    NOTAPPLICABLE_COMPLIANCESTATUS
    COMPLIANT_COMPLIANCESTATUS
    REMEDIATED_COMPLIANCESTATUS
    NONCOMPLIANT_COMPLIANCESTATUS
    ERROR_COMPLIANCESTATUS
    CONFLICT_COMPLIANCESTATUS
    NOTASSIGNED_COMPLIANCESTATUS
)

func (i ComplianceStatus) String() string {
    return []string{"UNKNOWN", "NOTAPPLICABLE", "COMPLIANT", "REMEDIATED", "NONCOMPLIANT", "ERROR", "CONFLICT", "NOTASSIGNED"}[i]
}
func ParseComplianceStatus(v string) (interface{}, error) {
    result := UNKNOWN_COMPLIANCESTATUS
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_COMPLIANCESTATUS
        case "NOTAPPLICABLE":
            result = NOTAPPLICABLE_COMPLIANCESTATUS
        case "COMPLIANT":
            result = COMPLIANT_COMPLIANCESTATUS
        case "REMEDIATED":
            result = REMEDIATED_COMPLIANCESTATUS
        case "NONCOMPLIANT":
            result = NONCOMPLIANT_COMPLIANCESTATUS
        case "ERROR":
            result = ERROR_COMPLIANCESTATUS
        case "CONFLICT":
            result = CONFLICT_COMPLIANCESTATUS
        case "NOTASSIGNED":
            result = NOTASSIGNED_COMPLIANCESTATUS
        default:
            return 0, errors.New("Unknown ComplianceStatus value: " + v)
    }
    return &result, nil
}
func SerializeComplianceStatus(values []ComplianceStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
