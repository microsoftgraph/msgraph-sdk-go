package graph
import (
    "strings"
    "errors"
)
// 
type SubjectRightsRequestStageStatus int

const (
    NOTSTARTED_SUBJECTRIGHTSREQUESTSTAGESTATUS SubjectRightsRequestStageStatus = iota
    CURRENT_SUBJECTRIGHTSREQUESTSTAGESTATUS
    COMPLETED_SUBJECTRIGHTSREQUESTSTAGESTATUS
    FAILED_SUBJECTRIGHTSREQUESTSTAGESTATUS
    UNKNOWNFUTUREVALUE_SUBJECTRIGHTSREQUESTSTAGESTATUS
)

func (i SubjectRightsRequestStageStatus) String() string {
    return []string{"NOTSTARTED", "CURRENT", "COMPLETED", "FAILED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseSubjectRightsRequestStageStatus(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "NOTSTARTED":
            return NOTSTARTED_SUBJECTRIGHTSREQUESTSTAGESTATUS, nil
        case "CURRENT":
            return CURRENT_SUBJECTRIGHTSREQUESTSTAGESTATUS, nil
        case "COMPLETED":
            return COMPLETED_SUBJECTRIGHTSREQUESTSTAGESTATUS, nil
        case "FAILED":
            return FAILED_SUBJECTRIGHTSREQUESTSTAGESTATUS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_SUBJECTRIGHTSREQUESTSTAGESTATUS, nil
    }
    return 0, errors.New("Unknown SubjectRightsRequestStageStatus value: " + v)
}
func SerializeSubjectRightsRequestStageStatus(values []SubjectRightsRequestStageStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
