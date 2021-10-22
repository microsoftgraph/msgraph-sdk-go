package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PlannerTaskDetails struct {
    Entity
    checklist *PlannerChecklistItems;
    description *string;
    previewType *PlannerPreviewType;
    references *PlannerExternalReferences;
}
func NewPlannerTaskDetails()(*PlannerTaskDetails) {
    m := &PlannerTaskDetails{
        Entity: *NewEntity(),
    }
    return m
}
func (m *PlannerTaskDetails) GetChecklist()(*PlannerChecklistItems) {
    if m == nil {
        return nil
    } else {
        return m.checklist
    }
}
func (m *PlannerTaskDetails) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *PlannerTaskDetails) GetPreviewType()(*PlannerPreviewType) {
    if m == nil {
        return nil
    } else {
        return m.previewType
    }
}
func (m *PlannerTaskDetails) GetReferences()(*PlannerExternalReferences) {
    if m == nil {
        return nil
    } else {
        return m.references
    }
}
func (m *PlannerTaskDetails) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["checklist"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerChecklistItems() })
        if err != nil {
            return err
        }
        m.SetChecklist(val.(*PlannerChecklistItems))
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["previewType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePlannerPreviewType)
        if err != nil {
            return err
        }
        cast := val.(PlannerPreviewType)
        m.SetPreviewType(&cast)
        return nil
    }
    res["references"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerExternalReferences() })
        if err != nil {
            return err
        }
        m.SetReferences(val.(*PlannerExternalReferences))
        return nil
    }
    return res
}
func (m *PlannerTaskDetails) IsNil()(bool) {
    return m == nil
}
func (m *PlannerTaskDetails) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("checklist", m.GetChecklist())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    if m.GetPreviewType() != nil {
        cast := m.GetPreviewType().String()
        err = writer.WriteStringValue("previewType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("references", m.GetReferences())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *PlannerTaskDetails) SetChecklist(value *PlannerChecklistItems)() {
    m.checklist = value
}
func (m *PlannerTaskDetails) SetDescription(value *string)() {
    m.description = value
}
func (m *PlannerTaskDetails) SetPreviewType(value *PlannerPreviewType)() {
    m.previewType = value
}
func (m *PlannerTaskDetails) SetReferences(value *PlannerExternalReferences)() {
    m.references = value
}
