package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WorkforceIntegrationEncryption struct {
    additionalData map[string]interface{};
    protocol *WorkforceIntegrationEncryptionProtocol;
    secret *string;
}
func NewWorkforceIntegrationEncryption()(*WorkforceIntegrationEncryption) {
    m := &WorkforceIntegrationEncryption{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *WorkforceIntegrationEncryption) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *WorkforceIntegrationEncryption) GetProtocol()(*WorkforceIntegrationEncryptionProtocol) {
    if m == nil {
        return nil
    } else {
        return m.protocol
    }
}
func (m *WorkforceIntegrationEncryption) GetSecret()(*string) {
    if m == nil {
        return nil
    } else {
        return m.secret
    }
}
func (m *WorkforceIntegrationEncryption) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["protocol"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseWorkforceIntegrationEncryptionProtocol)
        if err != nil {
            return err
        }
        cast := val.(WorkforceIntegrationEncryptionProtocol)
        m.SetProtocol(&cast)
        return nil
    }
    res["secret"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSecret(val)
        return nil
    }
    return res
}
func (m *WorkforceIntegrationEncryption) IsNil()(bool) {
    return m == nil
}
func (m *WorkforceIntegrationEncryption) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetProtocol() != nil {
        cast := m.GetProtocol().String()
        err := writer.WriteStringValue("protocol", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("secret", m.GetSecret())
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
func (m *WorkforceIntegrationEncryption) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *WorkforceIntegrationEncryption) SetProtocol(value *WorkforceIntegrationEncryptionProtocol)() {
    m.protocol = value
}
func (m *WorkforceIntegrationEncryption) SetSecret(value *string)() {
    m.secret = value
}
