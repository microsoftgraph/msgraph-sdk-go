package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ItemActivityStat provides operations to manage the collection of group entities.
type ItemActivityStat struct {
    Entity
    // Statistics about the access actions in this interval. Read-only.
    access ItemActionStatable;
    // Exposes the itemActivities represented in this itemActivityStat resource.
    activities []ItemActivityable;
    // Statistics about the create actions in this interval. Read-only.
    create ItemActionStatable;
    // Statistics about the delete actions in this interval. Read-only.
    delete ItemActionStatable;
    // Statistics about the edit actions in this interval. Read-only.
    edit ItemActionStatable;
    // When the interval ends. Read-only.
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Indicates that the statistics in this interval are based on incomplete data. Read-only.
    incompleteData IncompleteDataable;
    // Indicates whether the item is 'trending.' Read-only.
    isTrending *bool;
    // Statistics about the move actions in this interval. Read-only.
    move ItemActionStatable;
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
// CreateItemActivityStatFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemActivityStatFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewItemActivityStat(), nil
}
// GetAccess gets the access property value. Statistics about the access actions in this interval. Read-only.
func (m *ItemActivityStat) GetAccess()(ItemActionStatable) {
    if m == nil {
        return nil
    } else {
        return m.access
    }
}
// GetActivities gets the activities property value. Exposes the itemActivities represented in this itemActivityStat resource.
func (m *ItemActivityStat) GetActivities()([]ItemActivityable) {
    if m == nil {
        return nil
    } else {
        return m.activities
    }
}
// GetCreate gets the create property value. Statistics about the create actions in this interval. Read-only.
func (m *ItemActivityStat) GetCreate()(ItemActionStatable) {
    if m == nil {
        return nil
    } else {
        return m.create
    }
}
// GetDelete gets the delete property value. Statistics about the delete actions in this interval. Read-only.
func (m *ItemActivityStat) GetDelete()(ItemActionStatable) {
    if m == nil {
        return nil
    } else {
        return m.delete
    }
}
// GetEdit gets the edit property value. Statistics about the edit actions in this interval. Read-only.
func (m *ItemActivityStat) GetEdit()(ItemActionStatable) {
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
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemActivityStat) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["access"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemActionStatFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccess(val.(ItemActionStatable))
        }
        return nil
    }
    res["activities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateItemActivityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ItemActivityable, len(val))
            for i, v := range val {
                res[i] = v.(ItemActivityable)
            }
            m.SetActivities(res)
        }
        return nil
    }
    res["create"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemActionStatFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreate(val.(ItemActionStatable))
        }
        return nil
    }
    res["delete"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemActionStatFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDelete(val.(ItemActionStatable))
        }
        return nil
    }
    res["edit"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemActionStatFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdit(val.(ItemActionStatable))
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
        val, err := n.GetObjectValue(CreateIncompleteDataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncompleteData(val.(IncompleteDataable))
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
        val, err := n.GetObjectValue(CreateItemActionStatFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMove(val.(ItemActionStatable))
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
// GetIncompleteData gets the incompleteData property value. Indicates that the statistics in this interval are based on incomplete data. Read-only.
func (m *ItemActivityStat) GetIncompleteData()(IncompleteDataable) {
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
func (m *ItemActivityStat) GetMove()(ItemActionStatable) {
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
    if m.GetActivities() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetActivities()))
        for i, v := range m.GetActivities() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
func (m *ItemActivityStat) SetAccess(value ItemActionStatable)() {
    if m != nil {
        m.access = value
    }
}
// SetActivities sets the activities property value. Exposes the itemActivities represented in this itemActivityStat resource.
func (m *ItemActivityStat) SetActivities(value []ItemActivityable)() {
    if m != nil {
        m.activities = value
    }
}
// SetCreate sets the create property value. Statistics about the create actions in this interval. Read-only.
func (m *ItemActivityStat) SetCreate(value ItemActionStatable)() {
    if m != nil {
        m.create = value
    }
}
// SetDelete sets the delete property value. Statistics about the delete actions in this interval. Read-only.
func (m *ItemActivityStat) SetDelete(value ItemActionStatable)() {
    if m != nil {
        m.delete = value
    }
}
// SetEdit sets the edit property value. Statistics about the edit actions in this interval. Read-only.
func (m *ItemActivityStat) SetEdit(value ItemActionStatable)() {
    if m != nil {
        m.edit = value
    }
}
// SetEndDateTime sets the endDateTime property value. When the interval ends. Read-only.
func (m *ItemActivityStat) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.endDateTime = value
    }
}
// SetIncompleteData sets the incompleteData property value. Indicates that the statistics in this interval are based on incomplete data. Read-only.
func (m *ItemActivityStat) SetIncompleteData(value IncompleteDataable)() {
    if m != nil {
        m.incompleteData = value
    }
}
// SetIsTrending sets the isTrending property value. Indicates whether the item is 'trending.' Read-only.
func (m *ItemActivityStat) SetIsTrending(value *bool)() {
    if m != nil {
        m.isTrending = value
    }
}
// SetMove sets the move property value. Statistics about the move actions in this interval. Read-only.
func (m *ItemActivityStat) SetMove(value ItemActionStatable)() {
    if m != nil {
        m.move = value
    }
}
// SetStartDateTime sets the startDateTime property value. When the interval starts. Read-only.
func (m *ItemActivityStat) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.startDateTime = value
    }
}
