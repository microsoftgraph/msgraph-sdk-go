package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

type CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityEstimateStatisticsEstimateStatisticsPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewCasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityEstimateStatisticsEstimateStatisticsPostRequestBody instantiates a new CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityEstimateStatisticsEstimateStatisticsPostRequestBody and sets the default values.
func NewCasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityEstimateStatisticsEstimateStatisticsPostRequestBody()(*CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityEstimateStatisticsEstimateStatisticsPostRequestBody) {
    m := &CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityEstimateStatisticsEstimateStatisticsPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityEstimateStatisticsEstimateStatisticsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityEstimateStatisticsEstimateStatisticsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityEstimateStatisticsEstimateStatisticsPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityEstimateStatisticsEstimateStatisticsPostRequestBody) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityEstimateStatisticsEstimateStatisticsPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityEstimateStatisticsEstimateStatisticsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["statisticsOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.ParseStatisticsOptions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatisticsOptions(val.(*idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.StatisticsOptions))
        }
        return nil
    }
    return res
}
// GetStatisticsOptions gets the statisticsOptions property value. The statisticsOptions property
// returns a *StatisticsOptions when successful
func (m *CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityEstimateStatisticsEstimateStatisticsPostRequestBody) GetStatisticsOptions()(*idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.StatisticsOptions) {
    val, err := m.GetBackingStore().Get("statisticsOptions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.StatisticsOptions)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityEstimateStatisticsEstimateStatisticsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetStatisticsOptions() != nil {
        cast := (*m.GetStatisticsOptions()).String()
        err := writer.WriteStringValue("statisticsOptions", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityEstimateStatisticsEstimateStatisticsPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityEstimateStatisticsEstimateStatisticsPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetStatisticsOptions sets the statisticsOptions property value. The statisticsOptions property
func (m *CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityEstimateStatisticsEstimateStatisticsPostRequestBody) SetStatisticsOptions(value *idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.StatisticsOptions)() {
    err := m.GetBackingStore().Set("statisticsOptions", value)
    if err != nil {
        panic(err)
    }
}
type CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityEstimateStatisticsEstimateStatisticsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetStatisticsOptions()(*idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.StatisticsOptions)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetStatisticsOptions(value *idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.StatisticsOptions)()
}
