package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type EducationAssignmentSettings struct {
    Entity
    // Indicates whether turn-in celebration animation will be shown. A value of true indicates that the animation will not be shown. Default value is false.
    submissionAnimationDisabled *bool;
}
// Instantiates a new educationAssignmentSettings and sets the default values.
func NewEducationAssignmentSettings()(*EducationAssignmentSettings) {
    m := &EducationAssignmentSettings{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the submissionAnimationDisabled property value. Indicates whether turn-in celebration animation will be shown. A value of true indicates that the animation will not be shown. Default value is false.
func (m *EducationAssignmentSettings) GetSubmissionAnimationDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.submissionAnimationDisabled
    }
}
// The deserialization information for the current model
func (m *EducationAssignmentSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["submissionAnimationDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetSubmissionAnimationDisabled(val)
        return nil
    }
    return res
}
func (m *EducationAssignmentSettings) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *EducationAssignmentSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("submissionAnimationDisabled", m.GetSubmissionAnimationDisabled())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the submissionAnimationDisabled property value. Indicates whether turn-in celebration animation will be shown. A value of true indicates that the animation will not be shown. Default value is false.
// Parameters:
//  - value : Value to set for the submissionAnimationDisabled property.
func (m *EducationAssignmentSettings) SetSubmissionAnimationDisabled(value *bool)() {
    m.submissionAnimationDisabled = value
}
