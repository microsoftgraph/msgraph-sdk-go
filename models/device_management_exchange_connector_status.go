package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the deviceManagement singleton.
type DeviceManagementExchangeConnectorStatus int

const (
    NONE_DEVICEMANAGEMENTEXCHANGECONNECTORSTATUS DeviceManagementExchangeConnectorStatus = iota
    CONNECTIONPENDING_DEVICEMANAGEMENTEXCHANGECONNECTORSTATUS
    CONNECTED_DEVICEMANAGEMENTEXCHANGECONNECTORSTATUS
    DISCONNECTED_DEVICEMANAGEMENTEXCHANGECONNECTORSTATUS
)

func (i DeviceManagementExchangeConnectorStatus) String() string {
    return []string{"NONE", "CONNECTIONPENDING", "CONNECTED", "DISCONNECTED"}[i]
}
func ParseDeviceManagementExchangeConnectorStatus(v string) (interface{}, error) {
    result := NONE_DEVICEMANAGEMENTEXCHANGECONNECTORSTATUS
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_DEVICEMANAGEMENTEXCHANGECONNECTORSTATUS
        case "CONNECTIONPENDING":
            result = CONNECTIONPENDING_DEVICEMANAGEMENTEXCHANGECONNECTORSTATUS
        case "CONNECTED":
            result = CONNECTED_DEVICEMANAGEMENTEXCHANGECONNECTORSTATUS
        case "DISCONNECTED":
            result = DISCONNECTED_DEVICEMANAGEMENTEXCHANGECONNECTORSTATUS
        default:
            return 0, errors.New("Unknown DeviceManagementExchangeConnectorStatus value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementExchangeConnectorStatus(values []DeviceManagementExchangeConnectorStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
