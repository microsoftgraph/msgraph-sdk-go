package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerCategoryDescriptionsable 
type PlannerCategoryDescriptionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCategory1()(*string)
    GetCategory2()(*string)
    GetCategory3()(*string)
    GetCategory4()(*string)
    GetCategory5()(*string)
    GetCategory6()(*string)
    SetCategory1(value *string)()
    SetCategory2(value *string)()
    SetCategory3(value *string)()
    SetCategory4(value *string)()
    SetCategory5(value *string)()
    SetCategory6(value *string)()
}
