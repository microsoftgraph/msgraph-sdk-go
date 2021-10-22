package wipe

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WipeRequestBody struct {
    additionalData map[string]interface{};
    keepEnrollmentData *bool;
    keepUserData *bool;
    macOsUnlockCode *string;
    persistEsimDataPlan *bool;
}
func NewWipeRequestBody()(*WipeRequestBody) {
    m := &WipeRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *WipeRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *WipeRequestBody) GetKeepEnrollmentData()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.keepEnrollmentData
    }
}
func (m *WipeRequestBody) GetKeepUserData()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.keepUserData
    }
}
func (m *WipeRequestBody) GetMacOsUnlockCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.macOsUnlockCode
    }
}
func (m *WipeRequestBody) GetPersistEsimDataPlan()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.persistEsimDataPlan
    }
}
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
func (m *WipeRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *WipeRequestBody) SetKeepEnrollmentData(value *bool)() {
    m.keepEnrollmentData = value
}
func (m *WipeRequestBody) SetKeepUserData(value *bool)() {
    m.keepUserData = value
}
func (m *WipeRequestBody) SetMacOsUnlockCode(value *string)() {
    m.macOsUnlockCode = value
}
func (m *WipeRequestBody) SetPersistEsimDataPlan(value *bool)() {
    m.persistEsimDataPlan = value
}
