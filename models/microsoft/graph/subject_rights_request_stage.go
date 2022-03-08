package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the privacy singleton.
type SubjectRightsRequestStage int

const (
    CONTENTRETRIEVAL_SUBJECTRIGHTSREQUESTSTAGE SubjectRightsRequestStage = iota
    CONTENTREVIEW_SUBJECTRIGHTSREQUESTSTAGE
    GENERATEREPORT_SUBJECTRIGHTSREQUESTSTAGE
    CONTENTDELETION_SUBJECTRIGHTSREQUESTSTAGE
    CASERESOLVED_SUBJECTRIGHTSREQUESTSTAGE
    CONTENTESTIMATE_SUBJECTRIGHTSREQUESTSTAGE
    UNKNOWNFUTUREVALUE_SUBJECTRIGHTSREQUESTSTAGE
)

func (i SubjectRightsRequestStage) String() string {
    return []string{"CONTENTRETRIEVAL", "CONTENTREVIEW", "GENERATEREPORT", "CONTENTDELETION", "CASERESOLVED", "CONTENTESTIMATE", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseSubjectRightsRequestStage(v string) (interface{}, error) {
    result := CONTENTRETRIEVAL_SUBJECTRIGHTSREQUESTSTAGE
    switch strings.ToUpper(v) {
        case "CONTENTRETRIEVAL":
            result = CONTENTRETRIEVAL_SUBJECTRIGHTSREQUESTSTAGE
        case "CONTENTREVIEW":
            result = CONTENTREVIEW_SUBJECTRIGHTSREQUESTSTAGE
        case "GENERATEREPORT":
            result = GENERATEREPORT_SUBJECTRIGHTSREQUESTSTAGE
        case "CONTENTDELETION":
            result = CONTENTDELETION_SUBJECTRIGHTSREQUESTSTAGE
        case "CASERESOLVED":
            result = CASERESOLVED_SUBJECTRIGHTSREQUESTSTAGE
        case "CONTENTESTIMATE":
            result = CONTENTESTIMATE_SUBJECTRIGHTSREQUESTSTAGE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_SUBJECTRIGHTSREQUESTSTAGE
        default:
            return 0, errors.New("Unknown SubjectRightsRequestStage value: " + v)
    }
    return &result, nil
}
func SerializeSubjectRightsRequestStage(values []SubjectRightsRequestStage) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
