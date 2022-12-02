package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeOutlookSupportedLanguagesResponse provides operations to call the supportedLanguages method.
type MeOutlookSupportedLanguagesResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []LocaleInfoable
}
// NewMeOutlookSupportedLanguagesResponse instantiates a new MeOutlookSupportedLanguagesResponse and sets the default values.
func NewMeOutlookSupportedLanguagesResponse()(*MeOutlookSupportedLanguagesResponse) {
    m := &MeOutlookSupportedLanguagesResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateMeOutlookSupportedLanguagesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeOutlookSupportedLanguagesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeOutlookSupportedLanguagesResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeOutlookSupportedLanguagesResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLocaleInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LocaleInfoable, len(val))
            for i, v := range val {
                res[i] = v.(LocaleInfoable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *MeOutlookSupportedLanguagesResponse) GetValue()([]LocaleInfoable) {
    return m.value
}
// Serialize serializes information the current object
func (m *MeOutlookSupportedLanguagesResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *MeOutlookSupportedLanguagesResponse) SetValue(value []LocaleInfoable)() {
    m.value = value
}
