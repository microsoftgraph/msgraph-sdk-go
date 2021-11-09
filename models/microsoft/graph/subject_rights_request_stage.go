package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "CONTENTRETRIEVAL":
            return CONTENTRETRIEVAL_SUBJECTRIGHTSREQUESTSTAGE, nil
        case "CONTENTREVIEW":
            return CONTENTREVIEW_SUBJECTRIGHTSREQUESTSTAGE, nil
        case "GENERATEREPORT":
            return GENERATEREPORT_SUBJECTRIGHTSREQUESTSTAGE, nil
        case "CONTENTDELETION":
            return CONTENTDELETION_SUBJECTRIGHTSREQUESTSTAGE, nil
        case "CASERESOLVED":
            return CASERESOLVED_SUBJECTRIGHTSREQUESTSTAGE, nil
        case "CONTENTESTIMATE":
            return CONTENTESTIMATE_SUBJECTRIGHTSREQUESTSTAGE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_SUBJECTRIGHTSREQUESTSTAGE, nil
    }
    return 0, errors.New("Unknown SubjectRightsRequestStage value: " + v)
}
func SerializeSubjectRightsRequestStage(values []SubjectRightsRequestStage) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
