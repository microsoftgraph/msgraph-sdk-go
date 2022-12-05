package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsersItemTranslateExchangeIdsResponse provides operations to call the translateExchangeIds method.
type UsersItemTranslateExchangeIdsResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []ConvertIdResultable
}
// NewUsersItemTranslateExchangeIdsResponse instantiates a new UsersItemTranslateExchangeIdsResponse and sets the default values.
func NewUsersItemTranslateExchangeIdsResponse()(*UsersItemTranslateExchangeIdsResponse) {
    m := &UsersItemTranslateExchangeIdsResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateUsersItemTranslateExchangeIdsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsersItemTranslateExchangeIdsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsersItemTranslateExchangeIdsResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsersItemTranslateExchangeIdsResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateConvertIdResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConvertIdResultable, len(val))
            for i, v := range val {
                res[i] = v.(ConvertIdResultable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *UsersItemTranslateExchangeIdsResponse) GetValue()([]ConvertIdResultable) {
    return m.value
}
// Serialize serializes information the current object
func (m *UsersItemTranslateExchangeIdsResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UsersItemTranslateExchangeIdsResponse) SetValue(value []ConvertIdResultable)() {
    m.value = value
}
