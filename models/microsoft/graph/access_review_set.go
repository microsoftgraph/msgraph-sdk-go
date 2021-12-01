package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessReviewSet 
type AccessReviewSet struct {
    Entity
    // 
    definitions []AccessReviewScheduleDefinition;
}
// NewAccessReviewSet instantiates a new accessReviewSet and sets the default values.
func NewAccessReviewSet()(*AccessReviewSet) {
    m := &AccessReviewSet{
        Entity: *NewEntity(),
    }
    return m
}
// GetDefinitions gets the definitions property value. 
func (m *AccessReviewSet) GetDefinitions()([]AccessReviewScheduleDefinition) {
    if m == nil {
        return nil
    } else {
        return m.definitions
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessReviewSet) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["definitions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessReviewScheduleDefinition() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessReviewScheduleDefinition, len(val))
            for i, v := range val {
                res[i] = *(v.(*AccessReviewScheduleDefinition))
            }
            m.SetDefinitions(res)
        }
        return nil
    }
    return res
}
func (m *AccessReviewSet) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AccessReviewSet) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDefinitions()))
        for i, v := range m.GetDefinitions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("definitions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDefinitions sets the definitions property value. 
func (m *AccessReviewSet) SetDefinitions(value []AccessReviewScheduleDefinition)() {
    if m != nil {
        m.definitions = value
    }
}
