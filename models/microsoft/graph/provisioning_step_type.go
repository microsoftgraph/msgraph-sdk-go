package graph
import (
    "strings"
    "errors"
)
type ProvisioningStepType int

const (
    IMPORT_PROVISIONINGSTEPTYPE ProvisioningStepType = iota
    SCOPING_PROVISIONINGSTEPTYPE
    MATCHING_PROVISIONINGSTEPTYPE
    PROCESSING_PROVISIONINGSTEPTYPE
    REFERENCERESOLUTION_PROVISIONINGSTEPTYPE
    EXPORT_PROVISIONINGSTEPTYPE
    UNKNOWNFUTUREVALUE_PROVISIONINGSTEPTYPE
)

func (i ProvisioningStepType) String() string {
    return []string{"IMPORT", "SCOPING", "MATCHING", "PROCESSING", "REFERENCERESOLUTION", "EXPORT", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseProvisioningStepType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "IMPORT":
            return IMPORT_PROVISIONINGSTEPTYPE, nil
        case "SCOPING":
            return SCOPING_PROVISIONINGSTEPTYPE, nil
        case "MATCHING":
            return MATCHING_PROVISIONINGSTEPTYPE, nil
        case "PROCESSING":
            return PROCESSING_PROVISIONINGSTEPTYPE, nil
        case "REFERENCERESOLUTION":
            return REFERENCERESOLUTION_PROVISIONINGSTEPTYPE, nil
        case "EXPORT":
            return EXPORT_PROVISIONINGSTEPTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_PROVISIONINGSTEPTYPE, nil
    }
    return 0, errors.New("Unknown ProvisioningStepType value: " + v)
}
func SerializeProvisioningStepType(values []ProvisioningStepType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
