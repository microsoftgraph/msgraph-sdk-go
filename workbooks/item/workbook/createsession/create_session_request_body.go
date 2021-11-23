package createsession

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// createSessionRequestBody 
type CreateSessionRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    persistChanges *bool;
}
// NewCreateSessionRequestBody instantiates a new createSessionRequestBody and sets the default values.
func NewCreateSessionRequestBody()(*CreateSessionRequestBody) {
    m := &CreateSessionRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CreateSessionRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetPersistChanges gets the persistChanges property value. 
func (m *CreateSessionRequestBody) GetPersistChanges()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.persistChanges
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CreateSessionRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["persistChanges"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPersistChanges(val)
        }
        return nil
    }
    return res
}
func (m *CreateSessionRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CreateSessionRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("persistChanges", m.GetPersistChanges())
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
func (m *CreateSessionRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetPersistChanges sets the persistChanges property value. 
func (m *CreateSessionRequestBody) SetPersistChanges(value *bool)() {
    m.persistChanges = value
}
