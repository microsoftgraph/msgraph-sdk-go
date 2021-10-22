package graph
import (
    "strings"
    "errors"
)
type DeviceManagementPartnerTenantState int

const (
    UNKNOWN_DEVICEMANAGEMENTPARTNERTENANTSTATE DeviceManagementPartnerTenantState = iota
    UNAVAILABLE_DEVICEMANAGEMENTPARTNERTENANTSTATE
    ENABLED_DEVICEMANAGEMENTPARTNERTENANTSTATE
    TERMINATED_DEVICEMANAGEMENTPARTNERTENANTSTATE
    REJECTED_DEVICEMANAGEMENTPARTNERTENANTSTATE
    UNRESPONSIVE_DEVICEMANAGEMENTPARTNERTENANTSTATE
)

func (i DeviceManagementPartnerTenantState) String() string {
    return []string{"UNKNOWN", "UNAVAILABLE", "ENABLED", "TERMINATED", "REJECTED", "UNRESPONSIVE"}[i]
}
func ParseDeviceManagementPartnerTenantState(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_DEVICEMANAGEMENTPARTNERTENANTSTATE, nil
        case "UNAVAILABLE":
            return UNAVAILABLE_DEVICEMANAGEMENTPARTNERTENANTSTATE, nil
        case "ENABLED":
            return ENABLED_DEVICEMANAGEMENTPARTNERTENANTSTATE, nil
        case "TERMINATED":
            return TERMINATED_DEVICEMANAGEMENTPARTNERTENANTSTATE, nil
        case "REJECTED":
            return REJECTED_DEVICEMANAGEMENTPARTNERTENANTSTATE, nil
        case "UNRESPONSIVE":
            return UNRESPONSIVE_DEVICEMANAGEMENTPARTNERTENANTSTATE, nil
    }
    return 0, errors.New("Unknown DeviceManagementPartnerTenantState value: " + v)
}
func SerializeDeviceManagementPartnerTenantState(values []DeviceManagementPartnerTenantState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
