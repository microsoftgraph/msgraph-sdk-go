package graph
import (
    "strings"
    "errors"
)
// 
type DeviceManagementPartnerAppType int

const (
    UNKNOWN_DEVICEMANAGEMENTPARTNERAPPTYPE DeviceManagementPartnerAppType = iota
    SINGLETENANTAPP_DEVICEMANAGEMENTPARTNERAPPTYPE
    MULTITENANTAPP_DEVICEMANAGEMENTPARTNERAPPTYPE
)

func (i DeviceManagementPartnerAppType) String() string {
    return []string{"UNKNOWN", "SINGLETENANTAPP", "MULTITENANTAPP"}[i]
}
func ParseDeviceManagementPartnerAppType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_DEVICEMANAGEMENTPARTNERAPPTYPE, nil
        case "SINGLETENANTAPP":
            return SINGLETENANTAPP_DEVICEMANAGEMENTPARTNERAPPTYPE, nil
        case "MULTITENANTAPP":
            return MULTITENANTAPP_DEVICEMANAGEMENTPARTNERAPPTYPE, nil
    }
    return 0, errors.New("Unknown DeviceManagementPartnerAppType value: " + v)
}
func SerializeDeviceManagementPartnerAppType(values []DeviceManagementPartnerAppType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
