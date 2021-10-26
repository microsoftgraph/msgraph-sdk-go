package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ItemActivity struct {
    Entity
    // An item was accessed.
    access *AccessAction;
    // Details about when the activity took place. Read-only.
    activityDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Identity of who performed the action. Read-only.
    actor *IdentitySet;
    // Exposes the driveItem that was the target of this activity.
    driveItem *DriveItem;
}
// Instantiates a new itemActivity and sets the default values.
func NewItemActivity()(*ItemActivity) {
    m := &ItemActivity{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the access property value. An item was accessed.
func (m *ItemActivity) GetAccess()(*AccessAction) {
    if m == nil {
        return nil
    } else {
        return m.access
    }
}
// Gets the activityDateTime property value. Details about when the activity took place. Read-only.
func (m *ItemActivity) GetActivityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.activityDateTime
    }
}
// Gets the actor property value. Identity of who performed the action. Read-only.
func (m *ItemActivity) GetActor()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.actor
    }
}
// Gets the driveItem property value. Exposes the driveItem that was the target of this activity.
func (m *ItemActivity) GetDriveItem()(*DriveItem) {
    if m == nil {
        return nil
    } else {
        return m.driveItem
    }
}
// The deserialization information for the current model
func (m *ItemActivity) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["access"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessAction() })
        if err != nil {
            return err
        }
        m.SetAccess(val.(*AccessAction))
        return nil
    }
    res["activityDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetActivityDateTime(val)
        return nil
    }
    res["actor"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        m.SetActor(val.(*IdentitySet))
        return nil
    }
    res["driveItem"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDriveItem() })
        if err != nil {
            return err
        }
        m.SetDriveItem(val.(*DriveItem))
        return nil
    }
    return res
}
func (m *ItemActivity) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the access property value. An item was accessed.
// Parameters:
//  - value : Value to set for the access property.
func (m *ItemActivity) SetAccess(value *AccessAction)() {
    m.access = value
}
// Sets the activityDateTime property value. Details about when the activity took place. Read-only.
// Parameters:
//  - value : Value to set for the activityDateTime property.
func (m *ItemActivity) SetActivityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.activityDateTime = value
}
// Sets the actor property value. Identity of who performed the action. Read-only.
// Parameters:
//  - value : Value to set for the actor property.
func (m *ItemActivity) SetActor(value *IdentitySet)() {
    m.actor = value
}
// Sets the driveItem property value. Exposes the driveItem that was the target of this activity.
// Parameters:
//  - value : Value to set for the driveItem property.
func (m *ItemActivity) SetDriveItem(value *DriveItem)() {
    m.driveItem = value
}
