package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ContactFolderable 
type ContactFolderable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetChildFolders()([]ContactFolderable)
    GetContacts()([]Contactable)
    GetDisplayName()(*string)
    GetMultiValueExtendedProperties()([]MultiValueLegacyExtendedPropertyable)
    GetParentFolderId()(*string)
    GetSingleValueExtendedProperties()([]SingleValueLegacyExtendedPropertyable)
    SetChildFolders(value []ContactFolderable)()
    SetContacts(value []Contactable)()
    SetDisplayName(value *string)()
    SetMultiValueExtendedProperties(value []MultiValueLegacyExtendedPropertyable)()
    SetParentFolderId(value *string)()
    SetSingleValueExtendedProperties(value []SingleValueLegacyExtendedPropertyable)()
}
