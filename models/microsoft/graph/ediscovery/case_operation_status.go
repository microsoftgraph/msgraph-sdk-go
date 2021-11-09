package ediscovery
import (
    "strings"
    "errors"
)
// 
type CaseOperationStatus int

const (
    NOTSTARTED_CASEOPERATIONSTATUS CaseOperationStatus = iota
    SUBMISSIONFAILED_CASEOPERATIONSTATUS
    RUNNING_CASEOPERATIONSTATUS
    SUCCEEDED_CASEOPERATIONSTATUS
    PARTIALLYSUCCEEDED_CASEOPERATIONSTATUS
    FAILED_CASEOPERATIONSTATUS
)

func (i CaseOperationStatus) String() string {
    return []string{"NOTSTARTED", "SUBMISSIONFAILED", "RUNNING", "SUCCEEDED", "PARTIALLYSUCCEEDED", "FAILED"}[i]
}
func ParseCaseOperationStatus(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "NOTSTARTED":
            return NOTSTARTED_CASEOPERATIONSTATUS, nil
        case "SUBMISSIONFAILED":
            return SUBMISSIONFAILED_CASEOPERATIONSTATUS, nil
        case "RUNNING":
            return RUNNING_CASEOPERATIONSTATUS, nil
        case "SUCCEEDED":
            return SUCCEEDED_CASEOPERATIONSTATUS, nil
        case "PARTIALLYSUCCEEDED":
            return PARTIALLYSUCCEEDED_CASEOPERATIONSTATUS, nil
        case "FAILED":
            return FAILED_CASEOPERATIONSTATUS, nil
    }
    return 0, errors.New("Unknown CaseOperationStatus value: " + v)
}
func SerializeCaseOperationStatus(values []CaseOperationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
