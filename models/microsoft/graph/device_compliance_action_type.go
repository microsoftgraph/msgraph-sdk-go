package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "NOACTION":
            return NOACTION_DEVICECOMPLIANCEACTIONTYPE, nil
        case "NOTIFICATION":
            return NOTIFICATION_DEVICECOMPLIANCEACTIONTYPE, nil
        case "BLOCK":
            return BLOCK_DEVICECOMPLIANCEACTIONTYPE, nil
        case "RETIRE":
            return RETIRE_DEVICECOMPLIANCEACTIONTYPE, nil
        case "WIPE":
            return WIPE_DEVICECOMPLIANCEACTIONTYPE, nil
        case "REMOVERESOURCEACCESSPROFILES":
            return REMOVERESOURCEACCESSPROFILES_DEVICECOMPLIANCEACTIONTYPE, nil
        case "PUSHNOTIFICATION":
            return PUSHNOTIFICATION_DEVICECOMPLIANCEACTIONTYPE, nil
    }
    return 0, errors.New("Unknown DeviceComplianceActionType value: " + v)
}
func SerializeDeviceComplianceActionType(values []DeviceComplianceActionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
