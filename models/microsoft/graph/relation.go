package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    id62b8df0892707d421d6e0a5aefa589248c11f95794bf4122483a0ef812fad7d "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/termstore"
)

// Relation provides operations to manage the collection of group entities.
type Relation struct {
    Entity
    // The from [term] of the relation. The term from which the relationship is defined. A null value would indicate the relation is directly with the [set].
    fromTerm Termable;
    // The type of relation. Possible values are: pin, reuse.
    relationship *id62b8df0892707d421d6e0a5aefa589248c11f95794bf4122483a0ef812fad7d.RelationType;
    // The [set] in which the relation is relevant.
    set Setable;
    // The to [term] of the relation. The term to which the relationship is defined.
    toTerm Termable;
}
// NewRelation instantiates a new relation and sets the default values.
func NewRelation()(*Relation) {
    m := &Relation{
        Entity: *NewEntity(),
    }
    return m
}
// CreateRelationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRelationFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewRelation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Relation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["fromTerm"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateTermFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFromTerm(val.(Termable))
        }
        return nil
    }
    res["relationship"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(id62b8df0892707d421d6e0a5aefa589248c11f95794bf4122483a0ef812fad7d.ParseRelationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRelationship(val.(*id62b8df0892707d421d6e0a5aefa589248c11f95794bf4122483a0ef812fad7d.RelationType))
        }
        return nil
    }
    res["set"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSet(val.(Setable))
        }
        return nil
    }
    res["toTerm"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateTermFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetToTerm(val.(Termable))
        }
        return nil
    }
    return res
}
// GetFromTerm gets the fromTerm property value. The from [term] of the relation. The term from which the relationship is defined. A null value would indicate the relation is directly with the [set].
func (m *Relation) GetFromTerm()(Termable) {
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
func (m *Relation) GetSet()(Setable) {
    if m == nil {
        return nil
    } else {
        return m.set
    }
}
// GetToTerm gets the toTerm property value. The to [term] of the relation. The term to which the relationship is defined.
func (m *Relation) GetToTerm()(Termable) {
    if m == nil {
        return nil
    } else {
        return m.toTerm
    }
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
        cast := (*m.GetRelationship()).String()
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
func (m *Relation) SetFromTerm(value Termable)() {
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
func (m *Relation) SetSet(value Setable)() {
    if m != nil {
        m.set = value
    }
}
// SetToTerm sets the toTerm property value. The to [term] of the relation. The term to which the relationship is defined.
func (m *Relation) SetToTerm(value Termable)() {
    if m != nil {
        m.toTerm = value
    }
}
