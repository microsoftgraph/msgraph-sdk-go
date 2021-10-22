package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type MessageRuleActions struct {
    additionalData map[string]interface{};
    assignCategories []string;
    copyToFolder *string;
    delete *bool;
    forwardAsAttachmentTo []Recipient;
    forwardTo []Recipient;
    markAsRead *bool;
    markImportance *Importance;
    moveToFolder *string;
    permanentDelete *bool;
    redirectTo []Recipient;
    stopProcessingRules *bool;
}
func NewMessageRuleActions()(*MessageRuleActions) {
    m := &MessageRuleActions{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *MessageRuleActions) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *MessageRuleActions) GetAssignCategories()([]string) {
    if m == nil {
        return nil
    } else {
        return m.assignCategories
    }
}
func (m *MessageRuleActions) GetCopyToFolder()(*string) {
    if m == nil {
        return nil
    } else {
        return m.copyToFolder
    }
}
func (m *MessageRuleActions) GetDelete()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.delete
    }
}
func (m *MessageRuleActions) GetForwardAsAttachmentTo()([]Recipient) {
    if m == nil {
        return nil
    } else {
        return m.forwardAsAttachmentTo
    }
}
func (m *MessageRuleActions) GetForwardTo()([]Recipient) {
    if m == nil {
        return nil
    } else {
        return m.forwardTo
    }
}
func (m *MessageRuleActions) GetMarkAsRead()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.markAsRead
    }
}
func (m *MessageRuleActions) GetMarkImportance()(*Importance) {
    if m == nil {
        return nil
    } else {
        return m.markImportance
    }
}
func (m *MessageRuleActions) GetMoveToFolder()(*string) {
    if m == nil {
        return nil
    } else {
        return m.moveToFolder
    }
}
func (m *MessageRuleActions) GetPermanentDelete()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.permanentDelete
    }
}
func (m *MessageRuleActions) GetRedirectTo()([]Recipient) {
    if m == nil {
        return nil
    } else {
        return m.redirectTo
    }
}
func (m *MessageRuleActions) GetStopProcessingRules()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.stopProcessingRules
    }
}
func (m *MessageRuleActions) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["assignCategories"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetAssignCategories(res)
        return nil
    }
    res["copyToFolder"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCopyToFolder(val)
        return nil
    }
    res["delete"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetDelete(val)
        return nil
    }
    res["forwardAsAttachmentTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRecipient() })
        if err != nil {
            return err
        }
        res := make([]Recipient, len(val))
        for i, v := range val {
            res[i] = *(v.(*Recipient))
        }
        m.SetForwardAsAttachmentTo(res)
        return nil
    }
    res["forwardTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRecipient() })
        if err != nil {
            return err
        }
        res := make([]Recipient, len(val))
        for i, v := range val {
            res[i] = *(v.(*Recipient))
        }
        m.SetForwardTo(res)
        return nil
    }
    res["markAsRead"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetMarkAsRead(val)
        return nil
    }
    res["markImportance"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseImportance)
        if err != nil {
            return err
        }
        cast := val.(Importance)
        m.SetMarkImportance(&cast)
        return nil
    }
    res["moveToFolder"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMoveToFolder(val)
        return nil
    }
    res["permanentDelete"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetPermanentDelete(val)
        return nil
    }
    res["redirectTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRecipient() })
        if err != nil {
            return err
        }
        res := make([]Recipient, len(val))
        for i, v := range val {
            res[i] = *(v.(*Recipient))
        }
        m.SetRedirectTo(res)
        return nil
    }
    res["stopProcessingRules"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetStopProcessingRules(val)
        return nil
    }
    return res
}
func (m *MessageRuleActions) IsNil()(bool) {
    return m == nil
}
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
func (m *MessageRuleActions) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *MessageRuleActions) SetAssignCategories(value []string)() {
    m.assignCategories = value
}
func (m *MessageRuleActions) SetCopyToFolder(value *string)() {
    m.copyToFolder = value
}
func (m *MessageRuleActions) SetDelete(value *bool)() {
    m.delete = value
}
func (m *MessageRuleActions) SetForwardAsAttachmentTo(value []Recipient)() {
    m.forwardAsAttachmentTo = value
}
func (m *MessageRuleActions) SetForwardTo(value []Recipient)() {
    m.forwardTo = value
}
func (m *MessageRuleActions) SetMarkAsRead(value *bool)() {
    m.markAsRead = value
}
func (m *MessageRuleActions) SetMarkImportance(value *Importance)() {
    m.markImportance = value
}
func (m *MessageRuleActions) SetMoveToFolder(value *string)() {
    m.moveToFolder = value
}
func (m *MessageRuleActions) SetPermanentDelete(value *bool)() {
    m.permanentDelete = value
}
func (m *MessageRuleActions) SetRedirectTo(value []Recipient)() {
    m.redirectTo = value
}
func (m *MessageRuleActions) SetStopProcessingRules(value *bool)() {
    m.stopProcessingRules = value
}
