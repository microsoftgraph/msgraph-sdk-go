package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_DEVICEMANAGEMENTEXCHANGEACCESSSTATE, nil
        case "UNKNOWN":
            return UNKNOWN_DEVICEMANAGEMENTEXCHANGEACCESSSTATE, nil
        case "ALLOWED":
            return ALLOWED_DEVICEMANAGEMENTEXCHANGEACCESSSTATE, nil
        case "BLOCKED":
            return BLOCKED_DEVICEMANAGEMENTEXCHANGEACCESSSTATE, nil
        case "QUARANTINED":
            return QUARANTINED_DEVICEMANAGEMENTEXCHANGEACCESSSTATE, nil
    }
    return 0, errors.New("Unknown DeviceManagementExchangeAccessState value: " + v)
}
func SerializeDeviceManagementExchangeAccessState(values []DeviceManagementExchangeAccessState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
