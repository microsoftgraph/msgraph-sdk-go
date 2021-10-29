package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new verifiedDomain and sets the default values.
func NewVerifiedDomain()(*VerifiedDomain) {
    m := &VerifiedDomain{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VerifiedDomain) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the capabilities property value. For example, 'Email', 'OfficeCommunicationsOnline'.
func (m *VerifiedDomain) GetCapabilities()(*string) {
    if m == nil {
        return nil
    } else {
        return m.capabilities
    }
}
// Gets the isDefault property value. true if this is the default domain associated with the tenant; otherwise, false.
func (m *VerifiedDomain) GetIsDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDefault
    }
}
// Gets the isInitial property value. true if this is the initial domain associated with the tenant; otherwise, false
func (m *VerifiedDomain) GetIsInitial()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isInitial
    }
}
// Gets the name property value. The domain name; for example, 'contoso.onmicrosoft.com'
func (m *VerifiedDomain) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the type_escaped property value. For example, 'Managed'.
func (m *VerifiedDomain) GetType_escaped()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// The deserialization information for the current model
func (m *VerifiedDomain) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["capabilities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCapabilities(val)
        return nil
    }
    res["isDefault"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsDefault(val)
        return nil
    }
    res["isInitial"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsInitial(val)
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetType_escaped(val)
        return nil
    }
    return res
}
func (m *VerifiedDomain) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
        err := writer.WriteStringValue("type_escaped", m.GetType_escaped())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *VerifiedDomain) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the capabilities property value. For example, 'Email', 'OfficeCommunicationsOnline'.
// Parameters:
//  - value : Value to set for the capabilities property.
func (m *VerifiedDomain) SetCapabilities(value *string)() {
    m.capabilities = value
}
// Sets the isDefault property value. true if this is the default domain associated with the tenant; otherwise, false.
// Parameters:
//  - value : Value to set for the isDefault property.
func (m *VerifiedDomain) SetIsDefault(value *bool)() {
    m.isDefault = value
}
// Sets the isInitial property value. true if this is the initial domain associated with the tenant; otherwise, false
// Parameters:
//  - value : Value to set for the isInitial property.
func (m *VerifiedDomain) SetIsInitial(value *bool)() {
    m.isInitial = value
}
// Sets the name property value. The domain name; for example, 'contoso.onmicrosoft.com'
// Parameters:
//  - value : Value to set for the name property.
func (m *VerifiedDomain) SetName(value *string)() {
    m.name = value
}
// Sets the type_escaped property value. For example, 'Managed'.
// Parameters:
//  - value : Value to set for the type_escaped property.
func (m *VerifiedDomain) SetType_escaped(value *string)() {
    m.type_escaped = value
}
