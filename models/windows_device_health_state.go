package models
import (
    "errors"
    "strings"
)
// Computer endpoint protection state
type WindowsDeviceHealthState int

const (
    // Computer is clean and no action is required
    CLEAN_WINDOWSDEVICEHEALTHSTATE WindowsDeviceHealthState = iota
    // Computer is in pending full scan state
    FULLSCANPENDING_WINDOWSDEVICEHEALTHSTATE
    // Computer is in pending reboot state
    REBOOTPENDING_WINDOWSDEVICEHEALTHSTATE
    // Computer is in pending manual steps state
    MANUALSTEPSPENDING_WINDOWSDEVICEHEALTHSTATE
    // Computer is in pending offline scan state
    OFFLINESCANPENDING_WINDOWSDEVICEHEALTHSTATE
    // Computer is in critical failure state
    CRITICAL_WINDOWSDEVICEHEALTHSTATE
)

func (i WindowsDeviceHealthState) String() string {
    var values []string
    for p := WindowsDeviceHealthState(1); p <= CRITICAL_WINDOWSDEVICEHEALTHSTATE; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"clean", "fullScanPending", "rebootPending", "manualStepsPending", "offlineScanPending", "critical"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseWindowsDeviceHealthState(v string) (any, error) {
    var result WindowsDeviceHealthState
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "clean":
                result |= CLEAN_WINDOWSDEVICEHEALTHSTATE
            case "fullScanPending":
                result |= FULLSCANPENDING_WINDOWSDEVICEHEALTHSTATE
            case "rebootPending":
                result |= REBOOTPENDING_WINDOWSDEVICEHEALTHSTATE
            case "manualStepsPending":
                result |= MANUALSTEPSPENDING_WINDOWSDEVICEHEALTHSTATE
            case "offlineScanPending":
                result |= OFFLINESCANPENDING_WINDOWSDEVICEHEALTHSTATE
            case "critical":
                result |= CRITICAL_WINDOWSDEVICEHEALTHSTATE
            default:
                return 0, errors.New("Unknown WindowsDeviceHealthState value: " + v)
        }
    }
    return &result, nil
}
func SerializeWindowsDeviceHealthState(values []WindowsDeviceHealthState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WindowsDeviceHealthState) isMultiValue() bool {
    return true
}
