package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ImplicitGrantSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Specifies whether this web application can request an access token using the OAuth 2.0 implicit flow.
    enableAccessTokenIssuance *bool;
    // Specifies whether this web application can request an ID token using the OAuth 2.0 implicit flow.
    enableIdTokenIssuance *bool;
}
// Instantiates a new implicitGrantSettings and sets the default values.
func NewImplicitGrantSettings()(*ImplicitGrantSettings) {
    m := &ImplicitGrantSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ImplicitGrantSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the enableAccessTokenIssuance property value. Specifies whether this web application can request an access token using the OAuth 2.0 implicit flow.
func (m *ImplicitGrantSettings) GetEnableAccessTokenIssuance()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enableAccessTokenIssuance
    }
}
// Gets the enableIdTokenIssuance property value. Specifies whether this web application can request an ID token using the OAuth 2.0 implicit flow.
func (m *ImplicitGrantSettings) GetEnableIdTokenIssuance()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enableIdTokenIssuance
    }
}
// The deserialization information for the current model
func (m *ImplicitGrantSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["enableAccessTokenIssuance"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableAccessTokenIssuance(val)
        }
        return nil
    }
    res["enableIdTokenIssuance"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableIdTokenIssuance(val)
        }
        return nil
    }
    return res
}
func (m *ImplicitGrantSettings) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ImplicitGrantSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("enableAccessTokenIssuance", m.GetEnableAccessTokenIssuance())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enableIdTokenIssuance", m.GetEnableIdTokenIssuance())
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
func (m *ImplicitGrantSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the enableAccessTokenIssuance property value. Specifies whether this web application can request an access token using the OAuth 2.0 implicit flow.
// Parameters:
//  - value : Value to set for the enableAccessTokenIssuance property.
func (m *ImplicitGrantSettings) SetEnableAccessTokenIssuance(value *bool)() {
    m.enableAccessTokenIssuance = value
}
// Sets the enableIdTokenIssuance property value. Specifies whether this web application can request an ID token using the OAuth 2.0 implicit flow.
// Parameters:
//  - value : Value to set for the enableIdTokenIssuance property.
func (m *ImplicitGrantSettings) SetEnableIdTokenIssuance(value *bool)() {
    m.enableIdTokenIssuance = value
}
