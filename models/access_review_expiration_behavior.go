package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the identityGovernance singleton.
type AccessReviewExpirationBehavior int

const (
    KEEPACCESS_ACCESSREVIEWEXPIRATIONBEHAVIOR AccessReviewExpirationBehavior = iota
    REMOVEACCESS_ACCESSREVIEWEXPIRATIONBEHAVIOR
    ACCEPTACCESSRECOMMENDATION_ACCESSREVIEWEXPIRATIONBEHAVIOR
    UNKNOWNFUTUREVALUE_ACCESSREVIEWEXPIRATIONBEHAVIOR
)

func (i AccessReviewExpirationBehavior) String() string {
    return []string{"KEEPACCESS", "REMOVEACCESS", "ACCEPTACCESSRECOMMENDATION", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseAccessReviewExpirationBehavior(v string) (interface{}, error) {
    result := KEEPACCESS_ACCESSREVIEWEXPIRATIONBEHAVIOR
    switch strings.ToUpper(v) {
        case "KEEPACCESS":
            result = KEEPACCESS_ACCESSREVIEWEXPIRATIONBEHAVIOR
        case "REMOVEACCESS":
            result = REMOVEACCESS_ACCESSREVIEWEXPIRATIONBEHAVIOR
        case "ACCEPTACCESSRECOMMENDATION":
            result = ACCEPTACCESSRECOMMENDATION_ACCESSREVIEWEXPIRATIONBEHAVIOR
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_ACCESSREVIEWEXPIRATIONBEHAVIOR
        default:
            return 0, errors.New("Unknown AccessReviewExpirationBehavior value: " + v)
    }
    return &result, nil
}
func SerializeAccessReviewExpirationBehavior(values []AccessReviewExpirationBehavior) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
