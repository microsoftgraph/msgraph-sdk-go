package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Ediscoveryroot struct {
    Entity
    // 
    cases []Case_escaped;
}
// Instantiates a new ediscoveryroot and sets the default values.
func NewEdiscoveryroot()(*Ediscoveryroot) {
    m := &Ediscoveryroot{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the cases property value. 
func (m *Ediscoveryroot) GetCases()([]Case_escaped) {
    if m == nil {
        return nil
    } else {
        return m.cases
    }
}
// The deserialization information for the current model
func (m *Ediscoveryroot) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["cases"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCase_escaped() })
        if err != nil {
            return err
        }
        res := make([]Case_escaped, len(val))
        for i, v := range val {
            res[i] = *(v.(*Case_escaped))
        }
        m.SetCases(res)
        return nil
    }
    return res
}
func (m *Ediscoveryroot) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Ediscoveryroot) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCases()))
        for i, v := range m.GetCases() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("cases", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the cases property value. 
// Parameters:
//  - value : Value to set for the cases property.
func (m *Ediscoveryroot) SetCases(value []Case_escaped)() {
    m.cases = value
}
