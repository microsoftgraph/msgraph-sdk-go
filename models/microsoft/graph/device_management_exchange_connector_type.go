package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "ONPREMISES":
            return ONPREMISES_DEVICEMANAGEMENTEXCHANGECONNECTORTYPE, nil
        case "HOSTED":
            return HOSTED_DEVICEMANAGEMENTEXCHANGECONNECTORTYPE, nil
        case "SERVICETOSERVICE":
            return SERVICETOSERVICE_DEVICEMANAGEMENTEXCHANGECONNECTORTYPE, nil
        case "DEDICATED":
            return DEDICATED_DEVICEMANAGEMENTEXCHANGECONNECTORTYPE, nil
    }
    return 0, errors.New("Unknown DeviceManagementExchangeConnectorType value: " + v)
}
func SerializeDeviceManagementExchangeConnectorType(values []DeviceManagementExchangeConnectorType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
