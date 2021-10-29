package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "NOTREGISTERED":
            return NOTREGISTERED_DEVICEREGISTRATIONSTATE, nil
        case "REGISTERED":
            return REGISTERED_DEVICEREGISTRATIONSTATE, nil
        case "REVOKED":
            return REVOKED_DEVICEREGISTRATIONSTATE, nil
        case "KEYCONFLICT":
            return KEYCONFLICT_DEVICEREGISTRATIONSTATE, nil
        case "APPROVALPENDING":
            return APPROVALPENDING_DEVICEREGISTRATIONSTATE, nil
        case "CERTIFICATERESET":
            return CERTIFICATERESET_DEVICEREGISTRATIONSTATE, nil
        case "NOTREGISTEREDPENDINGENROLLMENT":
            return NOTREGISTEREDPENDINGENROLLMENT_DEVICEREGISTRATIONSTATE, nil
        case "UNKNOWN":
            return UNKNOWN_DEVICEREGISTRATIONSTATE, nil
    }
    return 0, errors.New("Unknown DeviceRegistrationState value: " + v)
}
func SerializeDeviceRegistrationState(values []DeviceRegistrationState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
