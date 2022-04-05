package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the deviceManagement singleton.
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
    result := UNKNOWN_DEVICEMANAGEMENTREPORTSTATUS
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_DEVICEMANAGEMENTREPORTSTATUS
        case "NOTSTARTED":
            result = NOTSTARTED_DEVICEMANAGEMENTREPORTSTATUS
        case "INPROGRESS":
            result = INPROGRESS_DEVICEMANAGEMENTREPORTSTATUS
        case "COMPLETED":
            result = COMPLETED_DEVICEMANAGEMENTREPORTSTATUS
        case "FAILED":
            result = FAILED_DEVICEMANAGEMENTREPORTSTATUS
        default:
            return 0, errors.New("Unknown DeviceManagementReportStatus value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementReportStatus(values []DeviceManagementReportStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
