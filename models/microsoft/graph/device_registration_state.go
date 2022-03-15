package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the drive singleton.
type DeviceRegistrationState int

const (
    NOTREGISTERED_DEVICEREGISTRATIONSTATE DeviceRegistrationState = iota
    REGISTERED_DEVICEREGISTRATIONSTATE
    REVOKED_DEVICEREGISTRATIONSTATE
    KEYCONFLICT_DEVICEREGISTRATIONSTATE
    APPROVALPENDING_DEVICEREGISTRATIONSTATE
    CERTIFICATERESET_DEVICEREGISTRATIONSTATE
    NOTREGISTEREDPENDINGENROLLMENT_DEVICEREGISTRATIONSTATE
    UNKNOWN_DEVICEREGISTRATIONSTATE
)

func (i DeviceRegistrationState) String() string {
    return []string{"NOTREGISTERED", "REGISTERED", "REVOKED", "KEYCONFLICT", "APPROVALPENDING", "CERTIFICATERESET", "NOTREGISTEREDPENDINGENROLLMENT", "UNKNOWN"}[i]
}
func ParseDeviceRegistrationState(v string) (interface{}, error) {
    result := NOTREGISTERED_DEVICEREGISTRATIONSTATE
    switch strings.ToUpper(v) {
        case "NOTREGISTERED":
            result = NOTREGISTERED_DEVICEREGISTRATIONSTATE
        case "REGISTERED":
            result = REGISTERED_DEVICEREGISTRATIONSTATE
        case "REVOKED":
            result = REVOKED_DEVICEREGISTRATIONSTATE
        case "KEYCONFLICT":
            result = KEYCONFLICT_DEVICEREGISTRATIONSTATE
        case "APPROVALPENDING":
            result = APPROVALPENDING_DEVICEREGISTRATIONSTATE
        case "CERTIFICATERESET":
            result = CERTIFICATERESET_DEVICEREGISTRATIONSTATE
        case "NOTREGISTEREDPENDINGENROLLMENT":
            result = NOTREGISTEREDPENDINGENROLLMENT_DEVICEREGISTRATIONSTATE
        case "UNKNOWN":
            result = UNKNOWN_DEVICEREGISTRATIONSTATE
        default:
            return 0, errors.New("Unknown DeviceRegistrationState value: " + v)
    }
    return &result, nil
}
func SerializeDeviceRegistrationState(values []DeviceRegistrationState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
