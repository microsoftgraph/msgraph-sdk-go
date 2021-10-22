package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "SUCCESS":
            return SUCCESS_APPLIEDCONDITIONALACCESSPOLICYRESULT, nil
        case "FAILURE":
            return FAILURE_APPLIEDCONDITIONALACCESSPOLICYRESULT, nil
        case "NOTAPPLIED":
            return NOTAPPLIED_APPLIEDCONDITIONALACCESSPOLICYRESULT, nil
        case "NOTENABLED":
            return NOTENABLED_APPLIEDCONDITIONALACCESSPOLICYRESULT, nil
        case "UNKNOWN":
            return UNKNOWN_APPLIEDCONDITIONALACCESSPOLICYRESULT, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_APPLIEDCONDITIONALACCESSPOLICYRESULT, nil
    }
    return 0, errors.New("Unknown AppliedConditionalAccessPolicyResult value: " + v)
}
func SerializeAppliedConditionalAccessPolicyResult(values []AppliedConditionalAccessPolicyResult) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
