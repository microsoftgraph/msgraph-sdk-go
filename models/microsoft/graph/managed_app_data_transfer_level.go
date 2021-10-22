package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "ALLAPPS":
            return ALLAPPS_MANAGEDAPPDATATRANSFERLEVEL, nil
        case "MANAGEDAPPS":
            return MANAGEDAPPS_MANAGEDAPPDATATRANSFERLEVEL, nil
        case "NONE":
            return NONE_MANAGEDAPPDATATRANSFERLEVEL, nil
    }
    return 0, errors.New("Unknown ManagedAppDataTransferLevel value: " + v)
}
func SerializeManagedAppDataTransferLevel(values []ManagedAppDataTransferLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
