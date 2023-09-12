package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SubjectRightsRequestEnumeratedSiteLocation 
type SubjectRightsRequestEnumeratedSiteLocation struct {
    SubjectRightsRequestSiteLocation
}
// NewSubjectRightsRequestEnumeratedSiteLocation instantiates a new subjectRightsRequestEnumeratedSiteLocation and sets the default values.
func NewSubjectRightsRequestEnumeratedSiteLocation()(*SubjectRightsRequestEnumeratedSiteLocation) {
    m := &SubjectRightsRequestEnumeratedSiteLocation{
        SubjectRightsRequestSiteLocation: *NewSubjectRightsRequestSiteLocation(),
    }
    odataTypeValue := "#microsoft.graph.subjectRightsRequestEnumeratedSiteLocation"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateSubjectRightsRequestEnumeratedSiteLocationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSubjectRightsRequestEnumeratedSiteLocationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSubjectRightsRequestEnumeratedSiteLocation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SubjectRightsRequestEnumeratedSiteLocation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.SubjectRightsRequestSiteLocation.GetFieldDeserializers()
    res["urls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetUrls(res)
        }
        return nil
    }
    return res
}
// GetUrls gets the urls property value. The urls property
func (m *SubjectRightsRequestEnumeratedSiteLocation) GetUrls()([]string) {
    val, err := m.GetBackingStore().Get("urls")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *SubjectRightsRequestEnumeratedSiteLocation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.SubjectRightsRequestSiteLocation.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetUrls() != nil {
        err = writer.WriteCollectionOfStringValues("urls", m.GetUrls())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetUrls sets the urls property value. The urls property
func (m *SubjectRightsRequestEnumeratedSiteLocation) SetUrls(value []string)() {
    err := m.GetBackingStore().Set("urls", value)
    if err != nil {
        panic(err)
    }
}
// SubjectRightsRequestEnumeratedSiteLocationable 
type SubjectRightsRequestEnumeratedSiteLocationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    SubjectRightsRequestSiteLocationable
    GetUrls()([]string)
    SetUrls(value []string)()
}
