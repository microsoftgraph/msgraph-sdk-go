package communications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19 "github.com/microsoftgraph/msgraph-sdk-go/models/callrecords"
)

type CallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeGetPstnCallsWithFromDateTimeWithToDateTimeGetResponse struct {
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BaseCollectionPaginationCountResponse
}
// NewCallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeGetPstnCallsWithFromDateTimeWithToDateTimeGetResponse instantiates a new CallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeGetPstnCallsWithFromDateTimeWithToDateTimeGetResponse and sets the default values.
func NewCallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeGetPstnCallsWithFromDateTimeWithToDateTimeGetResponse()(*CallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeGetPstnCallsWithFromDateTimeWithToDateTimeGetResponse) {
    m := &CallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeGetPstnCallsWithFromDateTimeWithToDateTimeGetResponse{
        BaseCollectionPaginationCountResponse: *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateCallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeGetPstnCallsWithFromDateTimeWithToDateTimeGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeGetPstnCallsWithFromDateTimeWithToDateTimeGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeGetPstnCallsWithFromDateTimeWithToDateTimeGetResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeGetPstnCallsWithFromDateTimeWithToDateTimeGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.CreatePstnCallLogRowFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.PstnCallLogRowable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.PstnCallLogRowable)
                }
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
// returns a []PstnCallLogRowable when successful
func (m *CallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeGetPstnCallsWithFromDateTimeWithToDateTimeGetResponse) GetValue()([]iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.PstnCallLogRowable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.PstnCallLogRowable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeGetPstnCallsWithFromDateTimeWithToDateTimeGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *CallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeGetPstnCallsWithFromDateTimeWithToDateTimeGetResponse) SetValue(value []iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.PstnCallLogRowable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
type CallrecordsMicrosoftgraphcallrecordsgetpstncallswithfromdatetimewithtodatetimeGetPstnCallsWithFromDateTimeWithToDateTimeGetResponseable interface {
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.PstnCallLogRowable)
    SetValue(value []iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.PstnCallLogRowable)()
}
