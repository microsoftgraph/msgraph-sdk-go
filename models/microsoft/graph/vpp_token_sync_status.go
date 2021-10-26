package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_VPPTOKENSYNCSTATUS, nil
        case "INPROGRESS":
            return INPROGRESS_VPPTOKENSYNCSTATUS, nil
        case "COMPLETED":
            return COMPLETED_VPPTOKENSYNCSTATUS, nil
        case "FAILED":
            return FAILED_VPPTOKENSYNCSTATUS, nil
    }
    return 0, errors.New("Unknown VppTokenSyncStatus value: " + v)
}
func SerializeVppTokenSyncStatus(values []VppTokenSyncStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
