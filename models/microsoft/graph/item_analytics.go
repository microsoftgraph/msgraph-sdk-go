package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ItemAnalytics provides operations to manage the collection of group entities.
type ItemAnalytics struct {
    Entity
    // 
    allTime ItemActivityStatable;
    // 
    itemActivityStats []ItemActivityStatable;
    // 
    lastSevenDays ItemActivityStatable;
}
// NewItemAnalytics instantiates a new itemAnalytics and sets the default values.
func NewItemAnalytics()(*ItemAnalytics) {
    m := &ItemAnalytics{
        Entity: *NewEntity(),
    }
    return m
}
// CreateItemAnalyticsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemAnalyticsFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewItemAnalytics(), nil
}
// GetAllTime gets the allTime property value. 
func (m *ItemAnalytics) GetAllTime()(ItemActivityStatable) {
    if m == nil {
        return nil
    } else {
        return m.allTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemAnalytics) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemActivityStatFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllTime(val.(ItemActivityStatable))
        }
        return nil
    }
    res["itemActivityStats"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateItemActivityStatFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ItemActivityStatable, len(val))
            for i, v := range val {
                res[i] = v.(ItemActivityStatable)
            }
            m.SetItemActivityStats(res)
        }
        return nil
    }
    res["lastSevenDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemActivityStatFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSevenDays(val.(ItemActivityStatable))
        }
        return nil
    }
    return res
}
// GetItemActivityStats gets the itemActivityStats property value. 
func (m *ItemAnalytics) GetItemActivityStats()([]ItemActivityStatable) {
    if m == nil {
        return nil
    } else {
        return m.itemActivityStats
    }
}
// GetLastSevenDays gets the lastSevenDays property value. 
func (m *ItemAnalytics) GetLastSevenDays()(ItemActivityStatable) {
    if m == nil {
        return nil
    } else {
        return m.lastSevenDays
    }
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
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
func (m *ItemAnalytics) SetAllTime(value ItemActivityStatable)() {
    if m != nil {
        m.allTime = value
    }
}
// SetItemActivityStats sets the itemActivityStats property value. 
func (m *ItemAnalytics) SetItemActivityStats(value []ItemActivityStatable)() {
    if m != nil {
        m.itemActivityStats = value
    }
}
// SetLastSevenDays sets the lastSevenDays property value. 
func (m *ItemAnalytics) SetLastSevenDays(value ItemActivityStatable)() {
    if m != nil {
        m.lastSevenDays = value
    }
}
