package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrintPrintersCreatePostRequestBodyable 
type PrintPrintersCreatePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCertificateSigningRequest()(PrintCertificateSigningRequestable)
    GetConnectorId()(*string)
    GetDisplayName()(*string)
    GetHasPhysicalDevice()(*bool)
    GetManufacturer()(*string)
    GetModel()(*string)
    GetPhysicalDeviceId()(*string)
    SetCertificateSigningRequest(value PrintCertificateSigningRequestable)()
    SetConnectorId(value *string)()
    SetDisplayName(value *string)()
    SetHasPhysicalDevice(value *bool)()
    SetManufacturer(value *string)()
    SetModel(value *string)()
    SetPhysicalDeviceId(value *string)()
}
