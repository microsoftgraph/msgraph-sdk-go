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
    switch strings.ToUpper(v) {
        case "ACTIVE":
            return ACTIVE_SUBJECTRIGHTSREQUESTSTATUS, nil
        case "CLOSED":
            return CLOSED_SUBJECTRIGHTSREQUESTSTATUS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_SUBJECTRIGHTSREQUESTSTATUS, nil
    }
    return 0, errors.New("Unknown SubjectRightsRequestStatus value: " + v)
}
func SerializeSubjectRightsRequestStatus(values []SubjectRightsRequestStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
