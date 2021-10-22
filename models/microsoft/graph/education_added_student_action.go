package graph
import (
    "strings"
    "errors"
)
type EducationAddedStudentAction int

const (
    NONE_EDUCATIONADDEDSTUDENTACTION EducationAddedStudentAction = iota
    ASSIGNIFOPEN_EDUCATIONADDEDSTUDENTACTION
    UNKNOWNFUTUREVALUE_EDUCATIONADDEDSTUDENTACTION
)

func (i EducationAddedStudentAction) String() string {
    return []string{"NONE", "ASSIGNIFOPEN", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseEducationAddedStudentAction(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_EDUCATIONADDEDSTUDENTACTION, nil
        case "ASSIGNIFOPEN":
            return ASSIGNIFOPEN_EDUCATIONADDEDSTUDENTACTION, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_EDUCATIONADDEDSTUDENTACTION, nil
    }
    return 0, errors.New("Unknown EducationAddedStudentAction value: " + v)
}
func SerializeEducationAddedStudentAction(values []EducationAddedStudentAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
