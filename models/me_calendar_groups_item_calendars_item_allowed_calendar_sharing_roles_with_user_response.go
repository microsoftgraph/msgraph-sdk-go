package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeCalendarGroupsItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse provides operations to call the allowedCalendarSharingRoles method.
type MeCalendarGroupsItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []CalendarRoleType
}
// NewMeCalendarGroupsItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse instantiates a new MeCalendarGroupsItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse and sets the default values.
func NewMeCalendarGroupsItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse()(*MeCalendarGroupsItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse) {
    m := &MeCalendarGroupsItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateMeCalendarGroupsItemCalendarsItemAllowedCalendarSharingRolesWithUserResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeCalendarGroupsItemCalendarsItemAllowedCalendarSharingRolesWithUserResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeCalendarGroupsItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeCalendarGroupsItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfEnumValues(ParseCalendarRoleType , m.SetValue)
    return res
}
// GetValue gets the value property value. The value property
func (m *MeCalendarGroupsItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse) GetValue()([]CalendarRoleType) {
    return m.value
}
// Serialize serializes information the current object
func (m *MeCalendarGroupsItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        err = writer.WriteCollectionOfStringValues("value", SerializeCalendarRoleType(m.GetValue()))
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *MeCalendarGroupsItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse) SetValue(value []CalendarRoleType)() {
    m.value = value
}
