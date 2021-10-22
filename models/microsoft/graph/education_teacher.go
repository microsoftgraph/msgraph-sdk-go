package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type EducationTeacher struct {
    additionalData map[string]interface{};
    externalId *string;
    teacherNumber *string;
}
func NewEducationTeacher()(*EducationTeacher) {
    m := &EducationTeacher{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *EducationTeacher) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *EducationTeacher) GetExternalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalId
    }
}
func (m *EducationTeacher) GetTeacherNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.teacherNumber
    }
}
func (m *EducationTeacher) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["externalId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExternalId(val)
        return nil
    }
    res["teacherNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTeacherNumber(val)
        return nil
    }
    return res
}
func (m *EducationTeacher) IsNil()(bool) {
    return m == nil
}
func (m *EducationTeacher) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("externalId", m.GetExternalId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("teacherNumber", m.GetTeacherNumber())
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
func (m *EducationTeacher) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *EducationTeacher) SetExternalId(value *string)() {
    m.externalId = value
}
func (m *EducationTeacher) SetTeacherNumber(value *string)() {
    m.teacherNumber = value
}
