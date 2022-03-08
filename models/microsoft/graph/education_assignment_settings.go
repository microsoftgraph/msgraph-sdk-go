package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EducationAssignmentSettings provides operations to manage the educationRoot singleton.
type EducationAssignmentSettings struct {
    Entity
    // Indicates whether turn-in celebration animation will be shown. A value of true indicates that the animation will not be shown. Default value is false.
    submissionAnimationDisabled *bool;
}
// NewEducationAssignmentSettings instantiates a new educationAssignmentSettings and sets the default values.
func NewEducationAssignmentSettings()(*EducationAssignmentSettings) {
    m := &EducationAssignmentSettings{
        Entity: *NewEntity(),
    }
    return m
}
// CreateEducationAssignmentSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationAssignmentSettingsFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewEducationAssignmentSettings(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationAssignmentSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["submissionAnimationDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubmissionAnimationDisabled(val)
        }
        return nil
    }
    return res
}
// GetSubmissionAnimationDisabled gets the submissionAnimationDisabled property value. Indicates whether turn-in celebration animation will be shown. A value of true indicates that the animation will not be shown. Default value is false.
func (m *EducationAssignmentSettings) GetSubmissionAnimationDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.submissionAnimationDisabled
    }
}
func (m *EducationAssignmentSettings) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetSubmissionAnimationDisabled sets the submissionAnimationDisabled property value. Indicates whether turn-in celebration animation will be shown. A value of true indicates that the animation will not be shown. Default value is false.
func (m *EducationAssignmentSettings) SetSubmissionAnimationDisabled(value *bool)() {
    if m != nil {
        m.submissionAnimationDisabled = value
    }
}
