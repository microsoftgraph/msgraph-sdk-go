package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the deviceManagement singleton.
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
    result := UNKNOWN_DEVICEMANAGEMENTPARTNERTENANTSTATE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_DEVICEMANAGEMENTPARTNERTENANTSTATE
        case "UNAVAILABLE":
            result = UNAVAILABLE_DEVICEMANAGEMENTPARTNERTENANTSTATE
        case "ENABLED":
            result = ENABLED_DEVICEMANAGEMENTPARTNERTENANTSTATE
        case "TERMINATED":
            result = TERMINATED_DEVICEMANAGEMENTPARTNERTENANTSTATE
        case "REJECTED":
            result = REJECTED_DEVICEMANAGEMENTPARTNERTENANTSTATE
        case "UNRESPONSIVE":
            result = UNRESPONSIVE_DEVICEMANAGEMENTPARTNERTENANTSTATE
        default:
            return 0, errors.New("Unknown DeviceManagementPartnerTenantState value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementPartnerTenantState(values []DeviceManagementPartnerTenantState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
