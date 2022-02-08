package graph
import (
    "strings"
    "errors"
)
// 
type AccessReviewHistoryDecisionFilter int

const (
    APPROVE_ACCESSREVIEWHISTORYDECISIONFILTER AccessReviewHistoryDecisionFilter = iota
    DENY_ACCESSREVIEWHISTORYDECISIONFILTER
    NOTREVIEWED_ACCESSREVIEWHISTORYDECISIONFILTER
    DONTKNOW_ACCESSREVIEWHISTORYDECISIONFILTER
    NOTNOTIFIED_ACCESSREVIEWHISTORYDECISIONFILTER
    UNKNOWNFUTUREVALUE_ACCESSREVIEWHISTORYDECISIONFILTER
)

func (i AccessReviewHistoryDecisionFilter) String() string {
    return []string{"APPROVE", "DENY", "NOTREVIEWED", "DONTKNOW", "NOTNOTIFIED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseAccessReviewHistoryDecisionFilter(v string) (interface{}, error) {
    result := APPROVE_ACCESSREVIEWHISTORYDECISIONFILTER
    switch strings.ToUpper(v) {
        case "APPROVE":
            result = APPROVE_ACCESSREVIEWHISTORYDECISIONFILTER
        case "DENY":
            result = DENY_ACCESSREVIEWHISTORYDECISIONFILTER
        case "NOTREVIEWED":
            result = NOTREVIEWED_ACCESSREVIEWHISTORYDECISIONFILTER
        case "DONTKNOW":
            result = DONTKNOW_ACCESSREVIEWHISTORYDECISIONFILTER
        case "NOTNOTIFIED":
            result = NOTNOTIFIED_ACCESSREVIEWHISTORYDECISIONFILTER
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_ACCESSREVIEWHISTORYDECISIONFILTER
        default:
            return 0, errors.New("Unknown AccessReviewHistoryDecisionFilter value: " + v)
    }
    return &result, nil
}
func SerializeAccessReviewHistoryDecisionFilter(values []AccessReviewHistoryDecisionFilter) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
