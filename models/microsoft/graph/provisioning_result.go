package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "SUCCESS":
            return SUCCESS_PROVISIONINGRESULT, nil
        case "FAILURE":
            return FAILURE_PROVISIONINGRESULT, nil
        case "SKIPPED":
            return SKIPPED_PROVISIONINGRESULT, nil
        case "WARNING":
            return WARNING_PROVISIONINGRESULT, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_PROVISIONINGRESULT, nil
    }
    return 0, errors.New("Unknown ProvisioningResult value: " + v)
}
func SerializeProvisioningResult(values []ProvisioningResult) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
