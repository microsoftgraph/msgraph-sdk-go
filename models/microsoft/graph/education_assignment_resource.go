package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type EducationAssignmentResource struct {
    Entity
    distributeForStudentWork *bool;
    resource *EducationResource;
}
func NewEducationAssignmentResource()(*EducationAssignmentResource) {
    m := &EducationAssignmentResource{
        Entity: *NewEntity(),
    }
    return m
}
func (m *EducationAssignmentResource) GetDistributeForStudentWork()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.distributeForStudentWork
    }
}
func (m *EducationAssignmentResource) GetResource()(*EducationResource) {
    if m == nil {
        return nil
    } else {
        return m.resource
    }
}
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
func (m *EducationAssignmentResource) SetDistributeForStudentWork(value *bool)() {
    m.distributeForStudentWork = value
}
func (m *EducationAssignmentResource) SetResource(value *EducationResource)() {
    m.resource = value
}
