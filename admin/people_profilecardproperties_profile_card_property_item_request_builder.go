package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// PeopleProfilecardpropertiesProfileCardPropertyItemRequestBuilder provides operations to manage the profileCardProperties property of the microsoft.graph.peopleAdminSettings entity.
type PeopleProfilecardpropertiesProfileCardPropertyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PeopleProfilecardpropertiesProfileCardPropertyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PeopleProfilecardpropertiesProfileCardPropertyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PeopleProfilecardpropertiesProfileCardPropertyItemRequestBuilderGetQueryParameters retrieve the properties of a profileCardProperty entity. The profileCardProperty is identified by its directoryPropertyName property.
type PeopleProfilecardpropertiesProfileCardPropertyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PeopleProfilecardpropertiesProfileCardPropertyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PeopleProfilecardpropertiesProfileCardPropertyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PeopleProfilecardpropertiesProfileCardPropertyItemRequestBuilderGetQueryParameters
}
// PeopleProfilecardpropertiesProfileCardPropertyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PeopleProfilecardpropertiesProfileCardPropertyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPeopleProfilecardpropertiesProfileCardPropertyItemRequestBuilderInternal instantiates a new PeopleProfilecardpropertiesProfileCardPropertyItemRequestBuilder and sets the default values.
func NewPeopleProfilecardpropertiesProfileCardPropertyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PeopleProfilecardpropertiesProfileCardPropertyItemRequestBuilder) {
    m := &PeopleProfilecardpropertiesProfileCardPropertyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/people/profileCardProperties/{profileCardProperty%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewPeopleProfilecardpropertiesProfileCardPropertyItemRequestBuilder instantiates a new PeopleProfilecardpropertiesProfileCardPropertyItemRequestBuilder and sets the default values.
func NewPeopleProfilecardpropertiesProfileCardPropertyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PeopleProfilecardpropertiesProfileCardPropertyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPeopleProfilecardpropertiesProfileCardPropertyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete the profileCardProperty object specified by its directoryPropertyName from the organization's profile card, and remove any localized customizations for that property.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/profilecardproperty-delete?view=graph-rest-1.0
func (m *PeopleProfilecardpropertiesProfileCardPropertyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *PeopleProfilecardpropertiesProfileCardPropertyItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// Get retrieve the properties of a profileCardProperty entity. The profileCardProperty is identified by its directoryPropertyName property.
// returns a ProfileCardPropertyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/profilecardproperty-get?view=graph-rest-1.0
func (m *PeopleProfilecardpropertiesProfileCardPropertyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *PeopleProfilecardpropertiesProfileCardPropertyItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ProfileCardPropertyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateProfileCardPropertyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ProfileCardPropertyable), nil
}
// Patch update the properties of a profileCardProperty object, identified by its directoryPropertyName property.
// returns a ProfileCardPropertyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/profilecardproperty-update?view=graph-rest-1.0
func (m *PeopleProfilecardpropertiesProfileCardPropertyItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ProfileCardPropertyable, requestConfiguration *PeopleProfilecardpropertiesProfileCardPropertyItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ProfileCardPropertyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateProfileCardPropertyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ProfileCardPropertyable), nil
}
// ToDeleteRequestInformation delete the profileCardProperty object specified by its directoryPropertyName from the organization's profile card, and remove any localized customizations for that property.
// returns a *RequestInformation when successful
func (m *PeopleProfilecardpropertiesProfileCardPropertyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *PeopleProfilecardpropertiesProfileCardPropertyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation retrieve the properties of a profileCardProperty entity. The profileCardProperty is identified by its directoryPropertyName property.
// returns a *RequestInformation when successful
func (m *PeopleProfilecardpropertiesProfileCardPropertyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PeopleProfilecardpropertiesProfileCardPropertyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the properties of a profileCardProperty object, identified by its directoryPropertyName property.
// returns a *RequestInformation when successful
func (m *PeopleProfilecardpropertiesProfileCardPropertyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ProfileCardPropertyable, requestConfiguration *PeopleProfilecardpropertiesProfileCardPropertyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *PeopleProfilecardpropertiesProfileCardPropertyItemRequestBuilder when successful
func (m *PeopleProfilecardpropertiesProfileCardPropertyItemRequestBuilder) WithUrl(rawUrl string)(*PeopleProfilecardpropertiesProfileCardPropertyItemRequestBuilder) {
    return NewPeopleProfilecardpropertiesProfileCardPropertyItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
