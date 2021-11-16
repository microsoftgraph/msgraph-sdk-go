package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "SUBMITTED":
            return SUBMITTED_ACCESSPACKAGEREQUESTSTATE, nil
        case "PENDINGAPPROVAL":
            return PENDINGAPPROVAL_ACCESSPACKAGEREQUESTSTATE, nil
        case "DELIVERING":
            return DELIVERING_ACCESSPACKAGEREQUESTSTATE, nil
        case "DELIVERED":
            return DELIVERED_ACCESSPACKAGEREQUESTSTATE, nil
        case "DELIVERYFAILED":
            return DELIVERYFAILED_ACCESSPACKAGEREQUESTSTATE, nil
        case "DENIED":
            return DENIED_ACCESSPACKAGEREQUESTSTATE, nil
        case "SCHEDULED":
            return SCHEDULED_ACCESSPACKAGEREQUESTSTATE, nil
        case "CANCELED":
            return CANCELED_ACCESSPACKAGEREQUESTSTATE, nil
        case "PARTIALLYDELIVERED":
            return PARTIALLYDELIVERED_ACCESSPACKAGEREQUESTSTATE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_ACCESSPACKAGEREQUESTSTATE, nil
    }
    return 0, errors.New("Unknown AccessPackageRequestState value: " + v)
}
func SerializeAccessPackageRequestState(values []AccessPackageRequestState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
