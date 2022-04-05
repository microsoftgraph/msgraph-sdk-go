package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the deviceManagement singleton.
type DeviceManagementSubscriptionState int

const (
    PENDING_DEVICEMANAGEMENTSUBSCRIPTIONSTATE DeviceManagementSubscriptionState = iota
    ACTIVE_DEVICEMANAGEMENTSUBSCRIPTIONSTATE
    WARNING_DEVICEMANAGEMENTSUBSCRIPTIONSTATE
    DISABLED_DEVICEMANAGEMENTSUBSCRIPTIONSTATE
    DELETED_DEVICEMANAGEMENTSUBSCRIPTIONSTATE
    BLOCKED_DEVICEMANAGEMENTSUBSCRIPTIONSTATE
    LOCKEDOUT_DEVICEMANAGEMENTSUBSCRIPTIONSTATE
)

func (i DeviceManagementSubscriptionState) String() string {
    return []string{"PENDING", "ACTIVE", "WARNING", "DISABLED", "DELETED", "BLOCKED", "LOCKEDOUT"}[i]
}
func ParseDeviceManagementSubscriptionState(v string) (interface{}, error) {
    result := PENDING_DEVICEMANAGEMENTSUBSCRIPTIONSTATE
    switch strings.ToUpper(v) {
        case "PENDING":
            result = PENDING_DEVICEMANAGEMENTSUBSCRIPTIONSTATE
        case "ACTIVE":
            result = ACTIVE_DEVICEMANAGEMENTSUBSCRIPTIONSTATE
        case "WARNING":
            result = WARNING_DEVICEMANAGEMENTSUBSCRIPTIONSTATE
        case "DISABLED":
            result = DISABLED_DEVICEMANAGEMENTSUBSCRIPTIONSTATE
        case "DELETED":
            result = DELETED_DEVICEMANAGEMENTSUBSCRIPTIONSTATE
        case "BLOCKED":
            result = BLOCKED_DEVICEMANAGEMENTSUBSCRIPTIONSTATE
        case "LOCKEDOUT":
            result = LOCKEDOUT_DEVICEMANAGEMENTSUBSCRIPTIONSTATE
        default:
            return 0, errors.New("Unknown DeviceManagementSubscriptionState value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementSubscriptionState(values []DeviceManagementSubscriptionState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
