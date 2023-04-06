package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BrowserSiteHistoryCollectionResponse 
type BrowserSiteHistoryCollectionResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewBrowserSiteHistoryCollectionResponse instantiates a new BrowserSiteHistoryCollectionResponse and sets the default values.
func NewBrowserSiteHistoryCollectionResponse()(*BrowserSiteHistoryCollectionResponse) {
    m := &BrowserSiteHistoryCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateBrowserSiteHistoryCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBrowserSiteHistoryCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBrowserSiteHistoryCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BrowserSiteHistoryCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBrowserSiteHistoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BrowserSiteHistoryable, len(val))
            for i, v := range val {
                res[i] = v.(BrowserSiteHistoryable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *BrowserSiteHistoryCollectionResponse) GetValue()([]BrowserSiteHistoryable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]BrowserSiteHistoryable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *BrowserSiteHistoryCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *BrowserSiteHistoryCollectionResponse) SetValue(value []BrowserSiteHistoryable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// BrowserSiteHistoryCollectionResponseable 
type BrowserSiteHistoryCollectionResponseable interface {
    BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]BrowserSiteHistoryable)
    SetValue(value []BrowserSiteHistoryable)()
}
