package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the drive singleton.
type ManagedDeviceOwnerType int

const (
    UNKNOWN_MANAGEDDEVICEOWNERTYPE ManagedDeviceOwnerType = iota
    COMPANY_MANAGEDDEVICEOWNERTYPE
    PERSONAL_MANAGEDDEVICEOWNERTYPE
)

func (i ManagedDeviceOwnerType) String() string {
    return []string{"UNKNOWN", "COMPANY", "PERSONAL"}[i]
}
func ParseManagedDeviceOwnerType(v string) (interface{}, error) {
    result := UNKNOWN_MANAGEDDEVICEOWNERTYPE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_MANAGEDDEVICEOWNERTYPE
        case "COMPANY":
            result = COMPANY_MANAGEDDEVICEOWNERTYPE
        case "PERSONAL":
            result = PERSONAL_MANAGEDDEVICEOWNERTYPE
        default:
            return 0, errors.New("Unknown ManagedDeviceOwnerType value: " + v)
    }
    return &result, nil
}
func SerializeManagedDeviceOwnerType(values []ManagedDeviceOwnerType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
