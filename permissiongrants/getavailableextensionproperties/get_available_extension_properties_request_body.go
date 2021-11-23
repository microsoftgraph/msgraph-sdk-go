package getavailableextensionproperties

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// getAvailableExtensionPropertiesRequestBody 
type GetAvailableExtensionPropertiesRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    isSyncedFromOnPremises *bool;
}
// NewGetAvailableExtensionPropertiesRequestBody instantiates a new getAvailableExtensionPropertiesRequestBody and sets the default values.
func NewGetAvailableExtensionPropertiesRequestBody()(*GetAvailableExtensionPropertiesRequestBody) {
    m := &GetAvailableExtensionPropertiesRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetAvailableExtensionPropertiesRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetIsSyncedFromOnPremises gets the isSyncedFromOnPremises property value. 
func (m *GetAvailableExtensionPropertiesRequestBody) GetIsSyncedFromOnPremises()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSyncedFromOnPremises
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetAvailableExtensionPropertiesRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["isSyncedFromOnPremises"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSyncedFromOnPremises(val)
        }
        return nil
    }
    return res
}
func (m *GetAvailableExtensionPropertiesRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GetAvailableExtensionPropertiesRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isSyncedFromOnPremises", m.GetIsSyncedFromOnPremises())
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
func (m *GetAvailableExtensionPropertiesRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetIsSyncedFromOnPremises sets the isSyncedFromOnPremises property value. 
func (m *GetAvailableExtensionPropertiesRequestBody) SetIsSyncedFromOnPremises(value *bool)() {
    m.isSyncedFromOnPremises = value
}
