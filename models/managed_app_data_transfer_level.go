package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the deviceAppManagement singleton.
type ManagedAppDataTransferLevel int

const (
    ALLAPPS_MANAGEDAPPDATATRANSFERLEVEL ManagedAppDataTransferLevel = iota
    MANAGEDAPPS_MANAGEDAPPDATATRANSFERLEVEL
    NONE_MANAGEDAPPDATATRANSFERLEVEL
)

func (i ManagedAppDataTransferLevel) String() string {
    return []string{"ALLAPPS", "MANAGEDAPPS", "NONE"}[i]
}
func ParseManagedAppDataTransferLevel(v string) (interface{}, error) {
    result := ALLAPPS_MANAGEDAPPDATATRANSFERLEVEL
    switch strings.ToUpper(v) {
        case "ALLAPPS":
            result = ALLAPPS_MANAGEDAPPDATATRANSFERLEVEL
        case "MANAGEDAPPS":
            result = MANAGEDAPPS_MANAGEDAPPDATATRANSFERLEVEL
        case "NONE":
            result = NONE_MANAGEDAPPDATATRANSFERLEVEL
        default:
            return 0, errors.New("Unknown ManagedAppDataTransferLevel value: " + v)
    }
    return &result, nil
}
func SerializeManagedAppDataTransferLevel(values []ManagedAppDataTransferLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
