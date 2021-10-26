package graph
import (
    "strings"
    "errors"
)
// 
type ColumnTypes int

const (
    NOTE_COLUMNTYPES ColumnTypes = iota
    TEXT_COLUMNTYPES
    CHOICE_COLUMNTYPES
    MULTICHOICE_COLUMNTYPES
    NUMBER_COLUMNTYPES
    CURRENCY_COLUMNTYPES
    DATETIME_COLUMNTYPES
    LOOKUP_COLUMNTYPES
    BOOLEAN_COLUMNTYPES
    USER_COLUMNTYPES
    URL_COLUMNTYPES
    CALCULATED_COLUMNTYPES
    LOCATION_COLUMNTYPES
    GEOLOCATION_COLUMNTYPES
    TERM_COLUMNTYPES
    MULTITERM_COLUMNTYPES
    THUMBNAIL_COLUMNTYPES
    APPROVALSTATUS_COLUMNTYPES
    UNKNOWNFUTUREVALUE_COLUMNTYPES
)

func (i ColumnTypes) String() string {
    return []string{"NOTE", "TEXT", "CHOICE", "MULTICHOICE", "NUMBER", "CURRENCY", "DATETIME", "LOOKUP", "BOOLEAN", "USER", "URL", "CALCULATED", "LOCATION", "GEOLOCATION", "TERM", "MULTITERM", "THUMBNAIL", "APPROVALSTATUS", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseColumnTypes(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "NOTE":
            return NOTE_COLUMNTYPES, nil
        case "TEXT":
            return TEXT_COLUMNTYPES, nil
        case "CHOICE":
            return CHOICE_COLUMNTYPES, nil
        case "MULTICHOICE":
            return MULTICHOICE_COLUMNTYPES, nil
        case "NUMBER":
            return NUMBER_COLUMNTYPES, nil
        case "CURRENCY":
            return CURRENCY_COLUMNTYPES, nil
        case "DATETIME":
            return DATETIME_COLUMNTYPES, nil
        case "LOOKUP":
            return LOOKUP_COLUMNTYPES, nil
        case "BOOLEAN":
            return BOOLEAN_COLUMNTYPES, nil
        case "USER":
            return USER_COLUMNTYPES, nil
        case "URL":
            return URL_COLUMNTYPES, nil
        case "CALCULATED":
            return CALCULATED_COLUMNTYPES, nil
        case "LOCATION":
            return LOCATION_COLUMNTYPES, nil
        case "GEOLOCATION":
            return GEOLOCATION_COLUMNTYPES, nil
        case "TERM":
            return TERM_COLUMNTYPES, nil
        case "MULTITERM":
            return MULTITERM_COLUMNTYPES, nil
        case "THUMBNAIL":
            return THUMBNAIL_COLUMNTYPES, nil
        case "APPROVALSTATUS":
            return APPROVALSTATUS_COLUMNTYPES, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_COLUMNTYPES, nil
    }
    return 0, errors.New("Unknown ColumnTypes value: " + v)
}
func SerializeColumnTypes(values []ColumnTypes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
