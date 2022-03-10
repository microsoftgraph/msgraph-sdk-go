package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EducationCourse provides operations to manage the educationRoot singleton.
type EducationCourse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Unique identifier for the course.
    courseNumber *string;
    // Description of the course.
    description *string;
    // Name of the course.
    displayName *string;
    // ID of the course from the syncing system.
    externalId *string;
    // Subject of the course.
    subject *string;
}
// NewEducationCourse instantiates a new educationCourse and sets the default values.
func NewEducationCourse()(*EducationCourse) {
    m := &EducationCourse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateEducationCourseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationCourseFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewEducationCourse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationCourse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCourseNumber gets the courseNumber property value. Unique identifier for the course.
func (m *EducationCourse) GetCourseNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.courseNumber
    }
}
// GetDescription gets the description property value. Description of the course.
func (m *EducationCourse) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. Name of the course.
func (m *EducationCourse) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetExternalId gets the externalId property value. ID of the course from the syncing system.
func (m *EducationCourse) GetExternalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationCourse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["courseNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCourseNumber(val)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
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
    res["subject"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubject(val)
        }
        return nil
    }
    return res
}
// GetSubject gets the subject property value. Subject of the course.
func (m *EducationCourse) GetSubject()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subject
    }
}
func (m *EducationCourse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *EducationCourse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("courseNumber", m.GetCourseNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
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
    {
        err := writer.WriteStringValue("subject", m.GetSubject())
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
func (m *EducationCourse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCourseNumber sets the courseNumber property value. Unique identifier for the course.
func (m *EducationCourse) SetCourseNumber(value *string)() {
    if m != nil {
        m.courseNumber = value
    }
}
// SetDescription sets the description property value. Description of the course.
func (m *EducationCourse) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. Name of the course.
func (m *EducationCourse) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetExternalId sets the externalId property value. ID of the course from the syncing system.
func (m *EducationCourse) SetExternalId(value *string)() {
    if m != nil {
        m.externalId = value
    }
}
// SetSubject sets the subject property value. Subject of the course.
func (m *EducationCourse) SetSubject(value *string)() {
    if m != nil {
        m.subject = value
    }
}
