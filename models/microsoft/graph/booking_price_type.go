package graph
import (
    "strings"
    "errors"
)
// 
type BookingPriceType int

const (
    UNDEFINED_BOOKINGPRICETYPE BookingPriceType = iota
    FIXEDPRICE_BOOKINGPRICETYPE
    STARTINGAT_BOOKINGPRICETYPE
    HOURLY_BOOKINGPRICETYPE
    FREE_BOOKINGPRICETYPE
    PRICEVARIES_BOOKINGPRICETYPE
    CALLUS_BOOKINGPRICETYPE
    NOTSET_BOOKINGPRICETYPE
    UNKNOWNFUTUREVALUE_BOOKINGPRICETYPE
)

func (i BookingPriceType) String() string {
    return []string{"UNDEFINED", "FIXEDPRICE", "STARTINGAT", "HOURLY", "FREE", "PRICEVARIES", "CALLUS", "NOTSET", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseBookingPriceType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "UNDEFINED":
            return UNDEFINED_BOOKINGPRICETYPE, nil
        case "FIXEDPRICE":
            return FIXEDPRICE_BOOKINGPRICETYPE, nil
        case "STARTINGAT":
            return STARTINGAT_BOOKINGPRICETYPE, nil
        case "HOURLY":
            return HOURLY_BOOKINGPRICETYPE, nil
        case "FREE":
            return FREE_BOOKINGPRICETYPE, nil
        case "PRICEVARIES":
            return PRICEVARIES_BOOKINGPRICETYPE, nil
        case "CALLUS":
            return CALLUS_BOOKINGPRICETYPE, nil
        case "NOTSET":
            return NOTSET_BOOKINGPRICETYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_BOOKINGPRICETYPE, nil
    }
    return 0, errors.New("Unknown BookingPriceType value: " + v)
}
func SerializeBookingPriceType(values []BookingPriceType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
