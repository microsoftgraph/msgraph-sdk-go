package externalconnectors
import (
    "strings"
    "errors"
)
// Provides operations to manage the collection of externalConnection entities.
type PropertyType int

const (
    STRING_PROPERTYTYPE PropertyType = iota
    INT64_PROPERTYTYPE
    DOUBLE_PROPERTYTYPE
    DATETIME_PROPERTYTYPE
    BOOLEAN_PROPERTYTYPE
    STRINGCOLLECTION_PROPERTYTYPE
    INT64COLLECTION_PROPERTYTYPE
    DOUBLECOLLECTION_PROPERTYTYPE
    DATETIMECOLLECTION_PROPERTYTYPE
    UNKNOWNFUTUREVALUE_PROPERTYTYPE
)

func (i PropertyType) String() string {
    return []string{"STRING", "INT64", "DOUBLE", "DATETIME", "BOOLEAN", "STRINGCOLLECTION", "INT64COLLECTION", "DOUBLECOLLECTION", "DATETIMECOLLECTION", "UNKNOWNFUTUREVALUE"}[i]
}
func ParsePropertyType(v string) (interface{}, error) {
    result := STRING_PROPERTYTYPE
    switch strings.ToUpper(v) {
        case "STRING":
            result = STRING_PROPERTYTYPE
        case "INT64":
            result = INT64_PROPERTYTYPE
        case "DOUBLE":
            result = DOUBLE_PROPERTYTYPE
        case "DATETIME":
            result = DATETIME_PROPERTYTYPE
        case "BOOLEAN":
            result = BOOLEAN_PROPERTYTYPE
        case "STRINGCOLLECTION":
            result = STRINGCOLLECTION_PROPERTYTYPE
        case "INT64COLLECTION":
            result = INT64COLLECTION_PROPERTYTYPE
        case "DOUBLECOLLECTION":
            result = DOUBLECOLLECTION_PROPERTYTYPE
        case "DATETIMECOLLECTION":
            result = DATETIMECOLLECTION_PROPERTYTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_PROPERTYTYPE
        default:
            return 0, errors.New("Unknown PropertyType value: " + v)
    }
    return &result, nil
}
func SerializePropertyType(values []PropertyType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
