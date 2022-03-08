package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the privacy singleton.
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
    result := NOTSTARTED_SUBJECTRIGHTSREQUESTSTAGESTATUS
    switch strings.ToUpper(v) {
        case "NOTSTARTED":
            result = NOTSTARTED_SUBJECTRIGHTSREQUESTSTAGESTATUS
        case "CURRENT":
            result = CURRENT_SUBJECTRIGHTSREQUESTSTAGESTATUS
        case "COMPLETED":
            result = COMPLETED_SUBJECTRIGHTSREQUESTSTAGESTATUS
        case "FAILED":
            result = FAILED_SUBJECTRIGHTSREQUESTSTAGESTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_SUBJECTRIGHTSREQUESTSTAGESTATUS
        default:
            return 0, errors.New("Unknown SubjectRightsRequestStageStatus value: " + v)
    }
    return &result, nil
}
func SerializeSubjectRightsRequestStageStatus(values []SubjectRightsRequestStageStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
