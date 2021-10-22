package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type EducationOnPremisesInfo struct {
    additionalData map[string]interface{};
    immutableId *string;
}
func NewEducationOnPremisesInfo()(*EducationOnPremisesInfo) {
    m := &EducationOnPremisesInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *EducationOnPremisesInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *EducationOnPremisesInfo) GetImmutableId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.immutableId
    }
}
func (m *EducationOnPremisesInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["immutableId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetImmutableId(val)
        return nil
    }
    return res
}
func (m *EducationOnPremisesInfo) IsNil()(bool) {
    return m == nil
}
func (m *EducationOnPremisesInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("immutableId", m.GetImmutableId())
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
func (m *EducationOnPremisesInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *EducationOnPremisesInfo) SetImmutableId(value *string)() {
    m.immutableId = value
}
