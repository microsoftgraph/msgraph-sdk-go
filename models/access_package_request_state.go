package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the identityGovernance singleton.
type AccessPackageRequestState int

const (
    SUBMITTED_ACCESSPACKAGEREQUESTSTATE AccessPackageRequestState = iota
    PENDINGAPPROVAL_ACCESSPACKAGEREQUESTSTATE
    DELIVERING_ACCESSPACKAGEREQUESTSTATE
    DELIVERED_ACCESSPACKAGEREQUESTSTATE
    DELIVERYFAILED_ACCESSPACKAGEREQUESTSTATE
    DENIED_ACCESSPACKAGEREQUESTSTATE
    SCHEDULED_ACCESSPACKAGEREQUESTSTATE
    CANCELED_ACCESSPACKAGEREQUESTSTATE
    PARTIALLYDELIVERED_ACCESSPACKAGEREQUESTSTATE
    UNKNOWNFUTUREVALUE_ACCESSPACKAGEREQUESTSTATE
)

func (i AccessPackageRequestState) String() string {
    return []string{"SUBMITTED", "PENDINGAPPROVAL", "DELIVERING", "DELIVERED", "DELIVERYFAILED", "DENIED", "SCHEDULED", "CANCELED", "PARTIALLYDELIVERED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseAccessPackageRequestState(v string) (interface{}, error) {
    result := SUBMITTED_ACCESSPACKAGEREQUESTSTATE
    switch strings.ToUpper(v) {
        case "SUBMITTED":
            result = SUBMITTED_ACCESSPACKAGEREQUESTSTATE
        case "PENDINGAPPROVAL":
            result = PENDINGAPPROVAL_ACCESSPACKAGEREQUESTSTATE
        case "DELIVERING":
            result = DELIVERING_ACCESSPACKAGEREQUESTSTATE
        case "DELIVERED":
            result = DELIVERED_ACCESSPACKAGEREQUESTSTATE
        case "DELIVERYFAILED":
            result = DELIVERYFAILED_ACCESSPACKAGEREQUESTSTATE
        case "DENIED":
            result = DENIED_ACCESSPACKAGEREQUESTSTATE
        case "SCHEDULED":
            result = SCHEDULED_ACCESSPACKAGEREQUESTSTATE
        case "CANCELED":
            result = CANCELED_ACCESSPACKAGEREQUESTSTATE
        case "PARTIALLYDELIVERED":
            result = PARTIALLYDELIVERED_ACCESSPACKAGEREQUESTSTATE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_ACCESSPACKAGEREQUESTSTATE
        default:
            return 0, errors.New("Unknown AccessPackageRequestState value: " + v)
    }
    return &result, nil
}
func SerializeAccessPackageRequestState(values []AccessPackageRequestState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
