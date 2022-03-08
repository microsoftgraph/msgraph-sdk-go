package getomasettingplaintextvaluewithsecretreferencevalueid

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GetOmaSettingPlainTextValueWithSecretReferenceValueIdResponse provides operations to call the getOmaSettingPlainTextValue method.
type GetOmaSettingPlainTextValueWithSecretReferenceValueIdResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    value *string;
}
// NewGetOmaSettingPlainTextValueWithSecretReferenceValueIdResponse instantiates a new getOmaSettingPlainTextValueWithSecretReferenceValueIdResponse and sets the default values.
func NewGetOmaSettingPlainTextValueWithSecretReferenceValueIdResponse()(*GetOmaSettingPlainTextValueWithSecretReferenceValueIdResponse) {
    m := &GetOmaSettingPlainTextValueWithSecretReferenceValueIdResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGetOmaSettingPlainTextValueWithSecretReferenceValueIdResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGetOmaSettingPlainTextValueWithSecretReferenceValueIdResponseFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewGetOmaSettingPlainTextValueWithSecretReferenceValueIdResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetOmaSettingPlainTextValueWithSecretReferenceValueIdResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetOmaSettingPlainTextValueWithSecretReferenceValueIdResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. 
func (m *GetOmaSettingPlainTextValueWithSecretReferenceValueIdResponse) GetValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
func (m *GetOmaSettingPlainTextValueWithSecretReferenceValueIdResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GetOmaSettingPlainTextValueWithSecretReferenceValueIdResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("value", m.GetValue())
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
func (m *GetOmaSettingPlainTextValueWithSecretReferenceValueIdResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetValue sets the value property value. 
func (m *GetOmaSettingPlainTextValueWithSecretReferenceValueIdResponse) SetValue(value *string)() {
    if m != nil {
        m.value = value
    }
}
