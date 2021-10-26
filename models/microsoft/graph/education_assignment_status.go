package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "DRAFT":
            return DRAFT_EDUCATIONASSIGNMENTSTATUS, nil
        case "PUBLISHED":
            return PUBLISHED_EDUCATIONASSIGNMENTSTATUS, nil
        case "ASSIGNED":
            return ASSIGNED_EDUCATIONASSIGNMENTSTATUS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_EDUCATIONASSIGNMENTSTATUS, nil
    }
    return 0, errors.New("Unknown EducationAssignmentStatus value: " + v)
}
func SerializeEducationAssignmentStatus(values []EducationAssignmentStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
