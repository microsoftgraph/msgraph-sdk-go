package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PlannerTaskDetails provides operations to manage the collection of drive entities.
type PlannerTaskDetails struct {
    Entity
    // The collection of checklist items on the task.
    checklist PlannerChecklistItemsable;
    // Description of the task.
    description *string;
    // This sets the type of preview that shows up on the task. The possible values are: automatic, noPreview, checklist, description, reference. When set to automatic the displayed preview is chosen by the app viewing the task.
    previewType *PlannerPreviewType;
    // The collection of references on the task.
    references PlannerExternalReferencesable;
}
// NewPlannerTaskDetails instantiates a new plannerTaskDetails and sets the default values.
func NewPlannerTaskDetails()(*PlannerTaskDetails) {
    m := &PlannerTaskDetails{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePlannerTaskDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerTaskDetailsFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewPlannerTaskDetails(), nil
}
// GetChecklist gets the checklist property value. The collection of checklist items on the task.
func (m *PlannerTaskDetails) GetChecklist()(PlannerChecklistItemsable) {
    if m == nil {
        return nil
    } else {
        return m.checklist
    }
}
// GetDescription gets the description property value. Description of the task.
func (m *PlannerTaskDetails) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerTaskDetails) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["checklist"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerChecklistItemsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChecklist(val.(PlannerChecklistItemsable))
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
            m.SetPreviewType(val.(*PlannerPreviewType))
        }
        return nil
    }
    res["references"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerExternalReferencesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReferences(val.(PlannerExternalReferencesable))
        }
        return nil
    }
    return res
}
// GetPreviewType gets the previewType property value. This sets the type of preview that shows up on the task. The possible values are: automatic, noPreview, checklist, description, reference. When set to automatic the displayed preview is chosen by the app viewing the task.
func (m *PlannerTaskDetails) GetPreviewType()(*PlannerPreviewType) {
    if m == nil {
        return nil
    } else {
        return m.previewType
    }
}
// GetReferences gets the references property value. The collection of references on the task.
func (m *PlannerTaskDetails) GetReferences()(PlannerExternalReferencesable) {
    if m == nil {
        return nil
    } else {
        return m.references
    }
}
func (m *PlannerTaskDetails) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
        cast := (*m.GetPreviewType()).String()
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
// SetChecklist sets the checklist property value. The collection of checklist items on the task.
func (m *PlannerTaskDetails) SetChecklist(value PlannerChecklistItemsable)() {
    if m != nil {
        m.checklist = value
    }
}
// SetDescription sets the description property value. Description of the task.
func (m *PlannerTaskDetails) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetPreviewType sets the previewType property value. This sets the type of preview that shows up on the task. The possible values are: automatic, noPreview, checklist, description, reference. When set to automatic the displayed preview is chosen by the app viewing the task.
func (m *PlannerTaskDetails) SetPreviewType(value *PlannerPreviewType)() {
    if m != nil {
        m.previewType = value
    }
}
// SetReferences sets the references property value. The collection of references on the task.
func (m *PlannerTaskDetails) SetReferences(value PlannerExternalReferencesable)() {
    if m != nil {
        m.references = value
    }
}
