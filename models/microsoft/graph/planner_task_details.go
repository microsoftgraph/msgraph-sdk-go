package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PlannerTaskDetails struct {
    Entity
    // The collection of checklist items on the task.
    checklist *PlannerChecklistItems;
    // Description of the task
    description *string;
    // This sets the type of preview that shows up on the task. The possible values are: automatic, noPreview, checklist, description, reference. When set to automatic the displayed preview is chosen by the app viewing the task.
    previewType *PlannerPreviewType;
    // The collection of references on the task.
    references *PlannerExternalReferences;
}
// Instantiates a new plannerTaskDetails and sets the default values.
func NewPlannerTaskDetails()(*PlannerTaskDetails) {
    m := &PlannerTaskDetails{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the checklist property value. The collection of checklist items on the task.
func (m *PlannerTaskDetails) GetChecklist()(*PlannerChecklistItems) {
    if m == nil {
        return nil
    } else {
        return m.checklist
    }
}
// Gets the description property value. Description of the task
func (m *PlannerTaskDetails) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the previewType property value. This sets the type of preview that shows up on the task. The possible values are: automatic, noPreview, checklist, description, reference. When set to automatic the displayed preview is chosen by the app viewing the task.
func (m *PlannerTaskDetails) GetPreviewType()(*PlannerPreviewType) {
    if m == nil {
        return nil
    } else {
        return m.previewType
    }
}
// Gets the references property value. The collection of references on the task.
func (m *PlannerTaskDetails) GetReferences()(*PlannerExternalReferences) {
    if m == nil {
        return nil
    } else {
        return m.references
    }
}
// The deserialization information for the current model
func (m *PlannerTaskDetails) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["checklist"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerChecklistItems() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChecklist(val.(*PlannerChecklistItems))
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
    res["previewType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePlannerPreviewType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(PlannerPreviewType)
            m.SetPreviewType(&cast)
        }
        return nil
    }
    res["references"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerExternalReferences() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReferences(val.(*PlannerExternalReferences))
        }
        return nil
    }
    return res
}
func (m *PlannerTaskDetails) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the checklist property value. The collection of checklist items on the task.
// Parameters:
//  - value : Value to set for the checklist property.
func (m *PlannerTaskDetails) SetChecklist(value *PlannerChecklistItems)() {
    m.checklist = value
}
// Sets the description property value. Description of the task
// Parameters:
//  - value : Value to set for the description property.
func (m *PlannerTaskDetails) SetDescription(value *string)() {
    m.description = value
}
// Sets the previewType property value. This sets the type of preview that shows up on the task. The possible values are: automatic, noPreview, checklist, description, reference. When set to automatic the displayed preview is chosen by the app viewing the task.
// Parameters:
//  - value : Value to set for the previewType property.
func (m *PlannerTaskDetails) SetPreviewType(value *PlannerPreviewType)() {
    m.previewType = value
}
// Sets the references property value. The collection of references on the task.
// Parameters:
//  - value : Value to set for the references property.
func (m *PlannerTaskDetails) SetReferences(value *PlannerExternalReferences)() {
    m.references = value
}
