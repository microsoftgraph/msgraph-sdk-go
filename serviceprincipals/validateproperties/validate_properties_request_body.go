package validateproperties

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ValidatePropertiesRequestBody struct {
    additionalData map[string]interface{};
    displayName *string;
    entityType *string;
    mailNickname *string;
    onBehalfOfUserId *string;
}
func NewValidatePropertiesRequestBody()(*ValidatePropertiesRequestBody) {
    m := &ValidatePropertiesRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ValidatePropertiesRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ValidatePropertiesRequestBody) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *ValidatePropertiesRequestBody) GetEntityType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.entityType
    }
}
func (m *ValidatePropertiesRequestBody) GetMailNickname()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mailNickname
    }
}
func (m *ValidatePropertiesRequestBody) GetOnBehalfOfUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onBehalfOfUserId
    }
}
func (m *ValidatePropertiesRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["entityType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEntityType(val)
        return nil
    }
    res["mailNickname"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMailNickname(val)
        return nil
    }
    res["onBehalfOfUserId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOnBehalfOfUserId(val)
        return nil
    }
    return res
}
func (m *ValidatePropertiesRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *ValidatePropertiesRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("entityType", m.GetEntityType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("mailNickname", m.GetMailNickname())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("onBehalfOfUserId", m.GetOnBehalfOfUserId())
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
func (m *ValidatePropertiesRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ValidatePropertiesRequestBody) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *ValidatePropertiesRequestBody) SetEntityType(value *string)() {
    m.entityType = value
}
func (m *ValidatePropertiesRequestBody) SetMailNickname(value *string)() {
    m.mailNickname = value
}
func (m *ValidatePropertiesRequestBody) SetOnBehalfOfUserId(value *string)() {
    m.onBehalfOfUserId = value
}
