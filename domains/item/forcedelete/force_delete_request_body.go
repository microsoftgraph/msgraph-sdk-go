package forcedelete

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ForceDeleteRequestBody 
type ForceDeleteRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    disableUserAccounts *bool;
}
// NewForceDeleteRequestBody instantiates a new forceDeleteRequestBody and sets the default values.
func NewForceDeleteRequestBody()(*ForceDeleteRequestBody) {
    m := &ForceDeleteRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ForceDeleteRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDisableUserAccounts gets the disableUserAccounts property value. 
func (m *ForceDeleteRequestBody) GetDisableUserAccounts()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.disableUserAccounts
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ForceDeleteRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["disableUserAccounts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisableUserAccounts(val)
        }
        return nil
    }
    return res
}
func (m *ForceDeleteRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ForceDeleteRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("disableUserAccounts", m.GetDisableUserAccounts())
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
func (m *ForceDeleteRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDisableUserAccounts sets the disableUserAccounts property value. 
func (m *ForceDeleteRequestBody) SetDisableUserAccounts(value *bool)() {
    if m != nil {
        m.disableUserAccounts = value
    }
}
