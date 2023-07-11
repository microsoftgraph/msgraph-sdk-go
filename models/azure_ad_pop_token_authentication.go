package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AzureAdPopTokenAuthentication 
type AzureAdPopTokenAuthentication struct {
    CustomExtensionAuthenticationConfiguration
    // The OdataType property
    OdataType *string
}
// NewAzureAdPopTokenAuthentication instantiates a new azureAdPopTokenAuthentication and sets the default values.
func NewAzureAdPopTokenAuthentication()(*AzureAdPopTokenAuthentication) {
    m := &AzureAdPopTokenAuthentication{
        CustomExtensionAuthenticationConfiguration: *NewCustomExtensionAuthenticationConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.azureAdPopTokenAuthentication"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAzureAdPopTokenAuthenticationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAzureAdPopTokenAuthenticationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAzureAdPopTokenAuthentication(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AzureAdPopTokenAuthentication) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CustomExtensionAuthenticationConfiguration.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *AzureAdPopTokenAuthentication) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CustomExtensionAuthenticationConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// AzureAdPopTokenAuthenticationable 
type AzureAdPopTokenAuthenticationable interface {
    CustomExtensionAuthenticationConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
