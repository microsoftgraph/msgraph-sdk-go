package models
import (
    "errors"
)
// 
type ObjectDefinitionMetadata int

const (
    PROPERTYNAMEACCOUNTENABLED_OBJECTDEFINITIONMETADATA ObjectDefinitionMetadata = iota
    PROPERTYNAMESOFTDELETED_OBJECTDEFINITIONMETADATA
    ISSOFTDELETIONSUPPORTED_OBJECTDEFINITIONMETADATA
    ISSYNCHRONIZEALLSUPPORTED_OBJECTDEFINITIONMETADATA
    CONNECTORDATASTORAGEREQUIRED_OBJECTDEFINITIONMETADATA
    EXTENSIONS_OBJECTDEFINITIONMETADATA
    BASEOBJECTNAME_OBJECTDEFINITIONMETADATA
)

func (i ObjectDefinitionMetadata) String() string {
    return []string{"PropertyNameAccountEnabled", "PropertyNameSoftDeleted", "IsSoftDeletionSupported", "IsSynchronizeAllSupported", "ConnectorDataStorageRequired", "Extensions", "BaseObjectName"}[i]
}
func ParseObjectDefinitionMetadata(v string) (any, error) {
    result := PROPERTYNAMEACCOUNTENABLED_OBJECTDEFINITIONMETADATA
    switch v {
        case "PropertyNameAccountEnabled":
            result = PROPERTYNAMEACCOUNTENABLED_OBJECTDEFINITIONMETADATA
        case "PropertyNameSoftDeleted":
            result = PROPERTYNAMESOFTDELETED_OBJECTDEFINITIONMETADATA
        case "IsSoftDeletionSupported":
            result = ISSOFTDELETIONSUPPORTED_OBJECTDEFINITIONMETADATA
        case "IsSynchronizeAllSupported":
            result = ISSYNCHRONIZEALLSUPPORTED_OBJECTDEFINITIONMETADATA
        case "ConnectorDataStorageRequired":
            result = CONNECTORDATASTORAGEREQUIRED_OBJECTDEFINITIONMETADATA
        case "Extensions":
            result = EXTENSIONS_OBJECTDEFINITIONMETADATA
        case "BaseObjectName":
            result = BASEOBJECTNAME_OBJECTDEFINITIONMETADATA
        default:
            return 0, errors.New("Unknown ObjectDefinitionMetadata value: " + v)
    }
    return &result, nil
}
func SerializeObjectDefinitionMetadata(values []ObjectDefinitionMetadata) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
