package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the drive singleton.
type TimeOffReasonIconType int

const (
    NONE_TIMEOFFREASONICONTYPE TimeOffReasonIconType = iota
    CAR_TIMEOFFREASONICONTYPE
    CALENDAR_TIMEOFFREASONICONTYPE
    RUNNING_TIMEOFFREASONICONTYPE
    PLANE_TIMEOFFREASONICONTYPE
    FIRSTAID_TIMEOFFREASONICONTYPE
    DOCTOR_TIMEOFFREASONICONTYPE
    NOTWORKING_TIMEOFFREASONICONTYPE
    CLOCK_TIMEOFFREASONICONTYPE
    JURYDUTY_TIMEOFFREASONICONTYPE
    GLOBE_TIMEOFFREASONICONTYPE
    CUP_TIMEOFFREASONICONTYPE
    PHONE_TIMEOFFREASONICONTYPE
    WEATHER_TIMEOFFREASONICONTYPE
    UMBRELLA_TIMEOFFREASONICONTYPE
    PIGGYBANK_TIMEOFFREASONICONTYPE
    DOG_TIMEOFFREASONICONTYPE
    CAKE_TIMEOFFREASONICONTYPE
    TRAFFICCONE_TIMEOFFREASONICONTYPE
    PIN_TIMEOFFREASONICONTYPE
    SUNNY_TIMEOFFREASONICONTYPE
    UNKNOWNFUTUREVALUE_TIMEOFFREASONICONTYPE
)

func (i TimeOffReasonIconType) String() string {
    return []string{"NONE", "CAR", "CALENDAR", "RUNNING", "PLANE", "FIRSTAID", "DOCTOR", "NOTWORKING", "CLOCK", "JURYDUTY", "GLOBE", "CUP", "PHONE", "WEATHER", "UMBRELLA", "PIGGYBANK", "DOG", "CAKE", "TRAFFICCONE", "PIN", "SUNNY", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseTimeOffReasonIconType(v string) (interface{}, error) {
    result := NONE_TIMEOFFREASONICONTYPE
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_TIMEOFFREASONICONTYPE
        case "CAR":
            result = CAR_TIMEOFFREASONICONTYPE
        case "CALENDAR":
            result = CALENDAR_TIMEOFFREASONICONTYPE
        case "RUNNING":
            result = RUNNING_TIMEOFFREASONICONTYPE
        case "PLANE":
            result = PLANE_TIMEOFFREASONICONTYPE
        case "FIRSTAID":
            result = FIRSTAID_TIMEOFFREASONICONTYPE
        case "DOCTOR":
            result = DOCTOR_TIMEOFFREASONICONTYPE
        case "NOTWORKING":
            result = NOTWORKING_TIMEOFFREASONICONTYPE
        case "CLOCK":
            result = CLOCK_TIMEOFFREASONICONTYPE
        case "JURYDUTY":
            result = JURYDUTY_TIMEOFFREASONICONTYPE
        case "GLOBE":
            result = GLOBE_TIMEOFFREASONICONTYPE
        case "CUP":
            result = CUP_TIMEOFFREASONICONTYPE
        case "PHONE":
            result = PHONE_TIMEOFFREASONICONTYPE
        case "WEATHER":
            result = WEATHER_TIMEOFFREASONICONTYPE
        case "UMBRELLA":
            result = UMBRELLA_TIMEOFFREASONICONTYPE
        case "PIGGYBANK":
            result = PIGGYBANK_TIMEOFFREASONICONTYPE
        case "DOG":
            result = DOG_TIMEOFFREASONICONTYPE
        case "CAKE":
            result = CAKE_TIMEOFFREASONICONTYPE
        case "TRAFFICCONE":
            result = TRAFFICCONE_TIMEOFFREASONICONTYPE
        case "PIN":
            result = PIN_TIMEOFFREASONICONTYPE
        case "SUNNY":
            result = SUNNY_TIMEOFFREASONICONTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_TIMEOFFREASONICONTYPE
        default:
            return 0, errors.New("Unknown TimeOffReasonIconType value: " + v)
    }
    return &result, nil
}
func SerializeTimeOffReasonIconType(values []TimeOffReasonIconType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
