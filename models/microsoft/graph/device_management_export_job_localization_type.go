package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the deviceManagement singleton.
type DeviceManagementExportJobLocalizationType int

const (
    LOCALIZEDVALUESASADDITIONALCOLUMN_DEVICEMANAGEMENTEXPORTJOBLOCALIZATIONTYPE DeviceManagementExportJobLocalizationType = iota
    REPLACELOCALIZABLEVALUES_DEVICEMANAGEMENTEXPORTJOBLOCALIZATIONTYPE
)

func (i DeviceManagementExportJobLocalizationType) String() string {
    return []string{"LOCALIZEDVALUESASADDITIONALCOLUMN", "REPLACELOCALIZABLEVALUES"}[i]
}
func ParseDeviceManagementExportJobLocalizationType(v string) (interface{}, error) {
    result := LOCALIZEDVALUESASADDITIONALCOLUMN_DEVICEMANAGEMENTEXPORTJOBLOCALIZATIONTYPE
    switch strings.ToUpper(v) {
        case "LOCALIZEDVALUESASADDITIONALCOLUMN":
            result = LOCALIZEDVALUESASADDITIONALCOLUMN_DEVICEMANAGEMENTEXPORTJOBLOCALIZATIONTYPE
        case "REPLACELOCALIZABLEVALUES":
            result = REPLACELOCALIZABLEVALUES_DEVICEMANAGEMENTEXPORTJOBLOCALIZATIONTYPE
        default:
            return 0, errors.New("Unknown DeviceManagementExportJobLocalizationType value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementExportJobLocalizationType(values []DeviceManagementExportJobLocalizationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
