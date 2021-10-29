package externalconnectors
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "STRING":
            return STRING_PROPERTYTYPE, nil
        case "INT64":
            return INT64_PROPERTYTYPE, nil
        case "DOUBLE":
            return DOUBLE_PROPERTYTYPE, nil
        case "DATETIME":
            return DATETIME_PROPERTYTYPE, nil
        case "BOOLEAN":
            return BOOLEAN_PROPERTYTYPE, nil
        case "STRINGCOLLECTION":
            return STRINGCOLLECTION_PROPERTYTYPE, nil
        case "INT64COLLECTION":
            return INT64COLLECTION_PROPERTYTYPE, nil
        case "DOUBLECOLLECTION":
            return DOUBLECOLLECTION_PROPERTYTYPE, nil
        case "DATETIMECOLLECTION":
            return DATETIMECOLLECTION_PROPERTYTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_PROPERTYTYPE, nil
    }
    return 0, errors.New("Unknown PropertyType value: " + v)
}
func SerializePropertyType(values []PropertyType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
