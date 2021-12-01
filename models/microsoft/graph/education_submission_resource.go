package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EducationSubmissionResource 
type EducationSubmissionResource struct {
    Entity
    // Pointer to the assignment from which this resource was copied. If this is null, the student uploaded the resource.
    assignmentResourceUrl *string;
    // Resource object.
    resource *EducationResource;
}
// NewEducationSubmissionResource instantiates a new educationSubmissionResource and sets the default values.
func NewEducationSubmissionResource()(*EducationSubmissionResource) {
    m := &EducationSubmissionResource{
        Entity: *NewEntity(),
    }
    return m
}
// GetAssignmentResourceUrl gets the assignmentResourceUrl property value. Pointer to the assignment from which this resource was copied. If this is null, the student uploaded the resource.
func (m *EducationSubmissionResource) GetAssignmentResourceUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assignmentResourceUrl
    }
}
// GetResource gets the resource property value. Resource object.
func (m *EducationSubmissionResource) GetResource()(*EducationResource) {
    if m == nil {
        return nil
    } else {
        return m.resource
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationSubmissionResource) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignmentResourceUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentResourceUrl(val)
        }
        return nil
    }
    res["resource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationResource() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val.(*EducationResource))
        }
        return nil
    }
    return res
}
func (m *EducationSubmissionResource) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *EducationSubmissionResource) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("assignmentResourceUrl", m.GetAssignmentResourceUrl())
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
// SetAssignmentResourceUrl sets the assignmentResourceUrl property value. Pointer to the assignment from which this resource was copied. If this is null, the student uploaded the resource.
func (m *EducationSubmissionResource) SetAssignmentResourceUrl(value *string)() {
    if m != nil {
        m.assignmentResourceUrl = value
    }
}
// SetResource sets the resource property value. Resource object.
func (m *EducationSubmissionResource) SetResource(value *EducationResource)() {
    if m != nil {
        m.resource = value
    }
}
