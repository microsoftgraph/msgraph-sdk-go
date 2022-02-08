package graph
import (
    "strings"
    "errors"
)
// 
type EducationUserRole int

const (
    STUDENT_EDUCATIONUSERROLE EducationUserRole = iota
    TEACHER_EDUCATIONUSERROLE
    NONE_EDUCATIONUSERROLE
    UNKNOWNFUTUREVALUE_EDUCATIONUSERROLE
)

func (i EducationUserRole) String() string {
    return []string{"STUDENT", "TEACHER", "NONE", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseEducationUserRole(v string) (interface{}, error) {
    result := STUDENT_EDUCATIONUSERROLE
    switch strings.ToUpper(v) {
        case "STUDENT":
            result = STUDENT_EDUCATIONUSERROLE
        case "TEACHER":
            result = TEACHER_EDUCATIONUSERROLE
        case "NONE":
            result = NONE_EDUCATIONUSERROLE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_EDUCATIONUSERROLE
        default:
            return 0, errors.New("Unknown EducationUserRole value: " + v)
    }
    return &result, nil
}
func SerializeEducationUserRole(values []EducationUserRole) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
