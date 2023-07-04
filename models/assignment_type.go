package models
import (
    "errors"
)
// 
type AssignmentType int

const (
    REQUIRED_ASSIGNMENTTYPE AssignmentType = iota
    RECOMMENDED_ASSIGNMENTTYPE
    UNKNOWNFUTUREVALUE_ASSIGNMENTTYPE
)

func (i AssignmentType) String() string {
    return []string{"required", "recommended", "unknownFutureValue"}[i]
}
func ParseAssignmentType(v string) (any, error) {
    result := REQUIRED_ASSIGNMENTTYPE
    switch v {
        case "required":
            result = REQUIRED_ASSIGNMENTTYPE
        case "recommended":
            result = RECOMMENDED_ASSIGNMENTTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ASSIGNMENTTYPE
        default:
            return 0, errors.New("Unknown AssignmentType value: " + v)
    }
    return &result, nil
}
func SerializeAssignmentType(values []AssignmentType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
