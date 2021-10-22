package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "PENDING":
            return PENDING_DEVICEMANAGEMENTSUBSCRIPTIONSTATE, nil
        case "ACTIVE":
            return ACTIVE_DEVICEMANAGEMENTSUBSCRIPTIONSTATE, nil
        case "WARNING":
            return WARNING_DEVICEMANAGEMENTSUBSCRIPTIONSTATE, nil
        case "DISABLED":
            return DISABLED_DEVICEMANAGEMENTSUBSCRIPTIONSTATE, nil
        case "DELETED":
            return DELETED_DEVICEMANAGEMENTSUBSCRIPTIONSTATE, nil
        case "BLOCKED":
            return BLOCKED_DEVICEMANAGEMENTSUBSCRIPTIONSTATE, nil
        case "LOCKEDOUT":
            return LOCKEDOUT_DEVICEMANAGEMENTSUBSCRIPTIONSTATE, nil
    }
    return 0, errors.New("Unknown DeviceManagementSubscriptionState value: " + v)
}
func SerializeDeviceManagementSubscriptionState(values []DeviceManagementSubscriptionState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
