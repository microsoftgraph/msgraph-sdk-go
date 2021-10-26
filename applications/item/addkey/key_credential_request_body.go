package addkey

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// 
type KeyCredentialRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    keyCredential *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.KeyCredential;
    // 
    passwordCredential *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PasswordCredential;
    // 
    proof *string;
}
// Instantiates a new KeyCredentialRequestBody and sets the default values.
func NewKeyCredentialRequestBody()(*KeyCredentialRequestBody) {
    m := &KeyCredentialRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *KeyCredentialRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the keyCredential property value. 
func (m *KeyCredentialRequestBody) GetKeyCredential()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.KeyCredential) {
    if m == nil {
        return nil
    } else {
        return m.keyCredential
    }
}
// Gets the passwordCredential property value. 
func (m *KeyCredentialRequestBody) GetPasswordCredential()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PasswordCredential) {
    if m == nil {
        return nil
    } else {
        return m.passwordCredential
    }
}
// Gets the proof property value. 
func (m *KeyCredentialRequestBody) GetProof()(*string) {
    if m == nil {
        return nil
    } else {
        return m.proof
    }
}
// The deserialization information for the current model
func (m *KeyCredentialRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["keyCredential"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewKeyCredential() })
        if err != nil {
            return err
        }
        m.SetKeyCredential(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.KeyCredential))
        return nil
    }
    res["passwordCredential"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewPasswordCredential() })
        if err != nil {
            return err
        }
        m.SetPasswordCredential(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PasswordCredential))
        return nil
    }
    res["proof"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetProof(val)
        return nil
    }
    return res
}
func (m *KeyCredentialRequestBody) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *KeyCredentialRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("keyCredential", m.GetKeyCredential())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("passwordCredential", m.GetPasswordCredential())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("proof", m.GetProof())
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
func (m *KeyCredentialRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the keyCredential property value. 
// Parameters:
//  - value : Value to set for the keyCredential property.
func (m *KeyCredentialRequestBody) SetKeyCredential(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.KeyCredential)() {
    m.keyCredential = value
}
// Sets the passwordCredential property value. 
// Parameters:
//  - value : Value to set for the passwordCredential property.
func (m *KeyCredentialRequestBody) SetPasswordCredential(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PasswordCredential)() {
    m.passwordCredential = value
}
// Sets the proof property value. 
// Parameters:
//  - value : Value to set for the proof property.
func (m *KeyCredentialRequestBody) SetProof(value *string)() {
    m.proof = value
}
