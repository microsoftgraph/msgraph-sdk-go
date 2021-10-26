package graph
import (
    "strings"
    "errors"
)
// 
type DeviceManagementReportStatus int

const (
    UNKNOWN_DEVICEMANAGEMENTREPORTSTATUS DeviceManagementReportStatus = iota
    NOTSTARTED_DEVICEMANAGEMENTREPORTSTATUS
    INPROGRESS_DEVICEMANAGEMENTREPORTSTATUS
    COMPLETED_DEVICEMANAGEMENTREPORTSTATUS
    FAILED_DEVICEMANAGEMENTREPORTSTATUS
)

func (i DeviceManagementReportStatus) String() string {
    return []string{"UNKNOWN", "NOTSTARTED", "INPROGRESS", "COMPLETED", "FAILED"}[i]
}
func ParseDeviceManagementReportStatus(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_DEVICEMANAGEMENTREPORTSTATUS, nil
        case "NOTSTARTED":
            return NOTSTARTED_DEVICEMANAGEMENTREPORTSTATUS, nil
        case "INPROGRESS":
            return INPROGRESS_DEVICEMANAGEMENTREPORTSTATUS, nil
        case "COMPLETED":
            return COMPLETED_DEVICEMANAGEMENTREPORTSTATUS, nil
        case "FAILED":
            return FAILED_DEVICEMANAGEMENTREPORTSTATUS, nil
    }
    return 0, errors.New("Unknown DeviceManagementReportStatus value: " + v)
}
func SerializeDeviceManagementReportStatus(values []DeviceManagementReportStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
