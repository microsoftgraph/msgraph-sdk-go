package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ItemAnalytics struct {
    Entity
    allTime *ItemActivityStat;
    itemActivityStats []ItemActivityStat;
    lastSevenDays *ItemActivityStat;
}
func NewItemAnalytics()(*ItemAnalytics) {
    m := &ItemAnalytics{
        Entity: *NewEntity(),
    }
    return m
}
func (m *ItemAnalytics) GetAllTime()(*ItemActivityStat) {
    if m == nil {
        return nil
    } else {
        return m.allTime
    }
}
func (m *ItemAnalytics) GetItemActivityStats()([]ItemActivityStat) {
    if m == nil {
        return nil
    } else {
        return m.itemActivityStats
    }
}
func (m *ItemAnalytics) GetLastSevenDays()(*ItemActivityStat) {
    if m == nil {
        return nil
    } else {
        return m.lastSevenDays
    }
}
func (m *ItemAnalytics) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemActivityStat() })
        if err != nil {
            return err
        }
        m.SetAllTime(val.(*ItemActivityStat))
        return nil
    }
    res["itemActivityStats"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemActivityStat() })
        if err != nil {
            return err
        }
        res := make([]ItemActivityStat, len(val))
        for i, v := range val {
            res[i] = *(v.(*ItemActivityStat))
        }
        m.SetItemActivityStats(res)
        return nil
    }
    res["lastSevenDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemActivityStat() })
        if err != nil {
            return err
        }
        m.SetLastSevenDays(val.(*ItemActivityStat))
        return nil
    }
    return res
}
func (m *ItemAnalytics) IsNil()(bool) {
    return m == nil
}
func (m *ItemAnalytics) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("allTime", m.GetAllTime())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetItemActivityStats()))
        for i, v := range m.GetItemActivityStats() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("itemActivityStats", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastSevenDays", m.GetLastSevenDays())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *ItemAnalytics) SetAllTime(value *ItemActivityStat)() {
    m.allTime = value
}
func (m *ItemAnalytics) SetItemActivityStats(value []ItemActivityStat)() {
    m.itemActivityStats = value
}
func (m *ItemAnalytics) SetLastSevenDays(value *ItemActivityStat)() {
    m.lastSevenDays = value
}
