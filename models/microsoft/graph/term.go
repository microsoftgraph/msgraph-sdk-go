package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Term struct {
    Entity
    // Children of current term.
    children []Term;
    // Date and time of term creation. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Description about term that is dependent on the languageTag.
    descriptions []LocalizedDescription;
    // Label metadata for a term.
    labels []LocalizedLabel;
    // Last date and time of term modification. Read-only.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Collection of properties on the term.
    properties []KeyValue;
    // To indicate which terms are related to the current term as either pinned or reused.
    relations []Relation;
    // The [set] in which the term is created.
    set *Set;
}
// Instantiates a new term and sets the default values.
func NewTerm()(*Term) {
    m := &Term{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the children property value. Children of current term.
func (m *Term) GetChildren()([]Term) {
    if m == nil {
        return nil
    } else {
        return m.children
    }
}
// Gets the createdDateTime property value. Date and time of term creation. Read-only.
func (m *Term) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the descriptions property value. Description about term that is dependent on the languageTag.
func (m *Term) GetDescriptions()([]LocalizedDescription) {
    if m == nil {
        return nil
    } else {
        return m.descriptions
    }
}
// Gets the labels property value. Label metadata for a term.
func (m *Term) GetLabels()([]LocalizedLabel) {
    if m == nil {
        return nil
    } else {
        return m.labels
    }
}
// Gets the lastModifiedDateTime property value. Last date and time of term modification. Read-only.
func (m *Term) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the properties property value. Collection of properties on the term.
func (m *Term) GetProperties()([]KeyValue) {
    if m == nil {
        return nil
    } else {
        return m.properties
    }
}
// Gets the relations property value. To indicate which terms are related to the current term as either pinned or reused.
func (m *Term) GetRelations()([]Relation) {
    if m == nil {
        return nil
    } else {
        return m.relations
    }
}
// Gets the set property value. The [set] in which the term is created.
func (m *Term) GetSet()(*Set) {
    if m == nil {
        return nil
    } else {
        return m.set
    }
}
// The deserialization information for the current model
func (m *Term) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["children"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTerm() })
        if err != nil {
            return err
        }
        res := make([]Term, len(val))
        for i, v := range val {
            res[i] = *(v.(*Term))
        }
        m.SetChildren(res)
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["descriptions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLocalizedDescription() })
        if err != nil {
            return err
        }
        res := make([]LocalizedDescription, len(val))
        for i, v := range val {
            res[i] = *(v.(*LocalizedDescription))
        }
        m.SetDescriptions(res)
        return nil
    }
    res["labels"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLocalizedLabel() })
        if err != nil {
            return err
        }
        res := make([]LocalizedLabel, len(val))
        for i, v := range val {
            res[i] = *(v.(*LocalizedLabel))
        }
        m.SetLabels(res)
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastModifiedDateTime(val)
        return nil
    }
    res["properties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewKeyValue() })
        if err != nil {
            return err
        }
        res := make([]KeyValue, len(val))
        for i, v := range val {
            res[i] = *(v.(*KeyValue))
        }
        m.SetProperties(res)
        return nil
    }
    res["relations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRelation() })
        if err != nil {
            return err
        }
        res := make([]Relation, len(val))
        for i, v := range val {
            res[i] = *(v.(*Relation))
        }
        m.SetRelations(res)
        return nil
    }
    res["set"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSet() })
        if err != nil {
            return err
        }
        m.SetSet(val.(*Set))
        return nil
    }
    return res
}
func (m *Term) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Term) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetChildren()))
        for i, v := range m.GetChildren() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("children", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDescriptions()))
        for i, v := range m.GetDescriptions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("descriptions", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetLabels()))
        for i, v := range m.GetLabels() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("labels", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetProperties()))
        for i, v := range m.GetProperties() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("properties", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRelations()))
        for i, v := range m.GetRelations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("relations", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("set", m.GetSet())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the children property value. Children of current term.
// Parameters:
//  - value : Value to set for the children property.
func (m *Term) SetChildren(value []Term)() {
    m.children = value
}
// Sets the createdDateTime property value. Date and time of term creation. Read-only.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *Term) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the descriptions property value. Description about term that is dependent on the languageTag.
// Parameters:
//  - value : Value to set for the descriptions property.
func (m *Term) SetDescriptions(value []LocalizedDescription)() {
    m.descriptions = value
}
// Sets the labels property value. Label metadata for a term.
// Parameters:
//  - value : Value to set for the labels property.
func (m *Term) SetLabels(value []LocalizedLabel)() {
    m.labels = value
}
// Sets the lastModifiedDateTime property value. Last date and time of term modification. Read-only.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *Term) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the properties property value. Collection of properties on the term.
// Parameters:
//  - value : Value to set for the properties property.
func (m *Term) SetProperties(value []KeyValue)() {
    m.properties = value
}
// Sets the relations property value. To indicate which terms are related to the current term as either pinned or reused.
// Parameters:
//  - value : Value to set for the relations property.
func (m *Term) SetRelations(value []Relation)() {
    m.relations = value
}
// Sets the set property value. The [set] in which the term is created.
// Parameters:
//  - value : Value to set for the set property.
func (m *Term) SetSet(value *Set)() {
    m.set = value
}
