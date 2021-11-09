package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new educationStudent and sets the default values.
func NewEducationStudent()(*EducationStudent) {
    m := &EducationStudent{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationStudent) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the birthDate property value. Birth date of the student.
func (m *EducationStudent) GetBirthDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.birthDate
    }
}
// Gets the externalId property value. ID of the student in the source system.
func (m *EducationStudent) GetExternalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalId
    }
}
// Gets the gender property value. The possible values are: female, male, other, unknownFutureValue.
func (m *EducationStudent) GetGender()(*EducationGender) {
    if m == nil {
        return nil
    } else {
        return m.gender
    }
}
// Gets the grade property value. Current grade level of the student.
func (m *EducationStudent) GetGrade()(*string) {
    if m == nil {
        return nil
    } else {
        return m.grade
    }
}
// Gets the graduationYear property value. Year the student is graduating from the school.
func (m *EducationStudent) GetGraduationYear()(*string) {
    if m == nil {
        return nil
    } else {
        return m.graduationYear
    }
}
// Gets the studentNumber property value. Student Number.
func (m *EducationStudent) GetStudentNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.studentNumber
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *EducationStudent) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the birthDate property value. Birth date of the student.
// Parameters:
//  - value : Value to set for the birthDate property.
func (m *EducationStudent) SetBirthDate(value *string)() {
    m.birthDate = value
}
// Sets the externalId property value. ID of the student in the source system.
// Parameters:
//  - value : Value to set for the externalId property.
func (m *EducationStudent) SetExternalId(value *string)() {
    m.externalId = value
}
// Sets the gender property value. The possible values are: female, male, other, unknownFutureValue.
// Parameters:
//  - value : Value to set for the gender property.
func (m *EducationStudent) SetGender(value *EducationGender)() {
    m.gender = value
}
// Sets the grade property value. Current grade level of the student.
// Parameters:
//  - value : Value to set for the grade property.
func (m *EducationStudent) SetGrade(value *string)() {
    m.grade = value
}
// Sets the graduationYear property value. Year the student is graduating from the school.
// Parameters:
//  - value : Value to set for the graduationYear property.
func (m *EducationStudent) SetGraduationYear(value *string)() {
    m.graduationYear = value
}
// Sets the studentNumber property value. Student Number.
// Parameters:
//  - value : Value to set for the studentNumber property.
func (m *EducationStudent) SetStudentNumber(value *string)() {
    m.studentNumber = value
}
