package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_COMPLIANCESTATUS, nil
        case "NOTAPPLICABLE":
            return NOTAPPLICABLE_COMPLIANCESTATUS, nil
        case "COMPLIANT":
            return COMPLIANT_COMPLIANCESTATUS, nil
        case "REMEDIATED":
            return REMEDIATED_COMPLIANCESTATUS, nil
        case "NONCOMPLIANT":
            return NONCOMPLIANT_COMPLIANCESTATUS, nil
        case "ERROR":
            return ERROR_COMPLIANCESTATUS, nil
        case "CONFLICT":
            return CONFLICT_COMPLIANCESTATUS, nil
        case "NOTASSIGNED":
            return NOTASSIGNED_COMPLIANCESTATUS, nil
    }
    return 0, errors.New("Unknown ComplianceStatus value: " + v)
}
func SerializeComplianceStatus(values []ComplianceStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
