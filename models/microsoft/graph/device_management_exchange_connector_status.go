package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_DEVICEMANAGEMENTEXCHANGECONNECTORSTATUS, nil
        case "CONNECTIONPENDING":
            return CONNECTIONPENDING_DEVICEMANAGEMENTEXCHANGECONNECTORSTATUS, nil
        case "CONNECTED":
            return CONNECTED_DEVICEMANAGEMENTEXCHANGECONNECTORSTATUS, nil
        case "DISCONNECTED":
            return DISCONNECTED_DEVICEMANAGEMENTEXCHANGECONNECTORSTATUS, nil
    }
    return 0, errors.New("Unknown DeviceManagementExchangeConnectorStatus value: " + v)
}
func SerializeDeviceManagementExchangeConnectorStatus(values []DeviceManagementExchangeConnectorStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
