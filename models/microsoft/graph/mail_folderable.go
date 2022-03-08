package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MailFolderable 
type MailFolderable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetChildFolderCount()(*int32)
    GetChildFolders()([]MailFolderable)
    GetDisplayName()(*string)
    GetIsHidden()(*bool)
    GetMessageRules()([]MessageRuleable)
    GetMessages()([]Messageable)
    GetMultiValueExtendedProperties()([]MultiValueLegacyExtendedPropertyable)
    GetParentFolderId()(*string)
    GetSingleValueExtendedProperties()([]SingleValueLegacyExtendedPropertyable)
    GetTotalItemCount()(*int32)
    GetUnreadItemCount()(*int32)
    SetChildFolderCount(value *int32)()
    SetChildFolders(value []MailFolderable)()
    SetDisplayName(value *string)()
    SetIsHidden(value *bool)()
    SetMessageRules(value []MessageRuleable)()
    SetMessages(value []Messageable)()
    SetMultiValueExtendedProperties(value []MultiValueLegacyExtendedPropertyable)()
    SetParentFolderId(value *string)()
    SetSingleValueExtendedProperties(value []SingleValueLegacyExtendedPropertyable)()
    SetTotalItemCount(value *int32)()
    SetUnreadItemCount(value *int32)()
}
