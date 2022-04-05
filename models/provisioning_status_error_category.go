package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the auditLogRoot singleton.
type ProvisioningStatusErrorCategory int

const (
    FAILURE_PROVISIONINGSTATUSERRORCATEGORY ProvisioningStatusErrorCategory = iota
    NONSERVICEFAILURE_PROVISIONINGSTATUSERRORCATEGORY
    SUCCESS_PROVISIONINGSTATUSERRORCATEGORY
    UNKNOWNFUTUREVALUE_PROVISIONINGSTATUSERRORCATEGORY
)

func (i ProvisioningStatusErrorCategory) String() string {
    return []string{"FAILURE", "NONSERVICEFAILURE", "SUCCESS", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseProvisioningStatusErrorCategory(v string) (interface{}, error) {
    result := FAILURE_PROVISIONINGSTATUSERRORCATEGORY
    switch strings.ToUpper(v) {
        case "FAILURE":
            result = FAILURE_PROVISIONINGSTATUSERRORCATEGORY
        case "NONSERVICEFAILURE":
            result = NONSERVICEFAILURE_PROVISIONINGSTATUSERRORCATEGORY
        case "SUCCESS":
            result = SUCCESS_PROVISIONINGSTATUSERRORCATEGORY
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_PROVISIONINGSTATUSERRORCATEGORY
        default:
            return 0, errors.New("Unknown ProvisioningStatusErrorCategory value: " + v)
    }
    return &result, nil
}
func SerializeProvisioningStatusErrorCategory(values []ProvisioningStatusErrorCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
