package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type EducationAssignmentSettings struct {
    Entity
    submissionAnimationDisabled *bool;
}
func NewEducationAssignmentSettings()(*EducationAssignmentSettings) {
    m := &EducationAssignmentSettings{
        Entity: *NewEntity(),
    }
    return m
}
func (m *EducationAssignmentSettings) GetSubmissionAnimationDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.submissionAnimationDisabled
    }
}
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
func (m *EducationAssignmentSettings) SetSubmissionAnimationDisabled(value *bool)() {
    m.submissionAnimationDisabled = value
}
