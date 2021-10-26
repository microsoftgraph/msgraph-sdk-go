package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type EducationAssignmentResource struct {
    Entity
    // Indicates whether this resource should be copied to each student submission for modification and submission. Required
    distributeForStudentWork *bool;
    // Resource object that has been associated with this assignment.
    resource *EducationResource;
}
// Instantiates a new educationAssignmentResource and sets the default values.
func NewEducationAssignmentResource()(*EducationAssignmentResource) {
    m := &EducationAssignmentResource{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the distributeForStudentWork property value. Indicates whether this resource should be copied to each student submission for modification and submission. Required
func (m *EducationAssignmentResource) GetDistributeForStudentWork()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.distributeForStudentWork
    }
}
// Gets the resource property value. Resource object that has been associated with this assignment.
func (m *EducationAssignmentResource) GetResource()(*EducationResource) {
    if m == nil {
        return nil
    } else {
        return m.resource
    }
}
// The deserialization information for the current model
func (m *EducationAssignmentResource) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["distributeForStudentWork"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetDistributeForStudentWork(val)
        return nil
    }
    res["resource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationResource() })
        if err != nil {
            return err
        }
        m.SetResource(val.(*EducationResource))
        return nil
    }
    return res
}
func (m *EducationAssignmentResource) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *EducationAssignmentResource) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("distributeForStudentWork", m.GetDistributeForStudentWork())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("resource", m.GetResource())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the distributeForStudentWork property value. Indicates whether this resource should be copied to each student submission for modification and submission. Required
// Parameters:
//  - value : Value to set for the distributeForStudentWork property.
func (m *EducationAssignmentResource) SetDistributeForStudentWork(value *bool)() {
    m.distributeForStudentWork = value
}
// Sets the resource property value. Resource object that has been associated with this assignment.
// Parameters:
//  - value : Value to set for the resource property.
func (m *EducationAssignmentResource) SetResource(value *EducationResource)() {
    m.resource = value
}
