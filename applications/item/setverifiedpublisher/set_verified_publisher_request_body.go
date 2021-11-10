package setverifiedpublisher

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SetVerifiedPublisherRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    verifiedPublisherId *string;
}
// Instantiates a new setVerifiedPublisherRequestBody and sets the default values.
func NewSetVerifiedPublisherRequestBody()(*SetVerifiedPublisherRequestBody) {
    m := &SetVerifiedPublisherRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SetVerifiedPublisherRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the verifiedPublisherId property value. 
func (m *SetVerifiedPublisherRequestBody) GetVerifiedPublisherId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.verifiedPublisherId
    }
}
// The deserialization information for the current model
func (m *SetVerifiedPublisherRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["verifiedPublisherId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerifiedPublisherId(val)
        }
        return nil
    }
    return res
}
func (m *SetVerifiedPublisherRequestBody) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SetVerifiedPublisherRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("verifiedPublisherId", m.GetVerifiedPublisherId())
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
func (m *SetVerifiedPublisherRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the verifiedPublisherId property value. 
// Parameters:
//  - value : Value to set for the verifiedPublisherId property.
func (m *SetVerifiedPublisherRequestBody) SetVerifiedPublisherId(value *string)() {
    m.verifiedPublisherId = value
}
