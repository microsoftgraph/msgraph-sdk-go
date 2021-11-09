package reports

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i01986228cfa34d663610609938a20fe1aa899eb3ae22c44d689bf3a7cc4fcec2 "github.com/microsoftgraph/msgraph-sdk-go/reports/getonedriveusageaccountdetailwithdate"
    i028952a5521bf457c8a8796489c7f5a1137cf7b0745234478ff5fd12bb087564 "github.com/microsoftgraph/msgraph-sdk-go/reports/getsharepointsiteusagestoragewithperiod"
    i037933f7467ae74beac89af9fadc6cf1a5a8272d8f21a5c2f98bc8f8e7c28ab6 "github.com/microsoftgraph/msgraph-sdk-go/reports/getskypeforbusinessdeviceusageuserdetailwithdate"
    i03d33a99775bdfdb348179d972066b208c7410da4d93f4776c48bb0c7619f539 "github.com/microsoftgraph/msgraph-sdk-go/reports/getyammerdeviceusagedistributionusercountswithperiod"
    i03fae197628d424772396da17b5102e4676ffef8400787778663d2f1e8b62054 "github.com/microsoftgraph/msgraph-sdk-go/reports/getemailappusageusercountswithperiod"
    i05c8d276b86beee65c00c115937a46c2753c7d328e142fe9afab6ef90b5a1a79 "github.com/microsoftgraph/msgraph-sdk-go/reports/getteamsdeviceusagedistributionusercountswithperiod"
    i05ca24f1e64f29c556c3c67a9024eea8b3aac4d350f05ff2c52f4511bf0b6f33 "github.com/microsoftgraph/msgraph-sdk-go/reports/getskypeforbusinessparticipantactivityusercountswithperiod"
    i066e5ea9f7e2ed1b047023128aab6caef0db934d50a71df44b7b6cf7bfb5a10e "github.com/microsoftgraph/msgraph-sdk-go/reports/getskypeforbusinessorganizeractivityusercountswithperiod"
    i069326085999a3926845312c56f68613c15a7dcb468bf7f0592670ea8f5896aa "github.com/microsoftgraph/msgraph-sdk-go/reports/getskypeforbusinessactivityusercountswithperiod"
    i0c4b2ea99f0789f3909bf8d293ade8bd9216c7010a8e0dc72ceea54c21c1d481 "github.com/microsoftgraph/msgraph-sdk-go/reports/getteamsuseractivityuserdetailwithdate"
    i0d6612898baf60cf3ad09e5ec9209afdfaffc25ddef2fe3fddf028f8f356d2e1 "github.com/microsoftgraph/msgraph-sdk-go/reports/getoffice365activeuserdetailwithdate"
    i0e6a5a79552b1c4826c9418d405e7420746bbd08e606bfd5c84ba1ad674aed8f "github.com/microsoftgraph/msgraph-sdk-go/reports/manageddeviceenrollmentfailuredetails"
    i13fef70f90226be585263be7022fe29fa384901ba2188c080017c9272ca23e99 "github.com/microsoftgraph/msgraph-sdk-go/reports/getyammeractivityusercountswithperiod"
    i1bfeeda14fdcd3c4bc1c9b02a4d925837c77654498184deaffbfeb8f496c10c5 "github.com/microsoftgraph/msgraph-sdk-go/reports/getsharepointactivityfilecountswithperiod"
    i1ccdf78d8963c250d79813dd0917eaaaee794135fc39e35801b01ea5f2005034 "github.com/microsoftgraph/msgraph-sdk-go/reports/getonedriveactivityusercountswithperiod"
    i1d9fef6f0fe1efedede6850168705608b3b5be81427b11c9c3d64b7dc8c126ed "github.com/microsoftgraph/msgraph-sdk-go/reports/dailyprintusagebyuser"
    i200c880280f3a34037634650d2cd9969f9a07037f773af21ebb6141e40dde23b "github.com/microsoftgraph/msgraph-sdk-go/reports/getskypeforbusinesspeertopeeractivitycountswithperiod"
    i21ef6a6d7bc99608d2279d5cc2343f5300ae76d3d739202a0c1bc4ff2d13a391 "github.com/microsoftgraph/msgraph-sdk-go/reports/getoffice365groupsactivityfilecountswithperiod"
    i2306eba962c786175114d00783028649abfdfb48f769241b2a7dca15db108361 "github.com/microsoftgraph/msgraph-sdk-go/reports/getoffice365servicesusercountswithperiod"
    i251d35954b406eaf53009c768700ddbf5c49cf2961b8eba3836e2905c19fe73a "github.com/microsoftgraph/msgraph-sdk-go/reports/deviceconfigurationuseractivity"
    i2699e05c2942728d3c9dbbec2f7da6269f18ec269bdc96532cb49ca45ea2e03f "github.com/microsoftgraph/msgraph-sdk-go/reports/getsharepointsiteusagefilecountswithperiod"
    i26d1169a1f986183917395a64e40bb9c1870a54b457336ef64e7efa21c214a15 "github.com/microsoftgraph/msgraph-sdk-go/reports/getoffice365groupsactivitystoragewithperiod"
    i285bf21395498d7467b7e1d3f62e43bbbab40d875429e2800566739471aaa586 "github.com/microsoftgraph/msgraph-sdk-go/reports/getemailappusageuserdetailwithperiod"
    i2e849a1f15987664ba1a6648f0c1d0eb9e2ca5965a632cc3761ad14bb91674bf "github.com/microsoftgraph/msgraph-sdk-go/reports/getskypeforbusinessorganizeractivitycountswithperiod"
    i2f04ee022b2f4dd0c31e5c7f89cc28c7d1e36585beb573c9c689634579fb8999 "github.com/microsoftgraph/msgraph-sdk-go/reports/getsharepointsiteusagepageswithperiod"
    i316dae3dc0a83b965993ba7f5a9d3b29d2b5dcbd706e8977d4a6cb34548c5275 "github.com/microsoftgraph/msgraph-sdk-go/reports/getonedriveusageaccountdetailwithperiod"
    i347d730506a06c407296d7557c653293a35a9c38b50c7256c1c7abb4c736ef9a "github.com/microsoftgraph/msgraph-sdk-go/reports/getsharepointsiteusagesitecountswithperiod"
    i34861f91c890c43257a1e1386ba24de1ffe44f842eadf95d13de30d27a852991 "github.com/microsoftgraph/msgraph-sdk-go/reports/getteamsuseractivityusercountswithperiod"
    i374eb5f5ea7befd0f519a8a1c2886282c8ae9910332bd6b507ad8b29eddfdf2d "github.com/microsoftgraph/msgraph-sdk-go/reports/manageddeviceenrollmentfailuredetailswithskipwithtopwithfilterwithskiptoken"
    i37cb418ef8dbdf59e2a836a306ab7fe4b5933e6211360de8f4f313f68496d242 "github.com/microsoftgraph/msgraph-sdk-go/reports/getemailactivityuserdetailwithdate"
    i37d455e922500fc7c5e44df5cc84a103ddbb8470a53860cb0c02e64fd6e8e775 "github.com/microsoftgraph/msgraph-sdk-go/reports/getmailboxusagequotastatusmailboxcountswithperiod"
    i37efe63b21f1e137ba9fdf350d68a656b504d390fe52651da7739ce398ec31c8 "github.com/microsoftgraph/msgraph-sdk-go/reports/getoffice365activationcounts"
    i39400581189e192e8c3927b2ac378c799b30d5cb2c7a932a3e2f9802d2a2f95d "github.com/microsoftgraph/msgraph-sdk-go/reports/getgrouparchivedprintjobswithgroupidwithstartdatetimewithenddatetime"
    i39d4923038966e0aa0bf413db480c947e66bb5fa71c83d893fbfb4aaacc6cfe9 "github.com/microsoftgraph/msgraph-sdk-go/reports/deviceconfigurationdeviceactivity"
    i3ed9439789e0857666ebca3cd16359b4745ffe979a682ea73e6cd26f8054ae21 "github.com/microsoftgraph/msgraph-sdk-go/reports/getprinterarchivedprintjobswithprinteridwithstartdatetimewithenddatetime"
    i3f12902a7f7457a60f97014c6db3b2bd89dac8cabe4d467b766677db3d6029d0 "github.com/microsoftgraph/msgraph-sdk-go/reports/monthlyprintusagebyuser"
    i3f4aebd4163df4cb0a4f38097f15a9f8ba73853688cf5a50b93a222eb4832abe "github.com/microsoftgraph/msgraph-sdk-go/reports/manageddeviceenrollmenttopfailureswithperiod"
    i40f65a5b54f01dfaea303a3582e82bcdaa71c2195cc00b7915431e77f38e9a0f "github.com/microsoftgraph/msgraph-sdk-go/reports/getsharepointactivityusercountswithperiod"
    i416e8a1993d00c5e5b47ab655f48ac49786eb1442a5459b7cb7023ac34c28093 "github.com/microsoftgraph/msgraph-sdk-go/reports/getskypeforbusinesspeertopeeractivityusercountswithperiod"
    i4319947bf5469fb8e87f0bc08088c383a13726c5be51ca3558867ea6abecc315 "github.com/microsoftgraph/msgraph-sdk-go/reports/getoffice365groupsactivitydetailwithdate"
    i4ea09032e0c277ad5c699279a1e67097201d217b8d0aac6b5962420831d787bd "github.com/microsoftgraph/msgraph-sdk-go/reports/getoffice365activeusercountswithperiod"
    i4ea8142505b6b02804cdd135578592c59e1ebd8c244878c8a27bac9df783863b "github.com/microsoftgraph/msgraph-sdk-go/reports/getemailappusageversionsusercountswithperiod"
    i5277312ad2cbf60385fe9448b03ae55d7aefafa14a9c09b135b745563e0dd271 "github.com/microsoftgraph/msgraph-sdk-go/reports/getyammergroupsactivitydetailwithperiod"
    i5445cb98613918d3a7a579fddf521b041c55f06714612e2097ed886a37be4286 "github.com/microsoftgraph/msgraph-sdk-go/reports/getoffice365groupsactivitygroupcountswithperiod"
    i54533b65ae17487eb09581591352cc93d62970fd3000b0af7a1abc12bf038091 "github.com/microsoftgraph/msgraph-sdk-go/reports/getemailactivityusercountswithperiod"
    i55b749bfea07392ccbd597dc8d76a153720e39963b1e66527cac5bb79f8c6b29 "github.com/microsoftgraph/msgraph-sdk-go/reports/getyammergroupsactivitygroupcountswithperiod"
    i56c28dc52ab4ecbfe37baebbfaf6f279f8d130872c0931a194264028d89e2c96 "github.com/microsoftgraph/msgraph-sdk-go/reports/getsharepointactivityuserdetailwithperiod"
    i5be64ff627afac918209c9e795b4c7ad94cf40584d7376fd0bddf4902f737286 "github.com/microsoftgraph/msgraph-sdk-go/reports/getoffice365activationsusercounts"
    i61a78ce1e0b6e2e60e6d50fe39c731a201868c12c59871387bf62c67fa03cc24 "github.com/microsoftgraph/msgraph-sdk-go/reports/getmailboxusagedetailwithperiod"
    i677bcf76c9544b6088f18b0d84351cd554f0537738c43ca67e647fffec66978c "github.com/microsoftgraph/msgraph-sdk-go/reports/getteamsdeviceusageuserdetailwithdate"
    i67d417cc0a32f095a4d9872596c73a5670bab5ca99f2c39b690a3dc3e53dacb6 "github.com/microsoftgraph/msgraph-sdk-go/reports/getoffice365activationsuserdetail"
    i6a26bcea122cbd3594cacac16b0715042c3ca45bb5f889d1fd0bd83d243a6fe3 "github.com/microsoftgraph/msgraph-sdk-go/reports/getskypeforbusinessactivityuserdetailwithdate"
    i6b3e5d0662b4c04bf1f969a3c62a48d50c4f7ecb57a3913782ce351d87a0035e "github.com/microsoftgraph/msgraph-sdk-go/reports/getsharepointactivitypageswithperiod"
    i6c58964b36e86bf47df5ddf5c8161ca0e4621aee61aaa9517f1dded3120cfc8d "github.com/microsoftgraph/msgraph-sdk-go/reports/getskypeforbusinessdeviceusageuserdetailwithperiod"
    i6e07bd3dadcc8aa2d52f07e77f1f35672ac7ce62aa2917c7a9e4e77b8f4c4db5 "github.com/microsoftgraph/msgraph-sdk-go/reports/getyammerdeviceusageuserdetailwithperiod"
    i6e8f025e6b5802865cf6c583d64d93f7e1c603ead045ca53bef1f6b9cde8a5b0 "github.com/microsoftgraph/msgraph-sdk-go/reports/getskypeforbusinessparticipantactivitycountswithperiod"
    i6edf15eaf951cb8d179890d37ee652592111a5bfeb9fd5709d3609c7593bed81 "github.com/microsoftgraph/msgraph-sdk-go/reports/getoffice365activeuserdetailwithperiod"
    i700e0ab2f1ea334038b3293b2396e12432ea3545b07ed127bb12d122bbb6bf90 "github.com/microsoftgraph/msgraph-sdk-go/reports/getmailboxusagemailboxcountswithperiod"
    i75a65b7f6cf316241684f8aeba4d08fb2c2e04274ef9390b6003ff2a3fe87418 "github.com/microsoftgraph/msgraph-sdk-go/reports/getonedriveusageaccountcountswithperiod"
    i7774a89bdd2b6bc2f70c1e7c58f236fd1f3aa68a40a72db4898fd9c5d4f66988 "github.com/microsoftgraph/msgraph-sdk-go/reports/getoffice365groupsactivitydetailwithperiod"
    i7c11427d5ee6b6eb1e984b209bd6cb9cd7704d1573048256590bf48fe10c95c2 "github.com/microsoftgraph/msgraph-sdk-go/reports/getemailappusageappsusercountswithperiod"
    i7ee204dcbd579a0c0d252ddbfcdd01bf9f244b2b38f671f02c4108f867a99c73 "github.com/microsoftgraph/msgraph-sdk-go/reports/getyammergroupsactivitycountswithperiod"
    i8ad77c0341e431729102d8b93f9066b82a7f4fc8ee3e748eeea007bf3d9d5b18 "github.com/microsoftgraph/msgraph-sdk-go/reports/getskypeforbusinessactivitycountswithperiod"
    i8d02e7a0b578c5b4465f96608d724ae82f61e67266f11ff2ab8aad96d55d1ac1 "github.com/microsoftgraph/msgraph-sdk-go/reports/getskypeforbusinessorganizeractivityminutecountswithperiod"
    i9b65430927e689a26a700312296e423fc6e6e6c0bb86f72287b75c7915e695e0 "github.com/microsoftgraph/msgraph-sdk-go/reports/getteamsuseractivitycountswithperiod"
    i9b77e3c3d15971fe31fa9e23d6bdc4c74a584c922585117306c4db805ac4131c "github.com/microsoftgraph/msgraph-sdk-go/reports/getyammerdeviceusageusercountswithperiod"
    i9d1992f642df096181dd0b243f5639c29682f838689705077b65251f267d222f "github.com/microsoftgraph/msgraph-sdk-go/reports/manageddeviceenrollmenttopfailures"
    i9de2dd5236ac703b10206d8522e5a8f5490288eadb62a3fbf88d9becd9f061bf "github.com/microsoftgraph/msgraph-sdk-go/reports/getskypeforbusinessdeviceusageusercountswithperiod"
    ia3da511007b4da9fa7b31d42bcef05972fa35df25142f53ebb0f2944e3932886 "github.com/microsoftgraph/msgraph-sdk-go/reports/getyammergroupsactivitydetailwithdate"
    ia5f41691c1fbeb62b6b1de131fa27f7306db515c8759a9e49ec88553c0e6d97d "github.com/microsoftgraph/msgraph-sdk-go/reports/getskypeforbusinessdeviceusagedistributionusercountswithperiod"
    ia79ac9c43dc0218a9c3f469a247120979433dd0e3e53de3949570224b3cf3936 "github.com/microsoftgraph/msgraph-sdk-go/reports/getskypeforbusinessactivityuserdetailwithperiod"
    iaaacf2b7aacf4d5f8107fa18d6f0d6cb33066e3335d8ab8cc1a2adfa1f4e3c42 "github.com/microsoftgraph/msgraph-sdk-go/reports/getteamsdeviceusageuserdetailwithperiod"
    iab84fe97b770990c561344b8be4caaca7e6d962145fd1ad9fc9e851313b38857 "github.com/microsoftgraph/msgraph-sdk-go/reports/getonedriveactivityuserdetailwithperiod"
    iafadefc9c12601a3611573bb976b73b251f9ea476a68e3bdc9ceab9e0d5cc308 "github.com/microsoftgraph/msgraph-sdk-go/reports/getuserarchivedprintjobswithuseridwithstartdatetimewithenddatetime"
    ib3bef6d1d1d843d8e20e58b41c29e9b82bb511ac97e381039f942ffee79a3b9a "github.com/microsoftgraph/msgraph-sdk-go/reports/getsharepointactivityuserdetailwithdate"
    ib9c6325caeea63e2dd57f99af77223cbb370ce799b27f8b8a3f6a3f987a9f9fe "github.com/microsoftgraph/msgraph-sdk-go/reports/getsharepointsiteusagedetailwithperiod"
    ibc84667b2a2b840609232b734d0a2873e66315fbe63aeadb3c3d468ad9f84d75 "github.com/microsoftgraph/msgraph-sdk-go/reports/getonedriveusagefilecountswithperiod"
    ibfd05176830665cf1dae7366e8b491572f134d04e12120a5499c283d6c8ae062 "github.com/microsoftgraph/msgraph-sdk-go/reports/getteamsuseractivityuserdetailwithperiod"
    ic5508dd2223dfecdf5814372064e28b262f8f8bfe997c49795b7f76e30889837 "github.com/microsoftgraph/msgraph-sdk-go/reports/getskypeforbusinesspeertopeeractivityminutecountswithperiod"
    ic5ac4372354eec289834ac2ae20a4b14244024792aa3c3ad01b0f0a204b1e81b "github.com/microsoftgraph/msgraph-sdk-go/reports/getsharepointsiteusagedetailwithdate"
    ic78e2d968f685321dafd3ca273401e484836d7ca84e4da07ae045edca4ca55a2 "github.com/microsoftgraph/msgraph-sdk-go/reports/dailyprintusagebyprinter"
    id05b47283d832e55b732685b30d918ee0e303ae639c6447cfff5c7a7089741c7 "github.com/microsoftgraph/msgraph-sdk-go/reports/getmailboxusagestoragewithperiod"
    id379a8afbc53e3411a43fbddcd71ea21f90a2a71ada609569a5c9040ebee5f51 "github.com/microsoftgraph/msgraph-sdk-go/reports/getyammeractivityuserdetailwithperiod"
    id7d2a94ac1ea97cc24d2cd9cb47182b168535107d0d09ad275c4e955f417e029 "github.com/microsoftgraph/msgraph-sdk-go/reports/getyammeractivityuserdetailwithdate"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    idcca069eb97aac5c90ca4ddb83fee2301a14619967cfdadbb5d0a81effc5cc59 "github.com/microsoftgraph/msgraph-sdk-go/reports/getskypeforbusinessparticipantactivityminutecountswithperiod"
    ie49c06630ffed10bc5667592819d2e6ba25594d178187e0c69a98a6cc5ca2e88 "github.com/microsoftgraph/msgraph-sdk-go/reports/getemailappusageuserdetailwithdate"
    ie6b56eee43945ce24ada2464a67b7217385822eb4143ed753098debeb7fa82eb "github.com/microsoftgraph/msgraph-sdk-go/reports/getemailactivityuserdetailwithperiod"
    iec7256d00a8e37c8dec628d1626f62f1224b45c25e3c24418830b8fede7bc006 "github.com/microsoftgraph/msgraph-sdk-go/reports/getonedriveactivityfilecountswithperiod"
    ieef21740e587e5f9d695c0f6753ee2fd4da00a0ced4e73355cae511e81a4e3cf "github.com/microsoftgraph/msgraph-sdk-go/reports/getemailactivitycountswithperiod"
    if1205fb664936127831759fefe4c2c43852474a3fc0419f51e89975f0f281d32 "github.com/microsoftgraph/msgraph-sdk-go/reports/monthlyprintusagebyprinter"
    if97af71691601e0f0046051c8e82fb71f8db2f1d3b5f8e7e40346d3fa47b196f "github.com/microsoftgraph/msgraph-sdk-go/reports/getonedriveusagestoragewithperiod"
    if97f8cb01f2d6a02697169f75d571d9a67fc56f2041590c23f83bf2114377bda "github.com/microsoftgraph/msgraph-sdk-go/reports/getyammerdeviceusageuserdetailwithdate"
    ifb70744a42eac9e3054301d8328a61f22ab480c9dcaaaa3f398cf2a26e1b242b "github.com/microsoftgraph/msgraph-sdk-go/reports/getoffice365groupsactivitycountswithperiod"
    ifd45221e0ff465bb6f37fa740578c697634181cf333e3c01e4db49019e2ba5eb "github.com/microsoftgraph/msgraph-sdk-go/reports/getyammeractivitycountswithperiod"
    ife6d24553cc77853635c06a75c3587d7c978531be356ad656116b7742f60e1f7 "github.com/microsoftgraph/msgraph-sdk-go/reports/getonedriveactivityuserdetailwithdate"
    ife99981e20f3609c2111e3e2aff666b997075ba1e7c98ff44f1950f1d6265ccb "github.com/microsoftgraph/msgraph-sdk-go/reports/getteamsdeviceusageusercountswithperiod"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i0fbd1879eec3f09756dabd53c9de402f839fcf5ed7e8be0da0d3e4135b124b0a "github.com/microsoftgraph/msgraph-sdk-go/reports/monthlyprintusagebyprinter/item"
    i48adffda977c4642ea9d1490fef3ceaa45cfab2defb01a97672b497f12eee0bc "github.com/microsoftgraph/msgraph-sdk-go/reports/dailyprintusagebyprinter/item"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i80d61621ea489f37484a66bdee6bdaa9bbb8a0b83a1cc326a225bedfc4c6dd91 "github.com/microsoftgraph/msgraph-sdk-go/reports/dailyprintusagebyuser/item"
    if4346a9921b148666613b317b233f49301e4eef2f130f6b9e02414250e5a2294 "github.com/microsoftgraph/msgraph-sdk-go/reports/monthlyprintusagebyuser/item"
)

// Builds and executes requests for operations under \reports
type ReportsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Get
type ReportsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ReportsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Get reports
type ReportsRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type ReportsRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ReportRoot;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new ReportsRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewReportsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ReportsRequestBuilder) {
    m := &ReportsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new ReportsRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewReportsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ReportsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get reports
// Parameters:
//  - options : Options for the request
func (m *ReportsRequestBuilder) CreateGetRequestInformation(options *ReportsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
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
// Update reports
// Parameters:
//  - options : Options for the request
func (m *ReportsRequestBuilder) CreatePatchRequestInformation(options *ReportsRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ReportsRequestBuilder) DailyPrintUsageByPrinter()(*ic78e2d968f685321dafd3ca273401e484836d7ca84e4da07ae045edca4ca55a2.DailyPrintUsageByPrinterRequestBuilder) {
    return ic78e2d968f685321dafd3ca273401e484836d7ca84e4da07ae045edca4ca55a2.NewDailyPrintUsageByPrinterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.reports.dailyPrintUsageByPrinter.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ReportsRequestBuilder) DailyPrintUsageByPrinterById(id string)(*i48adffda977c4642ea9d1490fef3ceaa45cfab2defb01a97672b497f12eee0bc.PrintUsageByPrinterRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByPrinter_id"] = id
    }
    return i48adffda977c4642ea9d1490fef3ceaa45cfab2defb01a97672b497f12eee0bc.NewPrintUsageByPrinterRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ReportsRequestBuilder) DailyPrintUsageByUser()(*i1d9fef6f0fe1efedede6850168705608b3b5be81427b11c9c3d64b7dc8c126ed.DailyPrintUsageByUserRequestBuilder) {
    return i1d9fef6f0fe1efedede6850168705608b3b5be81427b11c9c3d64b7dc8c126ed.NewDailyPrintUsageByUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.reports.dailyPrintUsageByUser.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ReportsRequestBuilder) DailyPrintUsageByUserById(id string)(*i80d61621ea489f37484a66bdee6bdaa9bbb8a0b83a1cc326a225bedfc4c6dd91.PrintUsageByUserRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByUser_id"] = id
    }
    return i80d61621ea489f37484a66bdee6bdaa9bbb8a0b83a1cc326a225bedfc4c6dd91.NewPrintUsageByUserRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Builds and executes requests for operations under \reports\microsoft.graph.deviceConfigurationDeviceActivity()
func (m *ReportsRequestBuilder) DeviceConfigurationDeviceActivity()(*i39d4923038966e0aa0bf413db480c947e66bb5fa71c83d893fbfb4aaacc6cfe9.DeviceConfigurationDeviceActivityRequestBuilder) {
    return i39d4923038966e0aa0bf413db480c947e66bb5fa71c83d893fbfb4aaacc6cfe9.NewDeviceConfigurationDeviceActivityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \reports\microsoft.graph.deviceConfigurationUserActivity()
func (m *ReportsRequestBuilder) DeviceConfigurationUserActivity()(*i251d35954b406eaf53009c768700ddbf5c49cf2961b8eba3836e2905c19fe73a.DeviceConfigurationUserActivityRequestBuilder) {
    return i251d35954b406eaf53009c768700ddbf5c49cf2961b8eba3836e2905c19fe73a.NewDeviceConfigurationUserActivityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get reports
// Parameters:
//  - options : Options for the request
func (m *ReportsRequestBuilder) Get(options *ReportsRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ReportRoot, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewReportRoot() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ReportRoot), nil
}
// Builds and executes requests for operations under \reports\microsoft.graph.getEmailActivityCounts(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetEmailActivityCountsWithPeriod(period *string)(*ieef21740e587e5f9d695c0f6753ee2fd4da00a0ced4e73355cae511e81a4e3cf.GetEmailActivityCountsWithPeriodRequestBuilder) {
    return ieef21740e587e5f9d695c0f6753ee2fd4da00a0ced4e73355cae511e81a4e3cf.NewGetEmailActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getEmailActivityUserCounts(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetEmailActivityUserCountsWithPeriod(period *string)(*i54533b65ae17487eb09581591352cc93d62970fd3000b0af7a1abc12bf038091.GetEmailActivityUserCountsWithPeriodRequestBuilder) {
    return i54533b65ae17487eb09581591352cc93d62970fd3000b0af7a1abc12bf038091.NewGetEmailActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getEmailActivityUserDetail(date={date})
// Parameters:
//  - date : Usage: date={date}
func (m *ReportsRequestBuilder) GetEmailActivityUserDetailWithDate(date *string)(*i37cb418ef8dbdf59e2a836a306ab7fe4b5933e6211360de8f4f313f68496d242.GetEmailActivityUserDetailWithDateRequestBuilder) {
    return i37cb418ef8dbdf59e2a836a306ab7fe4b5933e6211360de8f4f313f68496d242.NewGetEmailActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getEmailActivityUserDetail(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetEmailActivityUserDetailWithPeriod(period *string)(*ie6b56eee43945ce24ada2464a67b7217385822eb4143ed753098debeb7fa82eb.GetEmailActivityUserDetailWithPeriodRequestBuilder) {
    return ie6b56eee43945ce24ada2464a67b7217385822eb4143ed753098debeb7fa82eb.NewGetEmailActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getEmailAppUsageAppsUserCounts(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetEmailAppUsageAppsUserCountsWithPeriod(period *string)(*i7c11427d5ee6b6eb1e984b209bd6cb9cd7704d1573048256590bf48fe10c95c2.GetEmailAppUsageAppsUserCountsWithPeriodRequestBuilder) {
    return i7c11427d5ee6b6eb1e984b209bd6cb9cd7704d1573048256590bf48fe10c95c2.NewGetEmailAppUsageAppsUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getEmailAppUsageUserCounts(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetEmailAppUsageUserCountsWithPeriod(period *string)(*i03fae197628d424772396da17b5102e4676ffef8400787778663d2f1e8b62054.GetEmailAppUsageUserCountsWithPeriodRequestBuilder) {
    return i03fae197628d424772396da17b5102e4676ffef8400787778663d2f1e8b62054.NewGetEmailAppUsageUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getEmailAppUsageUserDetail(date={date})
// Parameters:
//  - date : Usage: date={date}
func (m *ReportsRequestBuilder) GetEmailAppUsageUserDetailWithDate(date *string)(*ie49c06630ffed10bc5667592819d2e6ba25594d178187e0c69a98a6cc5ca2e88.GetEmailAppUsageUserDetailWithDateRequestBuilder) {
    return ie49c06630ffed10bc5667592819d2e6ba25594d178187e0c69a98a6cc5ca2e88.NewGetEmailAppUsageUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getEmailAppUsageUserDetail(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetEmailAppUsageUserDetailWithPeriod(period *string)(*i285bf21395498d7467b7e1d3f62e43bbbab40d875429e2800566739471aaa586.GetEmailAppUsageUserDetailWithPeriodRequestBuilder) {
    return i285bf21395498d7467b7e1d3f62e43bbbab40d875429e2800566739471aaa586.NewGetEmailAppUsageUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getEmailAppUsageVersionsUserCounts(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetEmailAppUsageVersionsUserCountsWithPeriod(period *string)(*i4ea8142505b6b02804cdd135578592c59e1ebd8c244878c8a27bac9df783863b.GetEmailAppUsageVersionsUserCountsWithPeriodRequestBuilder) {
    return i4ea8142505b6b02804cdd135578592c59e1ebd8c244878c8a27bac9df783863b.NewGetEmailAppUsageVersionsUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getGroupArchivedPrintJobs(groupId='{groupId}',startDateTime={startDateTime},endDateTime={endDateTime})
// Parameters:
//  - endDateTime : Usage: endDateTime={endDateTime}
//  - groupId : Usage: groupId={groupId}
//  - startDateTime : Usage: startDateTime={startDateTime}
func (m *ReportsRequestBuilder) GetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTime(groupId *string, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*i39400581189e192e8c3927b2ac378c799b30d5cb2c7a932a3e2f9802d2a2f95d.GetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i39400581189e192e8c3927b2ac378c799b30d5cb2c7a932a3e2f9802d2a2f95d.NewGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, groupId, startDateTime, endDateTime);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getMailboxUsageDetail(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetMailboxUsageDetailWithPeriod(period *string)(*i61a78ce1e0b6e2e60e6d50fe39c731a201868c12c59871387bf62c67fa03cc24.GetMailboxUsageDetailWithPeriodRequestBuilder) {
    return i61a78ce1e0b6e2e60e6d50fe39c731a201868c12c59871387bf62c67fa03cc24.NewGetMailboxUsageDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getMailboxUsageMailboxCounts(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetMailboxUsageMailboxCountsWithPeriod(period *string)(*i700e0ab2f1ea334038b3293b2396e12432ea3545b07ed127bb12d122bbb6bf90.GetMailboxUsageMailboxCountsWithPeriodRequestBuilder) {
    return i700e0ab2f1ea334038b3293b2396e12432ea3545b07ed127bb12d122bbb6bf90.NewGetMailboxUsageMailboxCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getMailboxUsageQuotaStatusMailboxCounts(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetMailboxUsageQuotaStatusMailboxCountsWithPeriod(period *string)(*i37d455e922500fc7c5e44df5cc84a103ddbb8470a53860cb0c02e64fd6e8e775.GetMailboxUsageQuotaStatusMailboxCountsWithPeriodRequestBuilder) {
    return i37d455e922500fc7c5e44df5cc84a103ddbb8470a53860cb0c02e64fd6e8e775.NewGetMailboxUsageQuotaStatusMailboxCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getMailboxUsageStorage(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetMailboxUsageStorageWithPeriod(period *string)(*id05b47283d832e55b732685b30d918ee0e303ae639c6447cfff5c7a7089741c7.GetMailboxUsageStorageWithPeriodRequestBuilder) {
    return id05b47283d832e55b732685b30d918ee0e303ae639c6447cfff5c7a7089741c7.NewGetMailboxUsageStorageWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getOffice365ActivationCounts()
func (m *ReportsRequestBuilder) GetOffice365ActivationCounts()(*i37efe63b21f1e137ba9fdf350d68a656b504d390fe52651da7739ce398ec31c8.GetOffice365ActivationCountsRequestBuilder) {
    return i37efe63b21f1e137ba9fdf350d68a656b504d390fe52651da7739ce398ec31c8.NewGetOffice365ActivationCountsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getOffice365ActivationsUserCounts()
func (m *ReportsRequestBuilder) GetOffice365ActivationsUserCounts()(*i5be64ff627afac918209c9e795b4c7ad94cf40584d7376fd0bddf4902f737286.GetOffice365ActivationsUserCountsRequestBuilder) {
    return i5be64ff627afac918209c9e795b4c7ad94cf40584d7376fd0bddf4902f737286.NewGetOffice365ActivationsUserCountsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getOffice365ActivationsUserDetail()
func (m *ReportsRequestBuilder) GetOffice365ActivationsUserDetail()(*i67d417cc0a32f095a4d9872596c73a5670bab5ca99f2c39b690a3dc3e53dacb6.GetOffice365ActivationsUserDetailRequestBuilder) {
    return i67d417cc0a32f095a4d9872596c73a5670bab5ca99f2c39b690a3dc3e53dacb6.NewGetOffice365ActivationsUserDetailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getOffice365ActiveUserCounts(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetOffice365ActiveUserCountsWithPeriod(period *string)(*i4ea09032e0c277ad5c699279a1e67097201d217b8d0aac6b5962420831d787bd.GetOffice365ActiveUserCountsWithPeriodRequestBuilder) {
    return i4ea09032e0c277ad5c699279a1e67097201d217b8d0aac6b5962420831d787bd.NewGetOffice365ActiveUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getOffice365ActiveUserDetail(date={date})
// Parameters:
//  - date : Usage: date={date}
func (m *ReportsRequestBuilder) GetOffice365ActiveUserDetailWithDate(date *string)(*i0d6612898baf60cf3ad09e5ec9209afdfaffc25ddef2fe3fddf028f8f356d2e1.GetOffice365ActiveUserDetailWithDateRequestBuilder) {
    return i0d6612898baf60cf3ad09e5ec9209afdfaffc25ddef2fe3fddf028f8f356d2e1.NewGetOffice365ActiveUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getOffice365ActiveUserDetail(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetOffice365ActiveUserDetailWithPeriod(period *string)(*i6edf15eaf951cb8d179890d37ee652592111a5bfeb9fd5709d3609c7593bed81.GetOffice365ActiveUserDetailWithPeriodRequestBuilder) {
    return i6edf15eaf951cb8d179890d37ee652592111a5bfeb9fd5709d3609c7593bed81.NewGetOffice365ActiveUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getOffice365GroupsActivityCounts(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetOffice365GroupsActivityCountsWithPeriod(period *string)(*ifb70744a42eac9e3054301d8328a61f22ab480c9dcaaaa3f398cf2a26e1b242b.GetOffice365GroupsActivityCountsWithPeriodRequestBuilder) {
    return ifb70744a42eac9e3054301d8328a61f22ab480c9dcaaaa3f398cf2a26e1b242b.NewGetOffice365GroupsActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getOffice365GroupsActivityDetail(date={date})
// Parameters:
//  - date : Usage: date={date}
func (m *ReportsRequestBuilder) GetOffice365GroupsActivityDetailWithDate(date *string)(*i4319947bf5469fb8e87f0bc08088c383a13726c5be51ca3558867ea6abecc315.GetOffice365GroupsActivityDetailWithDateRequestBuilder) {
    return i4319947bf5469fb8e87f0bc08088c383a13726c5be51ca3558867ea6abecc315.NewGetOffice365GroupsActivityDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getOffice365GroupsActivityDetail(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetOffice365GroupsActivityDetailWithPeriod(period *string)(*i7774a89bdd2b6bc2f70c1e7c58f236fd1f3aa68a40a72db4898fd9c5d4f66988.GetOffice365GroupsActivityDetailWithPeriodRequestBuilder) {
    return i7774a89bdd2b6bc2f70c1e7c58f236fd1f3aa68a40a72db4898fd9c5d4f66988.NewGetOffice365GroupsActivityDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getOffice365GroupsActivityFileCounts(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetOffice365GroupsActivityFileCountsWithPeriod(period *string)(*i21ef6a6d7bc99608d2279d5cc2343f5300ae76d3d739202a0c1bc4ff2d13a391.GetOffice365GroupsActivityFileCountsWithPeriodRequestBuilder) {
    return i21ef6a6d7bc99608d2279d5cc2343f5300ae76d3d739202a0c1bc4ff2d13a391.NewGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getOffice365GroupsActivityGroupCounts(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetOffice365GroupsActivityGroupCountsWithPeriod(period *string)(*i5445cb98613918d3a7a579fddf521b041c55f06714612e2097ed886a37be4286.GetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder) {
    return i5445cb98613918d3a7a579fddf521b041c55f06714612e2097ed886a37be4286.NewGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getOffice365GroupsActivityStorage(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetOffice365GroupsActivityStorageWithPeriod(period *string)(*i26d1169a1f986183917395a64e40bb9c1870a54b457336ef64e7efa21c214a15.GetOffice365GroupsActivityStorageWithPeriodRequestBuilder) {
    return i26d1169a1f986183917395a64e40bb9c1870a54b457336ef64e7efa21c214a15.NewGetOffice365GroupsActivityStorageWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getOffice365ServicesUserCounts(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetOffice365ServicesUserCountsWithPeriod(period *string)(*i2306eba962c786175114d00783028649abfdfb48f769241b2a7dca15db108361.GetOffice365ServicesUserCountsWithPeriodRequestBuilder) {
    return i2306eba962c786175114d00783028649abfdfb48f769241b2a7dca15db108361.NewGetOffice365ServicesUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getOneDriveActivityFileCounts(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetOneDriveActivityFileCountsWithPeriod(period *string)(*iec7256d00a8e37c8dec628d1626f62f1224b45c25e3c24418830b8fede7bc006.GetOneDriveActivityFileCountsWithPeriodRequestBuilder) {
    return iec7256d00a8e37c8dec628d1626f62f1224b45c25e3c24418830b8fede7bc006.NewGetOneDriveActivityFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getOneDriveActivityUserCounts(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetOneDriveActivityUserCountsWithPeriod(period *string)(*i1ccdf78d8963c250d79813dd0917eaaaee794135fc39e35801b01ea5f2005034.GetOneDriveActivityUserCountsWithPeriodRequestBuilder) {
    return i1ccdf78d8963c250d79813dd0917eaaaee794135fc39e35801b01ea5f2005034.NewGetOneDriveActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getOneDriveActivityUserDetail(date={date})
// Parameters:
//  - date : Usage: date={date}
func (m *ReportsRequestBuilder) GetOneDriveActivityUserDetailWithDate(date *string)(*ife6d24553cc77853635c06a75c3587d7c978531be356ad656116b7742f60e1f7.GetOneDriveActivityUserDetailWithDateRequestBuilder) {
    return ife6d24553cc77853635c06a75c3587d7c978531be356ad656116b7742f60e1f7.NewGetOneDriveActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getOneDriveActivityUserDetail(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetOneDriveActivityUserDetailWithPeriod(period *string)(*iab84fe97b770990c561344b8be4caaca7e6d962145fd1ad9fc9e851313b38857.GetOneDriveActivityUserDetailWithPeriodRequestBuilder) {
    return iab84fe97b770990c561344b8be4caaca7e6d962145fd1ad9fc9e851313b38857.NewGetOneDriveActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getOneDriveUsageAccountCounts(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetOneDriveUsageAccountCountsWithPeriod(period *string)(*i75a65b7f6cf316241684f8aeba4d08fb2c2e04274ef9390b6003ff2a3fe87418.GetOneDriveUsageAccountCountsWithPeriodRequestBuilder) {
    return i75a65b7f6cf316241684f8aeba4d08fb2c2e04274ef9390b6003ff2a3fe87418.NewGetOneDriveUsageAccountCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getOneDriveUsageAccountDetail(date={date})
// Parameters:
//  - date : Usage: date={date}
func (m *ReportsRequestBuilder) GetOneDriveUsageAccountDetailWithDate(date *string)(*i01986228cfa34d663610609938a20fe1aa899eb3ae22c44d689bf3a7cc4fcec2.GetOneDriveUsageAccountDetailWithDateRequestBuilder) {
    return i01986228cfa34d663610609938a20fe1aa899eb3ae22c44d689bf3a7cc4fcec2.NewGetOneDriveUsageAccountDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getOneDriveUsageAccountDetail(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetOneDriveUsageAccountDetailWithPeriod(period *string)(*i316dae3dc0a83b965993ba7f5a9d3b29d2b5dcbd706e8977d4a6cb34548c5275.GetOneDriveUsageAccountDetailWithPeriodRequestBuilder) {
    return i316dae3dc0a83b965993ba7f5a9d3b29d2b5dcbd706e8977d4a6cb34548c5275.NewGetOneDriveUsageAccountDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getOneDriveUsageFileCounts(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetOneDriveUsageFileCountsWithPeriod(period *string)(*ibc84667b2a2b840609232b734d0a2873e66315fbe63aeadb3c3d468ad9f84d75.GetOneDriveUsageFileCountsWithPeriodRequestBuilder) {
    return ibc84667b2a2b840609232b734d0a2873e66315fbe63aeadb3c3d468ad9f84d75.NewGetOneDriveUsageFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getOneDriveUsageStorage(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetOneDriveUsageStorageWithPeriod(period *string)(*if97af71691601e0f0046051c8e82fb71f8db2f1d3b5f8e7e40346d3fa47b196f.GetOneDriveUsageStorageWithPeriodRequestBuilder) {
    return if97af71691601e0f0046051c8e82fb71f8db2f1d3b5f8e7e40346d3fa47b196f.NewGetOneDriveUsageStorageWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getPrinterArchivedPrintJobs(printerId='{printerId}',startDateTime={startDateTime},endDateTime={endDateTime})
// Parameters:
//  - endDateTime : Usage: endDateTime={endDateTime}
//  - printerId : Usage: printerId={printerId}
//  - startDateTime : Usage: startDateTime={startDateTime}
func (m *ReportsRequestBuilder) GetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTime(printerId *string, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*i3ed9439789e0857666ebca3cd16359b4745ffe979a682ea73e6cd26f8054ae21.GetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i3ed9439789e0857666ebca3cd16359b4745ffe979a682ea73e6cd26f8054ae21.NewGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, printerId, startDateTime, endDateTime);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getSharePointActivityFileCounts(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetSharePointActivityFileCountsWithPeriod(period *string)(*i1bfeeda14fdcd3c4bc1c9b02a4d925837c77654498184deaffbfeb8f496c10c5.GetSharePointActivityFileCountsWithPeriodRequestBuilder) {
    return i1bfeeda14fdcd3c4bc1c9b02a4d925837c77654498184deaffbfeb8f496c10c5.NewGetSharePointActivityFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getSharePointActivityPages(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetSharePointActivityPagesWithPeriod(period *string)(*i6b3e5d0662b4c04bf1f969a3c62a48d50c4f7ecb57a3913782ce351d87a0035e.GetSharePointActivityPagesWithPeriodRequestBuilder) {
    return i6b3e5d0662b4c04bf1f969a3c62a48d50c4f7ecb57a3913782ce351d87a0035e.NewGetSharePointActivityPagesWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getSharePointActivityUserCounts(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetSharePointActivityUserCountsWithPeriod(period *string)(*i40f65a5b54f01dfaea303a3582e82bcdaa71c2195cc00b7915431e77f38e9a0f.GetSharePointActivityUserCountsWithPeriodRequestBuilder) {
    return i40f65a5b54f01dfaea303a3582e82bcdaa71c2195cc00b7915431e77f38e9a0f.NewGetSharePointActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getSharePointActivityUserDetail(date={date})
// Parameters:
//  - date : Usage: date={date}
func (m *ReportsRequestBuilder) GetSharePointActivityUserDetailWithDate(date *string)(*ib3bef6d1d1d843d8e20e58b41c29e9b82bb511ac97e381039f942ffee79a3b9a.GetSharePointActivityUserDetailWithDateRequestBuilder) {
    return ib3bef6d1d1d843d8e20e58b41c29e9b82bb511ac97e381039f942ffee79a3b9a.NewGetSharePointActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getSharePointActivityUserDetail(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetSharePointActivityUserDetailWithPeriod(period *string)(*i56c28dc52ab4ecbfe37baebbfaf6f279f8d130872c0931a194264028d89e2c96.GetSharePointActivityUserDetailWithPeriodRequestBuilder) {
    return i56c28dc52ab4ecbfe37baebbfaf6f279f8d130872c0931a194264028d89e2c96.NewGetSharePointActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getSharePointSiteUsageDetail(date={date})
// Parameters:
//  - date : Usage: date={date}
func (m *ReportsRequestBuilder) GetSharePointSiteUsageDetailWithDate(date *string)(*ic5ac4372354eec289834ac2ae20a4b14244024792aa3c3ad01b0f0a204b1e81b.GetSharePointSiteUsageDetailWithDateRequestBuilder) {
    return ic5ac4372354eec289834ac2ae20a4b14244024792aa3c3ad01b0f0a204b1e81b.NewGetSharePointSiteUsageDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getSharePointSiteUsageDetail(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetSharePointSiteUsageDetailWithPeriod(period *string)(*ib9c6325caeea63e2dd57f99af77223cbb370ce799b27f8b8a3f6a3f987a9f9fe.GetSharePointSiteUsageDetailWithPeriodRequestBuilder) {
    return ib9c6325caeea63e2dd57f99af77223cbb370ce799b27f8b8a3f6a3f987a9f9fe.NewGetSharePointSiteUsageDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getSharePointSiteUsageFileCounts(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetSharePointSiteUsageFileCountsWithPeriod(period *string)(*i2699e05c2942728d3c9dbbec2f7da6269f18ec269bdc96532cb49ca45ea2e03f.GetSharePointSiteUsageFileCountsWithPeriodRequestBuilder) {
    return i2699e05c2942728d3c9dbbec2f7da6269f18ec269bdc96532cb49ca45ea2e03f.NewGetSharePointSiteUsageFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getSharePointSiteUsagePages(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetSharePointSiteUsagePagesWithPeriod(period *string)(*i2f04ee022b2f4dd0c31e5c7f89cc28c7d1e36585beb573c9c689634579fb8999.GetSharePointSiteUsagePagesWithPeriodRequestBuilder) {
    return i2f04ee022b2f4dd0c31e5c7f89cc28c7d1e36585beb573c9c689634579fb8999.NewGetSharePointSiteUsagePagesWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getSharePointSiteUsageSiteCounts(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetSharePointSiteUsageSiteCountsWithPeriod(period *string)(*i347d730506a06c407296d7557c653293a35a9c38b50c7256c1c7abb4c736ef9a.GetSharePointSiteUsageSiteCountsWithPeriodRequestBuilder) {
    return i347d730506a06c407296d7557c653293a35a9c38b50c7256c1c7abb4c736ef9a.NewGetSharePointSiteUsageSiteCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getSharePointSiteUsageStorage(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetSharePointSiteUsageStorageWithPeriod(period *string)(*i028952a5521bf457c8a8796489c7f5a1137cf7b0745234478ff5fd12bb087564.GetSharePointSiteUsageStorageWithPeriodRequestBuilder) {
    return i028952a5521bf457c8a8796489c7f5a1137cf7b0745234478ff5fd12bb087564.NewGetSharePointSiteUsageStorageWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getSkypeForBusinessActivityCounts(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetSkypeForBusinessActivityCountsWithPeriod(period *string)(*i8ad77c0341e431729102d8b93f9066b82a7f4fc8ee3e748eeea007bf3d9d5b18.GetSkypeForBusinessActivityCountsWithPeriodRequestBuilder) {
    return i8ad77c0341e431729102d8b93f9066b82a7f4fc8ee3e748eeea007bf3d9d5b18.NewGetSkypeForBusinessActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getSkypeForBusinessActivityUserCounts(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetSkypeForBusinessActivityUserCountsWithPeriod(period *string)(*i069326085999a3926845312c56f68613c15a7dcb468bf7f0592670ea8f5896aa.GetSkypeForBusinessActivityUserCountsWithPeriodRequestBuilder) {
    return i069326085999a3926845312c56f68613c15a7dcb468bf7f0592670ea8f5896aa.NewGetSkypeForBusinessActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getSkypeForBusinessActivityUserDetail(date={date})
// Parameters:
//  - date : Usage: date={date}
func (m *ReportsRequestBuilder) GetSkypeForBusinessActivityUserDetailWithDate(date *string)(*i6a26bcea122cbd3594cacac16b0715042c3ca45bb5f889d1fd0bd83d243a6fe3.GetSkypeForBusinessActivityUserDetailWithDateRequestBuilder) {
    return i6a26bcea122cbd3594cacac16b0715042c3ca45bb5f889d1fd0bd83d243a6fe3.NewGetSkypeForBusinessActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getSkypeForBusinessActivityUserDetail(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetSkypeForBusinessActivityUserDetailWithPeriod(period *string)(*ia79ac9c43dc0218a9c3f469a247120979433dd0e3e53de3949570224b3cf3936.GetSkypeForBusinessActivityUserDetailWithPeriodRequestBuilder) {
    return ia79ac9c43dc0218a9c3f469a247120979433dd0e3e53de3949570224b3cf3936.NewGetSkypeForBusinessActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getSkypeForBusinessDeviceUsageDistributionUserCounts(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriod(period *string)(*ia5f41691c1fbeb62b6b1de131fa27f7306db515c8759a9e49ec88553c0e6d97d.GetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) {
    return ia5f41691c1fbeb62b6b1de131fa27f7306db515c8759a9e49ec88553c0e6d97d.NewGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getSkypeForBusinessDeviceUsageUserCounts(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetSkypeForBusinessDeviceUsageUserCountsWithPeriod(period *string)(*i9de2dd5236ac703b10206d8522e5a8f5490288eadb62a3fbf88d9becd9f061bf.GetSkypeForBusinessDeviceUsageUserCountsWithPeriodRequestBuilder) {
    return i9de2dd5236ac703b10206d8522e5a8f5490288eadb62a3fbf88d9becd9f061bf.NewGetSkypeForBusinessDeviceUsageUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getSkypeForBusinessDeviceUsageUserDetail(date={date})
// Parameters:
//  - date : Usage: date={date}
func (m *ReportsRequestBuilder) GetSkypeForBusinessDeviceUsageUserDetailWithDate(date *string)(*i037933f7467ae74beac89af9fadc6cf1a5a8272d8f21a5c2f98bc8f8e7c28ab6.GetSkypeForBusinessDeviceUsageUserDetailWithDateRequestBuilder) {
    return i037933f7467ae74beac89af9fadc6cf1a5a8272d8f21a5c2f98bc8f8e7c28ab6.NewGetSkypeForBusinessDeviceUsageUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getSkypeForBusinessDeviceUsageUserDetail(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetSkypeForBusinessDeviceUsageUserDetailWithPeriod(period *string)(*i6c58964b36e86bf47df5ddf5c8161ca0e4621aee61aaa9517f1dded3120cfc8d.GetSkypeForBusinessDeviceUsageUserDetailWithPeriodRequestBuilder) {
    return i6c58964b36e86bf47df5ddf5c8161ca0e4621aee61aaa9517f1dded3120cfc8d.NewGetSkypeForBusinessDeviceUsageUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getSkypeForBusinessOrganizerActivityCounts(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetSkypeForBusinessOrganizerActivityCountsWithPeriod(period *string)(*i2e849a1f15987664ba1a6648f0c1d0eb9e2ca5965a632cc3761ad14bb91674bf.GetSkypeForBusinessOrganizerActivityCountsWithPeriodRequestBuilder) {
    return i2e849a1f15987664ba1a6648f0c1d0eb9e2ca5965a632cc3761ad14bb91674bf.NewGetSkypeForBusinessOrganizerActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getSkypeForBusinessOrganizerActivityMinuteCounts(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriod(period *string)(*i8d02e7a0b578c5b4465f96608d724ae82f61e67266f11ff2ab8aad96d55d1ac1.GetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriodRequestBuilder) {
    return i8d02e7a0b578c5b4465f96608d724ae82f61e67266f11ff2ab8aad96d55d1ac1.NewGetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getSkypeForBusinessOrganizerActivityUserCounts(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetSkypeForBusinessOrganizerActivityUserCountsWithPeriod(period *string)(*i066e5ea9f7e2ed1b047023128aab6caef0db934d50a71df44b7b6cf7bfb5a10e.GetSkypeForBusinessOrganizerActivityUserCountsWithPeriodRequestBuilder) {
    return i066e5ea9f7e2ed1b047023128aab6caef0db934d50a71df44b7b6cf7bfb5a10e.NewGetSkypeForBusinessOrganizerActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getSkypeForBusinessParticipantActivityCounts(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetSkypeForBusinessParticipantActivityCountsWithPeriod(period *string)(*i6e8f025e6b5802865cf6c583d64d93f7e1c603ead045ca53bef1f6b9cde8a5b0.GetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilder) {
    return i6e8f025e6b5802865cf6c583d64d93f7e1c603ead045ca53bef1f6b9cde8a5b0.NewGetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getSkypeForBusinessParticipantActivityMinuteCounts(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetSkypeForBusinessParticipantActivityMinuteCountsWithPeriod(period *string)(*idcca069eb97aac5c90ca4ddb83fee2301a14619967cfdadbb5d0a81effc5cc59.GetSkypeForBusinessParticipantActivityMinuteCountsWithPeriodRequestBuilder) {
    return idcca069eb97aac5c90ca4ddb83fee2301a14619967cfdadbb5d0a81effc5cc59.NewGetSkypeForBusinessParticipantActivityMinuteCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getSkypeForBusinessParticipantActivityUserCounts(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetSkypeForBusinessParticipantActivityUserCountsWithPeriod(period *string)(*i05ca24f1e64f29c556c3c67a9024eea8b3aac4d350f05ff2c52f4511bf0b6f33.GetSkypeForBusinessParticipantActivityUserCountsWithPeriodRequestBuilder) {
    return i05ca24f1e64f29c556c3c67a9024eea8b3aac4d350f05ff2c52f4511bf0b6f33.NewGetSkypeForBusinessParticipantActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getSkypeForBusinessPeerToPeerActivityCounts(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod(period *string)(*i200c880280f3a34037634650d2cd9969f9a07037f773af21ebb6141e40dde23b.GetSkypeForBusinessPeerToPeerActivityCountsWithPeriodRequestBuilder) {
    return i200c880280f3a34037634650d2cd9969f9a07037f773af21ebb6141e40dde23b.NewGetSkypeForBusinessPeerToPeerActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getSkypeForBusinessPeerToPeerActivityMinuteCounts(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriod(period *string)(*ic5508dd2223dfecdf5814372064e28b262f8f8bfe997c49795b7f76e30889837.GetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriodRequestBuilder) {
    return ic5508dd2223dfecdf5814372064e28b262f8f8bfe997c49795b7f76e30889837.NewGetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getSkypeForBusinessPeerToPeerActivityUserCounts(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriod(period *string)(*i416e8a1993d00c5e5b47ab655f48ac49786eb1442a5459b7cb7023ac34c28093.GetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriodRequestBuilder) {
    return i416e8a1993d00c5e5b47ab655f48ac49786eb1442a5459b7cb7023ac34c28093.NewGetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getTeamsDeviceUsageDistributionUserCounts(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetTeamsDeviceUsageDistributionUserCountsWithPeriod(period *string)(*i05c8d276b86beee65c00c115937a46c2753c7d328e142fe9afab6ef90b5a1a79.GetTeamsDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) {
    return i05c8d276b86beee65c00c115937a46c2753c7d328e142fe9afab6ef90b5a1a79.NewGetTeamsDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getTeamsDeviceUsageUserCounts(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetTeamsDeviceUsageUserCountsWithPeriod(period *string)(*ife99981e20f3609c2111e3e2aff666b997075ba1e7c98ff44f1950f1d6265ccb.GetTeamsDeviceUsageUserCountsWithPeriodRequestBuilder) {
    return ife99981e20f3609c2111e3e2aff666b997075ba1e7c98ff44f1950f1d6265ccb.NewGetTeamsDeviceUsageUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getTeamsDeviceUsageUserDetail(date={date})
// Parameters:
//  - date : Usage: date={date}
func (m *ReportsRequestBuilder) GetTeamsDeviceUsageUserDetailWithDate(date *string)(*i677bcf76c9544b6088f18b0d84351cd554f0537738c43ca67e647fffec66978c.GetTeamsDeviceUsageUserDetailWithDateRequestBuilder) {
    return i677bcf76c9544b6088f18b0d84351cd554f0537738c43ca67e647fffec66978c.NewGetTeamsDeviceUsageUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getTeamsDeviceUsageUserDetail(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetTeamsDeviceUsageUserDetailWithPeriod(period *string)(*iaaacf2b7aacf4d5f8107fa18d6f0d6cb33066e3335d8ab8cc1a2adfa1f4e3c42.GetTeamsDeviceUsageUserDetailWithPeriodRequestBuilder) {
    return iaaacf2b7aacf4d5f8107fa18d6f0d6cb33066e3335d8ab8cc1a2adfa1f4e3c42.NewGetTeamsDeviceUsageUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getTeamsUserActivityCounts(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetTeamsUserActivityCountsWithPeriod(period *string)(*i9b65430927e689a26a700312296e423fc6e6e6c0bb86f72287b75c7915e695e0.GetTeamsUserActivityCountsWithPeriodRequestBuilder) {
    return i9b65430927e689a26a700312296e423fc6e6e6c0bb86f72287b75c7915e695e0.NewGetTeamsUserActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getTeamsUserActivityUserCounts(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetTeamsUserActivityUserCountsWithPeriod(period *string)(*i34861f91c890c43257a1e1386ba24de1ffe44f842eadf95d13de30d27a852991.GetTeamsUserActivityUserCountsWithPeriodRequestBuilder) {
    return i34861f91c890c43257a1e1386ba24de1ffe44f842eadf95d13de30d27a852991.NewGetTeamsUserActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getTeamsUserActivityUserDetail(date={date})
// Parameters:
//  - date : Usage: date={date}
func (m *ReportsRequestBuilder) GetTeamsUserActivityUserDetailWithDate(date *string)(*i0c4b2ea99f0789f3909bf8d293ade8bd9216c7010a8e0dc72ceea54c21c1d481.GetTeamsUserActivityUserDetailWithDateRequestBuilder) {
    return i0c4b2ea99f0789f3909bf8d293ade8bd9216c7010a8e0dc72ceea54c21c1d481.NewGetTeamsUserActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getTeamsUserActivityUserDetail(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetTeamsUserActivityUserDetailWithPeriod(period *string)(*ibfd05176830665cf1dae7366e8b491572f134d04e12120a5499c283d6c8ae062.GetTeamsUserActivityUserDetailWithPeriodRequestBuilder) {
    return ibfd05176830665cf1dae7366e8b491572f134d04e12120a5499c283d6c8ae062.NewGetTeamsUserActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getUserArchivedPrintJobs(userId='{userId}',startDateTime={startDateTime},endDateTime={endDateTime})
// Parameters:
//  - endDateTime : Usage: endDateTime={endDateTime}
//  - startDateTime : Usage: startDateTime={startDateTime}
//  - userId : Usage: userId={userId}
func (m *ReportsRequestBuilder) GetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTime(userId *string, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*iafadefc9c12601a3611573bb976b73b251f9ea476a68e3bdc9ceab9e0d5cc308.GetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return iafadefc9c12601a3611573bb976b73b251f9ea476a68e3bdc9ceab9e0d5cc308.NewGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, userId, startDateTime, endDateTime);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getYammerActivityCounts(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetYammerActivityCountsWithPeriod(period *string)(*ifd45221e0ff465bb6f37fa740578c697634181cf333e3c01e4db49019e2ba5eb.GetYammerActivityCountsWithPeriodRequestBuilder) {
    return ifd45221e0ff465bb6f37fa740578c697634181cf333e3c01e4db49019e2ba5eb.NewGetYammerActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getYammerActivityUserCounts(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetYammerActivityUserCountsWithPeriod(period *string)(*i13fef70f90226be585263be7022fe29fa384901ba2188c080017c9272ca23e99.GetYammerActivityUserCountsWithPeriodRequestBuilder) {
    return i13fef70f90226be585263be7022fe29fa384901ba2188c080017c9272ca23e99.NewGetYammerActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getYammerActivityUserDetail(date={date})
// Parameters:
//  - date : Usage: date={date}
func (m *ReportsRequestBuilder) GetYammerActivityUserDetailWithDate(date *string)(*id7d2a94ac1ea97cc24d2cd9cb47182b168535107d0d09ad275c4e955f417e029.GetYammerActivityUserDetailWithDateRequestBuilder) {
    return id7d2a94ac1ea97cc24d2cd9cb47182b168535107d0d09ad275c4e955f417e029.NewGetYammerActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getYammerActivityUserDetail(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetYammerActivityUserDetailWithPeriod(period *string)(*id379a8afbc53e3411a43fbddcd71ea21f90a2a71ada609569a5c9040ebee5f51.GetYammerActivityUserDetailWithPeriodRequestBuilder) {
    return id379a8afbc53e3411a43fbddcd71ea21f90a2a71ada609569a5c9040ebee5f51.NewGetYammerActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getYammerDeviceUsageDistributionUserCounts(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetYammerDeviceUsageDistributionUserCountsWithPeriod(period *string)(*i03d33a99775bdfdb348179d972066b208c7410da4d93f4776c48bb0c7619f539.GetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) {
    return i03d33a99775bdfdb348179d972066b208c7410da4d93f4776c48bb0c7619f539.NewGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getYammerDeviceUsageUserCounts(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetYammerDeviceUsageUserCountsWithPeriod(period *string)(*i9b77e3c3d15971fe31fa9e23d6bdc4c74a584c922585117306c4db805ac4131c.GetYammerDeviceUsageUserCountsWithPeriodRequestBuilder) {
    return i9b77e3c3d15971fe31fa9e23d6bdc4c74a584c922585117306c4db805ac4131c.NewGetYammerDeviceUsageUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getYammerDeviceUsageUserDetail(date={date})
// Parameters:
//  - date : Usage: date={date}
func (m *ReportsRequestBuilder) GetYammerDeviceUsageUserDetailWithDate(date *string)(*if97f8cb01f2d6a02697169f75d571d9a67fc56f2041590c23f83bf2114377bda.GetYammerDeviceUsageUserDetailWithDateRequestBuilder) {
    return if97f8cb01f2d6a02697169f75d571d9a67fc56f2041590c23f83bf2114377bda.NewGetYammerDeviceUsageUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getYammerDeviceUsageUserDetail(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetYammerDeviceUsageUserDetailWithPeriod(period *string)(*i6e07bd3dadcc8aa2d52f07e77f1f35672ac7ce62aa2917c7a9e4e77b8f4c4db5.GetYammerDeviceUsageUserDetailWithPeriodRequestBuilder) {
    return i6e07bd3dadcc8aa2d52f07e77f1f35672ac7ce62aa2917c7a9e4e77b8f4c4db5.NewGetYammerDeviceUsageUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getYammerGroupsActivityCounts(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetYammerGroupsActivityCountsWithPeriod(period *string)(*i7ee204dcbd579a0c0d252ddbfcdd01bf9f244b2b38f671f02c4108f867a99c73.GetYammerGroupsActivityCountsWithPeriodRequestBuilder) {
    return i7ee204dcbd579a0c0d252ddbfcdd01bf9f244b2b38f671f02c4108f867a99c73.NewGetYammerGroupsActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getYammerGroupsActivityDetail(date={date})
// Parameters:
//  - date : Usage: date={date}
func (m *ReportsRequestBuilder) GetYammerGroupsActivityDetailWithDate(date *string)(*ia3da511007b4da9fa7b31d42bcef05972fa35df25142f53ebb0f2944e3932886.GetYammerGroupsActivityDetailWithDateRequestBuilder) {
    return ia3da511007b4da9fa7b31d42bcef05972fa35df25142f53ebb0f2944e3932886.NewGetYammerGroupsActivityDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getYammerGroupsActivityDetail(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetYammerGroupsActivityDetailWithPeriod(period *string)(*i5277312ad2cbf60385fe9448b03ae55d7aefafa14a9c09b135b745563e0dd271.GetYammerGroupsActivityDetailWithPeriodRequestBuilder) {
    return i5277312ad2cbf60385fe9448b03ae55d7aefafa14a9c09b135b745563e0dd271.NewGetYammerGroupsActivityDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.getYammerGroupsActivityGroupCounts(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) GetYammerGroupsActivityGroupCountsWithPeriod(period *string)(*i55b749bfea07392ccbd597dc8d76a153720e39963b1e66527cac5bb79f8c6b29.GetYammerGroupsActivityGroupCountsWithPeriodRequestBuilder) {
    return i55b749bfea07392ccbd597dc8d76a153720e39963b1e66527cac5bb79f8c6b29.NewGetYammerGroupsActivityGroupCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// Builds and executes requests for operations under \reports\microsoft.graph.managedDeviceEnrollmentFailureDetails()
func (m *ReportsRequestBuilder) ManagedDeviceEnrollmentFailureDetails()(*i0e6a5a79552b1c4826c9418d405e7420746bbd08e606bfd5c84ba1ad674aed8f.ManagedDeviceEnrollmentFailureDetailsRequestBuilder) {
    return i0e6a5a79552b1c4826c9418d405e7420746bbd08e606bfd5c84ba1ad674aed8f.NewManagedDeviceEnrollmentFailureDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \reports\microsoft.graph.managedDeviceEnrollmentFailureDetails(skip={skip},top={top},filter='{filter}',skipToken='{skipToken}')
// Parameters:
//  - filter : Usage: filter={filter}
//  - skip : Usage: skip={skip}
//  - skipToken : Usage: skipToken={skipToken}
//  - top : Usage: top={top}
func (m *ReportsRequestBuilder) ManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipToken(skip *int32, top *int32, filter *string, skipToken *string)(*i374eb5f5ea7befd0f519a8a1c2886282c8ae9910332bd6b507ad8b29eddfdf2d.ManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilder) {
    return i374eb5f5ea7befd0f519a8a1c2886282c8ae9910332bd6b507ad8b29eddfdf2d.NewManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top, filter, skipToken);
}
// Builds and executes requests for operations under \reports\microsoft.graph.managedDeviceEnrollmentTopFailures()
func (m *ReportsRequestBuilder) ManagedDeviceEnrollmentTopFailures()(*i9d1992f642df096181dd0b243f5639c29682f838689705077b65251f267d222f.ManagedDeviceEnrollmentTopFailuresRequestBuilder) {
    return i9d1992f642df096181dd0b243f5639c29682f838689705077b65251f267d222f.NewManagedDeviceEnrollmentTopFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \reports\microsoft.graph.managedDeviceEnrollmentTopFailures(period='{period}')
// Parameters:
//  - period : Usage: period={period}
func (m *ReportsRequestBuilder) ManagedDeviceEnrollmentTopFailuresWithPeriod(period *string)(*i3f4aebd4163df4cb0a4f38097f15a9f8ba73853688cf5a50b93a222eb4832abe.ManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder) {
    return i3f4aebd4163df4cb0a4f38097f15a9f8ba73853688cf5a50b93a222eb4832abe.NewManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
func (m *ReportsRequestBuilder) MonthlyPrintUsageByPrinter()(*if1205fb664936127831759fefe4c2c43852474a3fc0419f51e89975f0f281d32.MonthlyPrintUsageByPrinterRequestBuilder) {
    return if1205fb664936127831759fefe4c2c43852474a3fc0419f51e89975f0f281d32.NewMonthlyPrintUsageByPrinterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.reports.monthlyPrintUsageByPrinter.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ReportsRequestBuilder) MonthlyPrintUsageByPrinterById(id string)(*i0fbd1879eec3f09756dabd53c9de402f839fcf5ed7e8be0da0d3e4135b124b0a.PrintUsageByPrinterRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByPrinter_id"] = id
    }
    return i0fbd1879eec3f09756dabd53c9de402f839fcf5ed7e8be0da0d3e4135b124b0a.NewPrintUsageByPrinterRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ReportsRequestBuilder) MonthlyPrintUsageByUser()(*i3f12902a7f7457a60f97014c6db3b2bd89dac8cabe4d467b766677db3d6029d0.MonthlyPrintUsageByUserRequestBuilder) {
    return i3f12902a7f7457a60f97014c6db3b2bd89dac8cabe4d467b766677db3d6029d0.NewMonthlyPrintUsageByUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.reports.monthlyPrintUsageByUser.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ReportsRequestBuilder) MonthlyPrintUsageByUserById(id string)(*if4346a9921b148666613b317b233f49301e4eef2f130f6b9e02414250e5a2294.PrintUsageByUserRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByUser_id"] = id
    }
    return if4346a9921b148666613b317b233f49301e4eef2f130f6b9e02414250e5a2294.NewPrintUsageByUserRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Update reports
// Parameters:
//  - options : Options for the request
func (m *ReportsRequestBuilder) Patch(options *ReportsRequestBuilderPatchOptions)(error) {
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
