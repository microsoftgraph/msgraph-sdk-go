package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OfficeGraphInsights provides operations to manage the collection of application entities.
type OfficeGraphInsights struct {
    Entity
    // Access this property from the derived type itemInsights.
    shared []SharedInsightable
    // Access this property from the derived type itemInsights.
    trending []Trendingable
    // Access this property from the derived type itemInsights.
    used []UsedInsightable
}
// NewOfficeGraphInsights instantiates a new officeGraphInsights and sets the default values.
func NewOfficeGraphInsights()(*OfficeGraphInsights) {
    m := &OfficeGraphInsights{
        Entity: *NewEntity(),
    }
    return m
}
// CreateOfficeGraphInsightsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOfficeGraphInsightsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOfficeGraphInsights(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OfficeGraphInsights) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["shared"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSharedInsightFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SharedInsightable, len(val))
            for i, v := range val {
                res[i] = v.(SharedInsightable)
            }
            m.SetShared(res)
        }
        return nil
    }
    res["trending"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTrendingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Trendingable, len(val))
            for i, v := range val {
                res[i] = v.(Trendingable)
            }
            m.SetTrending(res)
        }
        return nil
    }
    res["used"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUsedInsightFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UsedInsightable, len(val))
            for i, v := range val {
                res[i] = v.(UsedInsightable)
            }
            m.SetUsed(res)
        }
        return nil
    }
    return res
}
// GetShared gets the shared property value. Access this property from the derived type itemInsights.
func (m *OfficeGraphInsights) GetShared()([]SharedInsightable) {
    if m == nil {
        return nil
    } else {
        return m.shared
    }
}
// GetTrending gets the trending property value. Access this property from the derived type itemInsights.
func (m *OfficeGraphInsights) GetTrending()([]Trendingable) {
    if m == nil {
        return nil
    } else {
        return m.trending
    }
}
// GetUsed gets the used property value. Access this property from the derived type itemInsights.
func (m *OfficeGraphInsights) GetUsed()([]UsedInsightable) {
    if m == nil {
        return nil
    } else {
        return m.used
    }
}
// Serialize serializes information the current object
func (m *OfficeGraphInsights) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetShared() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetShared()))
        for i, v := range m.GetShared() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("shared", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTrending() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTrending()))
        for i, v := range m.GetTrending() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("trending", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUsed() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUsed()))
        for i, v := range m.GetUsed() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("used", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetShared sets the shared property value. Access this property from the derived type itemInsights.
func (m *OfficeGraphInsights) SetShared(value []SharedInsightable)() {
    if m != nil {
        m.shared = value
    }
}
// SetTrending sets the trending property value. Access this property from the derived type itemInsights.
func (m *OfficeGraphInsights) SetTrending(value []Trendingable)() {
    if m != nil {
        m.trending = value
    }
}
// SetUsed sets the used property value. Access this property from the derived type itemInsights.
func (m *OfficeGraphInsights) SetUsed(value []UsedInsightable)() {
    if m != nil {
        m.used = value
    }
}
