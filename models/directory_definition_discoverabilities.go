package models
import (
    "errors"
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
    return []string{"None", "AttributeNames", "AttributeDataTypes", "AttributeReadOnly", "ReferenceAttributes", "UnknownFutureValue"}[i]
}
func ParseDirectoryDefinitionDiscoverabilities(v string) (any, error) {
    result := NONE_DIRECTORYDEFINITIONDISCOVERABILITIES
    switch v {
        case "None":
            result = NONE_DIRECTORYDEFINITIONDISCOVERABILITIES
        case "AttributeNames":
            result = ATTRIBUTENAMES_DIRECTORYDEFINITIONDISCOVERABILITIES
        case "AttributeDataTypes":
            result = ATTRIBUTEDATATYPES_DIRECTORYDEFINITIONDISCOVERABILITIES
        case "AttributeReadOnly":
            result = ATTRIBUTEREADONLY_DIRECTORYDEFINITIONDISCOVERABILITIES
        case "ReferenceAttributes":
            result = REFERENCEATTRIBUTES_DIRECTORYDEFINITIONDISCOVERABILITIES
        case "UnknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DIRECTORYDEFINITIONDISCOVERABILITIES
        default:
            return 0, errors.New("Unknown DirectoryDefinitionDiscoverabilities value: " + v)
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
