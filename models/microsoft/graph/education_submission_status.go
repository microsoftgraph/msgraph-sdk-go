package graph
import (
    "strings"
    "errors"
)
// 
type EducationSubmissionStatus int

const (
    WORKING_EDUCATIONSUBMISSIONSTATUS EducationSubmissionStatus = iota
    SUBMITTED_EDUCATIONSUBMISSIONSTATUS
    RELEASED_EDUCATIONSUBMISSIONSTATUS
    RETURNED_EDUCATIONSUBMISSIONSTATUS
    UNKNOWNFUTUREVALUE_EDUCATIONSUBMISSIONSTATUS
    REASSIGNED_EDUCATIONSUBMISSIONSTATUS
)

func (i EducationSubmissionStatus) String() string {
    return []string{"WORKING", "SUBMITTED", "RELEASED", "RETURNED", "UNKNOWNFUTUREVALUE", "REASSIGNED"}[i]
}
func ParseEducationSubmissionStatus(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "WORKING":
            return WORKING_EDUCATIONSUBMISSIONSTATUS, nil
        case "SUBMITTED":
            return SUBMITTED_EDUCATIONSUBMISSIONSTATUS, nil
        case "RELEASED":
            return RELEASED_EDUCATIONSUBMISSIONSTATUS, nil
        case "RETURNED":
            return RETURNED_EDUCATIONSUBMISSIONSTATUS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_EDUCATIONSUBMISSIONSTATUS, nil
        case "REASSIGNED":
            return REASSIGNED_EDUCATIONSUBMISSIONSTATUS, nil
    }
    return 0, errors.New("Unknown EducationSubmissionStatus value: " + v)
}
func SerializeEducationSubmissionStatus(values []EducationSubmissionStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
