package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DelegatedPermissionClassificationable 
type DelegatedPermissionClassificationable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetClassification()(*PermissionClassificationType)
    GetPermissionId()(*string)
    GetPermissionName()(*string)
    SetClassification(value *PermissionClassificationType)()
    SetPermissionId(value *string)()
    SetPermissionName(value *string)()
}
