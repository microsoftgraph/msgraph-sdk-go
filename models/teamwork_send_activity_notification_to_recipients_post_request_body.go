package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkSendActivityNotificationToRecipientsPostRequestBody provides operations to call the sendActivityNotificationToRecipients method.
type TeamworkSendActivityNotificationToRecipientsPostRequestBody struct {
    // The activityType property
    activityType *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The chainId property
    chainId *int64
    // The previewText property
    previewText ItemBodyable
    // The recipients property
    recipients []TeamworkNotificationRecipientable
    // The teamsAppId property
    teamsAppId *string
    // The templateParameters property
    templateParameters []KeyValuePairable
    // The topic property
    topic TeamworkActivityTopicable
}
// NewTeamworkSendActivityNotificationToRecipientsPostRequestBody instantiates a new TeamworkSendActivityNotificationToRecipientsPostRequestBody and sets the default values.
func NewTeamworkSendActivityNotificationToRecipientsPostRequestBody()(*TeamworkSendActivityNotificationToRecipientsPostRequestBody) {
    m := &TeamworkSendActivityNotificationToRecipientsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTeamworkSendActivityNotificationToRecipientsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkSendActivityNotificationToRecipientsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkSendActivityNotificationToRecipientsPostRequestBody(), nil
}
// GetActivityType gets the activityType property value. The activityType property
func (m *TeamworkSendActivityNotificationToRecipientsPostRequestBody) GetActivityType()(*string) {
    return m.activityType
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkSendActivityNotificationToRecipientsPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetChainId gets the chainId property value. The chainId property
func (m *TeamworkSendActivityNotificationToRecipientsPostRequestBody) GetChainId()(*int64) {
    return m.chainId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkSendActivityNotificationToRecipientsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["activityType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetActivityType)
    res["chainId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetChainId)
    res["previewText"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateItemBodyFromDiscriminatorValue , m.SetPreviewText)
    res["recipients"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateTeamworkNotificationRecipientFromDiscriminatorValue , m.SetRecipients)
    res["teamsAppId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTeamsAppId)
    res["templateParameters"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue , m.SetTemplateParameters)
    res["topic"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateTeamworkActivityTopicFromDiscriminatorValue , m.SetTopic)
    return res
}
// GetPreviewText gets the previewText property value. The previewText property
func (m *TeamworkSendActivityNotificationToRecipientsPostRequestBody) GetPreviewText()(ItemBodyable) {
    return m.previewText
}
// GetRecipients gets the recipients property value. The recipients property
func (m *TeamworkSendActivityNotificationToRecipientsPostRequestBody) GetRecipients()([]TeamworkNotificationRecipientable) {
    return m.recipients
}
// GetTeamsAppId gets the teamsAppId property value. The teamsAppId property
func (m *TeamworkSendActivityNotificationToRecipientsPostRequestBody) GetTeamsAppId()(*string) {
    return m.teamsAppId
}
// GetTemplateParameters gets the templateParameters property value. The templateParameters property
func (m *TeamworkSendActivityNotificationToRecipientsPostRequestBody) GetTemplateParameters()([]KeyValuePairable) {
    return m.templateParameters
}
// GetTopic gets the topic property value. The topic property
func (m *TeamworkSendActivityNotificationToRecipientsPostRequestBody) GetTopic()(TeamworkActivityTopicable) {
    return m.topic
}
// Serialize serializes information the current object
func (m *TeamworkSendActivityNotificationToRecipientsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("activityType", m.GetActivityType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("chainId", m.GetChainId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("previewText", m.GetPreviewText())
        if err != nil {
            return err
        }
    }
    if m.GetRecipients() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetRecipients())
        err := writer.WriteCollectionOfObjectValues("recipients", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("teamsAppId", m.GetTeamsAppId())
        if err != nil {
            return err
        }
    }
    if m.GetTemplateParameters() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTemplateParameters())
        err := writer.WriteCollectionOfObjectValues("templateParameters", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("topic", m.GetTopic())
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
// SetActivityType sets the activityType property value. The activityType property
func (m *TeamworkSendActivityNotificationToRecipientsPostRequestBody) SetActivityType(value *string)() {
    m.activityType = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkSendActivityNotificationToRecipientsPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetChainId sets the chainId property value. The chainId property
func (m *TeamworkSendActivityNotificationToRecipientsPostRequestBody) SetChainId(value *int64)() {
    m.chainId = value
}
// SetPreviewText sets the previewText property value. The previewText property
func (m *TeamworkSendActivityNotificationToRecipientsPostRequestBody) SetPreviewText(value ItemBodyable)() {
    m.previewText = value
}
// SetRecipients sets the recipients property value. The recipients property
func (m *TeamworkSendActivityNotificationToRecipientsPostRequestBody) SetRecipients(value []TeamworkNotificationRecipientable)() {
    m.recipients = value
}
// SetTeamsAppId sets the teamsAppId property value. The teamsAppId property
func (m *TeamworkSendActivityNotificationToRecipientsPostRequestBody) SetTeamsAppId(value *string)() {
    m.teamsAppId = value
}
// SetTemplateParameters sets the templateParameters property value. The templateParameters property
func (m *TeamworkSendActivityNotificationToRecipientsPostRequestBody) SetTemplateParameters(value []KeyValuePairable)() {
    m.templateParameters = value
}
// SetTopic sets the topic property value. The topic property
func (m *TeamworkSendActivityNotificationToRecipientsPostRequestBody) SetTopic(value TeamworkActivityTopicable)() {
    m.topic = value
}
