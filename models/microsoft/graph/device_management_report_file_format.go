package graph
import (
    "strings"
    "errors"
)
// 
type DeviceManagementReportFileFormat int

const (
    CSV_DEVICEMANAGEMENTREPORTFILEFORMAT DeviceManagementReportFileFormat = iota
    PDF_DEVICEMANAGEMENTREPORTFILEFORMAT
)

func (i DeviceManagementReportFileFormat) String() string {
    return []string{"CSV", "PDF"}[i]
}
func ParseDeviceManagementReportFileFormat(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "CSV":
            return CSV_DEVICEMANAGEMENTREPORTFILEFORMAT, nil
        case "PDF":
            return PDF_DEVICEMANAGEMENTREPORTFILEFORMAT, nil
    }
    return 0, errors.New("Unknown DeviceManagementReportFileFormat value: " + v)
}
func SerializeDeviceManagementReportFileFormat(values []DeviceManagementReportFileFormat) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
