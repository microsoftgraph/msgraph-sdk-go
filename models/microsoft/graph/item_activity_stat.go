package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ItemActivityStat struct {
    Entity
    access *ItemActionStat;
    activities []ItemActivity;
    create *ItemActionStat;
    delete *ItemActionStat;
    edit *ItemActionStat;
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    incompleteData *IncompleteData;
    isTrending *bool;
    move *ItemActionStat;
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
func NewItemActivityStat()(*ItemActivityStat) {
    m := &ItemActivityStat{
        Entity: *NewEntity(),
    }
    return m
}
func (m *ItemActivityStat) GetAccess()(*ItemActionStat) {
    if m == nil {
        return nil
    } else {
        return m.access
    }
}
func (m *ItemActivityStat) GetActivities()([]ItemActivity) {
    if m == nil {
        return nil
    } else {
        return m.activities
    }
}
func (m *ItemActivityStat) GetCreate()(*ItemActionStat) {
    if m == nil {
        return nil
    } else {
        return m.create
    }
}
func (m *ItemActivityStat) GetDelete()(*ItemActionStat) {
    if m == nil {
        return nil
    } else {
        return m.delete
    }
}
func (m *ItemActivityStat) GetEdit()(*ItemActionStat) {
    if m == nil {
        return nil
    } else {
        return m.edit
    }
}
func (m *ItemActivityStat) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.endDateTime
    }
}
func (m *ItemActivityStat) GetIncompleteData()(*IncompleteData) {
    if m == nil {
        return nil
    } else {
        return m.incompleteData
    }
}
func (m *ItemActivityStat) GetIsTrending()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isTrending
    }
}
func (m *ItemActivityStat) GetMove()(*ItemActionStat) {
    if m == nil {
        return nil
    } else {
        return m.move
    }
}
func (m *ItemActivityStat) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.startDateTime
    }
}
func (m *ItemActivityStat) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["access"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemActionStat() })
        if err != nil {
            return err
        }
        m.SetAccess(val.(*ItemActionStat))
        return nil
    }
    res["activities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemActivity() })
        if err != nil {
            return err
        }
        res := make([]ItemActivity, len(val))
        for i, v := range val {
            res[i] = *(v.(*ItemActivity))
        }
        m.SetActivities(res)
        return nil
    }
    res["create"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemActionStat() })
        if err != nil {
            return err
        }
        m.SetCreate(val.(*ItemActionStat))
        return nil
    }
    res["delete"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemActionStat() })
        if err != nil {
            return err
        }
        m.SetDelete(val.(*ItemActionStat))
        return nil
    }
    res["edit"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemActionStat() })
        if err != nil {
            return err
        }
        m.SetEdit(val.(*ItemActionStat))
        return nil
    }
    res["endDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetEndDateTime(val)
        return nil
    }
    res["incompleteData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIncompleteData() })
        if err != nil {
            return err
        }
        m.SetIncompleteData(val.(*IncompleteData))
        return nil
    }
    res["isTrending"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsTrending(val)
        return nil
    }
    res["move"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemActionStat() })
        if err != nil {
            return err
        }
        m.SetMove(val.(*ItemActionStat))
        return nil
    }
    res["startDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetStartDateTime(val)
        return nil
    }
    return res
}
func (m *ItemActivityStat) IsNil()(bool) {
    return m == nil
}
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
func (m *ItemActivityStat) SetAccess(value *ItemActionStat)() {
    m.access = value
}
func (m *ItemActivityStat) SetActivities(value []ItemActivity)() {
    m.activities = value
}
func (m *ItemActivityStat) SetCreate(value *ItemActionStat)() {
    m.create = value
}
func (m *ItemActivityStat) SetDelete(value *ItemActionStat)() {
    m.delete = value
}
func (m *ItemActivityStat) SetEdit(value *ItemActionStat)() {
    m.edit = value
}
func (m *ItemActivityStat) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endDateTime = value
}
func (m *ItemActivityStat) SetIncompleteData(value *IncompleteData)() {
    m.incompleteData = value
}
func (m *ItemActivityStat) SetIsTrending(value *bool)() {
    m.isTrending = value
}
func (m *ItemActivityStat) SetMove(value *ItemActionStat)() {
    m.move = value
}
func (m *ItemActivityStat) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
