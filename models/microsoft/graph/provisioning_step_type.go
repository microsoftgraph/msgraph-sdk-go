package graph
import (
    "strings"
    "errors"
)
// 
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
    result := IMPORT_PROVISIONINGSTEPTYPE
    switch strings.ToUpper(v) {
        case "IMPORT":
            result = IMPORT_PROVISIONINGSTEPTYPE
        case "SCOPING":
            result = SCOPING_PROVISIONINGSTEPTYPE
        case "MATCHING":
            result = MATCHING_PROVISIONINGSTEPTYPE
        case "PROCESSING":
            result = PROCESSING_PROVISIONINGSTEPTYPE
        case "REFERENCERESOLUTION":
            result = REFERENCERESOLUTION_PROVISIONINGSTEPTYPE
        case "EXPORT":
            result = EXPORT_PROVISIONINGSTEPTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_PROVISIONINGSTEPTYPE
        default:
            return 0, errors.New("Unknown ProvisioningStepType value: " + v)
    }
    return &result, nil
}
func SerializeProvisioningStepType(values []ProvisioningStepType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
