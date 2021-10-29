package graph
import (
    "strings"
    "errors"
)
// 
type DeviceManagementExchangeConnectorSyncType int

const (
    FULLSYNC_DEVICEMANAGEMENTEXCHANGECONNECTORSYNCTYPE DeviceManagementExchangeConnectorSyncType = iota
    DELTASYNC_DEVICEMANAGEMENTEXCHANGECONNECTORSYNCTYPE
)

func (i DeviceManagementExchangeConnectorSyncType) String() string {
    return []string{"FULLSYNC", "DELTASYNC"}[i]
}
func ParseDeviceManagementExchangeConnectorSyncType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "FULLSYNC":
            return FULLSYNC_DEVICEMANAGEMENTEXCHANGECONNECTORSYNCTYPE, nil
        case "DELTASYNC":
            return DELTASYNC_DEVICEMANAGEMENTEXCHANGECONNECTORSYNCTYPE, nil
    }
    return 0, errors.New("Unknown DeviceManagementExchangeConnectorSyncType value: " + v)
}
func SerializeDeviceManagementExchangeConnectorSyncType(values []DeviceManagementExchangeConnectorSyncType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
