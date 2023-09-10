package models
import (
    "errors"
    "strings"
)
// 
type DirectoryDefinitionDiscoverabilities int

const (
    NONE_DIRECTORYDEFINITIONDISCOVERABILITIES DirectoryDefinitionDiscoverabilities = iota
    ATTRIBUTENAMES_DIRECTORYDEFINITIONDISCOVERABILITIES
    ATTRIBUTEDATATYPES_DIRECTORYDEFINITIONDISCOVERABILITIES
    ATTRIBUTEREADONLY_DIRECTORYDEFINITIONDISCOVERABILITIES
    REFERENCEATTRIBUTES_DIRECTORYDEFINITIONDISCOVERABILITIES
    UNKNOWNFUTUREVALUE_DIRECTORYDEFINITIONDISCOVERABILITIES
)

func (i DirectoryDefinitionDiscoverabilities) String() string {
    var values []string
    for p := DirectoryDefinitionDiscoverabilities(1); p <= UNKNOWNFUTUREVALUE_DIRECTORYDEFINITIONDISCOVERABILITIES; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"None", "AttributeNames", "AttributeDataTypes", "AttributeReadOnly", "ReferenceAttributes", "UnknownFutureValue"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseDirectoryDefinitionDiscoverabilities(v string) (any, error) {
    var result DirectoryDefinitionDiscoverabilities
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "None":
                result |= NONE_DIRECTORYDEFINITIONDISCOVERABILITIES
            case "AttributeNames":
                result |= ATTRIBUTENAMES_DIRECTORYDEFINITIONDISCOVERABILITIES
            case "AttributeDataTypes":
                result |= ATTRIBUTEDATATYPES_DIRECTORYDEFINITIONDISCOVERABILITIES
            case "AttributeReadOnly":
                result |= ATTRIBUTEREADONLY_DIRECTORYDEFINITIONDISCOVERABILITIES
            case "ReferenceAttributes":
                result |= REFERENCEATTRIBUTES_DIRECTORYDEFINITIONDISCOVERABILITIES
            case "UnknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_DIRECTORYDEFINITIONDISCOVERABILITIES
            default:
                return 0, errors.New("Unknown DirectoryDefinitionDiscoverabilities value: " + v)
        }
    }
    return &result, nil
}
func SerializeDirectoryDefinitionDiscoverabilities(values []DirectoryDefinitionDiscoverabilities) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DirectoryDefinitionDiscoverabilities) isMultiValue() bool {
    return true
}
