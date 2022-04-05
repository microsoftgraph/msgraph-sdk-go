package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the solutionsRoot singleton.
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
    result := UNDEFINED_BOOKINGPRICETYPE
    switch strings.ToUpper(v) {
        case "UNDEFINED":
            result = UNDEFINED_BOOKINGPRICETYPE
        case "FIXEDPRICE":
            result = FIXEDPRICE_BOOKINGPRICETYPE
        case "STARTINGAT":
            result = STARTINGAT_BOOKINGPRICETYPE
        case "HOURLY":
            result = HOURLY_BOOKINGPRICETYPE
        case "FREE":
            result = FREE_BOOKINGPRICETYPE
        case "PRICEVARIES":
            result = PRICEVARIES_BOOKINGPRICETYPE
        case "CALLUS":
            result = CALLUS_BOOKINGPRICETYPE
        case "NOTSET":
            result = NOTSET_BOOKINGPRICETYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_BOOKINGPRICETYPE
        default:
            return 0, errors.New("Unknown BookingPriceType value: " + v)
    }
    return &result, nil
}
func SerializeBookingPriceType(values []BookingPriceType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
