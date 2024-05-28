package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ManageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilder provides operations to call the recoverPasscode method.
type ManageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilderInternal instantiates a new ManageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilder and sets the default values.
func NewManageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilder) {
    m := &ManageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}/recoverPasscode", pathParameters),
    }
    return m
}
// NewManageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilder instantiates a new ManageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilder and sets the default values.
func NewManageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilderInternal(urlParams, requestAdapter)
}
// Post recover passcode
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-devices-manageddevice-recoverpasscode?view=graph-rest-1.0
func (m *ManageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilder) Post(ctx context.Context, requestConfiguration *ManageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation recover passcode
// returns a *RequestInformation when successful
func (m *ManageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ManageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ManageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilder when successful
func (m *ManageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilder) WithUrl(rawUrl string)(*ManageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilder) {
    return NewManageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
