package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the educationRoot singleton.
type EducationAssignmentStatus int

const (
    DRAFT_EDUCATIONASSIGNMENTSTATUS EducationAssignmentStatus = iota
    PUBLISHED_EDUCATIONASSIGNMENTSTATUS
    ASSIGNED_EDUCATIONASSIGNMENTSTATUS
    UNKNOWNFUTUREVALUE_EDUCATIONASSIGNMENTSTATUS
)

func (i EducationAssignmentStatus) String() string {
    return []string{"DRAFT", "PUBLISHED", "ASSIGNED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseEducationAssignmentStatus(v string) (interface{}, error) {
    result := DRAFT_EDUCATIONASSIGNMENTSTATUS
    switch strings.ToUpper(v) {
        case "DRAFT":
            result = DRAFT_EDUCATIONASSIGNMENTSTATUS
        case "PUBLISHED":
            result = PUBLISHED_EDUCATIONASSIGNMENTSTATUS
        case "ASSIGNED":
            result = ASSIGNED_EDUCATIONASSIGNMENTSTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_EDUCATIONASSIGNMENTSTATUS
        default:
            return 0, errors.New("Unknown EducationAssignmentStatus value: " + v)
    }
    return &result, nil
}
func SerializeEducationAssignmentStatus(values []EducationAssignmentStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
