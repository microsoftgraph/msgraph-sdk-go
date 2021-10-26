package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_COMPLIANCESTATE, nil
        case "COMPLIANT":
            return COMPLIANT_COMPLIANCESTATE, nil
        case "NONCOMPLIANT":
            return NONCOMPLIANT_COMPLIANCESTATE, nil
        case "CONFLICT":
            return CONFLICT_COMPLIANCESTATE, nil
        case "ERROR":
            return ERROR_COMPLIANCESTATE, nil
        case "INGRACEPERIOD":
            return INGRACEPERIOD_COMPLIANCESTATE, nil
        case "CONFIGMANAGER":
            return CONFIGMANAGER_COMPLIANCESTATE, nil
    }
    return 0, errors.New("Unknown ComplianceState value: " + v)
}
func SerializeComplianceState(values []ComplianceState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
