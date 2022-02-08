package graph
import (
    "strings"
    "errors"
)
// 
type ImportedWindowsAutopilotDeviceIdentityImportStatus int

const (
    UNKNOWN_IMPORTEDWINDOWSAUTOPILOTDEVICEIDENTITYIMPORTSTATUS ImportedWindowsAutopilotDeviceIdentityImportStatus = iota
    PENDING_IMPORTEDWINDOWSAUTOPILOTDEVICEIDENTITYIMPORTSTATUS
    PARTIAL_IMPORTEDWINDOWSAUTOPILOTDEVICEIDENTITYIMPORTSTATUS
    COMPLETE_IMPORTEDWINDOWSAUTOPILOTDEVICEIDENTITYIMPORTSTATUS
    ERROR_IMPORTEDWINDOWSAUTOPILOTDEVICEIDENTITYIMPORTSTATUS
)

func (i ImportedWindowsAutopilotDeviceIdentityImportStatus) String() string {
    return []string{"UNKNOWN", "PENDING", "PARTIAL", "COMPLETE", "ERROR"}[i]
}
func ParseImportedWindowsAutopilotDeviceIdentityImportStatus(v string) (interface{}, error) {
    result := UNKNOWN_IMPORTEDWINDOWSAUTOPILOTDEVICEIDENTITYIMPORTSTATUS
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_IMPORTEDWINDOWSAUTOPILOTDEVICEIDENTITYIMPORTSTATUS
        case "PENDING":
            result = PENDING_IMPORTEDWINDOWSAUTOPILOTDEVICEIDENTITYIMPORTSTATUS
        case "PARTIAL":
            result = PARTIAL_IMPORTEDWINDOWSAUTOPILOTDEVICEIDENTITYIMPORTSTATUS
        case "COMPLETE":
            result = COMPLETE_IMPORTEDWINDOWSAUTOPILOTDEVICEIDENTITYIMPORTSTATUS
        case "ERROR":
            result = ERROR_IMPORTEDWINDOWSAUTOPILOTDEVICEIDENTITYIMPORTSTATUS
        default:
            return 0, errors.New("Unknown ImportedWindowsAutopilotDeviceIdentityImportStatus value: " + v)
    }
    return &result, nil
}
func SerializeImportedWindowsAutopilotDeviceIdentityImportStatus(values []ImportedWindowsAutopilotDeviceIdentityImportStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
