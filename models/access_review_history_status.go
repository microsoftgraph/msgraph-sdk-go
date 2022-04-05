package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the identityGovernance singleton.
type AccessReviewHistoryStatus int

const (
    DONE_ACCESSREVIEWHISTORYSTATUS AccessReviewHistoryStatus = iota
    INPROGRESS_ACCESSREVIEWHISTORYSTATUS
    ERROR_ACCESSREVIEWHISTORYSTATUS
    REQUESTED_ACCESSREVIEWHISTORYSTATUS
    UNKNOWNFUTUREVALUE_ACCESSREVIEWHISTORYSTATUS
)

func (i AccessReviewHistoryStatus) String() string {
    return []string{"DONE", "INPROGRESS", "ERROR", "REQUESTED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseAccessReviewHistoryStatus(v string) (interface{}, error) {
    result := DONE_ACCESSREVIEWHISTORYSTATUS
    switch strings.ToUpper(v) {
        case "DONE":
            result = DONE_ACCESSREVIEWHISTORYSTATUS
        case "INPROGRESS":
            result = INPROGRESS_ACCESSREVIEWHISTORYSTATUS
        case "ERROR":
            result = ERROR_ACCESSREVIEWHISTORYSTATUS
        case "REQUESTED":
            result = REQUESTED_ACCESSREVIEWHISTORYSTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_ACCESSREVIEWHISTORYSTATUS
        default:
            return 0, errors.New("Unknown AccessReviewHistoryStatus value: " + v)
    }
    return &result, nil
}
func SerializeAccessReviewHistoryStatus(values []AccessReviewHistoryStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
