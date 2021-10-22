package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type EducationStudent struct {
    additionalData map[string]interface{};
    birthDate *string;
    externalId *string;
    gender *EducationGender;
    grade *string;
    graduationYear *string;
    studentNumber *string;
}
func NewEducationStudent()(*EducationStudent) {
    m := &EducationStudent{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *EducationStudent) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *EducationStudent) GetBirthDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.birthDate
    }
}
func (m *EducationStudent) GetExternalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalId
    }
}
func (m *EducationStudent) GetGender()(*EducationGender) {
    if m == nil {
        return nil
    } else {
        return m.gender
    }
}
func (m *EducationStudent) GetGrade()(*string) {
    if m == nil {
        return nil
    } else {
        return m.grade
    }
}
func (m *EducationStudent) GetGraduationYear()(*string) {
    if m == nil {
        return nil
    } else {
        return m.graduationYear
    }
}
func (m *EducationStudent) GetStudentNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.studentNumber
    }
}
func (m *EducationStudent) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["birthDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetBirthDate(val)
        return nil
    }
    res["externalId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExternalId(val)
        return nil
    }
    res["gender"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseEducationGender)
        if err != nil {
            return err
        }
        cast := val.(EducationGender)
        m.SetGender(&cast)
        return nil
    }
    res["grade"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetGrade(val)
        return nil
    }
    res["graduationYear"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetGraduationYear(val)
        return nil
    }
    res["studentNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetStudentNumber(val)
        return nil
    }
    return res
}
func (m *EducationStudent) IsNil()(bool) {
    return m == nil
}
func (m *EducationStudent) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("birthDate", m.GetBirthDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("externalId", m.GetExternalId())
        if err != nil {
            return err
        }
    }
    if m.GetGender() != nil {
        cast := m.GetGender().String()
        err := writer.WriteStringValue("gender", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("grade", m.GetGrade())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("graduationYear", m.GetGraduationYear())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("studentNumber", m.GetStudentNumber())
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
func (m *EducationStudent) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *EducationStudent) SetBirthDate(value *string)() {
    m.birthDate = value
}
func (m *EducationStudent) SetExternalId(value *string)() {
    m.externalId = value
}
func (m *EducationStudent) SetGender(value *EducationGender)() {
    m.gender = value
}
func (m *EducationStudent) SetGrade(value *string)() {
    m.grade = value
}
func (m *EducationStudent) SetGraduationYear(value *string)() {
    m.graduationYear = value
}
func (m *EducationStudent) SetStudentNumber(value *string)() {
    m.studentNumber = value
}
