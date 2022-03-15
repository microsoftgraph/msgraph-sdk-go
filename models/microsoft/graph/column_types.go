package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the collection of drive entities.
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
    result := NOTE_COLUMNTYPES
    switch strings.ToUpper(v) {
        case "NOTE":
            result = NOTE_COLUMNTYPES
        case "TEXT":
            result = TEXT_COLUMNTYPES
        case "CHOICE":
            result = CHOICE_COLUMNTYPES
        case "MULTICHOICE":
            result = MULTICHOICE_COLUMNTYPES
        case "NUMBER":
            result = NUMBER_COLUMNTYPES
        case "CURRENCY":
            result = CURRENCY_COLUMNTYPES
        case "DATETIME":
            result = DATETIME_COLUMNTYPES
        case "LOOKUP":
            result = LOOKUP_COLUMNTYPES
        case "BOOLEAN":
            result = BOOLEAN_COLUMNTYPES
        case "USER":
            result = USER_COLUMNTYPES
        case "URL":
            result = URL_COLUMNTYPES
        case "CALCULATED":
            result = CALCULATED_COLUMNTYPES
        case "LOCATION":
            result = LOCATION_COLUMNTYPES
        case "GEOLOCATION":
            result = GEOLOCATION_COLUMNTYPES
        case "TERM":
            result = TERM_COLUMNTYPES
        case "MULTITERM":
            result = MULTITERM_COLUMNTYPES
        case "THUMBNAIL":
            result = THUMBNAIL_COLUMNTYPES
        case "APPROVALSTATUS":
            result = APPROVALSTATUS_COLUMNTYPES
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_COLUMNTYPES
        default:
            return 0, errors.New("Unknown ColumnTypes value: " + v)
    }
    return &result, nil
}
func SerializeColumnTypes(values []ColumnTypes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
