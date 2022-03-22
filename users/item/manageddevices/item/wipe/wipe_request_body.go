package wipe

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WipeRequestBody provides operations to call the wipe method.
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
// NewWipeRequestBody instantiates a new wipeRequestBody and sets the default values.
func NewWipeRequestBody()(*WipeRequestBody) {
    m := &WipeRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWipeRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWipeRequestBodyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewWipeRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WipeRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WipeRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["keepEnrollmentData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeepEnrollmentData(val)
        }
        return nil
    }
    res["keepUserData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeepUserData(val)
        }
        return nil
    }
    res["macOsUnlockCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMacOsUnlockCode(val)
        }
        return nil
    }
    res["persistEsimDataPlan"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPersistEsimDataPlan(val)
        }
        return nil
    }
    return res
}
// GetKeepEnrollmentData gets the keepEnrollmentData property value. 
func (m *WipeRequestBody) GetKeepEnrollmentData()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.keepEnrollmentData
    }
}
// GetKeepUserData gets the keepUserData property value. 
func (m *WipeRequestBody) GetKeepUserData()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.keepUserData
    }
}
// GetMacOsUnlockCode gets the macOsUnlockCode property value. 
func (m *WipeRequestBody) GetMacOsUnlockCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.macOsUnlockCode
    }
}
// GetPersistEsimDataPlan gets the persistEsimDataPlan property value. 
func (m *WipeRequestBody) GetPersistEsimDataPlan()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.persistEsimDataPlan
    }
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WipeRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetKeepEnrollmentData sets the keepEnrollmentData property value. 
func (m *WipeRequestBody) SetKeepEnrollmentData(value *bool)() {
    if m != nil {
        m.keepEnrollmentData = value
    }
}
// SetKeepUserData sets the keepUserData property value. 
func (m *WipeRequestBody) SetKeepUserData(value *bool)() {
    if m != nil {
        m.keepUserData = value
    }
}
// SetMacOsUnlockCode sets the macOsUnlockCode property value. 
func (m *WipeRequestBody) SetMacOsUnlockCode(value *string)() {
    if m != nil {
        m.macOsUnlockCode = value
    }
}
// SetPersistEsimDataPlan sets the persistEsimDataPlan property value. 
func (m *WipeRequestBody) SetPersistEsimDataPlan(value *bool)() {
    if m != nil {
        m.persistEsimDataPlan = value
    }
}
