package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the deviceManagement singleton.
type DeviceComplianceActionType int

const (
    NOACTION_DEVICECOMPLIANCEACTIONTYPE DeviceComplianceActionType = iota
    NOTIFICATION_DEVICECOMPLIANCEACTIONTYPE
    BLOCK_DEVICECOMPLIANCEACTIONTYPE
    RETIRE_DEVICECOMPLIANCEACTIONTYPE
    WIPE_DEVICECOMPLIANCEACTIONTYPE
    REMOVERESOURCEACCESSPROFILES_DEVICECOMPLIANCEACTIONTYPE
    PUSHNOTIFICATION_DEVICECOMPLIANCEACTIONTYPE
)

func (i DeviceComplianceActionType) String() string {
    return []string{"NOACTION", "NOTIFICATION", "BLOCK", "RETIRE", "WIPE", "REMOVERESOURCEACCESSPROFILES", "PUSHNOTIFICATION"}[i]
}
func ParseDeviceComplianceActionType(v string) (interface{}, error) {
    result := NOACTION_DEVICECOMPLIANCEACTIONTYPE
    switch strings.ToUpper(v) {
        case "NOACTION":
            result = NOACTION_DEVICECOMPLIANCEACTIONTYPE
        case "NOTIFICATION":
            result = NOTIFICATION_DEVICECOMPLIANCEACTIONTYPE
        case "BLOCK":
            result = BLOCK_DEVICECOMPLIANCEACTIONTYPE
        case "RETIRE":
            result = RETIRE_DEVICECOMPLIANCEACTIONTYPE
        case "WIPE":
            result = WIPE_DEVICECOMPLIANCEACTIONTYPE
        case "REMOVERESOURCEACCESSPROFILES":
            result = REMOVERESOURCEACCESSPROFILES_DEVICECOMPLIANCEACTIONTYPE
        case "PUSHNOTIFICATION":
            result = PUSHNOTIFICATION_DEVICECOMPLIANCEACTIONTYPE
        default:
            return 0, errors.New("Unknown DeviceComplianceActionType value: " + v)
    }
    return &result, nil
}
func SerializeDeviceComplianceActionType(values []DeviceComplianceActionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
