package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the privacy singleton.
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
    result := CUSTOMER_DATASUBJECTTYPE
    switch strings.ToUpper(v) {
        case "CUSTOMER":
            result = CUSTOMER_DATASUBJECTTYPE
        case "CURRENTEMPLOYEE":
            result = CURRENTEMPLOYEE_DATASUBJECTTYPE
        case "FORMEREMPLOYEE":
            result = FORMEREMPLOYEE_DATASUBJECTTYPE
        case "PROSPECTIVEEMPLOYEE":
            result = PROSPECTIVEEMPLOYEE_DATASUBJECTTYPE
        case "STUDENT":
            result = STUDENT_DATASUBJECTTYPE
        case "TEACHER":
            result = TEACHER_DATASUBJECTTYPE
        case "FACULTY":
            result = FACULTY_DATASUBJECTTYPE
        case "OTHER":
            result = OTHER_DATASUBJECTTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_DATASUBJECTTYPE
        default:
            return 0, errors.New("Unknown DataSubjectType value: " + v)
    }
    return &result, nil
}
func SerializeDataSubjectType(values []DataSubjectType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
