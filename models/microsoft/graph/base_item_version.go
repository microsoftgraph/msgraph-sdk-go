package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// BaseItemVersion 
type BaseItemVersion struct {
    Entity
    // Identity of the user which last modified the version. Read-only.
    lastModifiedBy *IdentitySet;
    // Date and time the version was last modified. Read-only.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Indicates the publication status of this particular version. Read-only.
    publication *PublicationFacet;
}
// NewBaseItemVersion instantiates a new baseItemVersion and sets the default values.
func NewBaseItemVersion()(*BaseItemVersion) {
    m := &BaseItemVersion{
        Entity: *NewEntity(),
    }
    return m
}
// GetLastModifiedBy gets the lastModifiedBy property value. Identity of the user which last modified the version. Read-only.
func (m *BaseItemVersion) GetLastModifiedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedBy
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Date and time the version was last modified. Read-only.
func (m *BaseItemVersion) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetPublication gets the publication property value. Indicates the publication status of this particular version. Read-only.
func (m *BaseItemVersion) GetPublication()(*PublicationFacet) {
    if m == nil {
        return nil
    } else {
        return m.publication
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BaseItemVersion) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["lastModifiedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(*IdentitySet))
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["publication"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPublicationFacet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublication(val.(*PublicationFacet))
        }
        return nil
    }
    return res
}
func (m *BaseItemVersion) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *BaseItemVersion) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("lastModifiedBy", m.GetLastModifiedBy())
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
        err = writer.WriteObjectValue("publication", m.GetPublication())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLastModifiedBy sets the lastModifiedBy property value. Identity of the user which last modified the version. Read-only.
func (m *BaseItemVersion) SetLastModifiedBy(value *IdentitySet)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Date and time the version was last modified. Read-only.
func (m *BaseItemVersion) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetPublication sets the publication property value. Indicates the publication status of this particular version. Read-only.
func (m *BaseItemVersion) SetPublication(value *PublicationFacet)() {
    m.publication = value
}
