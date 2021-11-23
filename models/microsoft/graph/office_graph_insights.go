package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OfficeGraphInsights 
type OfficeGraphInsights struct {
    Entity
    // Calculated relationship identifying documents shared with or by the user. This includes URLs, file attachments, and reference attachments to OneDrive for Business and SharePoint files found in Outlook messages and meetings. This also includes URLs and reference attachments to Teams conversations. Ordered by recency of share.
    shared []SharedInsight;
    // Calculated relationship identifying documents trending around a user. Trending documents are calculated based on activity of the user's closest network of people and include files stored in OneDrive for Business and SharePoint. Trending insights help the user to discover potentially useful content that the user has access to, but has never viewed before.
    trending []Trending;
    // Calculated relationship identifying the latest documents viewed or modified by a user, including OneDrive for Business and SharePoint documents, ranked by recency of use.
    used []UsedInsight;
}
// NewOfficeGraphInsights instantiates a new officeGraphInsights and sets the default values.
func NewOfficeGraphInsights()(*OfficeGraphInsights) {
    m := &OfficeGraphInsights{
        Entity: *NewEntity(),
    }
    return m
}
// GetShared gets the shared property value. Calculated relationship identifying documents shared with or by the user. This includes URLs, file attachments, and reference attachments to OneDrive for Business and SharePoint files found in Outlook messages and meetings. This also includes URLs and reference attachments to Teams conversations. Ordered by recency of share.
func (m *OfficeGraphInsights) GetShared()([]SharedInsight) {
    if m == nil {
        return nil
    } else {
        return m.shared
    }
}
// GetTrending gets the trending property value. Calculated relationship identifying documents trending around a user. Trending documents are calculated based on activity of the user's closest network of people and include files stored in OneDrive for Business and SharePoint. Trending insights help the user to discover potentially useful content that the user has access to, but has never viewed before.
func (m *OfficeGraphInsights) GetTrending()([]Trending) {
    if m == nil {
        return nil
    } else {
        return m.trending
    }
}
// GetUsed gets the used property value. Calculated relationship identifying the latest documents viewed or modified by a user, including OneDrive for Business and SharePoint documents, ranked by recency of use.
func (m *OfficeGraphInsights) GetUsed()([]UsedInsight) {
    if m == nil {
        return nil
    } else {
        return m.used
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OfficeGraphInsights) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["shared"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSharedInsight() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SharedInsight, len(val))
            for i, v := range val {
                res[i] = *(v.(*SharedInsight))
            }
            m.SetShared(res)
        }
        return nil
    }
    res["trending"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTrending() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Trending, len(val))
            for i, v := range val {
                res[i] = *(v.(*Trending))
            }
            m.SetTrending(res)
        }
        return nil
    }
    res["used"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUsedInsight() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UsedInsight, len(val))
            for i, v := range val {
                res[i] = *(v.(*UsedInsight))
            }
            m.SetUsed(res)
        }
        return nil
    }
    return res
}
func (m *OfficeGraphInsights) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *OfficeGraphInsights) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetShared()))
        for i, v := range m.GetShared() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("shared", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTrending()))
        for i, v := range m.GetTrending() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("trending", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUsed()))
        for i, v := range m.GetUsed() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("used", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetShared sets the shared property value. Calculated relationship identifying documents shared with or by the user. This includes URLs, file attachments, and reference attachments to OneDrive for Business and SharePoint files found in Outlook messages and meetings. This also includes URLs and reference attachments to Teams conversations. Ordered by recency of share.
func (m *OfficeGraphInsights) SetShared(value []SharedInsight)() {
    m.shared = value
}
// SetTrending sets the trending property value. Calculated relationship identifying documents trending around a user. Trending documents are calculated based on activity of the user's closest network of people and include files stored in OneDrive for Business and SharePoint. Trending insights help the user to discover potentially useful content that the user has access to, but has never viewed before.
func (m *OfficeGraphInsights) SetTrending(value []Trending)() {
    m.trending = value
}
// SetUsed sets the used property value. Calculated relationship identifying the latest documents viewed or modified by a user, including OneDrive for Business and SharePoint documents, ranked by recency of use.
func (m *OfficeGraphInsights) SetUsed(value []UsedInsight)() {
    m.used = value
}
