package graph
import (
    "strings"
    "errors"
)
// 
type DataSubjectType int

const (
    CUSTOMER_DATASUBJECTTYPE DataSubjectType = iota
    CURRENTEMPLOYEE_DATASUBJECTTYPE
    FORMEREMPLOYEE_DATASUBJECTTYPE
    PROSPECTIVEEMPLOYEE_DATASUBJECTTYPE
    STUDENT_DATASUBJECTTYPE
    TEACHER_DATASUBJECTTYPE
    FACULTY_DATASUBJECTTYPE
    OTHER_DATASUBJECTTYPE
    UNKNOWNFUTUREVALUE_DATASUBJECTTYPE
)

func (i DataSubjectType) String() string {
    return []string{"CUSTOMER", "CURRENTEMPLOYEE", "FORMEREMPLOYEE", "PROSPECTIVEEMPLOYEE", "STUDENT", "TEACHER", "FACULTY", "OTHER", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseDataSubjectType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "CUSTOMER":
            return CUSTOMER_DATASUBJECTTYPE, nil
        case "CURRENTEMPLOYEE":
            return CURRENTEMPLOYEE_DATASUBJECTTYPE, nil
        case "FORMEREMPLOYEE":
            return FORMEREMPLOYEE_DATASUBJECTTYPE, nil
        case "PROSPECTIVEEMPLOYEE":
            return PROSPECTIVEEMPLOYEE_DATASUBJECTTYPE, nil
        case "STUDENT":
            return STUDENT_DATASUBJECTTYPE, nil
        case "TEACHER":
            return TEACHER_DATASUBJECTTYPE, nil
        case "FACULTY":
            return FACULTY_DATASUBJECTTYPE, nil
        case "OTHER":
            return OTHER_DATASUBJECTTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_DATASUBJECTTYPE, nil
    }
    return 0, errors.New("Unknown DataSubjectType value: " + v)
}
func SerializeDataSubjectType(values []DataSubjectType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
