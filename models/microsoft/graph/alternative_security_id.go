package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AlternativeSecurityId 
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
// NewAlternativeSecurityId instantiates a new alternativeSecurityId and sets the default values.
func NewAlternativeSecurityId()(*AlternativeSecurityId) {
    m := &AlternativeSecurityId{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AlternativeSecurityId) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetIdentityProvider gets the identityProvider property value. For internal use only
func (m *AlternativeSecurityId) GetIdentityProvider()(*string) {
    if m == nil {
        return nil
    } else {
        return m.identityProvider
    }
}
// GetKey gets the key property value. For internal use only
func (m *AlternativeSecurityId) GetKey()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.key
    }
}
// GetType_escaped gets the type property value. For internal use only
func (m *AlternativeSecurityId) GetType_escaped()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AlternativeSecurityId) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["identityProvider"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityProvider(val)
        }
        return nil
    }
    res["key"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val)
        }
        return nil
    }
    res["type"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
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
func (m *AlternativeSecurityId) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
        err := writer.WriteInt32Value("type", m.GetType())
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
func (m *AlternativeSecurityId) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetIdentityProvider sets the identityProvider property value. For internal use only
func (m *AlternativeSecurityId) SetIdentityProvider(value *string)() {
    if m != nil {
        m.identityProvider = value
    }
}
// SetKey sets the key property value. For internal use only
func (m *AlternativeSecurityId) SetKey(value []byte)() {
    if m != nil {
        m.key = value
    }
}
// SetType_escaped sets the type property value. For internal use only
func (m *AlternativeSecurityId) SetType_escaped(value *int32)() {
    if m != nil {
        m.type_escaped = value
    }
}
