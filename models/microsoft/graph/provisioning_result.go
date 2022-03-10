package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the auditLogRoot singleton.
type ProvisioningResult int

const (
    SUCCESS_PROVISIONINGRESULT ProvisioningResult = iota
    FAILURE_PROVISIONINGRESULT
    SKIPPED_PROVISIONINGRESULT
    WARNING_PROVISIONINGRESULT
    UNKNOWNFUTUREVALUE_PROVISIONINGRESULT
)

func (i ProvisioningResult) String() string {
    return []string{"SUCCESS", "FAILURE", "SKIPPED", "WARNING", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseProvisioningResult(v string) (interface{}, error) {
    result := SUCCESS_PROVISIONINGRESULT
    switch strings.ToUpper(v) {
        case "SUCCESS":
            result = SUCCESS_PROVISIONINGRESULT
        case "FAILURE":
            result = FAILURE_PROVISIONINGRESULT
        case "SKIPPED":
            result = SKIPPED_PROVISIONINGRESULT
        case "WARNING":
            result = WARNING_PROVISIONINGRESULT
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_PROVISIONINGRESULT
        default:
            return 0, errors.New("Unknown ProvisioningResult value: " + v)
    }
    return &result, nil
}
func SerializeProvisioningResult(values []ProvisioningResult) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
