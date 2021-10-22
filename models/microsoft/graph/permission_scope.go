package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PermissionScope struct {
    additionalData map[string]interface{};
    adminConsentDescription *string;
    adminConsentDisplayName *string;
    id *string;
    isEnabled *bool;
    origin *string;
    type_escpaped *string;
    userConsentDescription *string;
    userConsentDisplayName *string;
    value *string;
}
func NewPermissionScope()(*PermissionScope) {
    m := &PermissionScope{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *PermissionScope) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *PermissionScope) GetAdminConsentDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.adminConsentDescription
    }
}
func (m *PermissionScope) GetAdminConsentDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.adminConsentDisplayName
    }
}
func (m *PermissionScope) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
func (m *PermissionScope) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
func (m *PermissionScope) GetOrigin()(*string) {
    if m == nil {
        return nil
    } else {
        return m.origin
    }
}
func (m *PermissionScope) GetType_escpaped()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escpaped
    }
}
func (m *PermissionScope) GetUserConsentDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userConsentDescription
    }
}
func (m *PermissionScope) GetUserConsentDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userConsentDisplayName
    }
}
func (m *PermissionScope) GetValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
func (m *PermissionScope) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["adminConsentDescription"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAdminConsentDescription(val)
        return nil
    }
    res["adminConsentDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAdminConsentDisplayName(val)
        return nil
    }
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetId(val)
        return nil
    }
    res["isEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsEnabled(val)
        return nil
    }
    res["origin"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOrigin(val)
        return nil
    }
    res["type_escpaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetType_escpaped(val)
        return nil
    }
    res["userConsentDescription"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserConsentDescription(val)
        return nil
    }
    res["userConsentDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserConsentDisplayName(val)
        return nil
    }
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetValue(val)
        return nil
    }
    return res
}
func (m *PermissionScope) IsNil()(bool) {
    return m == nil
}
func (m *PermissionScope) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("adminConsentDescription", m.GetAdminConsentDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("adminConsentDisplayName", m.GetAdminConsentDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("origin", m.GetOrigin())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type_escpaped", m.GetType_escpaped())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userConsentDescription", m.GetUserConsentDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userConsentDisplayName", m.GetUserConsentDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("value", m.GetValue())
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
func (m *PermissionScope) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *PermissionScope) SetAdminConsentDescription(value *string)() {
    m.adminConsentDescription = value
}
func (m *PermissionScope) SetAdminConsentDisplayName(value *string)() {
    m.adminConsentDisplayName = value
}
func (m *PermissionScope) SetId(value *string)() {
    m.id = value
}
func (m *PermissionScope) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
func (m *PermissionScope) SetOrigin(value *string)() {
    m.origin = value
}
func (m *PermissionScope) SetType_escpaped(value *string)() {
    m.type_escpaped = value
}
func (m *PermissionScope) SetUserConsentDescription(value *string)() {
    m.userConsentDescription = value
}
func (m *PermissionScope) SetUserConsentDisplayName(value *string)() {
    m.userConsentDisplayName = value
}
func (m *PermissionScope) SetValue(value *string)() {
    m.value = value
}
