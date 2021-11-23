package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ItemActivityStat 
type ItemActivityStat struct {
    Entity
    // Statistics about the access actions in this interval. Read-only.
    access *ItemActionStat;
    // Exposes the itemActivities represented in this itemActivityStat resource.
    activities []ItemActivity;
    // Statistics about the create actions in this interval. Read-only.
    create *ItemActionStat;
    // Statistics about the delete actions in this interval. Read-only.
    delete *ItemActionStat;
    // Statistics about the edit actions in this interval. Read-only.
    edit *ItemActionStat;
    // When the interval ends. Read-only.
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Indicates that the statistics in this interval are based on incomplete data. Read-only.
    incompleteData *IncompleteData;
    // Indicates whether the item is 'trending.' Read-only.
    isTrending *bool;
    // Statistics about the move actions in this interval. Read-only.
    move *ItemActionStat;
    // When the interval starts. Read-only.
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// NewItemActivityStat instantiates a new itemActivityStat and sets the default values.
func NewItemActivityStat()(*ItemActivityStat) {
    m := &ItemActivityStat{
        Entity: *NewEntity(),
    }
    return m
}
// GetAccess gets the access property value. Statistics about the access actions in this interval. Read-only.
func (m *ItemActivityStat) GetAccess()(*ItemActionStat) {
    if m == nil {
        return nil
    } else {
        return m.access
    }
}
// GetActivities gets the activities property value. Exposes the itemActivities represented in this itemActivityStat resource.
func (m *ItemActivityStat) GetActivities()([]ItemActivity) {
    if m == nil {
        return nil
    } else {
        return m.activities
    }
}
// GetCreate gets the create property value. Statistics about the create actions in this interval. Read-only.
func (m *ItemActivityStat) GetCreate()(*ItemActionStat) {
    if m == nil {
        return nil
    } else {
        return m.create
    }
}
// GetDelete gets the delete property value. Statistics about the delete actions in this interval. Read-only.
func (m *ItemActivityStat) GetDelete()(*ItemActionStat) {
    if m == nil {
        return nil
    } else {
        return m.delete
    }
}
// GetEdit gets the edit property value. Statistics about the edit actions in this interval. Read-only.
func (m *ItemActivityStat) GetEdit()(*ItemActionStat) {
    if m == nil {
        return nil
    } else {
        return m.edit
    }
}
// GetEndDateTime gets the endDateTime property value. When the interval ends. Read-only.
func (m *ItemActivityStat) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.endDateTime
    }
}
// GetIncompleteData gets the incompleteData property value. Indicates that the statistics in this interval are based on incomplete data. Read-only.
func (m *ItemActivityStat) GetIncompleteData()(*IncompleteData) {
    if m == nil {
        return nil
    } else {
        return m.incompleteData
    }
}
// GetIsTrending gets the isTrending property value. Indicates whether the item is 'trending.' Read-only.
func (m *ItemActivityStat) GetIsTrending()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isTrending
    }
}
// GetMove gets the move property value. Statistics about the move actions in this interval. Read-only.
func (m *ItemActivityStat) GetMove()(*ItemActionStat) {
    if m == nil {
        return nil
    } else {
        return m.move
    }
}
// GetStartDateTime gets the startDateTime property value. When the interval starts. Read-only.
func (m *ItemActivityStat) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.startDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemActivityStat) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["access"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemActionStat() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccess(val.(*ItemActionStat))
        }
        return nil
    }
    res["activities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemActivity() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ItemActivity, len(val))
            for i, v := range val {
                res[i] = *(v.(*ItemActivity))
            }
            m.SetActivities(res)
        }
        return nil
    }
    res["create"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemActionStat() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreate(val.(*ItemActionStat))
        }
        return nil
    }
    res["delete"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemActionStat() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDelete(val.(*ItemActionStat))
        }
        return nil
    }
    res["edit"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemActionStat() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdit(val.(*ItemActionStat))
        }
        return nil
    }
    res["endDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDateTime(val)
        }
        return nil
    }
    res["incompleteData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIncompleteData() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncompleteData(val.(*IncompleteData))
        }
        return nil
    }
    res["isTrending"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsTrending(val)
        }
        return nil
    }
    res["move"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemActionStat() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMove(val.(*ItemActionStat))
        }
        return nil
    }
    res["startDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
        }
        return nil
    }
    return res
}
func (m *ItemActivityStat) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ItemActivityStat) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetActivities()))
        for i, v := range m.GetActivities() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("activities", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("create", m.GetCreate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("delete", m.GetDelete())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("edit", m.GetEdit())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("incompleteData", m.GetIncompleteData())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isTrending", m.GetIsTrending())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("move", m.GetMove())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccess sets the access property value. Statistics about the access actions in this interval. Read-only.
func (m *ItemActivityStat) SetAccess(value *ItemActionStat)() {
    m.access = value
}
// SetActivities sets the activities property value. Exposes the itemActivities represented in this itemActivityStat resource.
func (m *ItemActivityStat) SetActivities(value []ItemActivity)() {
    m.activities = value
}
// SetCreate sets the create property value. Statistics about the create actions in this interval. Read-only.
func (m *ItemActivityStat) SetCreate(value *ItemActionStat)() {
    m.create = value
}
// SetDelete sets the delete property value. Statistics about the delete actions in this interval. Read-only.
func (m *ItemActivityStat) SetDelete(value *ItemActionStat)() {
    m.delete = value
}
// SetEdit sets the edit property value. Statistics about the edit actions in this interval. Read-only.
func (m *ItemActivityStat) SetEdit(value *ItemActionStat)() {
    m.edit = value
}
// SetEndDateTime sets the endDateTime property value. When the interval ends. Read-only.
func (m *ItemActivityStat) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endDateTime = value
}
// SetIncompleteData sets the incompleteData property value. Indicates that the statistics in this interval are based on incomplete data. Read-only.
func (m *ItemActivityStat) SetIncompleteData(value *IncompleteData)() {
    m.incompleteData = value
}
// SetIsTrending sets the isTrending property value. Indicates whether the item is 'trending.' Read-only.
func (m *ItemActivityStat) SetIsTrending(value *bool)() {
    m.isTrending = value
}
// SetMove sets the move property value. Statistics about the move actions in this interval. Read-only.
func (m *ItemActivityStat) SetMove(value *ItemActionStat)() {
    m.move = value
}
// SetStartDateTime sets the startDateTime property value. When the interval starts. Read-only.
func (m *ItemActivityStat) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
