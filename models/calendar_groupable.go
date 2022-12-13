package models

import (
    i2bacd9b8d8db2e77ee2b5c5ccb19d679c36f920b8fee9d786a0adafff458afcd "github.com/google/UUID"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CalendarGroupable 
type CalendarGroupable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCalendars()([]Calendarable)
    GetChangeKey()(*string)
    GetClassId()(*UUID)
    GetName()(*string)
    SetCalendars(value []Calendarable)()
    SetChangeKey(value *string)()
    SetClassId(value *UUID)()
    SetName(value *string)()
}
