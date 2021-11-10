package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type WorkforceIntegrationEncryption struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Possible values are: sharedSecret, unknownFutureValue.
    protocol *WorkforceIntegrationEncryptionProtocol;
    // Encryption shared secret.
    secret *string;
}
// Instantiates a new workforceIntegrationEncryption and sets the default values.
func NewWorkforceIntegrationEncryption()(*WorkforceIntegrationEncryption) {
    m := &WorkforceIntegrationEncryption{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkforceIntegrationEncryption) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the protocol property value. Possible values are: sharedSecret, unknownFutureValue.
func (m *WorkforceIntegrationEncryption) GetProtocol()(*WorkforceIntegrationEncryptionProtocol) {
    if m == nil {
        return nil
    } else {
        return m.protocol
    }
}
// Gets the secret property value. Encryption shared secret.
func (m *WorkforceIntegrationEncryption) GetSecret()(*string) {
    if m == nil {
        return nil
    } else {
        return m.secret
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *WorkforceIntegrationEncryption) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the protocol property value. Possible values are: sharedSecret, unknownFutureValue.
// Parameters:
//  - value : Value to set for the protocol property.
func (m *WorkforceIntegrationEncryption) SetProtocol(value *WorkforceIntegrationEncryptionProtocol)() {
    m.protocol = value
}
// Sets the secret property value. Encryption shared secret.
// Parameters:
//  - value : Value to set for the secret property.
func (m *WorkforceIntegrationEncryption) SetSecret(value *string)() {
    m.secret = value
}
