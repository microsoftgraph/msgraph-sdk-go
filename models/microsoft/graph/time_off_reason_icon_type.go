package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_TIMEOFFREASONICONTYPE, nil
        case "CAR":
            return CAR_TIMEOFFREASONICONTYPE, nil
        case "CALENDAR":
            return CALENDAR_TIMEOFFREASONICONTYPE, nil
        case "RUNNING":
            return RUNNING_TIMEOFFREASONICONTYPE, nil
        case "PLANE":
            return PLANE_TIMEOFFREASONICONTYPE, nil
        case "FIRSTAID":
            return FIRSTAID_TIMEOFFREASONICONTYPE, nil
        case "DOCTOR":
            return DOCTOR_TIMEOFFREASONICONTYPE, nil
        case "NOTWORKING":
            return NOTWORKING_TIMEOFFREASONICONTYPE, nil
        case "CLOCK":
            return CLOCK_TIMEOFFREASONICONTYPE, nil
        case "JURYDUTY":
            return JURYDUTY_TIMEOFFREASONICONTYPE, nil
        case "GLOBE":
            return GLOBE_TIMEOFFREASONICONTYPE, nil
        case "CUP":
            return CUP_TIMEOFFREASONICONTYPE, nil
        case "PHONE":
            return PHONE_TIMEOFFREASONICONTYPE, nil
        case "WEATHER":
            return WEATHER_TIMEOFFREASONICONTYPE, nil
        case "UMBRELLA":
            return UMBRELLA_TIMEOFFREASONICONTYPE, nil
        case "PIGGYBANK":
            return PIGGYBANK_TIMEOFFREASONICONTYPE, nil
        case "DOG":
            return DOG_TIMEOFFREASONICONTYPE, nil
        case "CAKE":
            return CAKE_TIMEOFFREASONICONTYPE, nil
        case "TRAFFICCONE":
            return TRAFFICCONE_TIMEOFFREASONICONTYPE, nil
        case "PIN":
            return PIN_TIMEOFFREASONICONTYPE, nil
        case "SUNNY":
            return SUNNY_TIMEOFFREASONICONTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_TIMEOFFREASONICONTYPE, nil
    }
    return 0, errors.New("Unknown TimeOffReasonIconType value: " + v)
}
func SerializeTimeOffReasonIconType(values []TimeOffReasonIconType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
