package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the educationRoot singleton.
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
    result := WORKING_EDUCATIONSUBMISSIONSTATUS
    switch strings.ToUpper(v) {
        case "WORKING":
            result = WORKING_EDUCATIONSUBMISSIONSTATUS
        case "SUBMITTED":
            result = SUBMITTED_EDUCATIONSUBMISSIONSTATUS
        case "RELEASED":
            result = RELEASED_EDUCATIONSUBMISSIONSTATUS
        case "RETURNED":
            result = RETURNED_EDUCATIONSUBMISSIONSTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_EDUCATIONSUBMISSIONSTATUS
        case "REASSIGNED":
            result = REASSIGNED_EDUCATIONSUBMISSIONSTATUS
        default:
            return 0, errors.New("Unknown EducationSubmissionStatus value: " + v)
    }
    return &result, nil
}
func SerializeEducationSubmissionStatus(values []EducationSubmissionStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
