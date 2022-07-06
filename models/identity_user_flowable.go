package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IdentityUserFlowable 
type IdentityUserFlowable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetType()(*string)
    GetUserFlowType()(*UserFlowType)
    GetUserFlowTypeVersion()(*float32)
    SetType(value *string)()
    SetUserFlowType(value *UserFlowType)()
    SetUserFlowTypeVersion(value *float32)()
}
