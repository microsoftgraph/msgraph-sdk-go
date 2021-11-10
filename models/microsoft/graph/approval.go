package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Approval struct {
    Entity
    // A collection of stages in the approval decision.
    stages []ApprovalStage;
}
// Instantiates a new approval and sets the default values.
func NewApproval()(*Approval) {
    m := &Approval{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the stages property value. A collection of stages in the approval decision.
func (m *Approval) GetStages()([]ApprovalStage) {
    if m == nil {
        return nil
    } else {
        return m.stages
    }
}
// The deserialization information for the current model
func (m *Approval) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["stages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewApprovalStage() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ApprovalStage, len(val))
            for i, v := range val {
                res[i] = *(v.(*ApprovalStage))
            }
            m.SetStages(res)
        }
        return nil
    }
    return res
}
func (m *Approval) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Approval) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetStages()))
        for i, v := range m.GetStages() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("stages", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the stages property value. A collection of stages in the approval decision.
// Parameters:
//  - value : Value to set for the stages property.
func (m *Approval) SetStages(value []ApprovalStage)() {
    m.stages = value
}
