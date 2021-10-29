package wipe

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type WipeRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    keepEnrollmentData *bool;
    // 
    keepUserData *bool;
    // 
    macOsUnlockCode *string;
    // 
    persistEsimDataPlan *bool;
}
// Instantiates a new wipeRequestBody and sets the default values.
func NewWipeRequestBody()(*WipeRequestBody) {
    m := &WipeRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WipeRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the keepEnrollmentData property value. 
func (m *WipeRequestBody) GetKeepEnrollmentData()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.keepEnrollmentData
    }
}
// Gets the keepUserData property value. 
func (m *WipeRequestBody) GetKeepUserData()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.keepUserData
    }
}
// Gets the macOsUnlockCode property value. 
func (m *WipeRequestBody) GetMacOsUnlockCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.macOsUnlockCode
    }
}
// Gets the persistEsimDataPlan property value. 
func (m *WipeRequestBody) GetPersistEsimDataPlan()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.persistEsimDataPlan
    }
}
// The deserialization information for the current model
func (m *WipeRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["keepEnrollmentData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetKeepEnrollmentData(val)
        return nil
    }
    res["keepUserData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetKeepUserData(val)
        return nil
    }
    res["macOsUnlockCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMacOsUnlockCode(val)
        return nil
    }
    res["persistEsimDataPlan"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetPersistEsimDataPlan(val)
        return nil
    }
    return res
}
func (m *WipeRequestBody) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *WipeRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("keepEnrollmentData", m.GetKeepEnrollmentData())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("keepUserData", m.GetKeepUserData())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("macOsUnlockCode", m.GetMacOsUnlockCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("persistEsimDataPlan", m.GetPersistEsimDataPlan())
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
func (m *WipeRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the keepEnrollmentData property value. 
// Parameters:
//  - value : Value to set for the keepEnrollmentData property.
func (m *WipeRequestBody) SetKeepEnrollmentData(value *bool)() {
    m.keepEnrollmentData = value
}
// Sets the keepUserData property value. 
// Parameters:
//  - value : Value to set for the keepUserData property.
func (m *WipeRequestBody) SetKeepUserData(value *bool)() {
    m.keepUserData = value
}
// Sets the macOsUnlockCode property value. 
// Parameters:
//  - value : Value to set for the macOsUnlockCode property.
func (m *WipeRequestBody) SetMacOsUnlockCode(value *string)() {
    m.macOsUnlockCode = value
}
// Sets the persistEsimDataPlan property value. 
// Parameters:
//  - value : Value to set for the persistEsimDataPlan property.
func (m *WipeRequestBody) SetPersistEsimDataPlan(value *bool)() {
    m.persistEsimDataPlan = value
}
