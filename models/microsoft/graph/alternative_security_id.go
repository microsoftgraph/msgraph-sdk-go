package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AlternativeSecurityId struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // For internal use only
    identityProvider *string;
    // For internal use only
    key []byte;
    // For internal use only
    type_escaped *int32;
}
// Instantiates a new alternativeSecurityId and sets the default values.
func NewAlternativeSecurityId()(*AlternativeSecurityId) {
    m := &AlternativeSecurityId{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AlternativeSecurityId) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the identityProvider property value. For internal use only
func (m *AlternativeSecurityId) GetIdentityProvider()(*string) {
    if m == nil {
        return nil
    } else {
        return m.identityProvider
    }
}
// Gets the key property value. For internal use only
func (m *AlternativeSecurityId) GetKey()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.key
    }
}
// Gets the type_escaped property value. For internal use only
func (m *AlternativeSecurityId) GetType_escaped()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// The deserialization information for the current model
func (m *AlternativeSecurityId) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["identityProvider"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetIdentityProvider(val)
        return nil
    }
    res["key"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetKey(val)
        return nil
    }
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetType_escaped(val)
        return nil
    }
    return res
}
func (m *AlternativeSecurityId) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AlternativeSecurityId) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("identityProvider", m.GetIdentityProvider())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteByteArrayValue("key", m.GetKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("type_escaped", m.GetType_escaped())
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
func (m *AlternativeSecurityId) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the identityProvider property value. For internal use only
// Parameters:
//  - value : Value to set for the identityProvider property.
func (m *AlternativeSecurityId) SetIdentityProvider(value *string)() {
    m.identityProvider = value
}
// Sets the key property value. For internal use only
// Parameters:
//  - value : Value to set for the key property.
func (m *AlternativeSecurityId) SetKey(value []byte)() {
    m.key = value
}
// Sets the type_escaped property value. For internal use only
// Parameters:
//  - value : Value to set for the type_escaped property.
func (m *AlternativeSecurityId) SetType_escaped(value *int32)() {
    m.type_escaped = value
}
