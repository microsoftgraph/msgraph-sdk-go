package graph
import (
    "strings"
    "errors"
)
// Provides operations to call the sync method.
type DeviceManagementExchangeConnectorSyncType int

const (
    FULLSYNC_DEVICEMANAGEMENTEXCHANGECONNECTORSYNCTYPE DeviceManagementExchangeConnectorSyncType = iota
    DELTASYNC_DEVICEMANAGEMENTEXCHANGECONNECTORSYNCTYPE
)

func (i DeviceManagementExchangeConnectorSyncType) String() string {
    return []string{"FULLSYNC", "DELTASYNC"}[i]
}
func ParseDeviceManagementExchangeConnectorSyncType(v string) (interface{}, error) {
    result := FULLSYNC_DEVICEMANAGEMENTEXCHANGECONNECTORSYNCTYPE
    switch strings.ToUpper(v) {
        case "FULLSYNC":
            result = FULLSYNC_DEVICEMANAGEMENTEXCHANGECONNECTORSYNCTYPE
        case "DELTASYNC":
            result = DELTASYNC_DEVICEMANAGEMENTEXCHANGECONNECTORSYNCTYPE
        default:
            return 0, errors.New("Unknown DeviceManagementExchangeConnectorSyncType value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementExchangeConnectorSyncType(values []DeviceManagementExchangeConnectorSyncType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
