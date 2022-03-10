package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the privacy singleton.
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
    result := EXPORT_SUBJECTRIGHTSREQUESTTYPE
    switch strings.ToUpper(v) {
        case "EXPORT":
            result = EXPORT_SUBJECTRIGHTSREQUESTTYPE
        case "DELETE":
            result = DELETE_SUBJECTRIGHTSREQUESTTYPE
        case "ACCESS":
            result = ACCESS_SUBJECTRIGHTSREQUESTTYPE
        case "TAGFORACTION":
            result = TAGFORACTION_SUBJECTRIGHTSREQUESTTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_SUBJECTRIGHTSREQUESTTYPE
        default:
            return 0, errors.New("Unknown SubjectRightsRequestType value: " + v)
    }
    return &result, nil
}
func SerializeSubjectRightsRequestType(values []SubjectRightsRequestType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
