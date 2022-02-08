package graph
import (
    "strings"
    "errors"
)
// 
type SubjectRightsRequestStatus int

const (
    ACTIVE_SUBJECTRIGHTSREQUESTSTATUS SubjectRightsRequestStatus = iota
    CLOSED_SUBJECTRIGHTSREQUESTSTATUS
    UNKNOWNFUTUREVALUE_SUBJECTRIGHTSREQUESTSTATUS
)

func (i SubjectRightsRequestStatus) String() string {
    return []string{"ACTIVE", "CLOSED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseSubjectRightsRequestStatus(v string) (interface{}, error) {
    result := ACTIVE_SUBJECTRIGHTSREQUESTSTATUS
    switch strings.ToUpper(v) {
        case "ACTIVE":
            result = ACTIVE_SUBJECTRIGHTSREQUESTSTATUS
        case "CLOSED":
            result = CLOSED_SUBJECTRIGHTSREQUESTSTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_SUBJECTRIGHTSREQUESTSTATUS
        default:
            return 0, errors.New("Unknown SubjectRightsRequestStatus value: " + v)
    }
    return &result, nil
}
func SerializeSubjectRightsRequestStatus(values []SubjectRightsRequestStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
