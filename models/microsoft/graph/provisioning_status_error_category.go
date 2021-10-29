package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "FAILURE":
            return FAILURE_PROVISIONINGSTATUSERRORCATEGORY, nil
        case "NONSERVICEFAILURE":
            return NONSERVICEFAILURE_PROVISIONINGSTATUSERRORCATEGORY, nil
        case "SUCCESS":
            return SUCCESS_PROVISIONINGSTATUSERRORCATEGORY, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_PROVISIONINGSTATUSERRORCATEGORY, nil
    }
    return 0, errors.New("Unknown ProvisioningStatusErrorCategory value: " + v)
}
func SerializeProvisioningStatusErrorCategory(values []ProvisioningStatusErrorCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
