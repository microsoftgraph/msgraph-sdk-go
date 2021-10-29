package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Set struct {
    Entity
    // Children terms of set in term [store].
    children []Term;
    // Date and time of set creation. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Description that gives details on the term usage.
    description *string;
    // Name of the set for each languageTag.
    localizedNames []LocalizedName;
    // 
    parentGroup *Group;
    // Custom properties for the set.
    properties []KeyValue;
    // Indicates which terms have been pinned or reused directly under the set.
    relations []Relation;
    // All the terms under the set.
    terms []Term;
}
// Instantiates a new set and sets the default values.
func NewSet()(*Set) {
    m := &Set{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the children property value. Children terms of set in term [store].
func (m *Set) GetChildren()([]Term) {
    if m == nil {
        return nil
    } else {
        return m.children
    }
}
// Gets the createdDateTime property value. Date and time of set creation. Read-only.
func (m *Set) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the description property value. Description that gives details on the term usage.
func (m *Set) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the localizedNames property value. Name of the set for each languageTag.
func (m *Set) GetLocalizedNames()([]LocalizedName) {
    if m == nil {
        return nil
    } else {
        return m.localizedNames
    }
}
// Gets the parentGroup property value. 
func (m *Set) GetParentGroup()(*Group) {
    if m == nil {
        return nil
    } else {
        return m.parentGroup
    }
}
// Gets the properties property value. Custom properties for the set.
func (m *Set) GetProperties()([]KeyValue) {
    if m == nil {
        return nil
    } else {
        return m.properties
    }
}
// Gets the relations property value. Indicates which terms have been pinned or reused directly under the set.
func (m *Set) GetRelations()([]Relation) {
    if m == nil {
        return nil
    } else {
        return m.relations
    }
}
// Gets the terms property value. All the terms under the set.
func (m *Set) GetTerms()([]Term) {
    if m == nil {
        return nil
    } else {
        return m.terms
    }
}
// The deserialization information for the current model
func (m *Set) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["localizedNames"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLocalizedName() })
        if err != nil {
            return err
        }
        res := make([]LocalizedName, len(val))
        for i, v := range val {
            res[i] = *(v.(*LocalizedName))
        }
        m.SetLocalizedNames(res)
        return nil
    }
    res["parentGroup"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroup() })
        if err != nil {
            return err
        }
        m.SetParentGroup(val.(*Group))
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
    res["terms"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTerm() })
        if err != nil {
            return err
        }
        res := make([]Term, len(val))
        for i, v := range val {
            res[i] = *(v.(*Term))
        }
        m.SetTerms(res)
        return nil
    }
    return res
}
func (m *Set) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Set) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetLocalizedNames()))
        for i, v := range m.GetLocalizedNames() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("localizedNames", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("parentGroup", m.GetParentGroup())
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTerms()))
        for i, v := range m.GetTerms() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("terms", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the children property value. Children terms of set in term [store].
// Parameters:
//  - value : Value to set for the children property.
func (m *Set) SetChildren(value []Term)() {
    m.children = value
}
// Sets the createdDateTime property value. Date and time of set creation. Read-only.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *Set) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the description property value. Description that gives details on the term usage.
// Parameters:
//  - value : Value to set for the description property.
func (m *Set) SetDescription(value *string)() {
    m.description = value
}
// Sets the localizedNames property value. Name of the set for each languageTag.
// Parameters:
//  - value : Value to set for the localizedNames property.
func (m *Set) SetLocalizedNames(value []LocalizedName)() {
    m.localizedNames = value
}
// Sets the parentGroup property value. 
// Parameters:
//  - value : Value to set for the parentGroup property.
func (m *Set) SetParentGroup(value *Group)() {
    m.parentGroup = value
}
// Sets the properties property value. Custom properties for the set.
// Parameters:
//  - value : Value to set for the properties property.
func (m *Set) SetProperties(value []KeyValue)() {
    m.properties = value
}
// Sets the relations property value. Indicates which terms have been pinned or reused directly under the set.
// Parameters:
//  - value : Value to set for the relations property.
func (m *Set) SetRelations(value []Relation)() {
    m.relations = value
}
// Sets the terms property value. All the terms under the set.
// Parameters:
//  - value : Value to set for the terms property.
func (m *Set) SetTerms(value []Term)() {
    m.terms = value
}
