package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ItemActivity provides operations to manage the drive singleton.
type ItemActivity struct {
    Entity
    // An item was accessed.
    access AccessActionable;
    // Details about when the activity took place. Read-only.
    activityDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Identity of who performed the action. Read-only.
    actor IdentitySetable;
    // Exposes the driveItem that was the target of this activity.
    driveItem DriveItemable;
}
// NewItemActivity instantiates a new itemActivity and sets the default values.
func NewItemActivity()(*ItemActivity) {
    m := &ItemActivity{
        Entity: *NewEntity(),
    }
    return m
}
// CreateItemActivityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemActivityFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewItemActivity(), nil
}
// GetAccess gets the access property value. An item was accessed.
func (m *ItemActivity) GetAccess()(AccessActionable) {
    if m == nil {
        return nil
    } else {
        return m.access
    }
}
// GetActivityDateTime gets the activityDateTime property value. Details about when the activity took place. Read-only.
func (m *ItemActivity) GetActivityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.activityDateTime
    }
}
// GetActor gets the actor property value. Identity of who performed the action. Read-only.
func (m *ItemActivity) GetActor()(IdentitySetable) {
    if m == nil {
        return nil
    } else {
        return m.actor
    }
}
// GetDriveItem gets the driveItem property value. Exposes the driveItem that was the target of this activity.
func (m *ItemActivity) GetDriveItem()(DriveItemable) {
    if m == nil {
        return nil
    } else {
        return m.driveItem
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemActivity) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["access"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccess(val.(AccessActionable))
        }
        return nil
    }
    res["activityDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivityDateTime(val)
        }
        return nil
    }
    res["actor"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActor(val.(IdentitySetable))
        }
        return nil
    }
    res["driveItem"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateDriveItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDriveItem(val.(DriveItemable))
        }
        return nil
    }
    return res
}
func (m *ItemActivity) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ItemActivity) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("access", m.GetAccess())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("activityDateTime", m.GetActivityDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("actor", m.GetActor())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("driveItem", m.GetDriveItem())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccess sets the access property value. An item was accessed.
func (m *ItemActivity) SetAccess(value AccessActionable)() {
    if m != nil {
        m.access = value
    }
}
// SetActivityDateTime sets the activityDateTime property value. Details about when the activity took place. Read-only.
func (m *ItemActivity) SetActivityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.activityDateTime = value
    }
}
// SetActor sets the actor property value. Identity of who performed the action. Read-only.
func (m *ItemActivity) SetActor(value IdentitySetable)() {
    if m != nil {
        m.actor = value
    }
}
// SetDriveItem sets the driveItem property value. Exposes the driveItem that was the target of this activity.
func (m *ItemActivity) SetDriveItem(value DriveItemable)() {
    if m != nil {
        m.driveItem = value
    }
}
