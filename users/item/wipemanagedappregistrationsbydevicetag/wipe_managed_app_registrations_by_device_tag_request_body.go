package wipemanagedappregistrationsbydevicetag

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WipeManagedAppRegistrationsByDeviceTagRequestBody provides operations to call the wipeManagedAppRegistrationsByDeviceTag method.
type WipeManagedAppRegistrationsByDeviceTagRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    deviceTag *string;
}
// NewWipeManagedAppRegistrationsByDeviceTagRequestBody instantiates a new wipeManagedAppRegistrationsByDeviceTagRequestBody and sets the default values.
func NewWipeManagedAppRegistrationsByDeviceTagRequestBody()(*WipeManagedAppRegistrationsByDeviceTagRequestBody) {
    m := &WipeManagedAppRegistrationsByDeviceTagRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWipeManagedAppRegistrationsByDeviceTagRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWipeManagedAppRegistrationsByDeviceTagRequestBodyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewWipeManagedAppRegistrationsByDeviceTagRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WipeManagedAppRegistrationsByDeviceTagRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDeviceTag gets the deviceTag property value. 
func (m *WipeManagedAppRegistrationsByDeviceTagRequestBody) GetDeviceTag()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceTag
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WipeManagedAppRegistrationsByDeviceTagRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["deviceTag"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceTag(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *WipeManagedAppRegistrationsByDeviceTagRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("deviceTag", m.GetDeviceTag())
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
func (m *WipeManagedAppRegistrationsByDeviceTagRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDeviceTag sets the deviceTag property value. 
func (m *WipeManagedAppRegistrationsByDeviceTagRequestBody) SetDeviceTag(value *string)() {
    if m != nil {
        m.deviceTag = value
    }
}
