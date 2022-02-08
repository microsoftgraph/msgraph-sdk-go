package graph
import (
    "strings"
    "errors"
)
// 
type DeviceManagementExchangeConnectorType int

const (
    ONPREMISES_DEVICEMANAGEMENTEXCHANGECONNECTORTYPE DeviceManagementExchangeConnectorType = iota
    HOSTED_DEVICEMANAGEMENTEXCHANGECONNECTORTYPE
    SERVICETOSERVICE_DEVICEMANAGEMENTEXCHANGECONNECTORTYPE
    DEDICATED_DEVICEMANAGEMENTEXCHANGECONNECTORTYPE
)

func (i DeviceManagementExchangeConnectorType) String() string {
    return []string{"ONPREMISES", "HOSTED", "SERVICETOSERVICE", "DEDICATED"}[i]
}
func ParseDeviceManagementExchangeConnectorType(v string) (interface{}, error) {
    result := ONPREMISES_DEVICEMANAGEMENTEXCHANGECONNECTORTYPE
    switch strings.ToUpper(v) {
        case "ONPREMISES":
            result = ONPREMISES_DEVICEMANAGEMENTEXCHANGECONNECTORTYPE
        case "HOSTED":
            result = HOSTED_DEVICEMANAGEMENTEXCHANGECONNECTORTYPE
        case "SERVICETOSERVICE":
            result = SERVICETOSERVICE_DEVICEMANAGEMENTEXCHANGECONNECTORTYPE
        case "DEDICATED":
            result = DEDICATED_DEVICEMANAGEMENTEXCHANGECONNECTORTYPE
        default:
            return 0, errors.New("Unknown DeviceManagementExchangeConnectorType value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementExchangeConnectorType(values []DeviceManagementExchangeConnectorType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
