package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// mailFolder 
type MailFolder struct {
    Entity
    // The number of immediate child mailFolders in the current mailFolder.
    childFolderCount *int32;
    // The collection of child folders in the mailFolder.
    childFolders []MailFolder;
    // The mailFolder's display name.
    displayName *string;
    // Indicates whether the mailFolder is hidden. This property can be set only when creating the folder. Find more information in Hidden mail folders.
    isHidden *bool;
    // The collection of rules that apply to the user's Inbox folder.
    messageRules []MessageRule;
    // The collection of messages in the mailFolder.
    messages []Message;
    // The collection of multi-value extended properties defined for the mailFolder. Read-only. Nullable.
    multiValueExtendedProperties []MultiValueLegacyExtendedProperty;
    // The unique identifier for the mailFolder's parent mailFolder.
    parentFolderId *string;
    // The collection of single-value extended properties defined for the mailFolder. Read-only. Nullable.
    singleValueExtendedProperties []SingleValueLegacyExtendedProperty;
    // The number of items in the mailFolder.
    totalItemCount *int32;
    // The number of items in the mailFolder marked as unread.
    unreadItemCount *int32;
}
// NewMailFolder instantiates a new mailFolder and sets the default values.
func NewMailFolder()(*MailFolder) {
    m := &MailFolder{
        Entity: *NewEntity(),
    }
    return m
}
// GetChildFolderCount gets the childFolderCount property value. The number of immediate child mailFolders in the current mailFolder.
func (m *MailFolder) GetChildFolderCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.childFolderCount
    }
}
// GetChildFolders gets the childFolders property value. The collection of child folders in the mailFolder.
func (m *MailFolder) GetChildFolders()([]MailFolder) {
    if m == nil {
        return nil
    } else {
        return m.childFolders
    }
}
// GetDisplayName gets the displayName property value. The mailFolder's display name.
func (m *MailFolder) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetIsHidden gets the isHidden property value. Indicates whether the mailFolder is hidden. This property can be set only when creating the folder. Find more information in Hidden mail folders.
func (m *MailFolder) GetIsHidden()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isHidden
    }
}
// GetMessageRules gets the messageRules property value. The collection of rules that apply to the user's Inbox folder.
func (m *MailFolder) GetMessageRules()([]MessageRule) {
    if m == nil {
        return nil
    } else {
        return m.messageRules
    }
}
// GetMessages gets the messages property value. The collection of messages in the mailFolder.
func (m *MailFolder) GetMessages()([]Message) {
    if m == nil {
        return nil
    } else {
        return m.messages
    }
}
// GetMultiValueExtendedProperties gets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the mailFolder. Read-only. Nullable.
func (m *MailFolder) GetMultiValueExtendedProperties()([]MultiValueLegacyExtendedProperty) {
    if m == nil {
        return nil
    } else {
        return m.multiValueExtendedProperties
    }
}
// GetParentFolderId gets the parentFolderId property value. The unique identifier for the mailFolder's parent mailFolder.
func (m *MailFolder) GetParentFolderId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.parentFolderId
    }
}
// GetSingleValueExtendedProperties gets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the mailFolder. Read-only. Nullable.
func (m *MailFolder) GetSingleValueExtendedProperties()([]SingleValueLegacyExtendedProperty) {
    if m == nil {
        return nil
    } else {
        return m.singleValueExtendedProperties
    }
}
// GetTotalItemCount gets the totalItemCount property value. The number of items in the mailFolder.
func (m *MailFolder) GetTotalItemCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalItemCount
    }
}
// GetUnreadItemCount gets the unreadItemCount property value. The number of items in the mailFolder marked as unread.
func (m *MailFolder) GetUnreadItemCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.unreadItemCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MailFolder) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["childFolderCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChildFolderCount(val)
        }
        return nil
    }
    res["childFolders"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMailFolder() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MailFolder, len(val))
            for i, v := range val {
                res[i] = *(v.(*MailFolder))
            }
            m.SetChildFolders(res)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["isHidden"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsHidden(val)
        }
        return nil
    }
    res["messageRules"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMessageRule() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MessageRule, len(val))
            for i, v := range val {
                res[i] = *(v.(*MessageRule))
            }
            m.SetMessageRules(res)
        }
        return nil
    }
    res["messages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMessage() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Message, len(val))
            for i, v := range val {
                res[i] = *(v.(*Message))
            }
            m.SetMessages(res)
        }
        return nil
    }
    res["multiValueExtendedProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMultiValueLegacyExtendedProperty() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MultiValueLegacyExtendedProperty, len(val))
            for i, v := range val {
                res[i] = *(v.(*MultiValueLegacyExtendedProperty))
            }
            m.SetMultiValueExtendedProperties(res)
        }
        return nil
    }
    res["parentFolderId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentFolderId(val)
        }
        return nil
    }
    res["singleValueExtendedProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSingleValueLegacyExtendedProperty() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SingleValueLegacyExtendedProperty, len(val))
            for i, v := range val {
                res[i] = *(v.(*SingleValueLegacyExtendedProperty))
            }
            m.SetSingleValueExtendedProperties(res)
        }
        return nil
    }
    res["totalItemCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalItemCount(val)
        }
        return nil
    }
    res["unreadItemCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnreadItemCount(val)
        }
        return nil
    }
    return res
}
func (m *MailFolder) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *MailFolder) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("childFolderCount", m.GetChildFolderCount())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetChildFolders()))
        for i, v := range m.GetChildFolders() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("childFolders", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isHidden", m.GetIsHidden())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMessageRules()))
        for i, v := range m.GetMessageRules() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("messageRules", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMessages()))
        for i, v := range m.GetMessages() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("messages", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMultiValueExtendedProperties()))
        for i, v := range m.GetMultiValueExtendedProperties() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("multiValueExtendedProperties", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("parentFolderId", m.GetParentFolderId())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSingleValueExtendedProperties()))
        for i, v := range m.GetSingleValueExtendedProperties() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("singleValueExtendedProperties", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalItemCount", m.GetTotalItemCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("unreadItemCount", m.GetUnreadItemCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetChildFolderCount sets the childFolderCount property value. The number of immediate child mailFolders in the current mailFolder.
func (m *MailFolder) SetChildFolderCount(value *int32)() {
    m.childFolderCount = value
}
// SetChildFolders sets the childFolders property value. The collection of child folders in the mailFolder.
func (m *MailFolder) SetChildFolders(value []MailFolder)() {
    m.childFolders = value
}
// SetDisplayName sets the displayName property value. The mailFolder's display name.
func (m *MailFolder) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIsHidden sets the isHidden property value. Indicates whether the mailFolder is hidden. This property can be set only when creating the folder. Find more information in Hidden mail folders.
func (m *MailFolder) SetIsHidden(value *bool)() {
    m.isHidden = value
}
// SetMessageRules sets the messageRules property value. The collection of rules that apply to the user's Inbox folder.
func (m *MailFolder) SetMessageRules(value []MessageRule)() {
    m.messageRules = value
}
// SetMessages sets the messages property value. The collection of messages in the mailFolder.
func (m *MailFolder) SetMessages(value []Message)() {
    m.messages = value
}
// SetMultiValueExtendedProperties sets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the mailFolder. Read-only. Nullable.
func (m *MailFolder) SetMultiValueExtendedProperties(value []MultiValueLegacyExtendedProperty)() {
    m.multiValueExtendedProperties = value
}
// SetParentFolderId sets the parentFolderId property value. The unique identifier for the mailFolder's parent mailFolder.
func (m *MailFolder) SetParentFolderId(value *string)() {
    m.parentFolderId = value
}
// SetSingleValueExtendedProperties sets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the mailFolder. Read-only. Nullable.
func (m *MailFolder) SetSingleValueExtendedProperties(value []SingleValueLegacyExtendedProperty)() {
    m.singleValueExtendedProperties = value
}
// SetTotalItemCount sets the totalItemCount property value. The number of items in the mailFolder.
func (m *MailFolder) SetTotalItemCount(value *int32)() {
    m.totalItemCount = value
}
// SetUnreadItemCount sets the unreadItemCount property value. The number of items in the mailFolder marked as unread.
func (m *MailFolder) SetUnreadItemCount(value *int32)() {
    m.unreadItemCount = value
}
