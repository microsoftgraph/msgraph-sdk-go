package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPeerOutlierRecommendationInsightSettings 
type GroupPeerOutlierRecommendationInsightSettings struct {
    AccessReviewRecommendationInsightSetting
}
// NewGroupPeerOutlierRecommendationInsightSettings instantiates a new groupPeerOutlierRecommendationInsightSettings and sets the default values.
func NewGroupPeerOutlierRecommendationInsightSettings()(*GroupPeerOutlierRecommendationInsightSettings) {
    m := &GroupPeerOutlierRecommendationInsightSettings{
        AccessReviewRecommendationInsightSetting: *NewAccessReviewRecommendationInsightSetting(),
    }
    odataTypeValue := "#microsoft.graph.groupPeerOutlierRecommendationInsightSettings"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateGroupPeerOutlierRecommendationInsightSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPeerOutlierRecommendationInsightSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPeerOutlierRecommendationInsightSettings(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPeerOutlierRecommendationInsightSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AccessReviewRecommendationInsightSetting.GetFieldDeserializers()
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *GroupPeerOutlierRecommendationInsightSettings) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *GroupPeerOutlierRecommendationInsightSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AccessReviewRecommendationInsightSetting.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *GroupPeerOutlierRecommendationInsightSettings) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// GroupPeerOutlierRecommendationInsightSettingsable 
type GroupPeerOutlierRecommendationInsightSettingsable interface {
    AccessReviewRecommendationInsightSettingable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    SetOdataType(value *string)()
}
