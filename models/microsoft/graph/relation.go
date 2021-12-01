package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    id62b8df0892707d421d6e0a5aefa589248c11f95794bf4122483a0ef812fad7d "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/termstore"
)

// Relation 
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
// NewRelation instantiates a new relation and sets the default values.
func NewRelation()(*Relation) {
    m := &Relation{
        Entity: *NewEntity(),
    }
    return m
}
// GetFromTerm gets the fromTerm property value. The from [term] of the relation. The term from which the relationship is defined. A null value would indicate the relation is directly with the [set].
func (m *Relation) GetFromTerm()(*Term) {
    if m == nil {
        return nil
    } else {
        return m.fromTerm
    }
}
// GetRelationship gets the relationship property value. The type of relation. Possible values are: pin, reuse.
func (m *Relation) GetRelationship()(*id62b8df0892707d421d6e0a5aefa589248c11f95794bf4122483a0ef812fad7d.RelationType) {
    if m == nil {
        return nil
    } else {
        return m.relationship
    }
}
// GetSet gets the set property value. The [set] in which the relation is relevant.
func (m *Relation) GetSet()(*Set) {
    if m == nil {
        return nil
    } else {
        return m.set
    }
}
// GetToTerm gets the toTerm property value. The to [term] of the relation. The term to which the relationship is defined.
func (m *Relation) GetToTerm()(*Term) {
    if m == nil {
        return nil
    } else {
        return m.toTerm
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Relation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["fromTerm"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTerm() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFromTerm(val.(*Term))
        }
        return nil
    }
    res["relationship"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(id62b8df0892707d421d6e0a5aefa589248c11f95794bf4122483a0ef812fad7d.ParseRelationType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(id62b8df0892707d421d6e0a5aefa589248c11f95794bf4122483a0ef812fad7d.RelationType)
            m.SetRelationship(&cast)
        }
        return nil
    }
    res["set"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSet(val.(*Set))
        }
        return nil
    }
    res["toTerm"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTerm() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetToTerm(val.(*Term))
        }
        return nil
    }
    return res
}
func (m *Relation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetFromTerm sets the fromTerm property value. The from [term] of the relation. The term from which the relationship is defined. A null value would indicate the relation is directly with the [set].
func (m *Relation) SetFromTerm(value *Term)() {
    if m != nil {
        m.fromTerm = value
    }
}
// SetRelationship sets the relationship property value. The type of relation. Possible values are: pin, reuse.
func (m *Relation) SetRelationship(value *id62b8df0892707d421d6e0a5aefa589248c11f95794bf4122483a0ef812fad7d.RelationType)() {
    if m != nil {
        m.relationship = value
    }
}
// SetSet sets the set property value. The [set] in which the relation is relevant.
func (m *Relation) SetSet(value *Set)() {
    if m != nil {
        m.set = value
    }
}
// SetToTerm sets the toTerm property value. The to [term] of the relation. The term to which the relationship is defined.
func (m *Relation) SetToTerm(value *Term)() {
    if m != nil {
        m.toTerm = value
    }
}
