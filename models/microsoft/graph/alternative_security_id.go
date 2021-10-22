package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AlternativeSecurityId struct {
    additionalData map[string]interface{};
    identityProvider *string;
    key []byte;
    type_escpaped *int32;
}
func NewAlternativeSecurityId()(*AlternativeSecurityId) {
    m := &AlternativeSecurityId{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AlternativeSecurityId) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AlternativeSecurityId) GetIdentityProvider()(*string) {
    if m == nil {
        return nil
    } else {
        return m.identityProvider
    }
}
func (m *AlternativeSecurityId) GetKey()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.key
    }
}
func (m *AlternativeSecurityId) GetType_escpaped()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.type_escpaped
    }
}
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
    res["type_escpaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetType_escpaped(val)
        return nil
    }
    return res
}
func (m *AlternativeSecurityId) IsNil()(bool) {
    return m == nil
}
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
        err := writer.WriteInt32Value("type_escpaped", m.GetType_escpaped())
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
func (m *AlternativeSecurityId) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AlternativeSecurityId) SetIdentityProvider(value *string)() {
    m.identityProvider = value
}
func (m *AlternativeSecurityId) SetKey(value []byte)() {
    m.key = value
}
func (m *AlternativeSecurityId) SetType_escpaped(value *int32)() {
    m.type_escpaped = value
}
