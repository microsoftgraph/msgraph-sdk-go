package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    id62b8df0892707d421d6e0a5aefa589248c11f95794bf4122483a0ef812fad7d "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/termstore"
)

// 
type Relation struct {
    Entity
    // The from [term] of the relation. The term from which the relationship is defined. A null value would indicate the relation is directly with the [set].
    fromTerm *Term;
    // The type of relation. Possible values are: pin, reuse.
    relationship *id62b8df0892707d421d6e0a5aefa589248c11f95794bf4122483a0ef812fad7d.RelationType;
    // The [set] in which the relation is relevant.
    set *Set;
    // The to [term] of the relation. The term to which the relationship is defined.
    toTerm *Term;
}
// Instantiates a new relation and sets the default values.
func NewRelation()(*Relation) {
    m := &Relation{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the fromTerm property value. The from [term] of the relation. The term from which the relationship is defined. A null value would indicate the relation is directly with the [set].
func (m *Relation) GetFromTerm()(*Term) {
    if m == nil {
        return nil
    } else {
        return m.fromTerm
    }
}
// Gets the relationship property value. The type of relation. Possible values are: pin, reuse.
func (m *Relation) GetRelationship()(*id62b8df0892707d421d6e0a5aefa589248c11f95794bf4122483a0ef812fad7d.RelationType) {
    if m == nil {
        return nil
    } else {
        return m.relationship
    }
}
// Gets the set property value. The [set] in which the relation is relevant.
func (m *Relation) GetSet()(*Set) {
    if m == nil {
        return nil
    } else {
        return m.set
    }
}
// Gets the toTerm property value. The to [term] of the relation. The term to which the relationship is defined.
func (m *Relation) GetToTerm()(*Term) {
    if m == nil {
        return nil
    } else {
        return m.toTerm
    }
}
// The deserialization information for the current model
func (m *Relation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["fromTerm"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTerm() })
        if err != nil {
            return err
        }
        m.SetFromTerm(val.(*Term))
        return nil
    }
    res["relationship"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(id62b8df0892707d421d6e0a5aefa589248c11f95794bf4122483a0ef812fad7d.ParseRelationType)
        if err != nil {
            return err
        }
        cast := val.(id62b8df0892707d421d6e0a5aefa589248c11f95794bf4122483a0ef812fad7d.RelationType)
        m.SetRelationship(&cast)
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
    res["toTerm"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTerm() })
        if err != nil {
            return err
        }
        m.SetToTerm(val.(*Term))
        return nil
    }
    return res
}
func (m *Relation) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Relation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("fromTerm", m.GetFromTerm())
        if err != nil {
            return err
        }
    }
    if m.GetRelationship() != nil {
        cast := m.GetRelationship().String()
        err = writer.WriteStringValue("relationship", &cast)
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
    {
        err = writer.WriteObjectValue("toTerm", m.GetToTerm())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the fromTerm property value. The from [term] of the relation. The term from which the relationship is defined. A null value would indicate the relation is directly with the [set].
// Parameters:
//  - value : Value to set for the fromTerm property.
func (m *Relation) SetFromTerm(value *Term)() {
    m.fromTerm = value
}
// Sets the relationship property value. The type of relation. Possible values are: pin, reuse.
// Parameters:
//  - value : Value to set for the relationship property.
func (m *Relation) SetRelationship(value *id62b8df0892707d421d6e0a5aefa589248c11f95794bf4122483a0ef812fad7d.RelationType)() {
    m.relationship = value
}
// Sets the set property value. The [set] in which the relation is relevant.
// Parameters:
//  - value : Value to set for the set property.
func (m *Relation) SetSet(value *Set)() {
    m.set = value
}
// Sets the toTerm property value. The to [term] of the relation. The term to which the relationship is defined.
// Parameters:
//  - value : Value to set for the toTerm property.
func (m *Relation) SetToTerm(value *Term)() {
    m.toTerm = value
}
