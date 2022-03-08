package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the drive singleton.
type ComplianceState int

const (
    UNKNOWN_COMPLIANCESTATE ComplianceState = iota
    COMPLIANT_COMPLIANCESTATE
    NONCOMPLIANT_COMPLIANCESTATE
    CONFLICT_COMPLIANCESTATE
    ERROR_COMPLIANCESTATE
    INGRACEPERIOD_COMPLIANCESTATE
    CONFIGMANAGER_COMPLIANCESTATE
)

func (i ComplianceState) String() string {
    return []string{"UNKNOWN", "COMPLIANT", "NONCOMPLIANT", "CONFLICT", "ERROR", "INGRACEPERIOD", "CONFIGMANAGER"}[i]
}
func ParseComplianceState(v string) (interface{}, error) {
    result := UNKNOWN_COMPLIANCESTATE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_COMPLIANCESTATE
        case "COMPLIANT":
            result = COMPLIANT_COMPLIANCESTATE
        case "NONCOMPLIANT":
            result = NONCOMPLIANT_COMPLIANCESTATE
        case "CONFLICT":
            result = CONFLICT_COMPLIANCESTATE
        case "ERROR":
            result = ERROR_COMPLIANCESTATE
        case "INGRACEPERIOD":
            result = INGRACEPERIOD_COMPLIANCESTATE
        case "CONFIGMANAGER":
            result = CONFIGMANAGER_COMPLIANCESTATE
        default:
            return 0, errors.New("Unknown ComplianceState value: " + v)
    }
    return &result, nil
}
func SerializeComplianceState(values []ComplianceState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
