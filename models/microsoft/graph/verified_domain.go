package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// VerifiedDomain 
type VerifiedDomain struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // For example, 'Email', 'OfficeCommunicationsOnline'.
    capabilities *string;
    // true if this is the default domain associated with the tenant; otherwise, false.
    isDefault *bool;
    // true if this is the initial domain associated with the tenant; otherwise, false
    isInitial *bool;
    // The domain name; for example, 'contoso.onmicrosoft.com'
    name *string;
    // For example, 'Managed'.
    type_escaped *string;
}
// NewVerifiedDomain instantiates a new verifiedDomain and sets the default values.
func NewVerifiedDomain()(*VerifiedDomain) {
    m := &VerifiedDomain{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VerifiedDomain) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCapabilities gets the capabilities property value. For example, 'Email', 'OfficeCommunicationsOnline'.
func (m *VerifiedDomain) GetCapabilities()(*string) {
    if m == nil {
        return nil
    } else {
        return m.capabilities
    }
}
// GetIsDefault gets the isDefault property value. true if this is the default domain associated with the tenant; otherwise, false.
func (m *VerifiedDomain) GetIsDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDefault
    }
}
// GetIsInitial gets the isInitial property value. true if this is the initial domain associated with the tenant; otherwise, false
func (m *VerifiedDomain) GetIsInitial()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isInitial
    }
}
// GetName gets the name property value. The domain name; for example, 'contoso.onmicrosoft.com'
func (m *VerifiedDomain) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetType gets the type property value. For example, 'Managed'.
func (m *VerifiedDomain) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VerifiedDomain) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["capabilities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCapabilities(val)
        }
        return nil
    }
    res["isDefault"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefault(val)
        }
        return nil
    }
    res["isInitial"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsInitial(val)
        }
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["type"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    return res
}
func (m *VerifiedDomain) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *VerifiedDomain) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("capabilities", m.GetCapabilities())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isDefault", m.GetIsDefault())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isInitial", m.GetIsInitial())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetType())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VerifiedDomain) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCapabilities sets the capabilities property value. For example, 'Email', 'OfficeCommunicationsOnline'.
func (m *VerifiedDomain) SetCapabilities(value *string)() {
    if m != nil {
        m.capabilities = value
    }
}
// SetIsDefault sets the isDefault property value. true if this is the default domain associated with the tenant; otherwise, false.
func (m *VerifiedDomain) SetIsDefault(value *bool)() {
    if m != nil {
        m.isDefault = value
    }
}
// SetIsInitial sets the isInitial property value. true if this is the initial domain associated with the tenant; otherwise, false
func (m *VerifiedDomain) SetIsInitial(value *bool)() {
    if m != nil {
        m.isInitial = value
    }
}
// SetName sets the name property value. The domain name; for example, 'contoso.onmicrosoft.com'
func (m *VerifiedDomain) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetType sets the type property value. For example, 'Managed'.
func (m *VerifiedDomain) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
