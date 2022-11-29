package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupsItemSitesItemListsItemItemsItemGetActivitiesByIntervalResponse provides operations to call the getActivitiesByInterval method.
type GroupsItemSitesItemListsItemItemsItemGetActivitiesByIntervalResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []ItemActivityStatable
}
// NewGroupsItemSitesItemListsItemItemsItemGetActivitiesByIntervalResponse instantiates a new GroupsItemSitesItemListsItemItemsItemGetActivitiesByIntervalResponse and sets the default values.
func NewGroupsItemSitesItemListsItemItemsItemGetActivitiesByIntervalResponse()(*GroupsItemSitesItemListsItemItemsItemGetActivitiesByIntervalResponse) {
    m := &GroupsItemSitesItemListsItemItemsItemGetActivitiesByIntervalResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateGroupsItemSitesItemListsItemItemsItemGetActivitiesByIntervalResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupsItemSitesItemListsItemItemsItemGetActivitiesByIntervalResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupsItemSitesItemListsItemItemsItemGetActivitiesByIntervalResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupsItemSitesItemListsItemItemsItemGetActivitiesByIntervalResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateItemActivityStatFromDiscriminatorValue , m.SetValue)
    return res
}
// GetValue gets the value property value. The value property
func (m *GroupsItemSitesItemListsItemItemsItemGetActivitiesByIntervalResponse) GetValue()([]ItemActivityStatable) {
    return m.value
}
// Serialize serializes information the current object
func (m *GroupsItemSitesItemListsItemItemsItemGetActivitiesByIntervalResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetValue())
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *GroupsItemSitesItemListsItemItemsItemGetActivitiesByIntervalResponse) SetValue(value []ItemActivityStatable)() {
    m.value = value
}
