package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the deviceAppManagement singleton.
type VppTokenSyncStatus int

const (
    NONE_VPPTOKENSYNCSTATUS VppTokenSyncStatus = iota
    INPROGRESS_VPPTOKENSYNCSTATUS
    COMPLETED_VPPTOKENSYNCSTATUS
    FAILED_VPPTOKENSYNCSTATUS
)

func (i VppTokenSyncStatus) String() string {
    return []string{"NONE", "INPROGRESS", "COMPLETED", "FAILED"}[i]
}
func ParseVppTokenSyncStatus(v string) (interface{}, error) {
    result := NONE_VPPTOKENSYNCSTATUS
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_VPPTOKENSYNCSTATUS
        case "INPROGRESS":
            result = INPROGRESS_VPPTOKENSYNCSTATUS
        case "COMPLETED":
            result = COMPLETED_VPPTOKENSYNCSTATUS
        case "FAILED":
            result = FAILED_VPPTOKENSYNCSTATUS
        default:
            return 0, errors.New("Unknown VppTokenSyncStatus value: " + v)
    }
    return &result, nil
}
func SerializeVppTokenSyncStatus(values []VppTokenSyncStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
