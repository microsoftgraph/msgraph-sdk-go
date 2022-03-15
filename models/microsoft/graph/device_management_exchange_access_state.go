package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the drive singleton.
type DeviceManagementExchangeAccessState int

const (
    NONE_DEVICEMANAGEMENTEXCHANGEACCESSSTATE DeviceManagementExchangeAccessState = iota
    UNKNOWN_DEVICEMANAGEMENTEXCHANGEACCESSSTATE
    ALLOWED_DEVICEMANAGEMENTEXCHANGEACCESSSTATE
    BLOCKED_DEVICEMANAGEMENTEXCHANGEACCESSSTATE
    QUARANTINED_DEVICEMANAGEMENTEXCHANGEACCESSSTATE
)

func (i DeviceManagementExchangeAccessState) String() string {
    return []string{"NONE", "UNKNOWN", "ALLOWED", "BLOCKED", "QUARANTINED"}[i]
}
func ParseDeviceManagementExchangeAccessState(v string) (interface{}, error) {
    result := NONE_DEVICEMANAGEMENTEXCHANGEACCESSSTATE
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_DEVICEMANAGEMENTEXCHANGEACCESSSTATE
        case "UNKNOWN":
            result = UNKNOWN_DEVICEMANAGEMENTEXCHANGEACCESSSTATE
        case "ALLOWED":
            result = ALLOWED_DEVICEMANAGEMENTEXCHANGEACCESSSTATE
        case "BLOCKED":
            result = BLOCKED_DEVICEMANAGEMENTEXCHANGEACCESSSTATE
        case "QUARANTINED":
            result = QUARANTINED_DEVICEMANAGEMENTEXCHANGEACCESSSTATE
        default:
            return 0, errors.New("Unknown DeviceManagementExchangeAccessState value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementExchangeAccessState(values []DeviceManagementExchangeAccessState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
