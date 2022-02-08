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
    result := DEFAULT_LOCATIONTYPE
    switch strings.ToUpper(v) {
        case "DEFAULT":
            result = DEFAULT_LOCATIONTYPE
        case "CONFERENCEROOM":
            result = CONFERENCEROOM_LOCATIONTYPE
        case "HOMEADDRESS":
            result = HOMEADDRESS_LOCATIONTYPE
        case "BUSINESSADDRESS":
            result = BUSINESSADDRESS_LOCATIONTYPE
        case "GEOCOORDINATES":
            result = GEOCOORDINATES_LOCATIONTYPE
        case "STREETADDRESS":
            result = STREETADDRESS_LOCATIONTYPE
        case "HOTEL":
            result = HOTEL_LOCATIONTYPE
        case "RESTAURANT":
            result = RESTAURANT_LOCATIONTYPE
        case "LOCALBUSINESS":
            result = LOCALBUSINESS_LOCATIONTYPE
        case "POSTALADDRESS":
            result = POSTALADDRESS_LOCATIONTYPE
        default:
            return 0, errors.New("Unknown LocationType value: " + v)
    }
    return &result, nil
}
func SerializeLocationType(values []LocationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
