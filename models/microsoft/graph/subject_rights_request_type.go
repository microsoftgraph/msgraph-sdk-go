package graph
import (
    "strings"
    "errors"
)
// 
type SubjectRightsRequestType int

const (
    EXPORT_SUBJECTRIGHTSREQUESTTYPE SubjectRightsRequestType = iota
    DELETE_SUBJECTRIGHTSREQUESTTYPE
    ACCESS_SUBJECTRIGHTSREQUESTTYPE
    TAGFORACTION_SUBJECTRIGHTSREQUESTTYPE
    UNKNOWNFUTUREVALUE_SUBJECTRIGHTSREQUESTTYPE
)

func (i SubjectRightsRequestType) String() string {
    return []string{"EXPORT", "DELETE", "ACCESS", "TAGFORACTION", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseSubjectRightsRequestType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "EXPORT":
            return EXPORT_SUBJECTRIGHTSREQUESTTYPE, nil
        case "DELETE":
            return DELETE_SUBJECTRIGHTSREQUESTTYPE, nil
        case "ACCESS":
            return ACCESS_SUBJECTRIGHTSREQUESTTYPE, nil
        case "TAGFORACTION":
            return TAGFORACTION_SUBJECTRIGHTSREQUESTTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_SUBJECTRIGHTSREQUESTTYPE, nil
    }
    return 0, errors.New("Unknown SubjectRightsRequestType value: " + v)
}
func SerializeSubjectRightsRequestType(values []SubjectRightsRequestType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
