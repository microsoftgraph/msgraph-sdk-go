package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EducationAssignmentResource 
type EducationAssignmentResource struct {
    Entity
    // Indicates whether this resource should be copied to each student submission for modification and submission. Required
    distributeForStudentWork *bool;
    // Resource object that has been associated with this assignment.
    resource EducationResourceable;
}
// NewEducationAssignmentResource instantiates a new educationAssignmentResource and sets the default values.
func NewEducationAssignmentResource()(*EducationAssignmentResource) {
    m := &EducationAssignmentResource{
        Entity: *NewEntity(),
    }
    return m
}
// CreateEducationAssignmentResourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationAssignmentResourceFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewEducationAssignmentResource(), nil
}
// GetDistributeForStudentWork gets the distributeForStudentWork property value. Indicates whether this resource should be copied to each student submission for modification and submission. Required
func (m *EducationAssignmentResource) GetDistributeForStudentWork()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.distributeForStudentWork
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationAssignmentResource) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["distributeForStudentWork"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDistributeForStudentWork(val)
        }
        return nil
    }
    res["resource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val.(EducationResourceable))
        }
        return nil
    }
    return res
}
// GetResource gets the resource property value. Resource object that has been associated with this assignment.
func (m *EducationAssignmentResource) GetResource()(EducationResourceable) {
    if m == nil {
        return nil
    } else {
        return m.resource
    }
}
// Serialize serializes information the current object
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
// SetDistributeForStudentWork sets the distributeForStudentWork property value. Indicates whether this resource should be copied to each student submission for modification and submission. Required
func (m *EducationAssignmentResource) SetDistributeForStudentWork(value *bool)() {
    if m != nil {
        m.distributeForStudentWork = value
    }
}
// SetResource sets the resource property value. Resource object that has been associated with this assignment.
func (m *EducationAssignmentResource) SetResource(value EducationResourceable)() {
    if m != nil {
        m.resource = value
    }
}
