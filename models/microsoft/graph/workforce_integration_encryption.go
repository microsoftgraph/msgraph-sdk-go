package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// workforceIntegrationEncryption 
type WorkforceIntegrationEncryption struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Possible values are: sharedSecret, unknownFutureValue.
    protocol *WorkforceIntegrationEncryptionProtocol;
    // Encryption shared secret.
    secret *string;
}
// NewWorkforceIntegrationEncryption instantiates a new workforceIntegrationEncryption and sets the default values.
func NewWorkforceIntegrationEncryption()(*WorkforceIntegrationEncryption) {
    m := &WorkforceIntegrationEncryption{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkforceIntegrationEncryption) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetProtocol gets the protocol property value. Possible values are: sharedSecret, unknownFutureValue.
func (m *WorkforceIntegrationEncryption) GetProtocol()(*WorkforceIntegrationEncryptionProtocol) {
    if m == nil {
        return nil
    } else {
        return m.protocol
    }
}
// GetSecret gets the secret property value. Encryption shared secret.
func (m *WorkforceIntegrationEncryption) GetSecret()(*string) {
    if m == nil {
        return nil
    } else {
        return m.secret
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkforceIntegrationEncryption) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["protocol"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseWorkforceIntegrationEncryptionProtocol)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(WorkforceIntegrationEncryptionProtocol)
            m.SetProtocol(&cast)
        }
        return nil
    }
    res["secret"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecret(val)
        }
        return nil
    }
    return res
}
func (m *WorkforceIntegrationEncryption) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkforceIntegrationEncryption) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetProtocol sets the protocol property value. Possible values are: sharedSecret, unknownFutureValue.
func (m *WorkforceIntegrationEncryption) SetProtocol(value *WorkforceIntegrationEncryptionProtocol)() {
    m.protocol = value
}
// SetSecret sets the secret property value. Encryption shared secret.
func (m *WorkforceIntegrationEncryption) SetSecret(value *string)() {
    m.secret = value
}
