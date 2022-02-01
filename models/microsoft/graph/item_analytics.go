package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ItemAnalytics 
type ItemAnalytics struct {
    Entity
    // 
    allTime *ItemActivityStat;
    // 
    itemActivityStats []ItemActivityStat;
    // 
    lastSevenDays *ItemActivityStat;
}
// NewItemAnalytics instantiates a new itemAnalytics and sets the default values.
func NewItemAnalytics()(*ItemAnalytics) {
    m := &ItemAnalytics{
        Entity: *NewEntity(),
    }
    return m
}
// GetAllTime gets the allTime property value. 
func (m *ItemAnalytics) GetAllTime()(*ItemActivityStat) {
    if m == nil {
        return nil
    } else {
        return m.allTime
    }
}
// GetItemActivityStats gets the itemActivityStats property value. 
func (m *ItemAnalytics) GetItemActivityStats()([]ItemActivityStat) {
    if m == nil {
        return nil
    } else {
        return m.itemActivityStats
    }
}
// GetLastSevenDays gets the lastSevenDays property value. 
func (m *ItemAnalytics) GetLastSevenDays()(*ItemActivityStat) {
    if m == nil {
        return nil
    } else {
        return m.lastSevenDays
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemAnalytics) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemActivityStat() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllTime(val.(*ItemActivityStat))
        }
        return nil
    }
    res["itemActivityStats"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemActivityStat() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ItemActivityStat, len(val))
            for i, v := range val {
                res[i] = *(v.(*ItemActivityStat))
            }
            m.SetItemActivityStats(res)
        }
        return nil
    }
    res["lastSevenDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemActivityStat() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSevenDays(val.(*ItemActivityStat))
        }
        return nil
    }
    return res
}
func (m *ItemAnalytics) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
    if m.GetItemActivityStats() != nil {
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
// SetAllTime sets the allTime property value. 
func (m *ItemAnalytics) SetAllTime(value *ItemActivityStat)() {
    if m != nil {
        m.allTime = value
    }
}
// SetItemActivityStats sets the itemActivityStats property value. 
func (m *ItemAnalytics) SetItemActivityStats(value []ItemActivityStat)() {
    if m != nil {
        m.itemActivityStats = value
    }
}
// SetLastSevenDays sets the lastSevenDays property value. 
func (m *ItemAnalytics) SetLastSevenDays(value *ItemActivityStat)() {
    if m != nil {
        m.lastSevenDays = value
    }
}
