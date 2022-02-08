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
    result := UNKNOWN_DEVICEMANAGEMENTPARTNERAPPTYPE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_DEVICEMANAGEMENTPARTNERAPPTYPE
        case "SINGLETENANTAPP":
            result = SINGLETENANTAPP_DEVICEMANAGEMENTPARTNERAPPTYPE
        case "MULTITENANTAPP":
            result = MULTITENANTAPP_DEVICEMANAGEMENTPARTNERAPPTYPE
        default:
            return 0, errors.New("Unknown DeviceManagementPartnerAppType value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementPartnerAppType(values []DeviceManagementPartnerAppType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
