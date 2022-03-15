package termstore

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// Group provides operations to manage the drive singleton.
type Group struct {
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Entity
    // Date and time of the group creation. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Description that gives details on the term usage.
    description *string;
    // Name of the group.
    displayName *string;
    // ID of the parent site of this group.
    parentSiteId *string;
    // Returns the type of the group. Possible values are global, system, and siteCollection.
    scope *TermGroupScope;
    // All sets under the group in a term [store].
    sets []Setable;
}
// NewGroup instantiates a new group and sets the default values.
func NewGroup()(*Group) {
    m := &Group{
        Entity: *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewEntity(),
    }
    return m
}
// CreateGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewGroup(), nil
}
// GetCreatedDateTime gets the createdDateTime property value. Date and time of the group creation. Read-only.
func (m *Group) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDescription gets the description property value. Description that gives details on the term usage.
func (m *Group) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. Name of the group.
func (m *Group) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Group) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
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
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["parentSiteId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentSiteId(val)
        }
        return nil
    }
    res["scope"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTermGroupScope)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScope(val.(*TermGroupScope))
        }
        return nil
    }
    res["sets"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Setable, len(val))
            for i, v := range val {
                res[i] = v.(Setable)
            }
            m.SetSets(res)
        }
        return nil
    }
    return res
}
// GetParentSiteId gets the parentSiteId property value. ID of the parent site of this group.
func (m *Group) GetParentSiteId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.parentSiteId
    }
}
// GetScope gets the scope property value. Returns the type of the group. Possible values are global, system, and siteCollection.
func (m *Group) GetScope()(*TermGroupScope) {
    if m == nil {
        return nil
    } else {
        return m.scope
    }
}
// GetSets gets the sets property value. All sets under the group in a term [store].
func (m *Group) GetSets()([]Setable) {
    if m == nil {
        return nil
    } else {
        return m.sets
    }
}
func (m *Group) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Group) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
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
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("parentSiteId", m.GetParentSiteId())
        if err != nil {
            return err
        }
    }
    if m.GetScope() != nil {
        cast := (*m.GetScope()).String()
        err = writer.WriteStringValue("scope", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSets() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSets()))
        for i, v := range m.GetSets() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("sets", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedDateTime sets the createdDateTime property value. Date and time of the group creation. Read-only.
func (m *Group) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDescription sets the description property value. Description that gives details on the term usage.
func (m *Group) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. Name of the group.
func (m *Group) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetParentSiteId sets the parentSiteId property value. ID of the parent site of this group.
func (m *Group) SetParentSiteId(value *string)() {
    if m != nil {
        m.parentSiteId = value
    }
}
// SetScope sets the scope property value. Returns the type of the group. Possible values are global, system, and siteCollection.
func (m *Group) SetScope(value *TermGroupScope)() {
    if m != nil {
        m.scope = value
    }
}
// SetSets sets the sets property value. All sets under the group in a term [store].
func (m *Group) SetSets(value []Setable)() {
    if m != nil {
        m.sets = value
    }
}
