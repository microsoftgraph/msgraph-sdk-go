package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EducationStudent 
type EducationStudent struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Birth date of the student.
    birthDate *string;
    // ID of the student in the source system.
    externalId *string;
    // The possible values are: female, male, other, unknownFutureValue.
    gender *EducationGender;
    // Current grade level of the student.
    grade *string;
    // Year the student is graduating from the school.
    graduationYear *string;
    // Student Number.
    studentNumber *string;
}
// NewEducationStudent instantiates a new educationStudent and sets the default values.
func NewEducationStudent()(*EducationStudent) {
    m := &EducationStudent{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationStudent) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetBirthDate gets the birthDate property value. Birth date of the student.
func (m *EducationStudent) GetBirthDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.birthDate
    }
}
// GetExternalId gets the externalId property value. ID of the student in the source system.
func (m *EducationStudent) GetExternalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalId
    }
}
// GetGender gets the gender property value. The possible values are: female, male, other, unknownFutureValue.
func (m *EducationStudent) GetGender()(*EducationGender) {
    if m == nil {
        return nil
    } else {
        return m.gender
    }
}
// GetGrade gets the grade property value. Current grade level of the student.
func (m *EducationStudent) GetGrade()(*string) {
    if m == nil {
        return nil
    } else {
        return m.grade
    }
}
// GetGraduationYear gets the graduationYear property value. Year the student is graduating from the school.
func (m *EducationStudent) GetGraduationYear()(*string) {
    if m == nil {
        return nil
    } else {
        return m.graduationYear
    }
}
// GetStudentNumber gets the studentNumber property value. Student Number.
func (m *EducationStudent) GetStudentNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.studentNumber
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationStudent) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["birthDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBirthDate(val)
        }
        return nil
    }
    res["externalId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalId(val)
        }
        return nil
    }
    res["gender"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseEducationGender)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(EducationGender)
            m.SetGender(&cast)
        }
        return nil
    }
    res["grade"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGrade(val)
        }
        return nil
    }
    res["graduationYear"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGraduationYear(val)
        }
        return nil
    }
    res["studentNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStudentNumber(val)
        }
        return nil
    }
    return res
}
func (m *EducationStudent) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationStudent) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetBirthDate sets the birthDate property value. Birth date of the student.
func (m *EducationStudent) SetBirthDate(value *string)() {
    m.birthDate = value
}
// SetExternalId sets the externalId property value. ID of the student in the source system.
func (m *EducationStudent) SetExternalId(value *string)() {
    m.externalId = value
}
// SetGender sets the gender property value. The possible values are: female, male, other, unknownFutureValue.
func (m *EducationStudent) SetGender(value *EducationGender)() {
    m.gender = value
}
// SetGrade sets the grade property value. Current grade level of the student.
func (m *EducationStudent) SetGrade(value *string)() {
    m.grade = value
}
// SetGraduationYear sets the graduationYear property value. Year the student is graduating from the school.
func (m *EducationStudent) SetGraduationYear(value *string)() {
    m.graduationYear = value
}
// SetStudentNumber sets the studentNumber property value. Student Number.
func (m *EducationStudent) SetStudentNumber(value *string)() {
    m.studentNumber = value
}
