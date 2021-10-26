package graph
import (
    "strings"
    "errors"
)
// 
type LocationType int

const (
    DEFAULT_LOCATIONTYPE LocationType = iota
    CONFERENCEROOM_LOCATIONTYPE
    HOMEADDRESS_LOCATIONTYPE
    BUSINESSADDRESS_LOCATIONTYPE
    GEOCOORDINATES_LOCATIONTYPE
    STREETADDRESS_LOCATIONTYPE
    HOTEL_LOCATIONTYPE
    RESTAURANT_LOCATIONTYPE
    LOCALBUSINESS_LOCATIONTYPE
    POSTALADDRESS_LOCATIONTYPE
)

func (i LocationType) String() string {
    return []string{"DEFAULT", "CONFERENCEROOM", "HOMEADDRESS", "BUSINESSADDRESS", "GEOCOORDINATES", "STREETADDRESS", "HOTEL", "RESTAURANT", "LOCALBUSINESS", "POSTALADDRESS"}[i]
}
func ParseLocationType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "DEFAULT":
            return DEFAULT_LOCATIONTYPE, nil
        case "CONFERENCEROOM":
            return CONFERENCEROOM_LOCATIONTYPE, nil
        case "HOMEADDRESS":
            return HOMEADDRESS_LOCATIONTYPE, nil
        case "BUSINESSADDRESS":
            return BUSINESSADDRESS_LOCATIONTYPE, nil
        case "GEOCOORDINATES":
            return GEOCOORDINATES_LOCATIONTYPE, nil
        case "STREETADDRESS":
            return STREETADDRESS_LOCATIONTYPE, nil
        case "HOTEL":
            return HOTEL_LOCATIONTYPE, nil
        case "RESTAURANT":
            return RESTAURANT_LOCATIONTYPE, nil
        case "LOCALBUSINESS":
            return LOCALBUSINESS_LOCATIONTYPE, nil
        case "POSTALADDRESS":
            return POSTALADDRESS_LOCATIONTYPE, nil
    }
    return 0, errors.New("Unknown LocationType value: " + v)
}
func SerializeLocationType(values []LocationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
