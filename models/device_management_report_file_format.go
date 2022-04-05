package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the deviceManagement singleton.
type DeviceManagementReportFileFormat int

const (
    CSV_DEVICEMANAGEMENTREPORTFILEFORMAT DeviceManagementReportFileFormat = iota
    PDF_DEVICEMANAGEMENTREPORTFILEFORMAT
)

func (i DeviceManagementReportFileFormat) String() string {
    return []string{"CSV", "PDF"}[i]
}
func ParseDeviceManagementReportFileFormat(v string) (interface{}, error) {
    result := CSV_DEVICEMANAGEMENTREPORTFILEFORMAT
    switch strings.ToUpper(v) {
        case "CSV":
            result = CSV_DEVICEMANAGEMENTREPORTFILEFORMAT
        case "PDF":
            result = PDF_DEVICEMANAGEMENTREPORTFILEFORMAT
        default:
            return 0, errors.New("Unknown DeviceManagementReportFileFormat value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementReportFileFormat(values []DeviceManagementReportFileFormat) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
