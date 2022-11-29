package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsersItemFindMeetingTimesPostRequestBody provides operations to call the findMeetingTimes method.
type UsersItemFindMeetingTimesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The attendees property
    attendees []AttendeeBaseable
    // The isOrganizerOptional property
    isOrganizerOptional *bool
    // The locationConstraint property
    locationConstraint LocationConstraintable
    // The maxCandidates property
    maxCandidates *int32
    // The meetingDuration property
    meetingDuration *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // The minimumAttendeePercentage property
    minimumAttendeePercentage *float64
    // The returnSuggestionReasons property
    returnSuggestionReasons *bool
    // The timeConstraint property
    timeConstraint TimeConstraintable
}
// NewUsersItemFindMeetingTimesPostRequestBody instantiates a new UsersItemFindMeetingTimesPostRequestBody and sets the default values.
func NewUsersItemFindMeetingTimesPostRequestBody()(*UsersItemFindMeetingTimesPostRequestBody) {
    m := &UsersItemFindMeetingTimesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUsersItemFindMeetingTimesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsersItemFindMeetingTimesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsersItemFindMeetingTimesPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UsersItemFindMeetingTimesPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAttendees gets the attendees property value. The attendees property
func (m *UsersItemFindMeetingTimesPostRequestBody) GetAttendees()([]AttendeeBaseable) {
    return m.attendees
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsersItemFindMeetingTimesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["attendees"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAttendeeBaseFromDiscriminatorValue , m.SetAttendees)
    res["isOrganizerOptional"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsOrganizerOptional)
    res["locationConstraint"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateLocationConstraintFromDiscriminatorValue , m.SetLocationConstraint)
    res["maxCandidates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetMaxCandidates)
    res["meetingDuration"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetISODurationValue(m.SetMeetingDuration)
    res["minimumAttendeePercentage"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetMinimumAttendeePercentage)
    res["returnSuggestionReasons"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetReturnSuggestionReasons)
    res["timeConstraint"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateTimeConstraintFromDiscriminatorValue , m.SetTimeConstraint)
    return res
}
// GetIsOrganizerOptional gets the isOrganizerOptional property value. The isOrganizerOptional property
func (m *UsersItemFindMeetingTimesPostRequestBody) GetIsOrganizerOptional()(*bool) {
    return m.isOrganizerOptional
}
// GetLocationConstraint gets the locationConstraint property value. The locationConstraint property
func (m *UsersItemFindMeetingTimesPostRequestBody) GetLocationConstraint()(LocationConstraintable) {
    return m.locationConstraint
}
// GetMaxCandidates gets the maxCandidates property value. The maxCandidates property
func (m *UsersItemFindMeetingTimesPostRequestBody) GetMaxCandidates()(*int32) {
    return m.maxCandidates
}
// GetMeetingDuration gets the meetingDuration property value. The meetingDuration property
func (m *UsersItemFindMeetingTimesPostRequestBody) GetMeetingDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.meetingDuration
}
// GetMinimumAttendeePercentage gets the minimumAttendeePercentage property value. The minimumAttendeePercentage property
func (m *UsersItemFindMeetingTimesPostRequestBody) GetMinimumAttendeePercentage()(*float64) {
    return m.minimumAttendeePercentage
}
// GetReturnSuggestionReasons gets the returnSuggestionReasons property value. The returnSuggestionReasons property
func (m *UsersItemFindMeetingTimesPostRequestBody) GetReturnSuggestionReasons()(*bool) {
    return m.returnSuggestionReasons
}
// GetTimeConstraint gets the timeConstraint property value. The timeConstraint property
func (m *UsersItemFindMeetingTimesPostRequestBody) GetTimeConstraint()(TimeConstraintable) {
    return m.timeConstraint
}
// Serialize serializes information the current object
func (m *UsersItemFindMeetingTimesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAttendees() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAttendees())
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
        err := writer.WriteFloat64Value("minimumAttendeePercentage", m.GetMinimumAttendeePercentage())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UsersItemFindMeetingTimesPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAttendees sets the attendees property value. The attendees property
func (m *UsersItemFindMeetingTimesPostRequestBody) SetAttendees(value []AttendeeBaseable)() {
    m.attendees = value
}
// SetIsOrganizerOptional sets the isOrganizerOptional property value. The isOrganizerOptional property
func (m *UsersItemFindMeetingTimesPostRequestBody) SetIsOrganizerOptional(value *bool)() {
    m.isOrganizerOptional = value
}
// SetLocationConstraint sets the locationConstraint property value. The locationConstraint property
func (m *UsersItemFindMeetingTimesPostRequestBody) SetLocationConstraint(value LocationConstraintable)() {
    m.locationConstraint = value
}
// SetMaxCandidates sets the maxCandidates property value. The maxCandidates property
func (m *UsersItemFindMeetingTimesPostRequestBody) SetMaxCandidates(value *int32)() {
    m.maxCandidates = value
}
// SetMeetingDuration sets the meetingDuration property value. The meetingDuration property
func (m *UsersItemFindMeetingTimesPostRequestBody) SetMeetingDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.meetingDuration = value
}
// SetMinimumAttendeePercentage sets the minimumAttendeePercentage property value. The minimumAttendeePercentage property
func (m *UsersItemFindMeetingTimesPostRequestBody) SetMinimumAttendeePercentage(value *float64)() {
    m.minimumAttendeePercentage = value
}
// SetReturnSuggestionReasons sets the returnSuggestionReasons property value. The returnSuggestionReasons property
func (m *UsersItemFindMeetingTimesPostRequestBody) SetReturnSuggestionReasons(value *bool)() {
    m.returnSuggestionReasons = value
}
// SetTimeConstraint sets the timeConstraint property value. The timeConstraint property
func (m *UsersItemFindMeetingTimesPostRequestBody) SetTimeConstraint(value TimeConstraintable)() {
    m.timeConstraint = value
}
