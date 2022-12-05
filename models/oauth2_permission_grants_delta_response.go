package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Oauth2PermissionGrantsDeltaResponse provides operations to call the delta method.
type Oauth2PermissionGrantsDeltaResponse struct {
    BaseDeltaFunctionResponse
    // The value property
    value []OAuth2PermissionGrantable
}
// NewOauth2PermissionGrantsDeltaResponse instantiates a new Oauth2PermissionGrantsDeltaResponse and sets the default values.
func NewOauth2PermissionGrantsDeltaResponse()(*Oauth2PermissionGrantsDeltaResponse) {
    m := &Oauth2PermissionGrantsDeltaResponse{
        BaseDeltaFunctionResponse: *NewBaseDeltaFunctionResponse(),
    }
    return m
}
// CreateOauth2PermissionGrantsDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOauth2PermissionGrantsDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOauth2PermissionGrantsDeltaResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Oauth2PermissionGrantsDeltaResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseDeltaFunctionResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOAuth2PermissionGrantFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OAuth2PermissionGrantable, len(val))
            for i, v := range val {
                res[i] = v.(OAuth2PermissionGrantable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *Oauth2PermissionGrantsDeltaResponse) GetValue()([]OAuth2PermissionGrantable) {
    return m.value
}
// Serialize serializes information the current object
func (m *Oauth2PermissionGrantsDeltaResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseDeltaFunctionResponse.Serialize(writer)
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
func (m *Oauth2PermissionGrantsDeltaResponse) SetValue(value []OAuth2PermissionGrantable)() {
    m.value = value
}
