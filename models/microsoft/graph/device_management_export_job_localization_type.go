package graph
import (
    "strings"
    "errors"
)
// 
type DeviceManagementExportJobLocalizationType int

const (
    LOCALIZEDVALUESASADDITIONALCOLUMN_DEVICEMANAGEMENTEXPORTJOBLOCALIZATIONTYPE DeviceManagementExportJobLocalizationType = iota
    REPLACELOCALIZABLEVALUES_DEVICEMANAGEMENTEXPORTJOBLOCALIZATIONTYPE
)

func (i DeviceManagementExportJobLocalizationType) String() string {
    return []string{"LOCALIZEDVALUESASADDITIONALCOLUMN", "REPLACELOCALIZABLEVALUES"}[i]
}
func ParseDeviceManagementExportJobLocalizationType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "LOCALIZEDVALUESASADDITIONALCOLUMN":
            return LOCALIZEDVALUESASADDITIONALCOLUMN_DEVICEMANAGEMENTEXPORTJOBLOCALIZATIONTYPE, nil
        case "REPLACELOCALIZABLEVALUES":
            return REPLACELOCALIZABLEVALUES_DEVICEMANAGEMENTEXPORTJOBLOCALIZATIONTYPE, nil
    }
    return 0, errors.New("Unknown DeviceManagementExportJobLocalizationType value: " + v)
}
func SerializeDeviceManagementExportJobLocalizationType(values []DeviceManagementExportJobLocalizationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
