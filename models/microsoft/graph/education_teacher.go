package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type EducationTeacher struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // ID of the teacher in the source system.
    externalId *string;
    // Teacher number.
    teacherNumber *string;
}
// Instantiates a new educationTeacher and sets the default values.
func NewEducationTeacher()(*EducationTeacher) {
    m := &EducationTeacher{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationTeacher) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the externalId property value. ID of the teacher in the source system.
func (m *EducationTeacher) GetExternalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalId
    }
}
// Gets the teacherNumber property value. Teacher number.
func (m *EducationTeacher) GetTeacherNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.teacherNumber
    }
}
// The deserialization information for the current model
func (m *EducationTeacher) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
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
    res["teacherNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeacherNumber(val)
        }
        return nil
    }
    return res
}
func (m *EducationTeacher) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *EducationTeacher) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the externalId property value. ID of the teacher in the source system.
// Parameters:
//  - value : Value to set for the externalId property.
func (m *EducationTeacher) SetExternalId(value *string)() {
    m.externalId = value
}
// Sets the teacherNumber property value. Teacher number.
// Parameters:
//  - value : Value to set for the teacherNumber property.
func (m *EducationTeacher) SetTeacherNumber(value *string)() {
    m.teacherNumber = value
}
