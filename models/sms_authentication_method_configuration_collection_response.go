package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SmsAuthenticationMethodConfigurationCollectionResponse 
type SmsAuthenticationMethodConfigurationCollectionResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewSmsAuthenticationMethodConfigurationCollectionResponse instantiates a new SmsAuthenticationMethodConfigurationCollectionResponse and sets the default values.
func NewSmsAuthenticationMethodConfigurationCollectionResponse()(*SmsAuthenticationMethodConfigurationCollectionResponse) {
    m := &SmsAuthenticationMethodConfigurationCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateSmsAuthenticationMethodConfigurationCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSmsAuthenticationMethodConfigurationCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSmsAuthenticationMethodConfigurationCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SmsAuthenticationMethodConfigurationCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSmsAuthenticationMethodConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SmsAuthenticationMethodConfigurationable, len(val))
            for i, v := range val {
                res[i] = v.(SmsAuthenticationMethodConfigurationable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *SmsAuthenticationMethodConfigurationCollectionResponse) GetValue()([]SmsAuthenticationMethodConfigurationable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SmsAuthenticationMethodConfigurationable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *SmsAuthenticationMethodConfigurationCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *SmsAuthenticationMethodConfigurationCollectionResponse) SetValue(value []SmsAuthenticationMethodConfigurationable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// SmsAuthenticationMethodConfigurationCollectionResponseable 
type SmsAuthenticationMethodConfigurationCollectionResponseable interface {
    BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]SmsAuthenticationMethodConfigurationable)
    SetValue(value []SmsAuthenticationMethodConfigurationable)()
}
