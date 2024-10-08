package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type ItemFindMeetingTimesPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// ItemFindMeetingTimesPostRequestBody_FindMeetingTimesPostRequestBody_minimumAttendeePercentage composed type wrapper for classes float64, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric, string
type ItemFindMeetingTimesPostRequestBody_FindMeetingTimesPostRequestBody_minimumAttendeePercentage struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewItemFindMeetingTimesPostRequestBody_FindMeetingTimesPostRequestBody_minimumAttendeePercentage instantiates a new ItemFindMeetingTimesPostRequestBody_FindMeetingTimesPostRequestBody_minimumAttendeePercentage and sets the default values.
func NewItemFindMeetingTimesPostRequestBody_FindMeetingTimesPostRequestBody_minimumAttendeePercentage()(*ItemFindMeetingTimesPostRequestBody_FindMeetingTimesPostRequestBody_minimumAttendeePercentage) {
    m := &ItemFindMeetingTimesPostRequestBody_FindMeetingTimesPostRequestBody_minimumAttendeePercentage{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateItemFindMeetingTimesPostRequestBody_FindMeetingTimesPostRequestBody_minimumAttendeePercentageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemFindMeetingTimesPostRequestBody_FindMeetingTimesPostRequestBody_minimumAttendeePercentageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewItemFindMeetingTimesPostRequestBody_FindMeetingTimesPostRequestBody_minimumAttendeePercentage()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    if val, err := parseNode.GetEnumValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ParseReferenceNumeric); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetReferenceNumeric(val)
    } else if val, err := parseNode.GetFloat64Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetDouble(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *ItemFindMeetingTimesPostRequestBody_FindMeetingTimesPostRequestBody_minimumAttendeePercentage) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *ItemFindMeetingTimesPostRequestBody_FindMeetingTimesPostRequestBody_minimumAttendeePercentage) GetDouble()(*float64) {
    val, err := m.GetBackingStore().Get("double")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemFindMeetingTimesPostRequestBody_FindMeetingTimesPostRequestBody_minimumAttendeePercentage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *ItemFindMeetingTimesPostRequestBody_FindMeetingTimesPostRequestBody_minimumAttendeePercentage) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *ItemFindMeetingTimesPostRequestBody_FindMeetingTimesPostRequestBody_minimumAttendeePercentage) GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric) {
    val, err := m.GetBackingStore().Get("referenceNumeric")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    }
    return nil
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *ItemFindMeetingTimesPostRequestBody_FindMeetingTimesPostRequestBody_minimumAttendeePercentage) GetString()(*string) {
    val, err := m.GetBackingStore().Get("string")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ItemFindMeetingTimesPostRequestBody_FindMeetingTimesPostRequestBody_minimumAttendeePercentage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetReferenceNumeric() != nil {
        cast := (*m.GetReferenceNumeric()).String()
        err := writer.WriteStringValue("", &cast)
        if err != nil {
            return err
        }
    } else if m.GetDouble() != nil {
        err := writer.WriteFloat64Value("", m.GetDouble())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *ItemFindMeetingTimesPostRequestBody_FindMeetingTimesPostRequestBody_minimumAttendeePercentage) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *ItemFindMeetingTimesPostRequestBody_FindMeetingTimesPostRequestBody_minimumAttendeePercentage) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric
func (m *ItemFindMeetingTimesPostRequestBody_FindMeetingTimesPostRequestBody_minimumAttendeePercentage) SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *ItemFindMeetingTimesPostRequestBody_FindMeetingTimesPostRequestBody_minimumAttendeePercentage) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
type ItemFindMeetingTimesPostRequestBody_FindMeetingTimesPostRequestBody_minimumAttendeePercentageable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDouble()(*float64)
    GetReferenceNumeric()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDouble(value *float64)()
    SetReferenceNumeric(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceNumeric)()
    SetString(value *string)()
}
// NewItemFindMeetingTimesPostRequestBody instantiates a new ItemFindMeetingTimesPostRequestBody and sets the default values.
func NewItemFindMeetingTimesPostRequestBody()(*ItemFindMeetingTimesPostRequestBody) {
    m := &ItemFindMeetingTimesPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemFindMeetingTimesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemFindMeetingTimesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemFindMeetingTimesPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemFindMeetingTimesPostRequestBody) GetAdditionalData()(map[string]any) {
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
// GetAttendees gets the attendees property value. The attendees property
// returns a []AttendeeBaseable when successful
func (m *ItemFindMeetingTimesPostRequestBody) GetAttendees()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttendeeBaseable) {
    val, err := m.GetBackingStore().Get("attendees")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttendeeBaseable)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *ItemFindMeetingTimesPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemFindMeetingTimesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["attendees"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAttendeeBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttendeeBaseable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttendeeBaseable)
                }
            }
            m.SetAttendees(res)
        }
        return nil
    }
    res["isOrganizerOptional"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsOrganizerOptional(val)
        }
        return nil
    }
    res["locationConstraint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateLocationConstraintFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocationConstraint(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LocationConstraintable))
        }
        return nil
    }
    res["maxCandidates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxCandidates(val)
        }
        return nil
    }
    res["meetingDuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeetingDuration(val)
        }
        return nil
    }
    res["minimumAttendeePercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemFindMeetingTimesPostRequestBody_FindMeetingTimesPostRequestBody_minimumAttendeePercentageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumAttendeePercentage(val.(*ItemFindMeetingTimesPostRequestBody_FindMeetingTimesPostRequestBody_minimumAttendeePercentage))
        }
        return nil
    }
    res["returnSuggestionReasons"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReturnSuggestionReasons(val)
        }
        return nil
    }
    res["timeConstraint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTimeConstraintFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeConstraint(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TimeConstraintable))
        }
        return nil
    }
    return res
}
// GetIsOrganizerOptional gets the isOrganizerOptional property value. The isOrganizerOptional property
// returns a *bool when successful
func (m *ItemFindMeetingTimesPostRequestBody) GetIsOrganizerOptional()(*bool) {
    val, err := m.GetBackingStore().Get("isOrganizerOptional")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLocationConstraint gets the locationConstraint property value. The locationConstraint property
// returns a LocationConstraintable when successful
func (m *ItemFindMeetingTimesPostRequestBody) GetLocationConstraint()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LocationConstraintable) {
    val, err := m.GetBackingStore().Get("locationConstraint")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LocationConstraintable)
    }
    return nil
}
// GetMaxCandidates gets the maxCandidates property value. The maxCandidates property
// returns a *int32 when successful
func (m *ItemFindMeetingTimesPostRequestBody) GetMaxCandidates()(*int32) {
    val, err := m.GetBackingStore().Get("maxCandidates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetMeetingDuration gets the meetingDuration property value. The meetingDuration property
// returns a *ISODuration when successful
func (m *ItemFindMeetingTimesPostRequestBody) GetMeetingDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    val, err := m.GetBackingStore().Get("meetingDuration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    }
    return nil
}
// GetMinimumAttendeePercentage gets the minimumAttendeePercentage property value. The minimumAttendeePercentage property
// returns a ItemFindMeetingTimesPostRequestBody_FindMeetingTimesPostRequestBody_minimumAttendeePercentageable when successful
func (m *ItemFindMeetingTimesPostRequestBody) GetMinimumAttendeePercentage()(ItemFindMeetingTimesPostRequestBody_FindMeetingTimesPostRequestBody_minimumAttendeePercentageable) {
    val, err := m.GetBackingStore().Get("minimumAttendeePercentage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ItemFindMeetingTimesPostRequestBody_FindMeetingTimesPostRequestBody_minimumAttendeePercentageable)
    }
    return nil
}
// GetReturnSuggestionReasons gets the returnSuggestionReasons property value. The returnSuggestionReasons property
// returns a *bool when successful
func (m *ItemFindMeetingTimesPostRequestBody) GetReturnSuggestionReasons()(*bool) {
    val, err := m.GetBackingStore().Get("returnSuggestionReasons")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetTimeConstraint gets the timeConstraint property value. The timeConstraint property
// returns a TimeConstraintable when successful
func (m *ItemFindMeetingTimesPostRequestBody) GetTimeConstraint()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TimeConstraintable) {
    val, err := m.GetBackingStore().Get("timeConstraint")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TimeConstraintable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ItemFindMeetingTimesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAttendees() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAttendees()))
        for i, v := range m.GetAttendees() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("attendees", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isOrganizerOptional", m.GetIsOrganizerOptional())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("locationConstraint", m.GetLocationConstraint())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("maxCandidates", m.GetMaxCandidates())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("meetingDuration", m.GetMeetingDuration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("minimumAttendeePercentage", m.GetMinimumAttendeePercentage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("returnSuggestionReasons", m.GetReturnSuggestionReasons())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("timeConstraint", m.GetTimeConstraint())
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
func (m *ItemFindMeetingTimesPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAttendees sets the attendees property value. The attendees property
func (m *ItemFindMeetingTimesPostRequestBody) SetAttendees(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttendeeBaseable)() {
    err := m.GetBackingStore().Set("attendees", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *ItemFindMeetingTimesPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetIsOrganizerOptional sets the isOrganizerOptional property value. The isOrganizerOptional property
func (m *ItemFindMeetingTimesPostRequestBody) SetIsOrganizerOptional(value *bool)() {
    err := m.GetBackingStore().Set("isOrganizerOptional", value)
    if err != nil {
        panic(err)
    }
}
// SetLocationConstraint sets the locationConstraint property value. The locationConstraint property
func (m *ItemFindMeetingTimesPostRequestBody) SetLocationConstraint(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LocationConstraintable)() {
    err := m.GetBackingStore().Set("locationConstraint", value)
    if err != nil {
        panic(err)
    }
}
// SetMaxCandidates sets the maxCandidates property value. The maxCandidates property
func (m *ItemFindMeetingTimesPostRequestBody) SetMaxCandidates(value *int32)() {
    err := m.GetBackingStore().Set("maxCandidates", value)
    if err != nil {
        panic(err)
    }
}
// SetMeetingDuration sets the meetingDuration property value. The meetingDuration property
func (m *ItemFindMeetingTimesPostRequestBody) SetMeetingDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    err := m.GetBackingStore().Set("meetingDuration", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumAttendeePercentage sets the minimumAttendeePercentage property value. The minimumAttendeePercentage property
func (m *ItemFindMeetingTimesPostRequestBody) SetMinimumAttendeePercentage(value ItemFindMeetingTimesPostRequestBody_FindMeetingTimesPostRequestBody_minimumAttendeePercentageable)() {
    err := m.GetBackingStore().Set("minimumAttendeePercentage", value)
    if err != nil {
        panic(err)
    }
}
// SetReturnSuggestionReasons sets the returnSuggestionReasons property value. The returnSuggestionReasons property
func (m *ItemFindMeetingTimesPostRequestBody) SetReturnSuggestionReasons(value *bool)() {
    err := m.GetBackingStore().Set("returnSuggestionReasons", value)
    if err != nil {
        panic(err)
    }
}
// SetTimeConstraint sets the timeConstraint property value. The timeConstraint property
func (m *ItemFindMeetingTimesPostRequestBody) SetTimeConstraint(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TimeConstraintable)() {
    err := m.GetBackingStore().Set("timeConstraint", value)
    if err != nil {
        panic(err)
    }
}
type ItemFindMeetingTimesPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAttendees()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttendeeBaseable)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetIsOrganizerOptional()(*bool)
    GetLocationConstraint()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LocationConstraintable)
    GetMaxCandidates()(*int32)
    GetMeetingDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetMinimumAttendeePercentage()(ItemFindMeetingTimesPostRequestBody_FindMeetingTimesPostRequestBody_minimumAttendeePercentageable)
    GetReturnSuggestionReasons()(*bool)
    GetTimeConstraint()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TimeConstraintable)
    SetAttendees(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttendeeBaseable)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetIsOrganizerOptional(value *bool)()
    SetLocationConstraint(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LocationConstraintable)()
    SetMaxCandidates(value *int32)()
    SetMeetingDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetMinimumAttendeePercentage(value ItemFindMeetingTimesPostRequestBody_FindMeetingTimesPostRequestBody_minimumAttendeePercentageable)()
    SetReturnSuggestionReasons(value *bool)()
    SetTimeConstraint(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TimeConstraintable)()
}
