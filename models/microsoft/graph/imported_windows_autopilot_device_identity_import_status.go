package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_IMPORTEDWINDOWSAUTOPILOTDEVICEIDENTITYIMPORTSTATUS, nil
        case "PENDING":
            return PENDING_IMPORTEDWINDOWSAUTOPILOTDEVICEIDENTITYIMPORTSTATUS, nil
        case "PARTIAL":
            return PARTIAL_IMPORTEDWINDOWSAUTOPILOTDEVICEIDENTITYIMPORTSTATUS, nil
        case "COMPLETE":
            return COMPLETE_IMPORTEDWINDOWSAUTOPILOTDEVICEIDENTITYIMPORTSTATUS, nil
        case "ERROR":
            return ERROR_IMPORTEDWINDOWSAUTOPILOTDEVICEIDENTITYIMPORTSTATUS, nil
    }
    return 0, errors.New("Unknown ImportedWindowsAutopilotDeviceIdentityImportStatus value: " + v)
}
func SerializeImportedWindowsAutopilotDeviceIdentityImportStatus(values []ImportedWindowsAutopilotDeviceIdentityImportStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
