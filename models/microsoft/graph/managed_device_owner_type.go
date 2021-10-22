package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_MANAGEDDEVICEOWNERTYPE, nil
        case "COMPANY":
            return COMPANY_MANAGEDDEVICEOWNERTYPE, nil
        case "PERSONAL":
            return PERSONAL_MANAGEDDEVICEOWNERTYPE, nil
    }
    return 0, errors.New("Unknown ManagedDeviceOwnerType value: " + v)
}
func SerializeManagedDeviceOwnerType(values []ManagedDeviceOwnerType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
