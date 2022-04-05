package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the auditLogRoot singleton.
type AppliedConditionalAccessPolicyResult int

const (
    SUCCESS_APPLIEDCONDITIONALACCESSPOLICYRESULT AppliedConditionalAccessPolicyResult = iota
    FAILURE_APPLIEDCONDITIONALACCESSPOLICYRESULT
    NOTAPPLIED_APPLIEDCONDITIONALACCESSPOLICYRESULT
    NOTENABLED_APPLIEDCONDITIONALACCESSPOLICYRESULT
    UNKNOWN_APPLIEDCONDITIONALACCESSPOLICYRESULT
    UNKNOWNFUTUREVALUE_APPLIEDCONDITIONALACCESSPOLICYRESULT
)

func (i AppliedConditionalAccessPolicyResult) String() string {
    return []string{"SUCCESS", "FAILURE", "NOTAPPLIED", "NOTENABLED", "UNKNOWN", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseAppliedConditionalAccessPolicyResult(v string) (interface{}, error) {
    result := SUCCESS_APPLIEDCONDITIONALACCESSPOLICYRESULT
    switch strings.ToUpper(v) {
        case "SUCCESS":
            result = SUCCESS_APPLIEDCONDITIONALACCESSPOLICYRESULT
        case "FAILURE":
            result = FAILURE_APPLIEDCONDITIONALACCESSPOLICYRESULT
        case "NOTAPPLIED":
            result = NOTAPPLIED_APPLIEDCONDITIONALACCESSPOLICYRESULT
        case "NOTENABLED":
            result = NOTENABLED_APPLIEDCONDITIONALACCESSPOLICYRESULT
        case "UNKNOWN":
            result = UNKNOWN_APPLIEDCONDITIONALACCESSPOLICYRESULT
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_APPLIEDCONDITIONALACCESSPOLICYRESULT
        default:
            return 0, errors.New("Unknown AppliedConditionalAccessPolicyResult value: " + v)
    }
    return &result, nil
}
func SerializeAppliedConditionalAccessPolicyResult(values []AppliedConditionalAccessPolicyResult) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
