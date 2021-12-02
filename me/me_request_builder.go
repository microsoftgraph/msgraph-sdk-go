package me

import (
    i010a9f35678b5a912313b7ffeba17283a4e8f5c5e57d7586280f5f222fde8365 "github.com/microsoftgraph/msgraph-sdk-go/me/settings"
    i043275ac95085702ea85fa0a3fb44f05cdf22167fe734032827b9f468ff2e66c "github.com/microsoftgraph/msgraph-sdk-go/me/followedsites"
    i0649e8b11b3ec8278bb4babf63ec9cafa0d10566ccd622ca5149bb1a04573075 "github.com/microsoftgraph/msgraph-sdk-go/me/manageddevices"
    i0aaf43e1d1874d48f4fefa1f3cb2018458e3bf8dbad402db3145aa32207019ca "github.com/microsoftgraph/msgraph-sdk-go/me/outlook"
    i0e15e1c0947fe4cf75784f5461116d0799880f431bc4828322b311b5c56abaed "github.com/microsoftgraph/msgraph-sdk-go/me/contacts"
    i10fecbef712fc029cbc38aae1061d11ef4ed3250124fe680ac304a3f2e9f0a63 "github.com/microsoftgraph/msgraph-sdk-go/me/directreports"
    i1156454f600e0ccae7e57d287a3b71ffac33356633440fdd599e5877247efb98 "github.com/microsoftgraph/msgraph-sdk-go/me/calendars"
    i163a59c87a433f4932bfb97703da0b1e5c2c913c9843ea865155805ec9ebcfbb "github.com/microsoftgraph/msgraph-sdk-go/me/onlinemeetings"
    i206cb0c8230584b060b2abec39bb01624ce85f18583f849d9062afe61183c36a "github.com/microsoftgraph/msgraph-sdk-go/me/changepassword"
    i28ffa81f7ece36b16e8ea949c5fffa192a8af291ec3bad0c76b169480823f154 "github.com/microsoftgraph/msgraph-sdk-go/me/transitivememberof"
    i2c49d98286eefcf44eb33ad01a55d610a378f0b72400964e65dcad0ef245bf34 "github.com/microsoftgraph/msgraph-sdk-go/me/drives"
    i3000cb547f984da6889114dc83c890ea62854b01b8390756a6c2a26c0c46762d "github.com/microsoftgraph/msgraph-sdk-go/me/chats"
    i3057974c2fd0cf7ae8fb427fde3c898513dde98528639b123e00f8faa4959af6 "github.com/microsoftgraph/msgraph-sdk-go/me/events"
    i3404e1dfd2d2c76555cdfc1dd24e427779cd8c615759f7b30fa14193d4b59213 "github.com/microsoftgraph/msgraph-sdk-go/me/removealldevicesfrommanagement"
    i3b9ec4bc140166bd4e87206b29fcc6d952c4572d8450b5035c8862d6f7922016 "github.com/microsoftgraph/msgraph-sdk-go/me/assignlicense"
    i3c451fb64f4e3998638de74865b1ab3e017353fb40ab4413a35316489b770d04 "github.com/microsoftgraph/msgraph-sdk-go/me/manager"
    i400df5b1ed0100dbcae47bf4ceadd56eb71bfe33cebe1e407f4d5bfb8117b817 "github.com/microsoftgraph/msgraph-sdk-go/me/activities"
    i5bd023ee7e5f4d7b9332adc367b19eb2bc99f765bb05c0574828b7f24051c924 "github.com/microsoftgraph/msgraph-sdk-go/me/translateexchangeids"
    i5fad4f46f08ef50fd16a3df6282037cd83e9eed629f5e5fcfc93ca7660a0bef4 "github.com/microsoftgraph/msgraph-sdk-go/me/sendmail"
    i60adae2bbf4cc3014c9132458e3c30e705d3562063f7ecaf9a17d471587d07cf "github.com/microsoftgraph/msgraph-sdk-go/me/createdobjects"
    i622694d3f9d305f2280b1a197d1be3543300e19e21977e6f8512ca71c207fc3b "github.com/microsoftgraph/msgraph-sdk-go/me/getmanagedapppolicies"
    i623b12e632bcd3b4f1cb07d67dd4c2c786876a6ac52e2a6f4bab0fc58c1e6e34 "github.com/microsoftgraph/msgraph-sdk-go/me/managedappregistrations"
    i623b9c23db2c6285bb1586e2570422826e72f619ad0cc7dc9ffd8dccc9c363be "github.com/microsoftgraph/msgraph-sdk-go/me/calendarview"
    i62ed0dc699525e42bbd5c7dfa30d19fdd5a8021fe52c792cf5693a26efb054e4 "github.com/microsoftgraph/msgraph-sdk-go/me/people"
    i63f42124c8863463eb5a7e814116ac3461bb25328a7ff946fbc6ad3a6561ab98 "github.com/microsoftgraph/msgraph-sdk-go/me/presence"
    i67b0755c468e050eae94b1218741505fa70f64cd77602bf39c9076f078a8ee50 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams"
    i6c0a6ed6de8e009fdea4dfe1e4bebefb8e3a9b7d7b00593a3c4993cdc95ecce4 "github.com/microsoftgraph/msgraph-sdk-go/me/licensedetails"
    i6c457664369f83f42fb70aa833daf4cca35b404c6deeb7d3d51e18dd61719e48 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication"
    i6f81259ff975e6273be10a39dd1ba9f46898b0ebf2fc15dc346754ed778b565a "github.com/microsoftgraph/msgraph-sdk-go/me/checkmemberobjects"
    i72454b2e896a5d46ccfa12b2d857be9f68cf4f97b5223742302c23e256370ecf "github.com/microsoftgraph/msgraph-sdk-go/me/getmemberobjects"
    i76e610fff791bfa8235b5911a528dc3f79ad06c1c1c1ee2d1b687504dcbe9474 "github.com/microsoftgraph/msgraph-sdk-go/me/getmailtips"
    i784e87d14910c0a7d5684a2351bbece9b82f859169e26ea19f4515728ce42db8 "github.com/microsoftgraph/msgraph-sdk-go/me/calendar"
    i79b93f38affc8286819ae2ed85e5d90ffbc5173ca9af7d386eb035af519e996a "github.com/microsoftgraph/msgraph-sdk-go/me/restore"
    i7c760ce34c8b15a60dd49909c50aa1a09d44d83d6ecbc77285333fcb80771ca3 "github.com/microsoftgraph/msgraph-sdk-go/me/extensions"
    i8a29b74e3d9357a073ac07e4373df7de12140089e29fff53089c9ea7c9da1be6 "github.com/microsoftgraph/msgraph-sdk-go/me/scopedrolememberof"
    i8fd348dda44ec8b066f0a5c3314ed97e4aba19fcc10c4c24274bec15a7d0fabd "github.com/microsoftgraph/msgraph-sdk-go/me/photo"
    i92e1bd77977f5be2fc73e3f2fbef2c3c51fe1de08b40c1d2b0d160c2fd3b1ec1 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups"
    i9887a3e336e586ff10eb9d0f624bbd9a4dd403d5586ca363a584ead190a55dd7 "github.com/microsoftgraph/msgraph-sdk-go/me/getmanagedappdiagnosticstatuses"
    i992c83c1628215a736c63aded31bdac51b6e05091f0a4780b9b6f37c699c24de "github.com/microsoftgraph/msgraph-sdk-go/me/getmembergroups"
    ia2c7815040bb4d52869b360e0b685da7f1a2717ac3e9decf8fccf4fdbbf31105 "github.com/microsoftgraph/msgraph-sdk-go/me/owneddevices"
    ia686aac8ef57825eafac8bf8956b37566d374128ec8d308a924e2493cce557c3 "github.com/microsoftgraph/msgraph-sdk-go/me/inferenceclassification"
    ia6c56ed70fd384389e50daa416af9b5762eca2c2398c3b08fe3a68b284ecb5ab "github.com/microsoftgraph/msgraph-sdk-go/me/onenote"
    ia995f6de67e6b602301051afac2755890a1af2e62b60ca01287f4a9f83b832cf "github.com/microsoftgraph/msgraph-sdk-go/me/agreementacceptances"
    ia9ef5226776ef05cff35a5a1a349d35741713774b97ff5f32ff743a2b019374c "github.com/microsoftgraph/msgraph-sdk-go/me/drive"
    iacc1ce00670f788919ef9a4db6beddef046cc7ece12addc6ba8af8fc04816e37 "github.com/microsoftgraph/msgraph-sdk-go/me/wipemanagedappregistrationsbydevicetag"
    ib34d54476bd7544828484c44dcea2ea9cb30a97c654291c38e3f683a739c8b88 "github.com/microsoftgraph/msgraph-sdk-go/me/ownedobjects"
    ib49902c927635d59f18b2fe0a7c9354fed84e6a801f167ab11920e706632d9bf "github.com/microsoftgraph/msgraph-sdk-go/me/reminderviewwithstartdatetimewithenddatetime"
    ibdc8482a4a0b860817f919057540f0ded62ef795b16107be11a838b262800802 "github.com/microsoftgraph/msgraph-sdk-go/me/devicemanagementtroubleshootingevents"
    ibff01a43a16df733abd3d1ac62583695fd73d1dd94feea5a43cad8fe202778cf "github.com/microsoftgraph/msgraph-sdk-go/me/findmeetingtimes"
    id62186df8bd24482efb58f5c6df32ea36d60861e1b0a5e18bb0fb8e93180ed0d "github.com/microsoftgraph/msgraph-sdk-go/me/contactfolders"
    id644801acb03b1bcc93277226e7d17fc3f0ab5b86cbcf0fc70a01d1b0446d745 "github.com/microsoftgraph/msgraph-sdk-go/me/reprocesslicenseassignment"
    id6eb6f70f32d9764e00ef9e87b0a8af8c1984d084fe87615f652ec10373b9325 "github.com/microsoftgraph/msgraph-sdk-go/me/todo"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    idb73949c839574cb1e040cbb4d45fd2e7af3c987f54f680fda5cc38f289fd606 "github.com/microsoftgraph/msgraph-sdk-go/me/teamwork"
    idea9ec8c60200f05e5423dbb896eb1bcf6a22107969d3eeab299cc6a76d874c3 "github.com/microsoftgraph/msgraph-sdk-go/me/exportpersonaldata"
    idec091ad007a3e12a149ec3f0e0dffee68ef12bce854b22de68e4463b4dc78df "github.com/microsoftgraph/msgraph-sdk-go/me/photos"
    ie5dfb4f09b7e53c1fe101d19d424cbec858c4a3b073bee55e2c974ad3c279614 "github.com/microsoftgraph/msgraph-sdk-go/me/insights"
    ieaba2dfa6d4e6328c4835d71f5a6ea14a337fbef194276a65c8f0e500bf02b11 "github.com/microsoftgraph/msgraph-sdk-go/me/memberof"
    iebf9e2b48049592a7c437cecc7853290076c9452eccfeb5b10f9f2bc0504fdfa "github.com/microsoftgraph/msgraph-sdk-go/me/checkmembergroups"
    iec29d828a5054453503e707bcd7facaa5d237b91dc87d7bfa7a5c7c2924a2feb "github.com/microsoftgraph/msgraph-sdk-go/me/messages"
    iefdd1942e52abc032856b0c8e7df1198ee77c343ca4b63b4334fe23327191948 "github.com/microsoftgraph/msgraph-sdk-go/me/approleassignments"
    if432e7eef537979da87e810b17f3a6bb87d25c6ee3487a4a9123658a7fabffea "github.com/microsoftgraph/msgraph-sdk-go/me/registereddevices"
    if574edc8d2af1f0fb214c0892ea6797d513abb833e76d156eb15607ade861039 "github.com/microsoftgraph/msgraph-sdk-go/me/planner"
    if7fed2ae20f77809a300e075348d5bea3d8904437f11a11dfb7b13fd95f43cfe "github.com/microsoftgraph/msgraph-sdk-go/me/mailfolders"
    ifb8ef527308910b38e34eeeff4d009a9151d9c282792f85326e91aa83c8fce74 "github.com/microsoftgraph/msgraph-sdk-go/me/revokesigninsessions"
    ifd5ccc5b410dff5303605b874d2eab1857dd0322743b8273f59c50f6cdc7eccc "github.com/microsoftgraph/msgraph-sdk-go/me/oauth2permissiongrants"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i0e5e8b891a423bb213a9ca59b700130801c575d572de2f19f7ef7c3dec985001 "github.com/microsoftgraph/msgraph-sdk-go/me/manageddevices/item"
    i23e227c47e59513561e1310dee7f275d72e661826989ad5fb79f0410e26e3d6d "github.com/microsoftgraph/msgraph-sdk-go/me/photos/item"
    i409a18a3611aab89ea3fdfbbfcd81dee9c098de84cb70a324adb81efe77cd9c9 "github.com/microsoftgraph/msgraph-sdk-go/me/events/item"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i50623afd83f34872b2858002ef89a796dbcb413e984eb8f2c393ab63dfe9f627 "github.com/microsoftgraph/msgraph-sdk-go/me/chats/item"
    i50dd9a7afd373a7096c01b2a6c52a0dc97b9ce9b8050e77ebbec5f2113eeedfb "github.com/microsoftgraph/msgraph-sdk-go/me/extensions/item"
    i6db87d2607e360763d6144f96a5539d25e575d719b1c6486ff32ca05194b469f "github.com/microsoftgraph/msgraph-sdk-go/me/calendars/item"
    i75076f3180c03011b591579c7879bb208b53bd7df65ecaec246e426b96222960 "github.com/microsoftgraph/msgraph-sdk-go/me/contactfolders/item"
    i752ffd2149afbb59609f4e3b13e45c1e9596654430b65978023435d0b0bf8870 "github.com/microsoftgraph/msgraph-sdk-go/me/approleassignments/item"
    i82519bf2607f9e93061888d9cc2d4e3080d112fa88638604197007e67f12d566 "github.com/microsoftgraph/msgraph-sdk-go/me/activities/item"
    i83de839160b1a76c562f0c577e0521385e7ededa8ac09093fcb0042ae0b34038 "github.com/microsoftgraph/msgraph-sdk-go/me/devicemanagementtroubleshootingevents/item"
    i97ba4ab89e2ae86eddef1645331bb1ff7664a5d8975c392bde50d7d40498c4d7 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item"
    i9c1bc246a87f6ba4dcf4c0f9c4f30c60d893d596cd2e949badd25f8199ed6cef "github.com/microsoftgraph/msgraph-sdk-go/me/onlinemeetings/item"
    ia1410212e65c13be6cba0cefbf2258deb7d5d955a0132ae6350cb69d39b4f134 "github.com/microsoftgraph/msgraph-sdk-go/me/people/item"
    iacc2566a5831cf5ddad547cea0842bc8e654bd4e9694e5687421afcbc7870697 "github.com/microsoftgraph/msgraph-sdk-go/me/calendargroups/item"
    ib6af81a46e1cd2294b1d59edaed878023b799ed71ffaeb0e543a0c7b2a9c45a0 "github.com/microsoftgraph/msgraph-sdk-go/me/messages/item"
    ib6fc3cd91dac93a9c6c6225979a445d5d12b9cbde994609403d5f99724571a99 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item"
    ic04fc81ee412d734a4c3dfd98f60894e358d83f93c390e56d414c05ff05e1c03 "github.com/microsoftgraph/msgraph-sdk-go/me/mailfolders/item"
    ic34e2b1716720766154a4200263bd99d00617b2dbbdacdf92c9bdd0f2abc909d "github.com/microsoftgraph/msgraph-sdk-go/me/calendarview/item"
    icfc73b24da43bea4379998bbe38e31b2617fa5b3e777573a24b5722fdae583b3 "github.com/microsoftgraph/msgraph-sdk-go/me/scopedrolememberof/item"
    iedeff048dd4d8b91b3d6b0f2ea6810bff050db01d964ad2520525eab286ef758 "github.com/microsoftgraph/msgraph-sdk-go/me/contacts/item"
)

// MeRequestBuilder builds and executes requests for operations under \me
type MeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// MeRequestBuilderGetOptions options for Get
type MeRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *MeRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MeRequestBuilderGetQueryParameters get me
type MeRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// MeRequestBuilderPatchOptions options for Patch
type MeRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.User;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *MeRequestBuilder) Activities()(*i400df5b1ed0100dbcae47bf4ceadd56eb71bfe33cebe1e407f4d5bfb8117b817.ActivitiesRequestBuilder) {
    return i400df5b1ed0100dbcae47bf4ceadd56eb71bfe33cebe1e407f4d5bfb8117b817.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.activities.item collection
func (m *MeRequestBuilder) ActivitiesById(id string)(*i82519bf2607f9e93061888d9cc2d4e3080d112fa88638604197007e67f12d566.UserActivityRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userActivity_id"] = id
    }
    return i82519bf2607f9e93061888d9cc2d4e3080d112fa88638604197007e67f12d566.NewUserActivityRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MeRequestBuilder) AgreementAcceptances()(*ia995f6de67e6b602301051afac2755890a1af2e62b60ca01287f4a9f83b832cf.AgreementAcceptancesRequestBuilder) {
    return ia995f6de67e6b602301051afac2755890a1af2e62b60ca01287f4a9f83b832cf.NewAgreementAcceptancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MeRequestBuilder) AppRoleAssignments()(*iefdd1942e52abc032856b0c8e7df1198ee77c343ca4b63b4334fe23327191948.AppRoleAssignmentsRequestBuilder) {
    return iefdd1942e52abc032856b0c8e7df1198ee77c343ca4b63b4334fe23327191948.NewAppRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppRoleAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.appRoleAssignments.item collection
func (m *MeRequestBuilder) AppRoleAssignmentsById(id string)(*i752ffd2149afbb59609f4e3b13e45c1e9596654430b65978023435d0b0bf8870.AppRoleAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appRoleAssignment_id"] = id
    }
    return i752ffd2149afbb59609f4e3b13e45c1e9596654430b65978023435d0b0bf8870.NewAppRoleAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MeRequestBuilder) AssignLicense()(*i3b9ec4bc140166bd4e87206b29fcc6d952c4572d8450b5035c8862d6f7922016.AssignLicenseRequestBuilder) {
    return i3b9ec4bc140166bd4e87206b29fcc6d952c4572d8450b5035c8862d6f7922016.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MeRequestBuilder) Authentication()(*i6c457664369f83f42fb70aa833daf4cca35b404c6deeb7d3d51e18dd61719e48.AuthenticationRequestBuilder) {
    return i6c457664369f83f42fb70aa833daf4cca35b404c6deeb7d3d51e18dd61719e48.NewAuthenticationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MeRequestBuilder) Calendar()(*i784e87d14910c0a7d5684a2351bbece9b82f859169e26ea19f4515728ce42db8.CalendarRequestBuilder) {
    return i784e87d14910c0a7d5684a2351bbece9b82f859169e26ea19f4515728ce42db8.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MeRequestBuilder) CalendarGroups()(*i92e1bd77977f5be2fc73e3f2fbef2c3c51fe1de08b40c1d2b0d160c2fd3b1ec1.CalendarGroupsRequestBuilder) {
    return i92e1bd77977f5be2fc73e3f2fbef2c3c51fe1de08b40c1d2b0d160c2fd3b1ec1.NewCalendarGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarGroupsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendarGroups.item collection
func (m *MeRequestBuilder) CalendarGroupsById(id string)(*iacc2566a5831cf5ddad547cea0842bc8e654bd4e9694e5687421afcbc7870697.CalendarGroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendarGroup_id"] = id
    }
    return iacc2566a5831cf5ddad547cea0842bc8e654bd4e9694e5687421afcbc7870697.NewCalendarGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MeRequestBuilder) Calendars()(*i1156454f600e0ccae7e57d287a3b71ffac33356633440fdd599e5877247efb98.CalendarsRequestBuilder) {
    return i1156454f600e0ccae7e57d287a3b71ffac33356633440fdd599e5877247efb98.NewCalendarsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendars.item collection
func (m *MeRequestBuilder) CalendarsById(id string)(*i6db87d2607e360763d6144f96a5539d25e575d719b1c6486ff32ca05194b469f.CalendarRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendar_id"] = id
    }
    return i6db87d2607e360763d6144f96a5539d25e575d719b1c6486ff32ca05194b469f.NewCalendarRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MeRequestBuilder) CalendarView()(*i623b9c23db2c6285bb1586e2570422826e72f619ad0cc7dc9ffd8dccc9c363be.CalendarViewRequestBuilder) {
    return i623b9c23db2c6285bb1586e2570422826e72f619ad0cc7dc9ffd8dccc9c363be.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarViewById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.calendarView.item collection
func (m *MeRequestBuilder) CalendarViewById(id string)(*ic34e2b1716720766154a4200263bd99d00617b2dbbdacdf92c9bdd0f2abc909d.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id"] = id
    }
    return ic34e2b1716720766154a4200263bd99d00617b2dbbdacdf92c9bdd0f2abc909d.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MeRequestBuilder) ChangePassword()(*i206cb0c8230584b060b2abec39bb01624ce85f18583f849d9062afe61183c36a.ChangePasswordRequestBuilder) {
    return i206cb0c8230584b060b2abec39bb01624ce85f18583f849d9062afe61183c36a.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MeRequestBuilder) Chats()(*i3000cb547f984da6889114dc83c890ea62854b01b8390756a6c2a26c0c46762d.ChatsRequestBuilder) {
    return i3000cb547f984da6889114dc83c890ea62854b01b8390756a6c2a26c0c46762d.NewChatsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChatsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.chats.item collection
func (m *MeRequestBuilder) ChatsById(id string)(*i50623afd83f34872b2858002ef89a796dbcb413e984eb8f2c393ab63dfe9f627.ChatRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chat_id"] = id
    }
    return i50623afd83f34872b2858002ef89a796dbcb413e984eb8f2c393ab63dfe9f627.NewChatRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MeRequestBuilder) CheckMemberGroups()(*iebf9e2b48049592a7c437cecc7853290076c9452eccfeb5b10f9f2bc0504fdfa.CheckMemberGroupsRequestBuilder) {
    return iebf9e2b48049592a7c437cecc7853290076c9452eccfeb5b10f9f2bc0504fdfa.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MeRequestBuilder) CheckMemberObjects()(*i6f81259ff975e6273be10a39dd1ba9f46898b0ebf2fc15dc346754ed778b565a.CheckMemberObjectsRequestBuilder) {
    return i6f81259ff975e6273be10a39dd1ba9f46898b0ebf2fc15dc346754ed778b565a.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMeRequestBuilderInternal instantiates a new MeRequestBuilder and sets the default values.
func NewMeRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MeRequestBuilder) {
    m := &MeRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeRequestBuilder instantiates a new MeRequestBuilder and sets the default values.
func NewMeRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *MeRequestBuilder) ContactFolders()(*id62186df8bd24482efb58f5c6df32ea36d60861e1b0a5e18bb0fb8e93180ed0d.ContactFoldersRequestBuilder) {
    return id62186df8bd24482efb58f5c6df32ea36d60861e1b0a5e18bb0fb8e93180ed0d.NewContactFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContactFoldersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.contactFolders.item collection
func (m *MeRequestBuilder) ContactFoldersById(id string)(*i75076f3180c03011b591579c7879bb208b53bd7df65ecaec246e426b96222960.ContactFolderRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contactFolder_id"] = id
    }
    return i75076f3180c03011b591579c7879bb208b53bd7df65ecaec246e426b96222960.NewContactFolderRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MeRequestBuilder) Contacts()(*i0e15e1c0947fe4cf75784f5461116d0799880f431bc4828322b311b5c56abaed.ContactsRequestBuilder) {
    return i0e15e1c0947fe4cf75784f5461116d0799880f431bc4828322b311b5c56abaed.NewContactsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContactsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.contacts.item collection
func (m *MeRequestBuilder) ContactsById(id string)(*iedeff048dd4d8b91b3d6b0f2ea6810bff050db01d964ad2520525eab286ef758.ContactRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contact_id"] = id
    }
    return iedeff048dd4d8b91b3d6b0f2ea6810bff050db01d964ad2520525eab286ef758.NewContactRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MeRequestBuilder) CreatedObjects()(*i60adae2bbf4cc3014c9132458e3c30e705d3562063f7ecaf9a17d471587d07cf.CreatedObjectsRequestBuilder) {
    return i60adae2bbf4cc3014c9132458e3c30e705d3562063f7ecaf9a17d471587d07cf.NewCreatedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get me
func (m *MeRequestBuilder) CreateGetRequestInformation(options *MeRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update me
func (m *MeRequestBuilder) CreatePatchRequestInformation(options *MeRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *MeRequestBuilder) DeviceManagementTroubleshootingEvents()(*ibdc8482a4a0b860817f919057540f0ded62ef795b16107be11a838b262800802.DeviceManagementTroubleshootingEventsRequestBuilder) {
    return ibdc8482a4a0b860817f919057540f0ded62ef795b16107be11a838b262800802.NewDeviceManagementTroubleshootingEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceManagementTroubleshootingEventsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.deviceManagementTroubleshootingEvents.item collection
func (m *MeRequestBuilder) DeviceManagementTroubleshootingEventsById(id string)(*i83de839160b1a76c562f0c577e0521385e7ededa8ac09093fcb0042ae0b34038.DeviceManagementTroubleshootingEventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementTroubleshootingEvent_id"] = id
    }
    return i83de839160b1a76c562f0c577e0521385e7ededa8ac09093fcb0042ae0b34038.NewDeviceManagementTroubleshootingEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MeRequestBuilder) DirectReports()(*i10fecbef712fc029cbc38aae1061d11ef4ed3250124fe680ac304a3f2e9f0a63.DirectReportsRequestBuilder) {
    return i10fecbef712fc029cbc38aae1061d11ef4ed3250124fe680ac304a3f2e9f0a63.NewDirectReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MeRequestBuilder) Drive()(*ia9ef5226776ef05cff35a5a1a349d35741713774b97ff5f32ff743a2b019374c.DriveRequestBuilder) {
    return ia9ef5226776ef05cff35a5a1a349d35741713774b97ff5f32ff743a2b019374c.NewDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MeRequestBuilder) Drives()(*i2c49d98286eefcf44eb33ad01a55d610a378f0b72400964e65dcad0ef245bf34.DrivesRequestBuilder) {
    return i2c49d98286eefcf44eb33ad01a55d610a378f0b72400964e65dcad0ef245bf34.NewDrivesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DrivesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.drives.item collection
func (m *MeRequestBuilder) DrivesById(id string)(*i97ba4ab89e2ae86eddef1645331bb1ff7664a5d8975c392bde50d7d40498c4d7.DriveRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["drive_id"] = id
    }
    return i97ba4ab89e2ae86eddef1645331bb1ff7664a5d8975c392bde50d7d40498c4d7.NewDriveRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MeRequestBuilder) Events()(*i3057974c2fd0cf7ae8fb427fde3c898513dde98528639b123e00f8faa4959af6.EventsRequestBuilder) {
    return i3057974c2fd0cf7ae8fb427fde3c898513dde98528639b123e00f8faa4959af6.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EventsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.events.item collection
func (m *MeRequestBuilder) EventsById(id string)(*i409a18a3611aab89ea3fdfbbfcd81dee9c098de84cb70a324adb81efe77cd9c9.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id"] = id
    }
    return i409a18a3611aab89ea3fdfbbfcd81dee9c098de84cb70a324adb81efe77cd9c9.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MeRequestBuilder) ExportPersonalData()(*idea9ec8c60200f05e5423dbb896eb1bcf6a22107969d3eeab299cc6a76d874c3.ExportPersonalDataRequestBuilder) {
    return idea9ec8c60200f05e5423dbb896eb1bcf6a22107969d3eeab299cc6a76d874c3.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MeRequestBuilder) Extensions()(*i7c760ce34c8b15a60dd49909c50aa1a09d44d83d6ecbc77285333fcb80771ca3.ExtensionsRequestBuilder) {
    return i7c760ce34c8b15a60dd49909c50aa1a09d44d83d6ecbc77285333fcb80771ca3.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.extensions.item collection
func (m *MeRequestBuilder) ExtensionsById(id string)(*i50dd9a7afd373a7096c01b2a6c52a0dc97b9ce9b8050e77ebbec5f2113eeedfb.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i50dd9a7afd373a7096c01b2a6c52a0dc97b9ce9b8050e77ebbec5f2113eeedfb.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MeRequestBuilder) FindMeetingTimes()(*ibff01a43a16df733abd3d1ac62583695fd73d1dd94feea5a43cad8fe202778cf.FindMeetingTimesRequestBuilder) {
    return ibff01a43a16df733abd3d1ac62583695fd73d1dd94feea5a43cad8fe202778cf.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MeRequestBuilder) FollowedSites()(*i043275ac95085702ea85fa0a3fb44f05cdf22167fe734032827b9f468ff2e66c.FollowedSitesRequestBuilder) {
    return i043275ac95085702ea85fa0a3fb44f05cdf22167fe734032827b9f468ff2e66c.NewFollowedSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get me
func (m *MeRequestBuilder) Get(options *MeRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.User, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewUser() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.User), nil
}
func (m *MeRequestBuilder) GetMailTips()(*i76e610fff791bfa8235b5911a528dc3f79ad06c1c1c1ee2d1b687504dcbe9474.GetMailTipsRequestBuilder) {
    return i76e610fff791bfa8235b5911a528dc3f79ad06c1c1c1ee2d1b687504dcbe9474.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses builds and executes requests for operations under \me\microsoft.graph.getManagedAppDiagnosticStatuses()
func (m *MeRequestBuilder) GetManagedAppDiagnosticStatuses()(*i9887a3e336e586ff10eb9d0f624bbd9a4dd403d5586ca363a584ead190a55dd7.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return i9887a3e336e586ff10eb9d0f624bbd9a4dd403d5586ca363a584ead190a55dd7.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies builds and executes requests for operations under \me\microsoft.graph.getManagedAppPolicies()
func (m *MeRequestBuilder) GetManagedAppPolicies()(*i622694d3f9d305f2280b1a197d1be3543300e19e21977e6f8512ca71c207fc3b.GetManagedAppPoliciesRequestBuilder) {
    return i622694d3f9d305f2280b1a197d1be3543300e19e21977e6f8512ca71c207fc3b.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MeRequestBuilder) GetMemberGroups()(*i992c83c1628215a736c63aded31bdac51b6e05091f0a4780b9b6f37c699c24de.GetMemberGroupsRequestBuilder) {
    return i992c83c1628215a736c63aded31bdac51b6e05091f0a4780b9b6f37c699c24de.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MeRequestBuilder) GetMemberObjects()(*i72454b2e896a5d46ccfa12b2d857be9f68cf4f97b5223742302c23e256370ecf.GetMemberObjectsRequestBuilder) {
    return i72454b2e896a5d46ccfa12b2d857be9f68cf4f97b5223742302c23e256370ecf.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MeRequestBuilder) InferenceClassification()(*ia686aac8ef57825eafac8bf8956b37566d374128ec8d308a924e2493cce557c3.InferenceClassificationRequestBuilder) {
    return ia686aac8ef57825eafac8bf8956b37566d374128ec8d308a924e2493cce557c3.NewInferenceClassificationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MeRequestBuilder) Insights()(*ie5dfb4f09b7e53c1fe101d19d424cbec858c4a3b073bee55e2c974ad3c279614.InsightsRequestBuilder) {
    return ie5dfb4f09b7e53c1fe101d19d424cbec858c4a3b073bee55e2c974ad3c279614.NewInsightsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MeRequestBuilder) JoinedTeams()(*i67b0755c468e050eae94b1218741505fa70f64cd77602bf39c9076f078a8ee50.JoinedTeamsRequestBuilder) {
    return i67b0755c468e050eae94b1218741505fa70f64cd77602bf39c9076f078a8ee50.NewJoinedTeamsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// JoinedTeamsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.joinedTeams.item collection
func (m *MeRequestBuilder) JoinedTeamsById(id string)(*ib6fc3cd91dac93a9c6c6225979a445d5d12b9cbde994609403d5f99724571a99.TeamRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["team_id"] = id
    }
    return ib6fc3cd91dac93a9c6c6225979a445d5d12b9cbde994609403d5f99724571a99.NewTeamRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MeRequestBuilder) LicenseDetails()(*i6c0a6ed6de8e009fdea4dfe1e4bebefb8e3a9b7d7b00593a3c4993cdc95ecce4.LicenseDetailsRequestBuilder) {
    return i6c0a6ed6de8e009fdea4dfe1e4bebefb8e3a9b7d7b00593a3c4993cdc95ecce4.NewLicenseDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LicenseDetailsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.licenseDetails.item collection
func (m *MeRequestBuilder) LicenseDetailsById(id string)(*i6c0a6ed6de8e009fdea4dfe1e4bebefb8e3a9b7d7b00593a3c4993cdc95ecce4.LicenseDetailsRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["licenseDetails_id"] = id
    }
    return i6c0a6ed6de8e009fdea4dfe1e4bebefb8e3a9b7d7b00593a3c4993cdc95ecce4.NewLicenseDetailsRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MeRequestBuilder) MailFolders()(*if7fed2ae20f77809a300e075348d5bea3d8904437f11a11dfb7b13fd95f43cfe.MailFoldersRequestBuilder) {
    return if7fed2ae20f77809a300e075348d5bea3d8904437f11a11dfb7b13fd95f43cfe.NewMailFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MailFoldersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.mailFolders.item collection
func (m *MeRequestBuilder) MailFoldersById(id string)(*ic04fc81ee412d734a4c3dfd98f60894e358d83f93c390e56d414c05ff05e1c03.MailFolderRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mailFolder_id"] = id
    }
    return ic04fc81ee412d734a4c3dfd98f60894e358d83f93c390e56d414c05ff05e1c03.NewMailFolderRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MeRequestBuilder) ManagedAppRegistrations()(*i623b12e632bcd3b4f1cb07d67dd4c2c786876a6ac52e2a6f4bab0fc58c1e6e34.ManagedAppRegistrationsRequestBuilder) {
    return i623b12e632bcd3b4f1cb07d67dd4c2c786876a6ac52e2a6f4bab0fc58c1e6e34.NewManagedAppRegistrationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MeRequestBuilder) ManagedDevices()(*i0649e8b11b3ec8278bb4babf63ec9cafa0d10566ccd622ca5149bb1a04573075.ManagedDevicesRequestBuilder) {
    return i0649e8b11b3ec8278bb4babf63ec9cafa0d10566ccd622ca5149bb1a04573075.NewManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDevicesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.managedDevices.item collection
func (m *MeRequestBuilder) ManagedDevicesById(id string)(*i0e5e8b891a423bb213a9ca59b700130801c575d572de2f19f7ef7c3dec985001.ManagedDeviceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDevice_id"] = id
    }
    return i0e5e8b891a423bb213a9ca59b700130801c575d572de2f19f7ef7c3dec985001.NewManagedDeviceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MeRequestBuilder) Manager()(*i3c451fb64f4e3998638de74865b1ab3e017353fb40ab4413a35316489b770d04.ManagerRequestBuilder) {
    return i3c451fb64f4e3998638de74865b1ab3e017353fb40ab4413a35316489b770d04.NewManagerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MeRequestBuilder) MemberOf()(*ieaba2dfa6d4e6328c4835d71f5a6ea14a337fbef194276a65c8f0e500bf02b11.MemberOfRequestBuilder) {
    return ieaba2dfa6d4e6328c4835d71f5a6ea14a337fbef194276a65c8f0e500bf02b11.NewMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MeRequestBuilder) Messages()(*iec29d828a5054453503e707bcd7facaa5d237b91dc87d7bfa7a5c7c2924a2feb.MessagesRequestBuilder) {
    return iec29d828a5054453503e707bcd7facaa5d237b91dc87d7bfa7a5c7c2924a2feb.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.messages.item collection
func (m *MeRequestBuilder) MessagesById(id string)(*ib6af81a46e1cd2294b1d59edaed878023b799ed71ffaeb0e543a0c7b2a9c45a0.MessageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["message_id"] = id
    }
    return ib6af81a46e1cd2294b1d59edaed878023b799ed71ffaeb0e543a0c7b2a9c45a0.NewMessageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MeRequestBuilder) Oauth2PermissionGrants()(*ifd5ccc5b410dff5303605b874d2eab1857dd0322743b8273f59c50f6cdc7eccc.Oauth2PermissionGrantsRequestBuilder) {
    return ifd5ccc5b410dff5303605b874d2eab1857dd0322743b8273f59c50f6cdc7eccc.NewOauth2PermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MeRequestBuilder) Onenote()(*ia6c56ed70fd384389e50daa416af9b5762eca2c2398c3b08fe3a68b284ecb5ab.OnenoteRequestBuilder) {
    return ia6c56ed70fd384389e50daa416af9b5762eca2c2398c3b08fe3a68b284ecb5ab.NewOnenoteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MeRequestBuilder) OnlineMeetings()(*i163a59c87a433f4932bfb97703da0b1e5c2c913c9843ea865155805ec9ebcfbb.OnlineMeetingsRequestBuilder) {
    return i163a59c87a433f4932bfb97703da0b1e5c2c913c9843ea865155805ec9ebcfbb.NewOnlineMeetingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OnlineMeetingsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.onlineMeetings.item collection
func (m *MeRequestBuilder) OnlineMeetingsById(id string)(*i9c1bc246a87f6ba4dcf4c0f9c4f30c60d893d596cd2e949badd25f8199ed6cef.OnlineMeetingRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onlineMeeting_id"] = id
    }
    return i9c1bc246a87f6ba4dcf4c0f9c4f30c60d893d596cd2e949badd25f8199ed6cef.NewOnlineMeetingRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MeRequestBuilder) Outlook()(*i0aaf43e1d1874d48f4fefa1f3cb2018458e3bf8dbad402db3145aa32207019ca.OutlookRequestBuilder) {
    return i0aaf43e1d1874d48f4fefa1f3cb2018458e3bf8dbad402db3145aa32207019ca.NewOutlookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MeRequestBuilder) OwnedDevices()(*ia2c7815040bb4d52869b360e0b685da7f1a2717ac3e9decf8fccf4fdbbf31105.OwnedDevicesRequestBuilder) {
    return ia2c7815040bb4d52869b360e0b685da7f1a2717ac3e9decf8fccf4fdbbf31105.NewOwnedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MeRequestBuilder) OwnedObjects()(*ib34d54476bd7544828484c44dcea2ea9cb30a97c654291c38e3f683a739c8b88.OwnedObjectsRequestBuilder) {
    return ib34d54476bd7544828484c44dcea2ea9cb30a97c654291c38e3f683a739c8b88.NewOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update me
func (m *MeRequestBuilder) Patch(options *MeRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *MeRequestBuilder) People()(*i62ed0dc699525e42bbd5c7dfa30d19fdd5a8021fe52c792cf5693a26efb054e4.PeopleRequestBuilder) {
    return i62ed0dc699525e42bbd5c7dfa30d19fdd5a8021fe52c792cf5693a26efb054e4.NewPeopleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PeopleById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.people.item collection
func (m *MeRequestBuilder) PeopleById(id string)(*ia1410212e65c13be6cba0cefbf2258deb7d5d955a0132ae6350cb69d39b4f134.PersonRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["person_id"] = id
    }
    return ia1410212e65c13be6cba0cefbf2258deb7d5d955a0132ae6350cb69d39b4f134.NewPersonRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MeRequestBuilder) Photo()(*i8fd348dda44ec8b066f0a5c3314ed97e4aba19fcc10c4c24274bec15a7d0fabd.PhotoRequestBuilder) {
    return i8fd348dda44ec8b066f0a5c3314ed97e4aba19fcc10c4c24274bec15a7d0fabd.NewPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MeRequestBuilder) Photos()(*idec091ad007a3e12a149ec3f0e0dffee68ef12bce854b22de68e4463b4dc78df.PhotosRequestBuilder) {
    return idec091ad007a3e12a149ec3f0e0dffee68ef12bce854b22de68e4463b4dc78df.NewPhotosRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PhotosById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.photos.item collection
func (m *MeRequestBuilder) PhotosById(id string)(*i23e227c47e59513561e1310dee7f275d72e661826989ad5fb79f0410e26e3d6d.ProfilePhotoRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["profilePhoto_id"] = id
    }
    return i23e227c47e59513561e1310dee7f275d72e661826989ad5fb79f0410e26e3d6d.NewProfilePhotoRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MeRequestBuilder) Planner()(*if574edc8d2af1f0fb214c0892ea6797d513abb833e76d156eb15607ade861039.PlannerRequestBuilder) {
    return if574edc8d2af1f0fb214c0892ea6797d513abb833e76d156eb15607ade861039.NewPlannerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MeRequestBuilder) Presence()(*i63f42124c8863463eb5a7e814116ac3461bb25328a7ff946fbc6ad3a6561ab98.PresenceRequestBuilder) {
    return i63f42124c8863463eb5a7e814116ac3461bb25328a7ff946fbc6ad3a6561ab98.NewPresenceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MeRequestBuilder) RegisteredDevices()(*if432e7eef537979da87e810b17f3a6bb87d25c6ee3487a4a9123658a7fabffea.RegisteredDevicesRequestBuilder) {
    return if432e7eef537979da87e810b17f3a6bb87d25c6ee3487a4a9123658a7fabffea.NewRegisteredDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReminderViewWithStartDateTimeWithEndDateTime builds and executes requests for operations under \me\microsoft.graph.reminderView(StartDateTime='{StartDateTime}',EndDateTime='{EndDateTime}')
func (m *MeRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(startDateTime *string, endDateTime *string)(*ib49902c927635d59f18b2fe0a7c9354fed84e6a801f167ab11920e706632d9bf.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return ib49902c927635d59f18b2fe0a7c9354fed84e6a801f167ab11920e706632d9bf.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, startDateTime, endDateTime);
}
func (m *MeRequestBuilder) RemoveAllDevicesFromManagement()(*i3404e1dfd2d2c76555cdfc1dd24e427779cd8c615759f7b30fa14193d4b59213.RemoveAllDevicesFromManagementRequestBuilder) {
    return i3404e1dfd2d2c76555cdfc1dd24e427779cd8c615759f7b30fa14193d4b59213.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MeRequestBuilder) ReprocessLicenseAssignment()(*id644801acb03b1bcc93277226e7d17fc3f0ab5b86cbcf0fc70a01d1b0446d745.ReprocessLicenseAssignmentRequestBuilder) {
    return id644801acb03b1bcc93277226e7d17fc3f0ab5b86cbcf0fc70a01d1b0446d745.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MeRequestBuilder) Restore()(*i79b93f38affc8286819ae2ed85e5d90ffbc5173ca9af7d386eb035af519e996a.RestoreRequestBuilder) {
    return i79b93f38affc8286819ae2ed85e5d90ffbc5173ca9af7d386eb035af519e996a.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MeRequestBuilder) RevokeSignInSessions()(*ifb8ef527308910b38e34eeeff4d009a9151d9c282792f85326e91aa83c8fce74.RevokeSignInSessionsRequestBuilder) {
    return ifb8ef527308910b38e34eeeff4d009a9151d9c282792f85326e91aa83c8fce74.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MeRequestBuilder) ScopedRoleMemberOf()(*i8a29b74e3d9357a073ac07e4373df7de12140089e29fff53089c9ea7c9da1be6.ScopedRoleMemberOfRequestBuilder) {
    return i8a29b74e3d9357a073ac07e4373df7de12140089e29fff53089c9ea7c9da1be6.NewScopedRoleMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ScopedRoleMemberOfById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.scopedRoleMemberOf.item collection
func (m *MeRequestBuilder) ScopedRoleMemberOfById(id string)(*icfc73b24da43bea4379998bbe38e31b2617fa5b3e777573a24b5722fdae583b3.ScopedRoleMembershipRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["scopedRoleMembership_id"] = id
    }
    return icfc73b24da43bea4379998bbe38e31b2617fa5b3e777573a24b5722fdae583b3.NewScopedRoleMembershipRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MeRequestBuilder) SendMail()(*i5fad4f46f08ef50fd16a3df6282037cd83e9eed629f5e5fcfc93ca7660a0bef4.SendMailRequestBuilder) {
    return i5fad4f46f08ef50fd16a3df6282037cd83e9eed629f5e5fcfc93ca7660a0bef4.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MeRequestBuilder) Settings()(*i010a9f35678b5a912313b7ffeba17283a4e8f5c5e57d7586280f5f222fde8365.SettingsRequestBuilder) {
    return i010a9f35678b5a912313b7ffeba17283a4e8f5c5e57d7586280f5f222fde8365.NewSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MeRequestBuilder) Teamwork()(*idb73949c839574cb1e040cbb4d45fd2e7af3c987f54f680fda5cc38f289fd606.TeamworkRequestBuilder) {
    return idb73949c839574cb1e040cbb4d45fd2e7af3c987f54f680fda5cc38f289fd606.NewTeamworkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MeRequestBuilder) Todo()(*id6eb6f70f32d9764e00ef9e87b0a8af8c1984d084fe87615f652ec10373b9325.TodoRequestBuilder) {
    return id6eb6f70f32d9764e00ef9e87b0a8af8c1984d084fe87615f652ec10373b9325.NewTodoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MeRequestBuilder) TransitiveMemberOf()(*i28ffa81f7ece36b16e8ea949c5fffa192a8af291ec3bad0c76b169480823f154.TransitiveMemberOfRequestBuilder) {
    return i28ffa81f7ece36b16e8ea949c5fffa192a8af291ec3bad0c76b169480823f154.NewTransitiveMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MeRequestBuilder) TranslateExchangeIds()(*i5bd023ee7e5f4d7b9332adc367b19eb2bc99f765bb05c0574828b7f24051c924.TranslateExchangeIdsRequestBuilder) {
    return i5bd023ee7e5f4d7b9332adc367b19eb2bc99f765bb05c0574828b7f24051c924.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MeRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*iacc1ce00670f788919ef9a4db6beddef046cc7ece12addc6ba8af8fc04816e37.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return iacc1ce00670f788919ef9a4db6beddef046cc7ece12addc6ba8af8fc04816e37.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
