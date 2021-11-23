package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// shared 
type Shared struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The identity of the owner of the shared item. Read-only.
    owner *IdentitySet;
    // Indicates the scope of how the item is shared: anonymous, organization, or users. Read-only.
    scope *string;
    // The identity of the user who shared the item. Read-only.
    sharedBy *IdentitySet;
    // The UTC date and time when the item was shared. Read-only.
    sharedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// NewShared instantiates a new shared and sets the default values.
func NewShared()(*Shared) {
    m := &Shared{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Shared) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetOwner gets the owner property value. The identity of the owner of the shared item. Read-only.
func (m *Shared) GetOwner()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.owner
    }
}
// GetScope gets the scope property value. Indicates the scope of how the item is shared: anonymous, organization, or users. Read-only.
func (m *Shared) GetScope()(*string) {
    if m == nil {
        return nil
    } else {
        return m.scope
    }
}
// GetSharedBy gets the sharedBy property value. The identity of the user who shared the item. Read-only.
func (m *Shared) GetSharedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.sharedBy
    }
}
// GetSharedDateTime gets the sharedDateTime property value. The UTC date and time when the item was shared. Read-only.
func (m *Shared) GetSharedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.sharedDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Shared) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["owner"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val.(*IdentitySet))
        }
        return nil
    }
    res["scope"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScope(val)
        }
        return nil
    }
    res["sharedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharedBy(val.(*IdentitySet))
        }
        return nil
    }
    res["sharedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharedDateTime(val)
        }
        return nil
    }
    return res
}
func (m *Shared) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Shared) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("owner", m.GetOwner())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("scope", m.GetScope())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("sharedBy", m.GetSharedBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("sharedDateTime", m.GetSharedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Shared) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOwner sets the owner property value. The identity of the owner of the shared item. Read-only.
func (m *Shared) SetOwner(value *IdentitySet)() {
    m.owner = value
}
// SetScope sets the scope property value. Indicates the scope of how the item is shared: anonymous, organization, or users. Read-only.
func (m *Shared) SetScope(value *string)() {
    m.scope = value
}
// SetSharedBy sets the sharedBy property value. The identity of the user who shared the item. Read-only.
func (m *Shared) SetSharedBy(value *IdentitySet)() {
    m.sharedBy = value
}
// SetSharedDateTime sets the sharedDateTime property value. The UTC date and time when the item was shared. Read-only.
func (m *Shared) SetSharedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.sharedDateTime = value
}
