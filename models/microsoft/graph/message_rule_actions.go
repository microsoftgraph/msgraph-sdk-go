package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MessageRuleActions 
type MessageRuleActions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // A list of categories to be assigned to a message.
    assignCategories []string;
    // The ID of a folder that a message is to be copied to.
    copyToFolder *string;
    // Indicates whether a message should be moved to the Deleted Items folder.
    delete *bool;
    // The email addresses of the recipients to which a message should be forwarded as an attachment.
    forwardAsAttachmentTo []Recipient;
    // The email addresses of the recipients to which a message should be forwarded.
    forwardTo []Recipient;
    // Indicates whether a message should be marked as read.
    markAsRead *bool;
    // Sets the importance of the message, which can be: low, normal, high.
    markImportance *Importance;
    // The ID of the folder that a message will be moved to.
    moveToFolder *string;
    // Indicates whether a message should be permanently deleted and not saved to the Deleted Items folder.
    permanentDelete *bool;
    // The email addresses to which a message should be redirected.
    redirectTo []Recipient;
    // Indicates whether subsequent rules should be evaluated.
    stopProcessingRules *bool;
}
// NewMessageRuleActions instantiates a new messageRuleActions and sets the default values.
func NewMessageRuleActions()(*MessageRuleActions) {
    m := &MessageRuleActions{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MessageRuleActions) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAssignCategories gets the assignCategories property value. A list of categories to be assigned to a message.
func (m *MessageRuleActions) GetAssignCategories()([]string) {
    if m == nil {
        return nil
    } else {
        return m.assignCategories
    }
}
// GetCopyToFolder gets the copyToFolder property value. The ID of a folder that a message is to be copied to.
func (m *MessageRuleActions) GetCopyToFolder()(*string) {
    if m == nil {
        return nil
    } else {
        return m.copyToFolder
    }
}
// GetDelete gets the delete property value. Indicates whether a message should be moved to the Deleted Items folder.
func (m *MessageRuleActions) GetDelete()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.delete
    }
}
// GetForwardAsAttachmentTo gets the forwardAsAttachmentTo property value. The email addresses of the recipients to which a message should be forwarded as an attachment.
func (m *MessageRuleActions) GetForwardAsAttachmentTo()([]Recipient) {
    if m == nil {
        return nil
    } else {
        return m.forwardAsAttachmentTo
    }
}
// GetForwardTo gets the forwardTo property value. The email addresses of the recipients to which a message should be forwarded.
func (m *MessageRuleActions) GetForwardTo()([]Recipient) {
    if m == nil {
        return nil
    } else {
        return m.forwardTo
    }
}
// GetMarkAsRead gets the markAsRead property value. Indicates whether a message should be marked as read.
func (m *MessageRuleActions) GetMarkAsRead()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.markAsRead
    }
}
// GetMarkImportance gets the markImportance property value. Sets the importance of the message, which can be: low, normal, high.
func (m *MessageRuleActions) GetMarkImportance()(*Importance) {
    if m == nil {
        return nil
    } else {
        return m.markImportance
    }
}
// GetMoveToFolder gets the moveToFolder property value. The ID of the folder that a message will be moved to.
func (m *MessageRuleActions) GetMoveToFolder()(*string) {
    if m == nil {
        return nil
    } else {
        return m.moveToFolder
    }
}
// GetPermanentDelete gets the permanentDelete property value. Indicates whether a message should be permanently deleted and not saved to the Deleted Items folder.
func (m *MessageRuleActions) GetPermanentDelete()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.permanentDelete
    }
}
// GetRedirectTo gets the redirectTo property value. The email addresses to which a message should be redirected.
func (m *MessageRuleActions) GetRedirectTo()([]Recipient) {
    if m == nil {
        return nil
    } else {
        return m.redirectTo
    }
}
// GetStopProcessingRules gets the stopProcessingRules property value. Indicates whether subsequent rules should be evaluated.
func (m *MessageRuleActions) GetStopProcessingRules()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.stopProcessingRules
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MessageRuleActions) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["assignCategories"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAssignCategories(res)
        }
        return nil
    }
    res["copyToFolder"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCopyToFolder(val)
        }
        return nil
    }
    res["delete"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDelete(val)
        }
        return nil
    }
    res["forwardAsAttachmentTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRecipient() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Recipient, len(val))
            for i, v := range val {
                res[i] = *(v.(*Recipient))
            }
            m.SetForwardAsAttachmentTo(res)
        }
        return nil
    }
    res["forwardTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRecipient() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Recipient, len(val))
            for i, v := range val {
                res[i] = *(v.(*Recipient))
            }
            m.SetForwardTo(res)
        }
        return nil
    }
    res["markAsRead"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMarkAsRead(val)
        }
        return nil
    }
    res["markImportance"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseImportance)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(Importance)
            m.SetMarkImportance(&cast)
        }
        return nil
    }
    res["moveToFolder"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMoveToFolder(val)
        }
        return nil
    }
    res["permanentDelete"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermanentDelete(val)
        }
        return nil
    }
    res["redirectTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRecipient() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Recipient, len(val))
            for i, v := range val {
                res[i] = *(v.(*Recipient))
            }
            m.SetRedirectTo(res)
        }
        return nil
    }
    res["stopProcessingRules"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStopProcessingRules(val)
        }
        return nil
    }
    return res
}
func (m *MessageRuleActions) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *MessageRuleActions) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("assignCategories", m.GetAssignCategories())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("copyToFolder", m.GetCopyToFolder())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("delete", m.GetDelete())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetForwardAsAttachmentTo()))
        for i, v := range m.GetForwardAsAttachmentTo() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("forwardAsAttachmentTo", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetForwardTo()))
        for i, v := range m.GetForwardTo() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("forwardTo", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("markAsRead", m.GetMarkAsRead())
        if err != nil {
            return err
        }
    }
    if m.GetMarkImportance() != nil {
        cast := m.GetMarkImportance().String()
        err := writer.WriteStringValue("markImportance", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("moveToFolder", m.GetMoveToFolder())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("permanentDelete", m.GetPermanentDelete())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRedirectTo()))
        for i, v := range m.GetRedirectTo() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("redirectTo", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("stopProcessingRules", m.GetStopProcessingRules())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MessageRuleActions) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAssignCategories sets the assignCategories property value. A list of categories to be assigned to a message.
func (m *MessageRuleActions) SetAssignCategories(value []string)() {
    if m != nil {
        m.assignCategories = value
    }
}
// SetCopyToFolder sets the copyToFolder property value. The ID of a folder that a message is to be copied to.
func (m *MessageRuleActions) SetCopyToFolder(value *string)() {
    if m != nil {
        m.copyToFolder = value
    }
}
// SetDelete sets the delete property value. Indicates whether a message should be moved to the Deleted Items folder.
func (m *MessageRuleActions) SetDelete(value *bool)() {
    if m != nil {
        m.delete = value
    }
}
// SetForwardAsAttachmentTo sets the forwardAsAttachmentTo property value. The email addresses of the recipients to which a message should be forwarded as an attachment.
func (m *MessageRuleActions) SetForwardAsAttachmentTo(value []Recipient)() {
    if m != nil {
        m.forwardAsAttachmentTo = value
    }
}
// SetForwardTo sets the forwardTo property value. The email addresses of the recipients to which a message should be forwarded.
func (m *MessageRuleActions) SetForwardTo(value []Recipient)() {
    if m != nil {
        m.forwardTo = value
    }
}
// SetMarkAsRead sets the markAsRead property value. Indicates whether a message should be marked as read.
func (m *MessageRuleActions) SetMarkAsRead(value *bool)() {
    if m != nil {
        m.markAsRead = value
    }
}
// SetMarkImportance sets the markImportance property value. Sets the importance of the message, which can be: low, normal, high.
func (m *MessageRuleActions) SetMarkImportance(value *Importance)() {
    if m != nil {
        m.markImportance = value
    }
}
// SetMoveToFolder sets the moveToFolder property value. The ID of the folder that a message will be moved to.
func (m *MessageRuleActions) SetMoveToFolder(value *string)() {
    if m != nil {
        m.moveToFolder = value
    }
}
// SetPermanentDelete sets the permanentDelete property value. Indicates whether a message should be permanently deleted and not saved to the Deleted Items folder.
func (m *MessageRuleActions) SetPermanentDelete(value *bool)() {
    if m != nil {
        m.permanentDelete = value
    }
}
// SetRedirectTo sets the redirectTo property value. The email addresses to which a message should be redirected.
func (m *MessageRuleActions) SetRedirectTo(value []Recipient)() {
    if m != nil {
        m.redirectTo = value
    }
}
// SetStopProcessingRules sets the stopProcessingRules property value. Indicates whether subsequent rules should be evaluated.
func (m *MessageRuleActions) SetStopProcessingRules(value *bool)() {
    if m != nil {
        m.stopProcessingRules = value
    }
}
