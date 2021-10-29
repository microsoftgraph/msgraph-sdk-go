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
    switch strings.ToUpper(v) {
        case "STUDENT":
            return STUDENT_EDUCATIONUSERROLE, nil
        case "TEACHER":
            return TEACHER_EDUCATIONUSERROLE, nil
        case "NONE":
            return NONE_EDUCATIONUSERROLE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_EDUCATIONUSERROLE, nil
    }
    return 0, errors.New("Unknown EducationUserRole value: " + v)
}
func SerializeEducationUserRole(values []EducationUserRole) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
