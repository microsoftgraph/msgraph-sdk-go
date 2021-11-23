package functions

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i0266785f33f4cc8c220e944b21b01c5947e87c01bcb41c020a5eddb02a9ed257 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/impower"
    i02bf3208bc8afd502281b1c8c45241c968d956cfa84b08a4a8c61ed099189b43 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/power"
    i0409f0a1cf64e0274eaf7de30693fa8cbba20462831b9d50c75dcb1110b2c48e "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/beta_inv"
    i0461f30ce1b6ad808a745d4c850e3b7750615a40428226bf45dd420249fa3875 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/imsub"
    i049e9d1c53e9c7d3f4e9ec70418d7aec06590e4019424eb68b3e6c1dec18a845 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/right"
    i0650ae4116ca9eb8a4f165acdbbeb4eb9b2a87b6df15c716bcff863b121e82c5 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/rept"
    i07ad831b4c905eef0f2b8147975bdc789d10f43b8261b423194792c65640c762 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/log10"
    i097b807941c5e18294aa5cf6621fcd0504002071465f2bf7987ff59de2aca92c "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/imcsch"
    i099b06710a4a149b79a3430055535b893ef0f2075c0f128716f0e326022844cc "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/ppmt"
    i0a886fcac7090403f01ae3a41ee46377ef9642eddf6fb911c9f8e55a4a6e35c7 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/randbetween"
    i0ab188e1c48548c3af73eb596d61e636fc8b4d8a715e840444b9622e63740fbe "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/replace"
    i0b3dc8bf2c5b69445e4aee0c44a1b80461aa78ff46d273df01cd9203922b28ca "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/dmax"
    i0b9eecb7ccfe792400a9c77664400f2d2a8d08f97db4c334d61a360325a50d8d "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/sech"
    i0c002f577eeb35c2aea092a0a697271289620508c6edc625840b274f2f02aa67 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/areas"
    i0dffa2d3e4f51c96e83725e18e43e2085197fe76a3f2cc370df8decb99554676 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/iserr"
    i0e120005f13ea8b2a2b93fdfd8a3730a0fa0fa7ac5360fc7b2c927444f41e1b7 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/imcot"
    i109ae362f2ade3a8826d4d477e2813f37a461350bf64445b7c53a3250ad89425 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/imdiv"
    i10f134b1fa93754f8715e8f54b26ea798cefb9f01e18cd0b9f5001dddbfb687b "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/ceiling_precise"
    i11766f797ebc2faf35ca5bd9fc2be9ca3099a81ea6ecfafb6c6678f6355ffc63 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/fixed"
    i124b15e94374881fc3ceb2d0b196066a5490324a97d0282f770abf019f5ebd72 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/intrate"
    i12c177d210a36c02a13c649a27eee3563c73121e988b53d1a5b0e07d943b1b5e "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/date"
    i12ce5b3d2a6128dd758ffd3fd51436cd30412f72f262fdbfc334abc06417e3f7 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/type_escaped"
    i13a285c284ba2120d7f47ea78b56d599c53ac40b7bf0f52dea113e716d2acb01 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/networkdays"
    i143e69103e0c49602fa219a4dd399818f412db7e7276d4263af8e6357a08e352 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/sign"
    i15a001f719e754b66f338d6f5feeb9021d4d0325bf7e573450ebf7407b0625c8 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/bitxor"
    i15bcb7ff9e078c7f041d0c54593e7b72a510b188ab533ff34d5acc37bbd04ed5 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/pv"
    i15f4c5ffb740886c3c9fa4ff008942d605f56f15a5a9ac16ddce2cbce6a6a174 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/poisson_dist"
    i16a2c7c1c23ba3e4a26a9034153b98e031a4d2505b415d8a6f93ab701c906928 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/gcd"
    i16adcfc7ead6bed8c6f0c25fa14d883aaa2b6874e7ff8bf7fc5b559ba0f5ada8 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/clean"
    i172364b817972ee8d188e239878e9d28a9730a745b39c8c7059c6177ad20ae4d "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/delta"
    i172d035bf5dab20ae0f3d85880acb671812d51c641eb99f52fdc34da268cf0fe "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/imsin"
    i1755304227922b0ac2038dc08b501d9ef710b978adb15d6211bfa3bd4baf16a2 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/isnumber"
    i179bb0c14bc435ddfa798d98f5b9ceed77f0d74d656eab292a71b5c2cdcefb31 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/day"
    i181dd5fee2d584c64321f0aa2fc010aa892dcc53eb3f40e83a6a3eb9c52f1e60 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/average"
    i18a2a6c3c741418f2ad549ddaa3c6adccf6ca9644ff9e7bf828710a10b7a6609 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/yielddisc"
    i196ecc451091a63b4e5a4d20f9956fe33210551341fd2b70b3f7c6f33bc90d0d "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/asin"
    i1a087e75f3b9161d8ce9481ff6539f40aa80c4d8fb841cc20f5841f312630609 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/weekday"
    i1a95ed2d2c1358a8e5a551f86342c22d27ead7fd9fea23c3e0f974d95289704c "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/nper"
    i1b3a3082854c5f5e5e7798d70e3ece3df8ecaaf43f8f08c2b5e4cd6cfc7516b0 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/xor"
    i1c1907b07e09c172dda3c49a4f313a14ca9651b2746e0bfeb1fbe37dbdd3e5c5 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/iserror"
    i1c7ecd5d9304af3cf2d1f18063bb065457b82f3762eef0998c8aa1365f22f800 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/pmt"
    i1d5ff915b0e62f1dbdf9af3a3d3a41f310dfde12a83f5753a8b75863457d468d "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/countifs"
    i1e19512606a6be24eefaac467f8c47c472b2dcdf8240b0e328a02af67678bdbe "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/ceiling_math"
    i1f4d3e0102ee9cebf20bdce62e14eec64e35b6cf5afbfae34c3eaad8d10a88e9 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/now"
    i1fc62d517918671eb2537fb2ec20d1e162fb4a5c9321c5e1d59d09cd12b56d55 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/dsum"
    i1fda0f7b146838f09b6022fc87d7d315f1ad404bbf5938aefc425a5484805cbd "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/imaginary"
    i20c624b914d32116e4852ae299f9489ecbb99fd83655afa4d29dbd67b41372ff "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/oddlprice"
    i20ffdeea2d03c7fc303f6b9e21d5f3c19cea0ae592ba935bfab22486e02a484d "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/vdb"
    i2114c9762c18b212ecaeaf82ff100b472674b7b5525cfd98d132aa1caa772c56 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/findb"
    i21b4bbaa6945c0f66867a02c17404afdb32643cba4e2e3a393bc38d19ded2c85 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/lower"
    i22a50a05268af02ad93b3e099195bd0aa10043cf850f934baaa779347ff912fd "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/maxa"
    i22ae119a3184e61ad2aae889ae7224bab4ea903395898b3eb5756b8488676892 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/kurt"
    i2479450dce1c1bcf5ee09162317898f4b82af362a855982508246d484edaaa8b "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/isref"
    i25706560466145b8e6f0ca13c02b864543dd157f7b8d5b9e8152f50ef7060f92 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/base"
    i25a82e0bc1d4b3db1874f6069a464756206a19345bb46a65038cb4d86d297e1a "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/coupdays"
    i26345570575fd85318c268057a084f3e7920644be05134a08f409b82ef2d087e "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/ecma_ceiling"
    i263f9c4268db82addce7e614d4c5702fa0452162ff16cd2a4b843da8d618b4e9 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/tbillyield"
    i27090be239364b444a136b5826b6896a10d7ecf32ebda209d4c91af0f55455bf "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/isnontext"
    i2736fbe644ba9c2b350316e1ff478ada2a4c58fc4a3d8aeefe7904ef5cf7f8d5 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/pi"
    i2766193728abd40f684dfaa8e25d37852148f793fdd189fd41e8836a8faf327b "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/rand"
    i27abe0dd8fef776e61d85d1af494b69c706d34811423987b8537ef7ac3a4065d "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/atanh"
    i27f60ca2023b9a4a5cc9d457ab7c1ebdefefd08ecd946d39b55f41e7c6bf329b "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/t_inv"
    i29591fecc3455d858a2d496b08a44cceb3934ca10d1a97e3e964fecf0ad89f7f "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/imsech"
    i297b0fb639f80c99fd017e82ff46d54acf62e7b366a48a1d74ec58773b7af6ff "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/vlookup"
    i2ae597988fb14a89eb97e83a4090774fb8bddfb370f4b82c600a1dfdbd31d5bf "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/imlog10"
    i2b383fba5c20322d304f2d869487c05e2693a7ce95255893aeb48eed972a2da9 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/hyperlink"
    i2cc67731e9c550a9cc653ccb9823cb407f72e05c8043783118760f0b7792187f "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/tbilleq"
    i2ec5919ffdb405f5ddc6809e4991195c3055085124325cfde9f900ea1a1816fc "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/chisq_inv_rt"
    i2f691a5ab87df22eea441e2516133fe7160951a84ac21da408c2c6fc21889c24 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/na"
    i30a090b5650b04cc93000bec9f08421b6363c6d560327127a84916a3b09f0086 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/t"
    i30dd8a3bec806587e9da967e397a243b809e20c4421c0454edac1b48ae5814a0 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/avedev"
    i31051972fdf67777e1c1e16a89b980d25bb524eee6cd715e3eccfe9e99387cad "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/year"
    i310b333bbd7ca74c312923d9e2339d29519809126287db8ffbd43d3ac2b3f32b "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/cumipmt"
    i32abbd433bb4aa4005a887aab37f79f1409aa21e4ac36cb2eda8bc7d3fff6fe4 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/pricedisc"
    i32e1f576fe43214f30cc8ffcd3f6853808eb7b877e9a328b9e8e2e647e6bf1c7 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/istext"
    i33353f01662e530553a98ee8936d864fcf8de48750304038df77208337f06742 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/standardize"
    i3457ae9a71aebc3d34cffda9ff5c133b2776a91cefac7fc05350402850687859 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/sumsq"
    i34c9271a0e30a434316e7cec8acb85de40a232501dd92511efc8c52ab3e0a708 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/oddfyield"
    i350fc4ebf814e7bfb9e3c38975155c78c6ea182f5af9de4f0e3d4f6a15b372d2 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/imcos"
    i3548428d79c006b6c84182290c7c63cb0c418cf298886e1351488dc98bb5f7d1 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/choose"
    i355b3d4ac7803c6a215bb8801502ba5fd3379a7c9cc52a9aa21240d59fa1d832 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/dec2oct"
    i35ba6be57981049c16486e05810aae9332e4f287c8a7fb56cba1de9585508ce4 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/lenb"
    i35bac92c54bdeaf46ce03cb02622ca9ef5b34152ddafea2898d92790e10f28ac "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/harmean"
    i3631de4628d7bc943b5b2b5b9a5ae5b6f6b176596f019b2abe9c2cdba1aedea7 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/norm_s_dist"
    i368eb1aa0b23f0d4e8267b91adca60d8fa7015f90d191cc4cc826cd78864f82a "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/f_dist_rt"
    i36b8593f4e7696b829be2b3641edeb4aa7090a5297244e7329a8d9ebabd11c98 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/imexp"
    i3715210d20a7eb43e392d3a551304d689734e97b74fc573d51234daae97ece0c "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/binom_dist"
    i3767ab8e1ea6498037301289a27a8c90b420e3097e0ed91c5eef9ffc3dcfc2a8 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/leftb"
    i382d1b83df58a9e122bb42db7279c4e8cc39d16055184c25db824773341d6d04 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/devsq"
    i38fc3cb72c39e10c123242af337410c10810bac285ded0e88d03eea4d7630c1c "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/text"
    i39f4cd87e36f2f3a0b33d803ee7df3c500c56510f3a7e138e50274650d2d65de "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/hour"
    i3a6c51172f28c572ffc43c7fd7e49866fd2b0c97b32fe94e0e28d58d09fc4276 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/skew"
    i3aa7f5e4a9ad281863916a5dae9bb0fb53bb4941b8b2664332317740e6a4c15b "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/if_escaped"
    i3b46708bc24ffb2b82689186a5a6f750836ea124768d99b4a4fbc31efa01d9c5 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/skew_p"
    i3b6f6c7ed33e73600236ac505d6d8ddfd90f8b9d72a97c577cf41f7f5bd9bad5 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/imcosh"
    i3bfcffb921445a5837441582c023df3ed7a99ae389fbed7fda482943e3fd5121 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/weeknum"
    i3f6bd1bd30b3d4eae2b33020bfff18913d339b66027e66bb84d38211e36bed13 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/today"
    i42ce97198842ba0a177d6ed68d50d81eafa58190eb0f5ed32a2b12a2a29d3f1f "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/gammaln_precise"
    i434657a2f6458186aeeb593186f3d5915f1bf0677f364b0cecbf804a6d4c5ee9 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/fisher"
    i43db37f22ff8e93c72e7ee1e05631ed0e75db3f46614da86fd2f4edc442ba8bb "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/syd"
    i445b80f059057d65d647c9ef7d80680e79c644d587f2eb3b38fa0efdb98a56dd "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/f_dist"
    i447958a9355da6b66054ac0c3105f155673cc9c53508f382a4b0343ab49fe403 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/convert"
    i448ab36b74b378173a5767a9d3dc2209b2ec52bd89af54860a8c353a1249eabb "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/percentrank_exc"
    i44c1d271b0ded00fc624c0e85f818ca370a3df578fdbd135e4c10a90d904bc6d "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/log"
    i466df9e270b2d70578bc2a6d4c9a524fa41dc1c297ee378d2acb0ad54c38b68c "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/acosh"
    i46e8d5dfc629fbad06935065311cb4fed88bfdbff165a8f8e2f720b918a700ea "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/coth"
    i4754a47a3b6303895ed69638f39124035edcdda6699250aad1531a3c94d76c8e "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/dcounta"
    i4884add5aa3eb5601d03ab509d95736ec1fc8844dae31e2af280f26ca05632a5 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/cumprinc"
    i492970e7012b0e3e57541d3e1fb52c45d9e86f9cf9d45f1f8e4d3b498fef0ec7 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/imsum"
    i4a28b15112d3eb1c9daf437595d1c8ed680301669d4c3b6bb1a4c5d82e8f2a3c "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/asinh"
    i4a4e247f672df0066887e62056a41a014663030a786167984fe09b18a220f24f "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/percentile_exc"
    i4a6097157b72180db7ddd2c7f18b5bcf5fcb552868620e4eb3e0cabb9c634aa6 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/xnpv"
    i4a96440b3734be843c4b51c553f8a21fce777610999e21a0c062b60b80230a9a "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/dvar"
    i4ace284db0c47bd0aede99134978f5c929609ce4e35b202bdc8f6500f9efbdbc "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/tan"
    i4ad15c91bffdd6a7acc7559f7a10a51b66d05a9e3af49cebecda7381579555b3 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/dstdevp"
    i4c5ae7091bd09004549bef0c9652ff1ae3fb0bf193fe9ded298576b924e200a1 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/len"
    i4cfbd3520b5e74b2b8045a64e640927da5eac9de8112b4e237992f1e226e4814 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/acot"
    i4d961ff236da4359d12fbdfa7eee89fd03184cb0e49a2e2752677085a30472d5 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/left"
    i50381e463138dff6b12cd8560121ead89cc7ccadaf433cf12d03fb404c64c9bf "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/floor_math"
    i505aa17044c18b82f550f25288473c9ef4eb4704713abd4d25f343671d631204 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/ispmt"
    i5071f5260556cf090013e5064eda236cca516fed4d302085452c10f4c14ba7fc "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/hypgeom_dist"
    i50e91789dd90831d9488dc1a97a93610c341349db57691f090a7769ebc37e4a8 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/countblank"
    i53dc72af898ff473ccd17dd44368585359f3b67d1ff8aa7e80aaf3102680c763 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/z_test"
    i559163cfcd0e3e255e4a9a51f12ccc72dfd4e1513c60909da25bbe744c7b7eed "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/quartile_inc"
    i55f82606d2062de63a09cfa317d19a38bc44cf9caca3f84be3b452b7888fc07f "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/replaceb"
    i56028c0a8d8e19dc408b210807675ce31e98253a9bb7a1e6d13cd087691dd0f2 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/chisq_inv"
    i566e7cb1af1639c06f9ebeba3ccb8a1fdc6c345eeefbfe6acac09d1582c12126 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/sln"
    i5693d5ef4e4534c8934b5902523fea06b2534f831e4373120963f1d7186f3fc3 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/rri"
    i56c64d431ad2709306c509a9f9f0ec8e4d894d5c4bff5eeefdd4cbda7a7b7587 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/not"
    i58e768b6d2fa089950374be07f39abc3adf5e6b058aa73f7d55895cba00cf8f9 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/daverage"
    i5947d51544c1c2349b970821fbd25c451d4082bdae4c87a7b4d3b30efa11c91d "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/edate"
    i59aff93824d19ad9146734be1350addc911717282bdfb6e5e50619b006989eaa "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/iso_ceiling"
    i5a73a60fca271f2a6e4ea5782835bbeb6b9e46a1ff77cbe422456e88a3a280f4 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/mina"
    i5b15e4d00f30aeae062247f8166cfc2cb8d59c9df116fe7e2f059fd95d9901f3 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/complex"
    i5b1782155c5c405bee8b8f5a33cba1abcf8e62c7784228ab4eb0aafed5551b0c "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/oddfprice"
    i5d2b89aed9f287e850042083336337a009be43b6571b38dd4a953ae60b0ecfd6 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/sec"
    i5dc621c21d96565fe810946c44544466dd6248d89e84a7626b434803d541c4b5 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/sumif"
    i6014adf79a8085e3e438be1d03d71d6a94be7c73a6f54d011025b6869298a304 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/imabs"
    i60cd4be8a98f401f373ef391d0d720d59a0175eb2da787338a8c47ca72230756 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/int"
    i60e394d3b7133e81856f6d4d53d288fe70c710fd2154a56322e7b421951fd182 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/combina"
    i61ffd8641cbb3c8a3866e6c4bb21d96c551dbadf0ed8c3dd0be6f104e33403b3 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/besselj"
    i6233f1d5ffabf3630193ff6919645c7ac85a1e1635e8136ae522a7bfd3ffb273 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/rounddown"
    i62f9d809bb32318e6fa706fa90f69b0be2c4f5ce7ab88f2190db66f8d8b74fe4 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/npv"
    i64a2584bf10ad8de049d14b396033b7516bde2e1c43d45bdfe01e57fa63fd156 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/exp"
    i64cd0ff82c53728496e0d6426872ae2a123ca37929a7d2775137213719fa0377 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/coupdaybs"
    i64dcf16df745efb2a277f9da7811014fb9e46d69a69d3d2c0d98c40271d60c41 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/lookup"
    i65dce0cd1c63cecf0ab371edc274c37bdc073b35276cfb5d1568c31a942277f7 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/bin2dec"
    i6635c9b9d1a681ece3749da33e90e966d20a0993e382e510bbd2789e2741ea27 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/stdevpa"
    i66a6fa74a3368903bb93f6617cd38885b5e2b4f3fa3102aac5ac15e66b5cace5 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/atan2"
    i670b8d255cccdd20e535f3c9f26bfa95cb9e18d53062194370910cf773a878f3 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/phi"
    i6856ff2d3a83f0fc27c438aed87b8b7bc8df33e54d4639861cbd0e840cdab9bc "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/averagea"
    i68deb2e05608a96df2e0ca27f79d5923264ab2ba69b0b423931ee7b0a6582c6a "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/trimmean"
    i68ed9c8d4e625536900643b1d40f8ccd5aea6c7e5a3262f769fa0aa3ceeeace0 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/atan"
    i69701c6709d9985805f83da725f7ec822215ad3364b0f2d80da6e08ef7c08a42 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/dproduct"
    i6a5180477c05d488f65a77d747e4059176532e7fdb70500facf4d07dd8761a64 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/islogical"
    i6af25ac649d63cffd177e7d031b0347ab699ff6d27d2fc95425b75f23c4bf93d "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/fvschedule"
    i6b66ab4ec1d212e8495dd66f10b6a82a906af5bdd344e286a87d49fdb988d600 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/varpa"
    i6b81f000dd827003c2cce3b132f6286772878283234a01c0e2cfcbf8a5a1c9b9 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/csc"
    i6bc19e2249e93a4bf23661983c03ec617bb683cbdbfdd722b26080d00ca1a73e "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/confidence_t"
    i6bce0181254fb93d7930572c11d9dd1b75a83e5b3645600893fa1fff1d12d7b0 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/trunc"
    i6c60f2a52e3ae8c843199907b1e7e75d8fa9cd0c6d4b44f2a1d0ef8436c1c17d "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/dvarp"
    i6c856566e659f6bb93261ea7e1b9d0d666d90bcc1fcdb924a50ddaecb8791e2a "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/and"
    i6d1c5d20d33fc36fa7f2b3544243b168c8398ac6bad21b9c78e284fdbb93b93f "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/combin"
    i6d3601e4620b7b6f10fbca1f300891246e6557070abaf6d0386bb594b467848b "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/columns"
    i6d38b4a0b1957aac95edf4d270f1233268dd9b75f70f69210e4ae5ea32629ed4 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/bitrshift"
    i6f2d58290fdd3849b9d1a88c1c57213b6d67e7f72feb084b24b9cb19e5f4a580 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/count"
    i6f8f30774095daf20cfddf1d05be778001542ade98513f0407f2bcd5da65dea9 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/code"
    i70a53c457b4f61c31d118ea80495b0701da5b8be62fedf3afd3cc50d4b715424 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/expon_dist"
    i7452bb10748075477cba95fcbc0c13c2d89ffdcbff9f76e0d57843f166d47c09 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/imconjugate"
    i75be6168523edf64b90022fdf9bec646c1bfc8ef1efa5fa3ac44b4afd2e3eb37 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/proper"
    i763ba95b5f563556a70f24c79a10763e911fab1b34f6028e7d5edb86cba65062 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/degrees"
    i765e2dc1e4db8a288a0bfd1170c6f962898bca87ac6757fc74b6fbb622678bc2 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/fact"
    i79e82baba0b6ad85021449e22432171eb56e9fdf4f2f7704c3aa88bcd6c58ad1 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/f_inv"
    i7a405238184cd8f69856daa2e7423da2c3d858797d6f153b711887b5aff45587 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/exact"
    i7afce34f3f2fd116a978ece18049d97b8250d2f8d24f54c2b7244b820740b61b "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/tbillprice"
    i7b102aaeba7077cd75bd4b4829916459ef28397accd5982c90e6a3a60d059865 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/sqrtpi"
    i7dfe45d82e3dda0f3316e137bedc1b70710e63e18987a0fb781c89837a0b5260 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/t_dist"
    i7f3bdefbf8fd1669aacf56b2cc408855c56c01593e6510244dd00a4c06dac54f "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/concatenate"
    i80c0df997d334f32be5b33abf7bb1214ab8f06300beaacc75d9acbd254bd7fed "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/nominal"
    i82621983f40ab0734ef85bfbb618ffaa7cad9f67cfb178f33096ebf6673889f7 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/mid"
    i82f8a69830531d76eab554d0589b70ba84e9ea54e49ddab273bff7c4693966a5 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/bahttext"
    i82fdfa05294be1ba769b77abf428e73287bc5a10a4c911c9625b2030af94e9e2 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/upper"
    i83006d85d2ec9e9c821e0df31373485e9bd4b920829e4a4980282767b0a72eb0 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/var_p"
    i8384057cf305481cf65d4ea1b5ad309efd1eb6cbf2e43f5d6937c31a076dc3c2 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/mirr"
    i83ea1fbcbefdac4f93b0345b4d5d2a5e65ba70d393aa82950e87b9d5048a334e "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/lognorm_inv"
    i83ff9d1937f074321dfcd778ed52ef8deac4669fd191114eb80bf60660c99517 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/pduration"
    i8451f7c3341a403be03d74fa7ff4bce0b6cff9a7985aa2bcc16a5eb72d4eb911 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/value"
    i851e9134195868314efcca697670811addf31221342640ba1a490891597a3664 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/oct2bin"
    i869a657b2ae7ad505dee44529f6b7b72249331958fbba1421c6c7ce183f865db "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/timevalue"
    i86ef0e159279f96f70079c0b4493ddddbd65f6f451676e153884ef85a549febf "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/stdev_s"
    i88b361cc03b1e56e044d4445f4d89a0ecb754efe4ea824b27420646e43020419 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/unichar"
    i8a5aa2ca50c45a9aaa4a64a26c03d619e6ac6eb2ce39d9d7e0d0d75791d41d27 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/lcm"
    i8ac88088cb1eeb8d34fde38032a24e268fd5ea5f3d5b1d1dd6920dd78c5f2f01 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/sumifs"
    i8aefbb33c5c2a864ea5a1d50f2b1f89ae7bb4a02e2f6b68c31f1c22d54dd7840 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/bin2oct"
    i8bdd663dc6e442ec62f88e5eb50ebbc8aa5436c55981ab70bfd57082bc01a470 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/bitor"
    i8c9b12150110dff23884f98804b25416e782ee4cf9818fd9a1b45f85ceb50ed4 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/coupnum"
    i8caef61af8f8a889701fec9d0da76ceb2fe1d50c922be6a77f1c055a873a66a4 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/gamma_inv"
    i8d3d59cd13ac07770cca89a1e57d61c44887692d79a8cf032c27110ec3ce63c7 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/dollar"
    i8d50d793930b499a674af3fec201df5db71eda2abab0233872f8eb5111ba3aff "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/days"
    i8d5e244078615b973dc2a759b4d35335b31f7071c851da4dbc55fe0d0f7dc9f6 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/erfc"
    i8dd17ce26e31c69bf75ded2876e67eb6f10cccbfddce7006913070ca62dabbe8 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/dec2bin"
    i8e4cda1c8d4d879c336f3412bf3cf660300b4382767348a71dd7fb969b2a6192 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/quotient"
    i8e5d631c1427fcf55be915b53173e402b28a1ae7010a4e4eb883defce512ba55 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/bessely"
    i8f663f1f1cfc333d0c8519ce44a252b1dc321a2509021d5cea45983899f75c0f "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/irr"
    i8fa46174c81b311f05fab89ad00c2e5a54e905ed5d900848b66b8c55f033334f "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/gammaln"
    i91014446a234d7aee31201c0b75a75f41c56ac2352caaafbe59f81ced3df23eb "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/minute"
    i93785e3a673fb9020e0f0ce85ee7cae614300dd743e9dcceb38cb99e8cecde42 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/disc"
    i938a669433496fcdc36cb0d50f11ce775a2b7fbce462efe42da5f705832dcc26 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/stdeva"
    i93a333ba632bc00534645305713d91f78a439c34f2de5bf13e758d28fac8103a "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/imtan"
    i93c21983a784b61538963bd9d57c7dcbb5585d5cb15f477bb3890722edf15131 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/factdouble"
    i93f8f23d8d5723f4c27ac12031f23090341ad1da31ec8769d014d1dde1ec5739 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/imcsc"
    i963416bc83cf905778043344d0bab1da3fca59fc6c8c051b4575a104cbdf28e3 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/t_dist_2t"
    i96dec6eeb36d6bbc03b666f68b3bdc9d1d43aa8b2bb484895dc9a732ab7e3f99 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/arabic"
    i97090a4a7e2dd73810bd3d01799a0104da063a747c3ee8971844b46528a76b40 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/eomonth"
    i985e10a00be70c0d4c56e4933e1f13b0e6e96eafb096e52d812287d89fc04a26 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/small"
    i986d90630b8bd4704c82b2fdf90293b3ab0e120575e89da616037e92bf0059d1 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/isoweeknum"
    i9998c115b39ccdcda164097ad8c6e64a01279ccc04c7659a8bad4f83c7156268 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/counta"
    i99c3f166a423758c32dbc701aa5590ca9233e7d0fe1ac471c4373b792487f0c9 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/max"
    i9a3f732bd7241183f682437f2f058d6f652e9ba5cc72ff3e03857f0788befb86 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/min"
    i9a47a986544d6a5f06b03d7cb7453c0870eb8874b2f226ff9499ae252d8e65e5 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/rank_eq"
    i9a6dbb59c9ed734bc1782519bdf4a1d44e774d2a6f4e474bceb9ef2ca4833392 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/gamma"
    i9a7860e17c1a21c709fe83d415604d69e827b6a98f94305171b5ec3b01f769bd "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/var_s"
    i9a8c3c1caa6ef8ce9239ac05447a657a5c1df8e3fc86936c6a80fe3fde5af182 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/xirr"
    i9aa03510a433102dd03798716d94e40befa72b313c5eea5043a03f9c60bcf6fc "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/or"
    i9b6def06d0eea432c15b352f850397af2fc381e9ba3df44c71d1d2fcf5f68ff6 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/char"
    i9c29a5a8283fefc59b3748875d07ee0baed2a86e2ed47ff8a3aaf9229d514305 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/dollarfr"
    i9d7b75a1b7e470f87e877b53d428a3510edf27826fd7887599769cfb74e9254d "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/unicode"
    i9e7061d8c4b334646d3b9442f20b3d0f0139b2c32cf905f7ec5cf4e479c36327 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/month"
    i9e941f3832437c523ffcd8acd8a92118e25b7a78ba5e92934af586e216183102 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/floor_precise"
    ia08221458d3ce92ffa340c9aa72763f26ef30b9af826a69b3fb3297753f6d1e3 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/roundup"
    ia08e5cb104669018ca6e491458a452542cef1b63353541697e68648e7b1031e8 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/vara"
    ia0c51d44b14541036535a1c7d7dab0e5e713dc165de517410a1f12c4f862c574 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/imln"
    ia152b34300806c4b7358ab8944b42af0de7aac3c22ad861b524beee369bdcd32 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/averageifs"
    ia177a9f5c18691cc4406e66e13883d49fbdd9680bfb701c43311e417d1161835 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/quartile_exc"
    ia2368438ab65c6ddbe0b26f6ba1296d19fcdb897869c0f0744ce722675daad33 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/trim"
    ia251f0aa1f0b2cfbdbb1576a91aaa4d1136f52468bb576fcb08dddeaa0499903 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/asc"
    ia480020e4a6c2e08476997d30d7775a6b23fd3fef49c44bfe41efb91bfc5482e "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/hlookup"
    ia4b182a5c215c546905ed2917110a5553ab0fd097232bd6b3a079abcd7d85723 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/substitute"
    ia4c89f7e9370bce5de71aa5a04bba1cb2dc0defa1923953b9aba644ff6115f29 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/isodd"
    ia4d363d4f1ca7b24dcccc1ef2cd633f884a8d61f1acf1699f6d8a72ebf14737b "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/oddlyield"
    ia4f2ee6d2f82f9a2d3b1aa5510ad49e1c0c19d30e2756cb5032f0f376a908870 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/even"
    ia4f9e47d4e53e3eac9bb21e885d76281217b5ccfe774b3e87b0ebc017a58d50f "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/hex2bin"
    ia72c98a4751729a1e0614fc8ed05a45d4d39faa0ff7481aa4c31bae0ade72188 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/stdev_p"
    ia8bfcef4032b0b7b3cfd0aa9712c6695c4c4baef57a47bff0df8c2d3e3cbe8de "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/lognorm_dist"
    iaa84aa290eae65b463a1ae1a6c3c4c4bc756a998a754e7118ceb8eae9ff2d6e6 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/effect"
    iaa87c9aca478301cc253d3d7e5bfe18a8cd7a4b5b78b7a47fe04bd32cf9f8517 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/received"
    iaacdbe1e397058fb6b88798374843ea2675a9a2f3a341db474e4363dbb71793f "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/abs"
    iac60b2c91ba5e9073d79b2404aa84140abc5eb1fa5031e6fbf82ce642af5d310 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/time"
    iacf0ea4e9ca20f4848adbfde01a685454f860c4de3f8013934409c3cbb368e88 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/imreal"
    iad8c455c383d91e2ddc21056c2cf1ad006674858b6e9d1b318fb40a4d1ead44a "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/product"
    iaf300ff360e1f70c118ab1245a8805547261654f42fcf5e927329cef9c94b60e "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/odd"
    iaf8b5876f5af8a9f7a5ef8e7f8b3e0fc749df12943de859ec1580bb6cb234ea9 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/mod"
    iaf8e3d23027fdaf5b8ca992293bed72048279fb31fa871cff3a2f3f1a0da7144 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/improduct"
    ib213aa115a5783efc6a93a4547d27e9aba61ae37bd2a46ec2535a3ed61aa0537 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/tanh"
    ib24fa4f1e36f60866106cd4336f027a6eed101dc3872e5373ab1775c0d8f56af "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/csch"
    ib2e975e6222afb294e23e89e8673480e7656c04e27ac37f0248bbf50598b984c "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/confidence_norm"
    ib316263ec23cf028a1c9c43f97a6c87006bd31b72eab6e609a69453c7b1e8249 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/round"
    ib6550e75fbfdd9e37216fc4794caa6c4e81f25f92c381a6ffa3d58fe2873c6f5 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/ipmt"
    ib673860838d2a67b4e7dff1294ba96e0d0aec3a3a68b8c5213ab6aa777bf03ba "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/gauss"
    ib7a0edb5aa7958aa836226b0f07db012d5ff202d020c5bb7ab64fbce74978422 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/countif"
    ib7f7db4b356c3f0edd4a9082a87e69d8998d0582eb2b0584623289ad5701c513 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/chisq_dist"
    ib8295a41b8411110d323251838d9fb6d72cec878915644fa9149be58e9f19c3e "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/seriessum"
    ib8859f36db353950f680ccddd5f4852712454947f41ab65661b9a34c8eeacb14 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/ln"
    ib8c96c5eea945469ccedea2b68135503234eeb7add02c3da84f01264bb80d249 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/sqrt"
    ib986f65cf08c1509d5d5863841d24bbdbb8a670c65cc079e60805b955053434f "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/erf_precise"
    ibb028562b184f1114c572f1bfd4ecdd09c376808379f82fea18ab98530345d6a "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/sin"
    ibb4b60e00ae41cdee94346a18b7e427da2d7d055714ab6e431a83dd2d89ce743 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/ddb"
    ibbb26bea14a287ff49e53453caa33dbd62106e41fcc9daafdd036da1c9ae6549 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/percentile_inc"
    ibc437c6427620d2d6935f429015bf0c10a2efc14ea03389317f69042cb0a5675 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/couppcd"
    ibd0b49ef1c39b06635846dd9dcb0e6f0a9011de6063d0aa2c4c2b1767d01c418 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/sinh"
    ibd26ce201da12322ee502415fc2557dcb546bff2f58bd0a0a66b118743d29263 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/bitlshift"
    ibdc76d0292afb9a985b5244ce64cb6b69cc4ba0b43b1f44ce50759619239f7a2 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/dcount"
    ibe16069a015d9a991398b696604a0d7a50b8d4a0a5173021c47dc7d31b8488e7 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/cos"
    ibf6a6d809a25a2a33016ecf8fdf760c94d91711708465fac845a89db9d8f556e "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/besselk"
    ibff0f8dc02f42b95aaf558016884991d09301c80f0c47f7362135a26671457f8 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/find"
    ic1f613dc5187698d819b0aa08d2ad012fde20b02f33d821c2fec65ef88ee79f6 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/imsec"
    ic252910cd61d1027a295d59661b3f0559e1293642f79081245cdb05f9599f7a4 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/match"
    ic36d8e74f949e6d32934a7cfcdaf4efea58b9682c075b49628f1cea48acc72be "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/bitand"
    ic3951154c4dbc4f18015ccc1d7414fc2a172a85714344787fb1275900b85bb78 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/gamma_dist"
    ic425e6efc51044806a510230c7303eb3179d638506cc5fedd86de91216089b21 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/dget"
    ic4442d05ae33e3592c99a59c7ea34f74b3a63364fcd7ceffa6b4032cb787e84c "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/hex2dec"
    ic499eeaaef51669c1c8054f66f2ecd14ee3a2d2184057dacca6a48b2811bb6d2 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/chisq_dist_rt"
    ic502dc6126b618cf966cc3eda56369ef70a539cb821b815eeb940cd5ac10e249 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/acoth"
    ic6288d7f3b52cf2063447b3f9794fc8b3dc4e6f34dca8e1ecb05955a7ab2ea9c "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/cot"
    ic6d083c78589d498dc0b395084aa2bd1401f2fd9a6f2d03336fcfea7e6c65e16 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/error_type"
    ic75fb62de06bcf50e6f7ce850a4840cf36b8bfb7d2e49f295b2badc4c46f1fc2 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/dbcs"
    ic790acc862cd6e1a97f1f9148e01bd57da0a487382e1add1eeef1153d85ff9a5 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/duration"
    ic8b44e7573fd7c941987e253f26d7389a87bfca781b04b5bcbf2c27fb25f8616 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/t_inv_2t"
    ic932478eb02695cb9c1766e2919fb21f692b10fd800981f7c8590cab8a333b56 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/yield"
    ica5cc4cfc9127bfa086a9705314366a7ad715f4e0686b90fcaa0b3125318914e "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/t_dist_rt"
    icaef813c1a68be08058cf4c8982fec7f193c7f4ce4d210432c546739b2e837da "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/percentrank_inc"
    icb260c694333510d9258bb8d576e81be9fc30faa8929963d3be12b97ee63bef0 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/binom_dist_range"
    icc8e84210ed02bc389729a7a57e6183f9d3583c7b675c2b73874160538e26f9d "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/decimal"
    icd0ec1f1466ecbe7ab9a65dbd6e31713844263d6c0983e7b7574d98d9f910a00 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/days360"
    icd7b760eebee4cd326831f82d28e124ec9d530ad0cf6bca53e53ce3ff1dbb1af "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/imlog2"
    ice2b692e5c82f169b0bfa2f37f0215a2eb1020ee060f49a315caa4420c43b6ce "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/workday_intl"
    ice907567aec7422264bd904682d942a146136dd2e818b14bc812458e91dad04f "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/accrint"
    icedf2711d1b11d1597a25ff1dceea21dc8397e613cff25b8ed7a4ff1cdfaf381 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/large"
    id0142a852a9c647fb622b73201977e57aa86e47ce78f59bbc8799e24ad96c296 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/multinomial"
    id0ab7dda529fcb6b25781c1d2cd57da3924ae8281797dc28cc48d2c8ba39a2b7 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/oct2hex"
    id15ae092e5e71c8fd67b0d9dc90012d5d94c1951c2f73267ff19ab3047707f55 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/amordegrc"
    id18fc1c9e53e811409aa175b1903d8a4231f5e1b21c1f6c08ae8bbd9b4fd2c3e "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/second"
    id3cc5e976c999c1f47e4185f7ca5a05df40cc09f816869ec70b3c32edbfe4cc2 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/coupncd"
    id446936ca56486e000b19e23dbd525b966a48aee1300a88134721dd5cf9efd9e "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/rows"
    id45c75ee8baefe98f15057c7666cb289b4eff355de908057521641b52f602dad "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/f_inv_rt"
    id5f48c79968fe4a26b921d06e11cdbc22451895ec74a40ff23bd2af6f9f6bfe4 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/price"
    id608eda20ad5780629a89de437920cd619ad7ea1ab54a8a2d38c169436e5221f "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/dec2hex"
    id64721fa9455c034fbb573e2a0b0bc6d698db5ab904ebe920ddffa534e0e1e6f "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/coupdaysnc"
    id6488ecb4a4e5c1f26048e3228034097bd61dd097baee10c729251d53ef89185 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/amorlinc"
    id6dcff57797b11fba1fcdb52de657107a8c156f9ad8d483dc42b2c441152fd77 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/median"
    id719949da934cf20c3c44bdc0722b2534c7d0c91c0572433ba43b92dba4765b2 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/hex2oct"
    id7d817f874340b589502b9ff0c20273a5aab557e903fb1df341b0912f3a2a025 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/db"
    id822393bcd6685a69bedfef7aa06f206b88d195277a7bf0d4a1c044e4cc0e270 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/dmin"
    id94bd29f13aaee26ed2f3e6fdbd875e4cac32863214435c3ed3eb09d14fa6d8d "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/dollarde"
    id969821a72a1cf354ea78cfd2ddce9e4fa199548dfe0cb362443d8cf26d23c22 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/averageif"
    id9cec7065b42565aeb87054ce216fafd8b4934b7a69129ec555ac31cda884b1a "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/dstdev"
    idab688590eba2c52ae8ce39e7c029e05705eaf3846f9f3661da25fe15b15a2e0 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/permut"
    idc8fd3f08e7ffc89e34663dc85de24b296c593e9aa28ecfd0c5c1fde908d1caa "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/isna"
    idcca0ede7f1179661a6261ae372a93f01e7276d27ae334bed6b68ed754c93500 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/beta_dist"
    iddcfbe2f11f04c4acb3d6005a2592baa86f0985fc04742f4eb3a30ff2b5f82e0 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/imsqrt"
    ide699df9204c2fb4c05840413b3fef64453c8456166397ced94412afbcd7572c "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/midb"
    idf32f364f7c0dd785101c2eb2dfdc4ae095700b99137c793ce79d9e482f5684b "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/bin2hex"
    idfce0e11b43129d348eab5e2d598a5e343fe5d432d4225dc1aec75c1017df6e4 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/permutationa"
    ie2bdf3ef3c78f4f44af114a402c97bbb77df98d9cddab9ba7301caf4fff6fb66 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/imsinh"
    ie3218b4af599d768c57d5aa1130d1ec1c6779febe7e16e6499a9456d46c026a4 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/usdollar"
    ie3a60eaa52f0e2c83224f957f6c5fd743ef6922115a77b7a30b636c2ddd0f617 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/subtotal"
    ie45871831f52f50388688cb40e938774f60c309d595c6da228650c05e48e2f2d "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/erf"
    ie487f292c394041439201b1836fb377bbaf737623fb82bed4786600999e0a332 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/sheets"
    ie4c7b2ccc0fefd3b2166097de5afc1341fbe3df95ae3077cfb857b85eee03a24 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/mduration"
    ie54b5add0586442f2594c622471e354c7e7e13086776a6587d0de11b8f32c685 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/fisherinv"
    ie5716615358214f3b3b2e47c740f34ff277d58709f634223426fe9032bf24b6f "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/n"
    ie5f4e649a2430cb9ba32512c7cd00b5175d24789f7845a39081aa6e917d6ad49 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/numbervalue"
    ie6397e87d1cd0c2613c317e8cf61006e73a296a728d9390fcd75bd30517971fe "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/geomean"
    ie650d4648a9f8150fff1faca1b39152f94ae4bef583ce99ac183fd9ff7f6f362 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/yieldmat"
    ie84b383822ebdf2fb50939420eb43921caadfdced5cd685bff8fae6d1fef9dd6 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/erfc_precise"
    ie8c1476b60a276425ece63c534f4e3c1ee0409af598fc7be59a08023d52ccacd "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/mround"
    ie9f3949526768b26584d54e2031840cce1419e5425128f5260fc078d3eacfe3d "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/norm_dist"
    iea0b5aa5ad4cf5c06d500264f61a91dfa902e71561c07c9f328fdfa6874a7725 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/besseli"
    iead39b340d7b27388c9912da5044632d1d33a0c34dc34ae5dcb1aa93456f1f70 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/rightb"
    ieb5426682c782d2a15d99495fa6ea6438abe7e1e528e9fccdcc79bdf4e31105b "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/weibull_dist"
    ied0c2d1a639b42db012bda18c0f9b2f7c4ebab2f9b1e3028b9209e5df07386ae "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/pricemat"
    iefb8b5dc5714d55e2e4537e8bdb401e5c22894cefd038573046698e064963b32 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/norm_s_inv"
    if2cc99747d2a61f6ea9c0ec7aee2987e560c1e63342dd98f49246f082274726c "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/norm_inv"
    if314cc65ab505afff71197b6b58adf355ac2583ec88039f4833d4245565cdb45 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/oct2dec"
    if344cf49bccda079366198f2bd231540bf5bff866fbbb61dd98ac2e8433d1341 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/binom_inv"
    if40f99f52f24942945f4b9ca132698b8658c20b8be734bd51f2f36d554686efd "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/cosh"
    if44efbdfc9b5c14c1fc3f5d0ddb7ed3bd98a90400a4e96b1d18314d4d7765659 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/sum"
    if45416702cd1d18be7686d76f0e214785248198fdbc18dd4876ecf03112da042 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/radians"
    if4be56e795a06bb1c9ca704768f0c7e99112bcaee942408865b012e925ee96d7 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/false"
    if4c493403752863f1ada69246676520b0ede3e19a8ca2643caf3fd5dcf0a07d5 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/accrintm"
    if672ced31e76bd2dec896594956640ddc503e702f15673b496157f5d13b489b0 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/gestep"
    if6ca53a02d3f7eb4a0ef9b12ddc8dde0d6b2cc96ae1a41b91c028d7daf9ead56 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/yearfrac"
    if6db57cb9f25d39cd5d2bd50eb111ef58f546d0d2d9da25b255d33ab5c29ea31 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/imargument"
    if71db3cbdf59abfdb3c8e61da5a55a3813b1e3cd7365c27838a7f35de2d04665 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/acos"
    if7e680f3499d68a27fe79595542c43bfb9cdcc74173f7ab77e839125313ae4a2 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/rank_avg"
    if85e354c04136d477faa4e0f4230b6aa3f937d7654bcc233236b241bf7b1be3a "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/negbinom_dist"
    if8e737917131ca982223412916f7e20cc3b2fdc7e0c7fa5cac8c2dbcf066536a "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/workday"
    if90a89d555d1426169b10dbb9ac2554cda2ce1f4f2f7839205cd2fcbfa03444a "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/rate"
    if9338e79dad08cd3805e50da66aea478338f153a9d4d999956c67ff2071bfadd "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/datevalue"
    if9ff5147ed19f76469c0737eaf1588b869564fa419d964e6acd6ddb95f1f65a3 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/roman"
    ifba3e217442e7b422f933a4732577709c619487368acff3f8aaa56e855079820 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/networkdays_intl"
    ifc0b96049e408483dcda28b9da9af8b3b9d0e3ab22420ca3ca0c69580846792f "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/sheet"
    ifc654a64ea750853647dab0efb222301689322f2c0f6756df87f42142737aff3 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/isformula"
    ifc77bdc03a4d73e8f9f6ecd5a111eaf0822b7692d8ac2b7bfec7261049d4dee9 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/iseven"
    ife1070a90db9ee740de63da00f5e400585fdf3b2d8073fdc24415dc354e42ced "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/fv"
    ifec202f6e2f8b4f71b4533ff606b32733c1244f51a6f1c2e6d45eabf4b17a8c5 "github.com/microsoftgraph/msgraph-sdk-go/workbooks/item/workbook/functions/true"
)

// FunctionsRequestBuilder builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\functions
type FunctionsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// FunctionsRequestBuilderDeleteOptions options for Delete
type FunctionsRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// FunctionsRequestBuilderGetOptions options for Get
type FunctionsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *FunctionsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// FunctionsRequestBuilderGetQueryParameters get functions from workbooks
type FunctionsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// FunctionsRequestBuilderPatchOptions options for Patch
type FunctionsRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookFunctions;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *FunctionsRequestBuilder) Abs()(*iaacdbe1e397058fb6b88798374843ea2675a9a2f3a341db474e4363dbb71793f.AbsRequestBuilder) {
    return iaacdbe1e397058fb6b88798374843ea2675a9a2f3a341db474e4363dbb71793f.NewAbsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) AccrInt()(*ice907567aec7422264bd904682d942a146136dd2e818b14bc812458e91dad04f.AccrIntRequestBuilder) {
    return ice907567aec7422264bd904682d942a146136dd2e818b14bc812458e91dad04f.NewAccrIntRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) AccrIntM()(*if4c493403752863f1ada69246676520b0ede3e19a8ca2643caf3fd5dcf0a07d5.AccrIntMRequestBuilder) {
    return if4c493403752863f1ada69246676520b0ede3e19a8ca2643caf3fd5dcf0a07d5.NewAccrIntMRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Acos()(*if71db3cbdf59abfdb3c8e61da5a55a3813b1e3cd7365c27838a7f35de2d04665.AcosRequestBuilder) {
    return if71db3cbdf59abfdb3c8e61da5a55a3813b1e3cd7365c27838a7f35de2d04665.NewAcosRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Acosh()(*i466df9e270b2d70578bc2a6d4c9a524fa41dc1c297ee378d2acb0ad54c38b68c.AcoshRequestBuilder) {
    return i466df9e270b2d70578bc2a6d4c9a524fa41dc1c297ee378d2acb0ad54c38b68c.NewAcoshRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Acot()(*i4cfbd3520b5e74b2b8045a64e640927da5eac9de8112b4e237992f1e226e4814.AcotRequestBuilder) {
    return i4cfbd3520b5e74b2b8045a64e640927da5eac9de8112b4e237992f1e226e4814.NewAcotRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Acoth()(*ic502dc6126b618cf966cc3eda56369ef70a539cb821b815eeb940cd5ac10e249.AcothRequestBuilder) {
    return ic502dc6126b618cf966cc3eda56369ef70a539cb821b815eeb940cd5ac10e249.NewAcothRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) AmorDegrc()(*id15ae092e5e71c8fd67b0d9dc90012d5d94c1951c2f73267ff19ab3047707f55.AmorDegrcRequestBuilder) {
    return id15ae092e5e71c8fd67b0d9dc90012d5d94c1951c2f73267ff19ab3047707f55.NewAmorDegrcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) AmorLinc()(*id6488ecb4a4e5c1f26048e3228034097bd61dd097baee10c729251d53ef89185.AmorLincRequestBuilder) {
    return id6488ecb4a4e5c1f26048e3228034097bd61dd097baee10c729251d53ef89185.NewAmorLincRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) And()(*i6c856566e659f6bb93261ea7e1b9d0d666d90bcc1fcdb924a50ddaecb8791e2a.AndRequestBuilder) {
    return i6c856566e659f6bb93261ea7e1b9d0d666d90bcc1fcdb924a50ddaecb8791e2a.NewAndRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Arabic()(*i96dec6eeb36d6bbc03b666f68b3bdc9d1d43aa8b2bb484895dc9a732ab7e3f99.ArabicRequestBuilder) {
    return i96dec6eeb36d6bbc03b666f68b3bdc9d1d43aa8b2bb484895dc9a732ab7e3f99.NewArabicRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Areas()(*i0c002f577eeb35c2aea092a0a697271289620508c6edc625840b274f2f02aa67.AreasRequestBuilder) {
    return i0c002f577eeb35c2aea092a0a697271289620508c6edc625840b274f2f02aa67.NewAreasRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Asc()(*ia251f0aa1f0b2cfbdbb1576a91aaa4d1136f52468bb576fcb08dddeaa0499903.AscRequestBuilder) {
    return ia251f0aa1f0b2cfbdbb1576a91aaa4d1136f52468bb576fcb08dddeaa0499903.NewAscRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Asin()(*i196ecc451091a63b4e5a4d20f9956fe33210551341fd2b70b3f7c6f33bc90d0d.AsinRequestBuilder) {
    return i196ecc451091a63b4e5a4d20f9956fe33210551341fd2b70b3f7c6f33bc90d0d.NewAsinRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Asinh()(*i4a28b15112d3eb1c9daf437595d1c8ed680301669d4c3b6bb1a4c5d82e8f2a3c.AsinhRequestBuilder) {
    return i4a28b15112d3eb1c9daf437595d1c8ed680301669d4c3b6bb1a4c5d82e8f2a3c.NewAsinhRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Atan()(*i68ed9c8d4e625536900643b1d40f8ccd5aea6c7e5a3262f769fa0aa3ceeeace0.AtanRequestBuilder) {
    return i68ed9c8d4e625536900643b1d40f8ccd5aea6c7e5a3262f769fa0aa3ceeeace0.NewAtanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Atan2()(*i66a6fa74a3368903bb93f6617cd38885b5e2b4f3fa3102aac5ac15e66b5cace5.Atan2RequestBuilder) {
    return i66a6fa74a3368903bb93f6617cd38885b5e2b4f3fa3102aac5ac15e66b5cace5.NewAtan2RequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Atanh()(*i27abe0dd8fef776e61d85d1af494b69c706d34811423987b8537ef7ac3a4065d.AtanhRequestBuilder) {
    return i27abe0dd8fef776e61d85d1af494b69c706d34811423987b8537ef7ac3a4065d.NewAtanhRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) AveDev()(*i30dd8a3bec806587e9da967e397a243b809e20c4421c0454edac1b48ae5814a0.AveDevRequestBuilder) {
    return i30dd8a3bec806587e9da967e397a243b809e20c4421c0454edac1b48ae5814a0.NewAveDevRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Average()(*i181dd5fee2d584c64321f0aa2fc010aa892dcc53eb3f40e83a6a3eb9c52f1e60.AverageRequestBuilder) {
    return i181dd5fee2d584c64321f0aa2fc010aa892dcc53eb3f40e83a6a3eb9c52f1e60.NewAverageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) AverageA()(*i6856ff2d3a83f0fc27c438aed87b8b7bc8df33e54d4639861cbd0e840cdab9bc.AverageARequestBuilder) {
    return i6856ff2d3a83f0fc27c438aed87b8b7bc8df33e54d4639861cbd0e840cdab9bc.NewAverageARequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) AverageIf()(*id969821a72a1cf354ea78cfd2ddce9e4fa199548dfe0cb362443d8cf26d23c22.AverageIfRequestBuilder) {
    return id969821a72a1cf354ea78cfd2ddce9e4fa199548dfe0cb362443d8cf26d23c22.NewAverageIfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) AverageIfs()(*ia152b34300806c4b7358ab8944b42af0de7aac3c22ad861b524beee369bdcd32.AverageIfsRequestBuilder) {
    return ia152b34300806c4b7358ab8944b42af0de7aac3c22ad861b524beee369bdcd32.NewAverageIfsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) BahtText()(*i82f8a69830531d76eab554d0589b70ba84e9ea54e49ddab273bff7c4693966a5.BahtTextRequestBuilder) {
    return i82f8a69830531d76eab554d0589b70ba84e9ea54e49ddab273bff7c4693966a5.NewBahtTextRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Base()(*i25706560466145b8e6f0ca13c02b864543dd157f7b8d5b9e8152f50ef7060f92.BaseRequestBuilder) {
    return i25706560466145b8e6f0ca13c02b864543dd157f7b8d5b9e8152f50ef7060f92.NewBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) BesselI()(*iea0b5aa5ad4cf5c06d500264f61a91dfa902e71561c07c9f328fdfa6874a7725.BesselIRequestBuilder) {
    return iea0b5aa5ad4cf5c06d500264f61a91dfa902e71561c07c9f328fdfa6874a7725.NewBesselIRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) BesselJ()(*i61ffd8641cbb3c8a3866e6c4bb21d96c551dbadf0ed8c3dd0be6f104e33403b3.BesselJRequestBuilder) {
    return i61ffd8641cbb3c8a3866e6c4bb21d96c551dbadf0ed8c3dd0be6f104e33403b3.NewBesselJRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) BesselK()(*ibf6a6d809a25a2a33016ecf8fdf760c94d91711708465fac845a89db9d8f556e.BesselKRequestBuilder) {
    return ibf6a6d809a25a2a33016ecf8fdf760c94d91711708465fac845a89db9d8f556e.NewBesselKRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) BesselY()(*i8e5d631c1427fcf55be915b53173e402b28a1ae7010a4e4eb883defce512ba55.BesselYRequestBuilder) {
    return i8e5d631c1427fcf55be915b53173e402b28a1ae7010a4e4eb883defce512ba55.NewBesselYRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Beta_Dist()(*idcca0ede7f1179661a6261ae372a93f01e7276d27ae334bed6b68ed754c93500.Beta_DistRequestBuilder) {
    return idcca0ede7f1179661a6261ae372a93f01e7276d27ae334bed6b68ed754c93500.NewBeta_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Beta_Inv()(*i0409f0a1cf64e0274eaf7de30693fa8cbba20462831b9d50c75dcb1110b2c48e.Beta_InvRequestBuilder) {
    return i0409f0a1cf64e0274eaf7de30693fa8cbba20462831b9d50c75dcb1110b2c48e.NewBeta_InvRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Bin2Dec()(*i65dce0cd1c63cecf0ab371edc274c37bdc073b35276cfb5d1568c31a942277f7.Bin2DecRequestBuilder) {
    return i65dce0cd1c63cecf0ab371edc274c37bdc073b35276cfb5d1568c31a942277f7.NewBin2DecRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Bin2Hex()(*idf32f364f7c0dd785101c2eb2dfdc4ae095700b99137c793ce79d9e482f5684b.Bin2HexRequestBuilder) {
    return idf32f364f7c0dd785101c2eb2dfdc4ae095700b99137c793ce79d9e482f5684b.NewBin2HexRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Bin2Oct()(*i8aefbb33c5c2a864ea5a1d50f2b1f89ae7bb4a02e2f6b68c31f1c22d54dd7840.Bin2OctRequestBuilder) {
    return i8aefbb33c5c2a864ea5a1d50f2b1f89ae7bb4a02e2f6b68c31f1c22d54dd7840.NewBin2OctRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Binom_Dist()(*i3715210d20a7eb43e392d3a551304d689734e97b74fc573d51234daae97ece0c.Binom_DistRequestBuilder) {
    return i3715210d20a7eb43e392d3a551304d689734e97b74fc573d51234daae97ece0c.NewBinom_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Binom_Dist_Range()(*icb260c694333510d9258bb8d576e81be9fc30faa8929963d3be12b97ee63bef0.Binom_Dist_RangeRequestBuilder) {
    return icb260c694333510d9258bb8d576e81be9fc30faa8929963d3be12b97ee63bef0.NewBinom_Dist_RangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Binom_Inv()(*if344cf49bccda079366198f2bd231540bf5bff866fbbb61dd98ac2e8433d1341.Binom_InvRequestBuilder) {
    return if344cf49bccda079366198f2bd231540bf5bff866fbbb61dd98ac2e8433d1341.NewBinom_InvRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Bitand()(*ic36d8e74f949e6d32934a7cfcdaf4efea58b9682c075b49628f1cea48acc72be.BitandRequestBuilder) {
    return ic36d8e74f949e6d32934a7cfcdaf4efea58b9682c075b49628f1cea48acc72be.NewBitandRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Bitlshift()(*ibd26ce201da12322ee502415fc2557dcb546bff2f58bd0a0a66b118743d29263.BitlshiftRequestBuilder) {
    return ibd26ce201da12322ee502415fc2557dcb546bff2f58bd0a0a66b118743d29263.NewBitlshiftRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Bitor()(*i8bdd663dc6e442ec62f88e5eb50ebbc8aa5436c55981ab70bfd57082bc01a470.BitorRequestBuilder) {
    return i8bdd663dc6e442ec62f88e5eb50ebbc8aa5436c55981ab70bfd57082bc01a470.NewBitorRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Bitrshift()(*i6d38b4a0b1957aac95edf4d270f1233268dd9b75f70f69210e4ae5ea32629ed4.BitrshiftRequestBuilder) {
    return i6d38b4a0b1957aac95edf4d270f1233268dd9b75f70f69210e4ae5ea32629ed4.NewBitrshiftRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Bitxor()(*i15a001f719e754b66f338d6f5feeb9021d4d0325bf7e573450ebf7407b0625c8.BitxorRequestBuilder) {
    return i15a001f719e754b66f338d6f5feeb9021d4d0325bf7e573450ebf7407b0625c8.NewBitxorRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Ceiling_Math()(*i1e19512606a6be24eefaac467f8c47c472b2dcdf8240b0e328a02af67678bdbe.Ceiling_MathRequestBuilder) {
    return i1e19512606a6be24eefaac467f8c47c472b2dcdf8240b0e328a02af67678bdbe.NewCeiling_MathRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Ceiling_Precise()(*i10f134b1fa93754f8715e8f54b26ea798cefb9f01e18cd0b9f5001dddbfb687b.Ceiling_PreciseRequestBuilder) {
    return i10f134b1fa93754f8715e8f54b26ea798cefb9f01e18cd0b9f5001dddbfb687b.NewCeiling_PreciseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Char()(*i9b6def06d0eea432c15b352f850397af2fc381e9ba3df44c71d1d2fcf5f68ff6.CharRequestBuilder) {
    return i9b6def06d0eea432c15b352f850397af2fc381e9ba3df44c71d1d2fcf5f68ff6.NewCharRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ChiSq_Dist()(*ib7f7db4b356c3f0edd4a9082a87e69d8998d0582eb2b0584623289ad5701c513.ChiSq_DistRequestBuilder) {
    return ib7f7db4b356c3f0edd4a9082a87e69d8998d0582eb2b0584623289ad5701c513.NewChiSq_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ChiSq_Dist_RT()(*ic499eeaaef51669c1c8054f66f2ecd14ee3a2d2184057dacca6a48b2811bb6d2.ChiSq_Dist_RTRequestBuilder) {
    return ic499eeaaef51669c1c8054f66f2ecd14ee3a2d2184057dacca6a48b2811bb6d2.NewChiSq_Dist_RTRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ChiSq_Inv()(*i56028c0a8d8e19dc408b210807675ce31e98253a9bb7a1e6d13cd087691dd0f2.ChiSq_InvRequestBuilder) {
    return i56028c0a8d8e19dc408b210807675ce31e98253a9bb7a1e6d13cd087691dd0f2.NewChiSq_InvRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ChiSq_Inv_RT()(*i2ec5919ffdb405f5ddc6809e4991195c3055085124325cfde9f900ea1a1816fc.ChiSq_Inv_RTRequestBuilder) {
    return i2ec5919ffdb405f5ddc6809e4991195c3055085124325cfde9f900ea1a1816fc.NewChiSq_Inv_RTRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Choose()(*i3548428d79c006b6c84182290c7c63cb0c418cf298886e1351488dc98bb5f7d1.ChooseRequestBuilder) {
    return i3548428d79c006b6c84182290c7c63cb0c418cf298886e1351488dc98bb5f7d1.NewChooseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Clean()(*i16adcfc7ead6bed8c6f0c25fa14d883aaa2b6874e7ff8bf7fc5b559ba0f5ada8.CleanRequestBuilder) {
    return i16adcfc7ead6bed8c6f0c25fa14d883aaa2b6874e7ff8bf7fc5b559ba0f5ada8.NewCleanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Code()(*i6f8f30774095daf20cfddf1d05be778001542ade98513f0407f2bcd5da65dea9.CodeRequestBuilder) {
    return i6f8f30774095daf20cfddf1d05be778001542ade98513f0407f2bcd5da65dea9.NewCodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Columns()(*i6d3601e4620b7b6f10fbca1f300891246e6557070abaf6d0386bb594b467848b.ColumnsRequestBuilder) {
    return i6d3601e4620b7b6f10fbca1f300891246e6557070abaf6d0386bb594b467848b.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Combin()(*i6d1c5d20d33fc36fa7f2b3544243b168c8398ac6bad21b9c78e284fdbb93b93f.CombinRequestBuilder) {
    return i6d1c5d20d33fc36fa7f2b3544243b168c8398ac6bad21b9c78e284fdbb93b93f.NewCombinRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Combina()(*i60e394d3b7133e81856f6d4d53d288fe70c710fd2154a56322e7b421951fd182.CombinaRequestBuilder) {
    return i60e394d3b7133e81856f6d4d53d288fe70c710fd2154a56322e7b421951fd182.NewCombinaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Complex()(*i5b15e4d00f30aeae062247f8166cfc2cb8d59c9df116fe7e2f059fd95d9901f3.ComplexRequestBuilder) {
    return i5b15e4d00f30aeae062247f8166cfc2cb8d59c9df116fe7e2f059fd95d9901f3.NewComplexRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Concatenate()(*i7f3bdefbf8fd1669aacf56b2cc408855c56c01593e6510244dd00a4c06dac54f.ConcatenateRequestBuilder) {
    return i7f3bdefbf8fd1669aacf56b2cc408855c56c01593e6510244dd00a4c06dac54f.NewConcatenateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Confidence_Norm()(*ib2e975e6222afb294e23e89e8673480e7656c04e27ac37f0248bbf50598b984c.Confidence_NormRequestBuilder) {
    return ib2e975e6222afb294e23e89e8673480e7656c04e27ac37f0248bbf50598b984c.NewConfidence_NormRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Confidence_T()(*i6bc19e2249e93a4bf23661983c03ec617bb683cbdbfdd722b26080d00ca1a73e.Confidence_TRequestBuilder) {
    return i6bc19e2249e93a4bf23661983c03ec617bb683cbdbfdd722b26080d00ca1a73e.NewConfidence_TRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewFunctionsRequestBuilderInternal instantiates a new FunctionsRequestBuilder and sets the default values.
func NewFunctionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*FunctionsRequestBuilder) {
    m := &FunctionsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/workbooks/{driveItem_id}/workbook/functions{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewFunctionsRequestBuilder instantiates a new FunctionsRequestBuilder and sets the default values.
func NewFunctionsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*FunctionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFunctionsRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *FunctionsRequestBuilder) Convert()(*i447958a9355da6b66054ac0c3105f155673cc9c53508f382a4b0343ab49fe403.ConvertRequestBuilder) {
    return i447958a9355da6b66054ac0c3105f155673cc9c53508f382a4b0343ab49fe403.NewConvertRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Cos()(*ibe16069a015d9a991398b696604a0d7a50b8d4a0a5173021c47dc7d31b8488e7.CosRequestBuilder) {
    return ibe16069a015d9a991398b696604a0d7a50b8d4a0a5173021c47dc7d31b8488e7.NewCosRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Cosh()(*if40f99f52f24942945f4b9ca132698b8658c20b8be734bd51f2f36d554686efd.CoshRequestBuilder) {
    return if40f99f52f24942945f4b9ca132698b8658c20b8be734bd51f2f36d554686efd.NewCoshRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Cot()(*ic6288d7f3b52cf2063447b3f9794fc8b3dc4e6f34dca8e1ecb05955a7ab2ea9c.CotRequestBuilder) {
    return ic6288d7f3b52cf2063447b3f9794fc8b3dc4e6f34dca8e1ecb05955a7ab2ea9c.NewCotRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Coth()(*i46e8d5dfc629fbad06935065311cb4fed88bfdbff165a8f8e2f720b918a700ea.CothRequestBuilder) {
    return i46e8d5dfc629fbad06935065311cb4fed88bfdbff165a8f8e2f720b918a700ea.NewCothRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Count()(*i6f2d58290fdd3849b9d1a88c1c57213b6d67e7f72feb084b24b9cb19e5f4a580.CountRequestBuilder) {
    return i6f2d58290fdd3849b9d1a88c1c57213b6d67e7f72feb084b24b9cb19e5f4a580.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) CountA()(*i9998c115b39ccdcda164097ad8c6e64a01279ccc04c7659a8bad4f83c7156268.CountARequestBuilder) {
    return i9998c115b39ccdcda164097ad8c6e64a01279ccc04c7659a8bad4f83c7156268.NewCountARequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) CountBlank()(*i50e91789dd90831d9488dc1a97a93610c341349db57691f090a7769ebc37e4a8.CountBlankRequestBuilder) {
    return i50e91789dd90831d9488dc1a97a93610c341349db57691f090a7769ebc37e4a8.NewCountBlankRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) CountIf()(*ib7a0edb5aa7958aa836226b0f07db012d5ff202d020c5bb7ab64fbce74978422.CountIfRequestBuilder) {
    return ib7a0edb5aa7958aa836226b0f07db012d5ff202d020c5bb7ab64fbce74978422.NewCountIfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) CountIfs()(*i1d5ff915b0e62f1dbdf9af3a3d3a41f310dfde12a83f5753a8b75863457d468d.CountIfsRequestBuilder) {
    return i1d5ff915b0e62f1dbdf9af3a3d3a41f310dfde12a83f5753a8b75863457d468d.NewCountIfsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) CoupDayBs()(*i64cd0ff82c53728496e0d6426872ae2a123ca37929a7d2775137213719fa0377.CoupDayBsRequestBuilder) {
    return i64cd0ff82c53728496e0d6426872ae2a123ca37929a7d2775137213719fa0377.NewCoupDayBsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) CoupDays()(*i25a82e0bc1d4b3db1874f6069a464756206a19345bb46a65038cb4d86d297e1a.CoupDaysRequestBuilder) {
    return i25a82e0bc1d4b3db1874f6069a464756206a19345bb46a65038cb4d86d297e1a.NewCoupDaysRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) CoupDaysNc()(*id64721fa9455c034fbb573e2a0b0bc6d698db5ab904ebe920ddffa534e0e1e6f.CoupDaysNcRequestBuilder) {
    return id64721fa9455c034fbb573e2a0b0bc6d698db5ab904ebe920ddffa534e0e1e6f.NewCoupDaysNcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) CoupNcd()(*id3cc5e976c999c1f47e4185f7ca5a05df40cc09f816869ec70b3c32edbfe4cc2.CoupNcdRequestBuilder) {
    return id3cc5e976c999c1f47e4185f7ca5a05df40cc09f816869ec70b3c32edbfe4cc2.NewCoupNcdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) CoupNum()(*i8c9b12150110dff23884f98804b25416e782ee4cf9818fd9a1b45f85ceb50ed4.CoupNumRequestBuilder) {
    return i8c9b12150110dff23884f98804b25416e782ee4cf9818fd9a1b45f85ceb50ed4.NewCoupNumRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) CoupPcd()(*ibc437c6427620d2d6935f429015bf0c10a2efc14ea03389317f69042cb0a5675.CoupPcdRequestBuilder) {
    return ibc437c6427620d2d6935f429015bf0c10a2efc14ea03389317f69042cb0a5675.NewCoupPcdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property functions for workbooks
func (m *FunctionsRequestBuilder) CreateDeleteRequestInformation(options *FunctionsRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
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
// CreateGetRequestInformation get functions from workbooks
func (m *FunctionsRequestBuilder) CreateGetRequestInformation(options *FunctionsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property functions in workbooks
func (m *FunctionsRequestBuilder) CreatePatchRequestInformation(options *FunctionsRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *FunctionsRequestBuilder) Csc()(*i6b81f000dd827003c2cce3b132f6286772878283234a01c0e2cfcbf8a5a1c9b9.CscRequestBuilder) {
    return i6b81f000dd827003c2cce3b132f6286772878283234a01c0e2cfcbf8a5a1c9b9.NewCscRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Csch()(*ib24fa4f1e36f60866106cd4336f027a6eed101dc3872e5373ab1775c0d8f56af.CschRequestBuilder) {
    return ib24fa4f1e36f60866106cd4336f027a6eed101dc3872e5373ab1775c0d8f56af.NewCschRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) CumIPmt()(*i310b333bbd7ca74c312923d9e2339d29519809126287db8ffbd43d3ac2b3f32b.CumIPmtRequestBuilder) {
    return i310b333bbd7ca74c312923d9e2339d29519809126287db8ffbd43d3ac2b3f32b.NewCumIPmtRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) CumPrinc()(*i4884add5aa3eb5601d03ab509d95736ec1fc8844dae31e2af280f26ca05632a5.CumPrincRequestBuilder) {
    return i4884add5aa3eb5601d03ab509d95736ec1fc8844dae31e2af280f26ca05632a5.NewCumPrincRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Date()(*i12c177d210a36c02a13c649a27eee3563c73121e988b53d1a5b0e07d943b1b5e.DateRequestBuilder) {
    return i12c177d210a36c02a13c649a27eee3563c73121e988b53d1a5b0e07d943b1b5e.NewDateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Datevalue()(*if9338e79dad08cd3805e50da66aea478338f153a9d4d999956c67ff2071bfadd.DatevalueRequestBuilder) {
    return if9338e79dad08cd3805e50da66aea478338f153a9d4d999956c67ff2071bfadd.NewDatevalueRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Daverage()(*i58e768b6d2fa089950374be07f39abc3adf5e6b058aa73f7d55895cba00cf8f9.DaverageRequestBuilder) {
    return i58e768b6d2fa089950374be07f39abc3adf5e6b058aa73f7d55895cba00cf8f9.NewDaverageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Day()(*i179bb0c14bc435ddfa798d98f5b9ceed77f0d74d656eab292a71b5c2cdcefb31.DayRequestBuilder) {
    return i179bb0c14bc435ddfa798d98f5b9ceed77f0d74d656eab292a71b5c2cdcefb31.NewDayRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Days()(*i8d50d793930b499a674af3fec201df5db71eda2abab0233872f8eb5111ba3aff.DaysRequestBuilder) {
    return i8d50d793930b499a674af3fec201df5db71eda2abab0233872f8eb5111ba3aff.NewDaysRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Days360()(*icd0ec1f1466ecbe7ab9a65dbd6e31713844263d6c0983e7b7574d98d9f910a00.Days360RequestBuilder) {
    return icd0ec1f1466ecbe7ab9a65dbd6e31713844263d6c0983e7b7574d98d9f910a00.NewDays360RequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Db()(*id7d817f874340b589502b9ff0c20273a5aab557e903fb1df341b0912f3a2a025.DbRequestBuilder) {
    return id7d817f874340b589502b9ff0c20273a5aab557e903fb1df341b0912f3a2a025.NewDbRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Dbcs()(*ic75fb62de06bcf50e6f7ce850a4840cf36b8bfb7d2e49f295b2badc4c46f1fc2.DbcsRequestBuilder) {
    return ic75fb62de06bcf50e6f7ce850a4840cf36b8bfb7d2e49f295b2badc4c46f1fc2.NewDbcsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Dcount()(*ibdc76d0292afb9a985b5244ce64cb6b69cc4ba0b43b1f44ce50759619239f7a2.DcountRequestBuilder) {
    return ibdc76d0292afb9a985b5244ce64cb6b69cc4ba0b43b1f44ce50759619239f7a2.NewDcountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) DcountA()(*i4754a47a3b6303895ed69638f39124035edcdda6699250aad1531a3c94d76c8e.DcountARequestBuilder) {
    return i4754a47a3b6303895ed69638f39124035edcdda6699250aad1531a3c94d76c8e.NewDcountARequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Ddb()(*ibb4b60e00ae41cdee94346a18b7e427da2d7d055714ab6e431a83dd2d89ce743.DdbRequestBuilder) {
    return ibb4b60e00ae41cdee94346a18b7e427da2d7d055714ab6e431a83dd2d89ce743.NewDdbRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Dec2Bin()(*i8dd17ce26e31c69bf75ded2876e67eb6f10cccbfddce7006913070ca62dabbe8.Dec2BinRequestBuilder) {
    return i8dd17ce26e31c69bf75ded2876e67eb6f10cccbfddce7006913070ca62dabbe8.NewDec2BinRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Dec2Hex()(*id608eda20ad5780629a89de437920cd619ad7ea1ab54a8a2d38c169436e5221f.Dec2HexRequestBuilder) {
    return id608eda20ad5780629a89de437920cd619ad7ea1ab54a8a2d38c169436e5221f.NewDec2HexRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Dec2Oct()(*i355b3d4ac7803c6a215bb8801502ba5fd3379a7c9cc52a9aa21240d59fa1d832.Dec2OctRequestBuilder) {
    return i355b3d4ac7803c6a215bb8801502ba5fd3379a7c9cc52a9aa21240d59fa1d832.NewDec2OctRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Decimal()(*icc8e84210ed02bc389729a7a57e6183f9d3583c7b675c2b73874160538e26f9d.DecimalRequestBuilder) {
    return icc8e84210ed02bc389729a7a57e6183f9d3583c7b675c2b73874160538e26f9d.NewDecimalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Degrees()(*i763ba95b5f563556a70f24c79a10763e911fab1b34f6028e7d5edb86cba65062.DegreesRequestBuilder) {
    return i763ba95b5f563556a70f24c79a10763e911fab1b34f6028e7d5edb86cba65062.NewDegreesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property functions for workbooks
func (m *FunctionsRequestBuilder) Delete(options *FunctionsRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *FunctionsRequestBuilder) Delta()(*i172364b817972ee8d188e239878e9d28a9730a745b39c8c7059c6177ad20ae4d.DeltaRequestBuilder) {
    return i172364b817972ee8d188e239878e9d28a9730a745b39c8c7059c6177ad20ae4d.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) DevSq()(*i382d1b83df58a9e122bb42db7279c4e8cc39d16055184c25db824773341d6d04.DevSqRequestBuilder) {
    return i382d1b83df58a9e122bb42db7279c4e8cc39d16055184c25db824773341d6d04.NewDevSqRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Dget()(*ic425e6efc51044806a510230c7303eb3179d638506cc5fedd86de91216089b21.DgetRequestBuilder) {
    return ic425e6efc51044806a510230c7303eb3179d638506cc5fedd86de91216089b21.NewDgetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Disc()(*i93785e3a673fb9020e0f0ce85ee7cae614300dd743e9dcceb38cb99e8cecde42.DiscRequestBuilder) {
    return i93785e3a673fb9020e0f0ce85ee7cae614300dd743e9dcceb38cb99e8cecde42.NewDiscRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Dmax()(*i0b3dc8bf2c5b69445e4aee0c44a1b80461aa78ff46d273df01cd9203922b28ca.DmaxRequestBuilder) {
    return i0b3dc8bf2c5b69445e4aee0c44a1b80461aa78ff46d273df01cd9203922b28ca.NewDmaxRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Dmin()(*id822393bcd6685a69bedfef7aa06f206b88d195277a7bf0d4a1c044e4cc0e270.DminRequestBuilder) {
    return id822393bcd6685a69bedfef7aa06f206b88d195277a7bf0d4a1c044e4cc0e270.NewDminRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Dollar()(*i8d3d59cd13ac07770cca89a1e57d61c44887692d79a8cf032c27110ec3ce63c7.DollarRequestBuilder) {
    return i8d3d59cd13ac07770cca89a1e57d61c44887692d79a8cf032c27110ec3ce63c7.NewDollarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) DollarDe()(*id94bd29f13aaee26ed2f3e6fdbd875e4cac32863214435c3ed3eb09d14fa6d8d.DollarDeRequestBuilder) {
    return id94bd29f13aaee26ed2f3e6fdbd875e4cac32863214435c3ed3eb09d14fa6d8d.NewDollarDeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) DollarFr()(*i9c29a5a8283fefc59b3748875d07ee0baed2a86e2ed47ff8a3aaf9229d514305.DollarFrRequestBuilder) {
    return i9c29a5a8283fefc59b3748875d07ee0baed2a86e2ed47ff8a3aaf9229d514305.NewDollarFrRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Dproduct()(*i69701c6709d9985805f83da725f7ec822215ad3364b0f2d80da6e08ef7c08a42.DproductRequestBuilder) {
    return i69701c6709d9985805f83da725f7ec822215ad3364b0f2d80da6e08ef7c08a42.NewDproductRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) DstDev()(*id9cec7065b42565aeb87054ce216fafd8b4934b7a69129ec555ac31cda884b1a.DstDevRequestBuilder) {
    return id9cec7065b42565aeb87054ce216fafd8b4934b7a69129ec555ac31cda884b1a.NewDstDevRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) DstDevP()(*i4ad15c91bffdd6a7acc7559f7a10a51b66d05a9e3af49cebecda7381579555b3.DstDevPRequestBuilder) {
    return i4ad15c91bffdd6a7acc7559f7a10a51b66d05a9e3af49cebecda7381579555b3.NewDstDevPRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Dsum()(*i1fc62d517918671eb2537fb2ec20d1e162fb4a5c9321c5e1d59d09cd12b56d55.DsumRequestBuilder) {
    return i1fc62d517918671eb2537fb2ec20d1e162fb4a5c9321c5e1d59d09cd12b56d55.NewDsumRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Duration()(*ic790acc862cd6e1a97f1f9148e01bd57da0a487382e1add1eeef1153d85ff9a5.DurationRequestBuilder) {
    return ic790acc862cd6e1a97f1f9148e01bd57da0a487382e1add1eeef1153d85ff9a5.NewDurationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Dvar()(*i4a96440b3734be843c4b51c553f8a21fce777610999e21a0c062b60b80230a9a.DvarRequestBuilder) {
    return i4a96440b3734be843c4b51c553f8a21fce777610999e21a0c062b60b80230a9a.NewDvarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) DvarP()(*i6c60f2a52e3ae8c843199907b1e7e75d8fa9cd0c6d4b44f2a1d0ef8436c1c17d.DvarPRequestBuilder) {
    return i6c60f2a52e3ae8c843199907b1e7e75d8fa9cd0c6d4b44f2a1d0ef8436c1c17d.NewDvarPRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Ecma_Ceiling()(*i26345570575fd85318c268057a084f3e7920644be05134a08f409b82ef2d087e.Ecma_CeilingRequestBuilder) {
    return i26345570575fd85318c268057a084f3e7920644be05134a08f409b82ef2d087e.NewEcma_CeilingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Edate()(*i5947d51544c1c2349b970821fbd25c451d4082bdae4c87a7b4d3b30efa11c91d.EdateRequestBuilder) {
    return i5947d51544c1c2349b970821fbd25c451d4082bdae4c87a7b4d3b30efa11c91d.NewEdateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Effect()(*iaa84aa290eae65b463a1ae1a6c3c4c4bc756a998a754e7118ceb8eae9ff2d6e6.EffectRequestBuilder) {
    return iaa84aa290eae65b463a1ae1a6c3c4c4bc756a998a754e7118ceb8eae9ff2d6e6.NewEffectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) EoMonth()(*i97090a4a7e2dd73810bd3d01799a0104da063a747c3ee8971844b46528a76b40.EoMonthRequestBuilder) {
    return i97090a4a7e2dd73810bd3d01799a0104da063a747c3ee8971844b46528a76b40.NewEoMonthRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Erf()(*ie45871831f52f50388688cb40e938774f60c309d595c6da228650c05e48e2f2d.ErfRequestBuilder) {
    return ie45871831f52f50388688cb40e938774f60c309d595c6da228650c05e48e2f2d.NewErfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Erf_Precise()(*ib986f65cf08c1509d5d5863841d24bbdbb8a670c65cc079e60805b955053434f.Erf_PreciseRequestBuilder) {
    return ib986f65cf08c1509d5d5863841d24bbdbb8a670c65cc079e60805b955053434f.NewErf_PreciseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ErfC()(*i8d5e244078615b973dc2a759b4d35335b31f7071c851da4dbc55fe0d0f7dc9f6.ErfCRequestBuilder) {
    return i8d5e244078615b973dc2a759b4d35335b31f7071c851da4dbc55fe0d0f7dc9f6.NewErfCRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ErfC_Precise()(*ie84b383822ebdf2fb50939420eb43921caadfdced5cd685bff8fae6d1fef9dd6.ErfC_PreciseRequestBuilder) {
    return ie84b383822ebdf2fb50939420eb43921caadfdced5cd685bff8fae6d1fef9dd6.NewErfC_PreciseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Error_Type()(*ic6d083c78589d498dc0b395084aa2bd1401f2fd9a6f2d03336fcfea7e6c65e16.Error_TypeRequestBuilder) {
    return ic6d083c78589d498dc0b395084aa2bd1401f2fd9a6f2d03336fcfea7e6c65e16.NewError_TypeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Even()(*ia4f2ee6d2f82f9a2d3b1aa5510ad49e1c0c19d30e2756cb5032f0f376a908870.EvenRequestBuilder) {
    return ia4f2ee6d2f82f9a2d3b1aa5510ad49e1c0c19d30e2756cb5032f0f376a908870.NewEvenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Exact()(*i7a405238184cd8f69856daa2e7423da2c3d858797d6f153b711887b5aff45587.ExactRequestBuilder) {
    return i7a405238184cd8f69856daa2e7423da2c3d858797d6f153b711887b5aff45587.NewExactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Exp()(*i64a2584bf10ad8de049d14b396033b7516bde2e1c43d45bdfe01e57fa63fd156.ExpRequestBuilder) {
    return i64a2584bf10ad8de049d14b396033b7516bde2e1c43d45bdfe01e57fa63fd156.NewExpRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Expon_Dist()(*i70a53c457b4f61c31d118ea80495b0701da5b8be62fedf3afd3cc50d4b715424.Expon_DistRequestBuilder) {
    return i70a53c457b4f61c31d118ea80495b0701da5b8be62fedf3afd3cc50d4b715424.NewExpon_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) F_Dist()(*i445b80f059057d65d647c9ef7d80680e79c644d587f2eb3b38fa0efdb98a56dd.F_DistRequestBuilder) {
    return i445b80f059057d65d647c9ef7d80680e79c644d587f2eb3b38fa0efdb98a56dd.NewF_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) F_Dist_RT()(*i368eb1aa0b23f0d4e8267b91adca60d8fa7015f90d191cc4cc826cd78864f82a.F_Dist_RTRequestBuilder) {
    return i368eb1aa0b23f0d4e8267b91adca60d8fa7015f90d191cc4cc826cd78864f82a.NewF_Dist_RTRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) F_Inv()(*i79e82baba0b6ad85021449e22432171eb56e9fdf4f2f7704c3aa88bcd6c58ad1.F_InvRequestBuilder) {
    return i79e82baba0b6ad85021449e22432171eb56e9fdf4f2f7704c3aa88bcd6c58ad1.NewF_InvRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) F_Inv_RT()(*id45c75ee8baefe98f15057c7666cb289b4eff355de908057521641b52f602dad.F_Inv_RTRequestBuilder) {
    return id45c75ee8baefe98f15057c7666cb289b4eff355de908057521641b52f602dad.NewF_Inv_RTRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Fact()(*i765e2dc1e4db8a288a0bfd1170c6f962898bca87ac6757fc74b6fbb622678bc2.FactRequestBuilder) {
    return i765e2dc1e4db8a288a0bfd1170c6f962898bca87ac6757fc74b6fbb622678bc2.NewFactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) FactDouble()(*i93c21983a784b61538963bd9d57c7dcbb5585d5cb15f477bb3890722edf15131.FactDoubleRequestBuilder) {
    return i93c21983a784b61538963bd9d57c7dcbb5585d5cb15f477bb3890722edf15131.NewFactDoubleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) False()(*if4be56e795a06bb1c9ca704768f0c7e99112bcaee942408865b012e925ee96d7.FalseRequestBuilder) {
    return if4be56e795a06bb1c9ca704768f0c7e99112bcaee942408865b012e925ee96d7.NewFalseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Find()(*ibff0f8dc02f42b95aaf558016884991d09301c80f0c47f7362135a26671457f8.FindRequestBuilder) {
    return ibff0f8dc02f42b95aaf558016884991d09301c80f0c47f7362135a26671457f8.NewFindRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) FindB()(*i2114c9762c18b212ecaeaf82ff100b472674b7b5525cfd98d132aa1caa772c56.FindBRequestBuilder) {
    return i2114c9762c18b212ecaeaf82ff100b472674b7b5525cfd98d132aa1caa772c56.NewFindBRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Fisher()(*i434657a2f6458186aeeb593186f3d5915f1bf0677f364b0cecbf804a6d4c5ee9.FisherRequestBuilder) {
    return i434657a2f6458186aeeb593186f3d5915f1bf0677f364b0cecbf804a6d4c5ee9.NewFisherRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) FisherInv()(*ie54b5add0586442f2594c622471e354c7e7e13086776a6587d0de11b8f32c685.FisherInvRequestBuilder) {
    return ie54b5add0586442f2594c622471e354c7e7e13086776a6587d0de11b8f32c685.NewFisherInvRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Fixed()(*i11766f797ebc2faf35ca5bd9fc2be9ca3099a81ea6ecfafb6c6678f6355ffc63.FixedRequestBuilder) {
    return i11766f797ebc2faf35ca5bd9fc2be9ca3099a81ea6ecfafb6c6678f6355ffc63.NewFixedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Floor_Math()(*i50381e463138dff6b12cd8560121ead89cc7ccadaf433cf12d03fb404c64c9bf.Floor_MathRequestBuilder) {
    return i50381e463138dff6b12cd8560121ead89cc7ccadaf433cf12d03fb404c64c9bf.NewFloor_MathRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Floor_Precise()(*i9e941f3832437c523ffcd8acd8a92118e25b7a78ba5e92934af586e216183102.Floor_PreciseRequestBuilder) {
    return i9e941f3832437c523ffcd8acd8a92118e25b7a78ba5e92934af586e216183102.NewFloor_PreciseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Fv()(*ife1070a90db9ee740de63da00f5e400585fdf3b2d8073fdc24415dc354e42ced.FvRequestBuilder) {
    return ife1070a90db9ee740de63da00f5e400585fdf3b2d8073fdc24415dc354e42ced.NewFvRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Fvschedule()(*i6af25ac649d63cffd177e7d031b0347ab699ff6d27d2fc95425b75f23c4bf93d.FvscheduleRequestBuilder) {
    return i6af25ac649d63cffd177e7d031b0347ab699ff6d27d2fc95425b75f23c4bf93d.NewFvscheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Gamma()(*i9a6dbb59c9ed734bc1782519bdf4a1d44e774d2a6f4e474bceb9ef2ca4833392.GammaRequestBuilder) {
    return i9a6dbb59c9ed734bc1782519bdf4a1d44e774d2a6f4e474bceb9ef2ca4833392.NewGammaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Gamma_Dist()(*ic3951154c4dbc4f18015ccc1d7414fc2a172a85714344787fb1275900b85bb78.Gamma_DistRequestBuilder) {
    return ic3951154c4dbc4f18015ccc1d7414fc2a172a85714344787fb1275900b85bb78.NewGamma_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Gamma_Inv()(*i8caef61af8f8a889701fec9d0da76ceb2fe1d50c922be6a77f1c055a873a66a4.Gamma_InvRequestBuilder) {
    return i8caef61af8f8a889701fec9d0da76ceb2fe1d50c922be6a77f1c055a873a66a4.NewGamma_InvRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) GammaLn()(*i8fa46174c81b311f05fab89ad00c2e5a54e905ed5d900848b66b8c55f033334f.GammaLnRequestBuilder) {
    return i8fa46174c81b311f05fab89ad00c2e5a54e905ed5d900848b66b8c55f033334f.NewGammaLnRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) GammaLn_Precise()(*i42ce97198842ba0a177d6ed68d50d81eafa58190eb0f5ed32a2b12a2a29d3f1f.GammaLn_PreciseRequestBuilder) {
    return i42ce97198842ba0a177d6ed68d50d81eafa58190eb0f5ed32a2b12a2a29d3f1f.NewGammaLn_PreciseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Gauss()(*ib673860838d2a67b4e7dff1294ba96e0d0aec3a3a68b8c5213ab6aa777bf03ba.GaussRequestBuilder) {
    return ib673860838d2a67b4e7dff1294ba96e0d0aec3a3a68b8c5213ab6aa777bf03ba.NewGaussRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Gcd()(*i16a2c7c1c23ba3e4a26a9034153b98e031a4d2505b415d8a6f93ab701c906928.GcdRequestBuilder) {
    return i16a2c7c1c23ba3e4a26a9034153b98e031a4d2505b415d8a6f93ab701c906928.NewGcdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) GeoMean()(*ie6397e87d1cd0c2613c317e8cf61006e73a296a728d9390fcd75bd30517971fe.GeoMeanRequestBuilder) {
    return ie6397e87d1cd0c2613c317e8cf61006e73a296a728d9390fcd75bd30517971fe.NewGeoMeanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) GeStep()(*if672ced31e76bd2dec896594956640ddc503e702f15673b496157f5d13b489b0.GeStepRequestBuilder) {
    return if672ced31e76bd2dec896594956640ddc503e702f15673b496157f5d13b489b0.NewGeStepRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get functions from workbooks
func (m *FunctionsRequestBuilder) Get(options *FunctionsRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookFunctions, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewWorkbookFunctions() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WorkbookFunctions), nil
}
func (m *FunctionsRequestBuilder) HarMean()(*i35bac92c54bdeaf46ce03cb02622ca9ef5b34152ddafea2898d92790e10f28ac.HarMeanRequestBuilder) {
    return i35bac92c54bdeaf46ce03cb02622ca9ef5b34152ddafea2898d92790e10f28ac.NewHarMeanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Hex2Bin()(*ia4f9e47d4e53e3eac9bb21e885d76281217b5ccfe774b3e87b0ebc017a58d50f.Hex2BinRequestBuilder) {
    return ia4f9e47d4e53e3eac9bb21e885d76281217b5ccfe774b3e87b0ebc017a58d50f.NewHex2BinRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Hex2Dec()(*ic4442d05ae33e3592c99a59c7ea34f74b3a63364fcd7ceffa6b4032cb787e84c.Hex2DecRequestBuilder) {
    return ic4442d05ae33e3592c99a59c7ea34f74b3a63364fcd7ceffa6b4032cb787e84c.NewHex2DecRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Hex2Oct()(*id719949da934cf20c3c44bdc0722b2534c7d0c91c0572433ba43b92dba4765b2.Hex2OctRequestBuilder) {
    return id719949da934cf20c3c44bdc0722b2534c7d0c91c0572433ba43b92dba4765b2.NewHex2OctRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Hlookup()(*ia480020e4a6c2e08476997d30d7775a6b23fd3fef49c44bfe41efb91bfc5482e.HlookupRequestBuilder) {
    return ia480020e4a6c2e08476997d30d7775a6b23fd3fef49c44bfe41efb91bfc5482e.NewHlookupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Hour()(*i39f4cd87e36f2f3a0b33d803ee7df3c500c56510f3a7e138e50274650d2d65de.HourRequestBuilder) {
    return i39f4cd87e36f2f3a0b33d803ee7df3c500c56510f3a7e138e50274650d2d65de.NewHourRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Hyperlink()(*i2b383fba5c20322d304f2d869487c05e2693a7ce95255893aeb48eed972a2da9.HyperlinkRequestBuilder) {
    return i2b383fba5c20322d304f2d869487c05e2693a7ce95255893aeb48eed972a2da9.NewHyperlinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) HypGeom_Dist()(*i5071f5260556cf090013e5064eda236cca516fed4d302085452c10f4c14ba7fc.HypGeom_DistRequestBuilder) {
    return i5071f5260556cf090013e5064eda236cca516fed4d302085452c10f4c14ba7fc.NewHypGeom_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) If_escaped()(*i3aa7f5e4a9ad281863916a5dae9bb0fb53bb4941b8b2664332317740e6a4c15b.IfRequestBuilder) {
    return i3aa7f5e4a9ad281863916a5dae9bb0fb53bb4941b8b2664332317740e6a4c15b.NewIfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImAbs()(*i6014adf79a8085e3e438be1d03d71d6a94be7c73a6f54d011025b6869298a304.ImAbsRequestBuilder) {
    return i6014adf79a8085e3e438be1d03d71d6a94be7c73a6f54d011025b6869298a304.NewImAbsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Imaginary()(*i1fda0f7b146838f09b6022fc87d7d315f1ad404bbf5938aefc425a5484805cbd.ImaginaryRequestBuilder) {
    return i1fda0f7b146838f09b6022fc87d7d315f1ad404bbf5938aefc425a5484805cbd.NewImaginaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImArgument()(*if6db57cb9f25d39cd5d2bd50eb111ef58f546d0d2d9da25b255d33ab5c29ea31.ImArgumentRequestBuilder) {
    return if6db57cb9f25d39cd5d2bd50eb111ef58f546d0d2d9da25b255d33ab5c29ea31.NewImArgumentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImConjugate()(*i7452bb10748075477cba95fcbc0c13c2d89ffdcbff9f76e0d57843f166d47c09.ImConjugateRequestBuilder) {
    return i7452bb10748075477cba95fcbc0c13c2d89ffdcbff9f76e0d57843f166d47c09.NewImConjugateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImCos()(*i350fc4ebf814e7bfb9e3c38975155c78c6ea182f5af9de4f0e3d4f6a15b372d2.ImCosRequestBuilder) {
    return i350fc4ebf814e7bfb9e3c38975155c78c6ea182f5af9de4f0e3d4f6a15b372d2.NewImCosRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImCosh()(*i3b6f6c7ed33e73600236ac505d6d8ddfd90f8b9d72a97c577cf41f7f5bd9bad5.ImCoshRequestBuilder) {
    return i3b6f6c7ed33e73600236ac505d6d8ddfd90f8b9d72a97c577cf41f7f5bd9bad5.NewImCoshRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImCot()(*i0e120005f13ea8b2a2b93fdfd8a3730a0fa0fa7ac5360fc7b2c927444f41e1b7.ImCotRequestBuilder) {
    return i0e120005f13ea8b2a2b93fdfd8a3730a0fa0fa7ac5360fc7b2c927444f41e1b7.NewImCotRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImCsc()(*i93f8f23d8d5723f4c27ac12031f23090341ad1da31ec8769d014d1dde1ec5739.ImCscRequestBuilder) {
    return i93f8f23d8d5723f4c27ac12031f23090341ad1da31ec8769d014d1dde1ec5739.NewImCscRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImCsch()(*i097b807941c5e18294aa5cf6621fcd0504002071465f2bf7987ff59de2aca92c.ImCschRequestBuilder) {
    return i097b807941c5e18294aa5cf6621fcd0504002071465f2bf7987ff59de2aca92c.NewImCschRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImDiv()(*i109ae362f2ade3a8826d4d477e2813f37a461350bf64445b7c53a3250ad89425.ImDivRequestBuilder) {
    return i109ae362f2ade3a8826d4d477e2813f37a461350bf64445b7c53a3250ad89425.NewImDivRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImExp()(*i36b8593f4e7696b829be2b3641edeb4aa7090a5297244e7329a8d9ebabd11c98.ImExpRequestBuilder) {
    return i36b8593f4e7696b829be2b3641edeb4aa7090a5297244e7329a8d9ebabd11c98.NewImExpRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImLn()(*ia0c51d44b14541036535a1c7d7dab0e5e713dc165de517410a1f12c4f862c574.ImLnRequestBuilder) {
    return ia0c51d44b14541036535a1c7d7dab0e5e713dc165de517410a1f12c4f862c574.NewImLnRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImLog10()(*i2ae597988fb14a89eb97e83a4090774fb8bddfb370f4b82c600a1dfdbd31d5bf.ImLog10RequestBuilder) {
    return i2ae597988fb14a89eb97e83a4090774fb8bddfb370f4b82c600a1dfdbd31d5bf.NewImLog10RequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImLog2()(*icd7b760eebee4cd326831f82d28e124ec9d530ad0cf6bca53e53ce3ff1dbb1af.ImLog2RequestBuilder) {
    return icd7b760eebee4cd326831f82d28e124ec9d530ad0cf6bca53e53ce3ff1dbb1af.NewImLog2RequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImPower()(*i0266785f33f4cc8c220e944b21b01c5947e87c01bcb41c020a5eddb02a9ed257.ImPowerRequestBuilder) {
    return i0266785f33f4cc8c220e944b21b01c5947e87c01bcb41c020a5eddb02a9ed257.NewImPowerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImProduct()(*iaf8e3d23027fdaf5b8ca992293bed72048279fb31fa871cff3a2f3f1a0da7144.ImProductRequestBuilder) {
    return iaf8e3d23027fdaf5b8ca992293bed72048279fb31fa871cff3a2f3f1a0da7144.NewImProductRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImReal()(*iacf0ea4e9ca20f4848adbfde01a685454f860c4de3f8013934409c3cbb368e88.ImRealRequestBuilder) {
    return iacf0ea4e9ca20f4848adbfde01a685454f860c4de3f8013934409c3cbb368e88.NewImRealRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImSec()(*ic1f613dc5187698d819b0aa08d2ad012fde20b02f33d821c2fec65ef88ee79f6.ImSecRequestBuilder) {
    return ic1f613dc5187698d819b0aa08d2ad012fde20b02f33d821c2fec65ef88ee79f6.NewImSecRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImSech()(*i29591fecc3455d858a2d496b08a44cceb3934ca10d1a97e3e964fecf0ad89f7f.ImSechRequestBuilder) {
    return i29591fecc3455d858a2d496b08a44cceb3934ca10d1a97e3e964fecf0ad89f7f.NewImSechRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImSin()(*i172d035bf5dab20ae0f3d85880acb671812d51c641eb99f52fdc34da268cf0fe.ImSinRequestBuilder) {
    return i172d035bf5dab20ae0f3d85880acb671812d51c641eb99f52fdc34da268cf0fe.NewImSinRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImSinh()(*ie2bdf3ef3c78f4f44af114a402c97bbb77df98d9cddab9ba7301caf4fff6fb66.ImSinhRequestBuilder) {
    return ie2bdf3ef3c78f4f44af114a402c97bbb77df98d9cddab9ba7301caf4fff6fb66.NewImSinhRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImSqrt()(*iddcfbe2f11f04c4acb3d6005a2592baa86f0985fc04742f4eb3a30ff2b5f82e0.ImSqrtRequestBuilder) {
    return iddcfbe2f11f04c4acb3d6005a2592baa86f0985fc04742f4eb3a30ff2b5f82e0.NewImSqrtRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImSub()(*i0461f30ce1b6ad808a745d4c850e3b7750615a40428226bf45dd420249fa3875.ImSubRequestBuilder) {
    return i0461f30ce1b6ad808a745d4c850e3b7750615a40428226bf45dd420249fa3875.NewImSubRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImSum()(*i492970e7012b0e3e57541d3e1fb52c45d9e86f9cf9d45f1f8e4d3b498fef0ec7.ImSumRequestBuilder) {
    return i492970e7012b0e3e57541d3e1fb52c45d9e86f9cf9d45f1f8e4d3b498fef0ec7.NewImSumRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImTan()(*i93a333ba632bc00534645305713d91f78a439c34f2de5bf13e758d28fac8103a.ImTanRequestBuilder) {
    return i93a333ba632bc00534645305713d91f78a439c34f2de5bf13e758d28fac8103a.NewImTanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Int()(*i60cd4be8a98f401f373ef391d0d720d59a0175eb2da787338a8c47ca72230756.IntRequestBuilder) {
    return i60cd4be8a98f401f373ef391d0d720d59a0175eb2da787338a8c47ca72230756.NewIntRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) IntRate()(*i124b15e94374881fc3ceb2d0b196066a5490324a97d0282f770abf019f5ebd72.IntRateRequestBuilder) {
    return i124b15e94374881fc3ceb2d0b196066a5490324a97d0282f770abf019f5ebd72.NewIntRateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Ipmt()(*ib6550e75fbfdd9e37216fc4794caa6c4e81f25f92c381a6ffa3d58fe2873c6f5.IpmtRequestBuilder) {
    return ib6550e75fbfdd9e37216fc4794caa6c4e81f25f92c381a6ffa3d58fe2873c6f5.NewIpmtRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Irr()(*i8f663f1f1cfc333d0c8519ce44a252b1dc321a2509021d5cea45983899f75c0f.IrrRequestBuilder) {
    return i8f663f1f1cfc333d0c8519ce44a252b1dc321a2509021d5cea45983899f75c0f.NewIrrRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) IsErr()(*i0dffa2d3e4f51c96e83725e18e43e2085197fe76a3f2cc370df8decb99554676.IsErrRequestBuilder) {
    return i0dffa2d3e4f51c96e83725e18e43e2085197fe76a3f2cc370df8decb99554676.NewIsErrRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) IsError()(*i1c1907b07e09c172dda3c49a4f313a14ca9651b2746e0bfeb1fbe37dbdd3e5c5.IsErrorRequestBuilder) {
    return i1c1907b07e09c172dda3c49a4f313a14ca9651b2746e0bfeb1fbe37dbdd3e5c5.NewIsErrorRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) IsEven()(*ifc77bdc03a4d73e8f9f6ecd5a111eaf0822b7692d8ac2b7bfec7261049d4dee9.IsEvenRequestBuilder) {
    return ifc77bdc03a4d73e8f9f6ecd5a111eaf0822b7692d8ac2b7bfec7261049d4dee9.NewIsEvenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) IsFormula()(*ifc654a64ea750853647dab0efb222301689322f2c0f6756df87f42142737aff3.IsFormulaRequestBuilder) {
    return ifc654a64ea750853647dab0efb222301689322f2c0f6756df87f42142737aff3.NewIsFormulaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) IsLogical()(*i6a5180477c05d488f65a77d747e4059176532e7fdb70500facf4d07dd8761a64.IsLogicalRequestBuilder) {
    return i6a5180477c05d488f65a77d747e4059176532e7fdb70500facf4d07dd8761a64.NewIsLogicalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) IsNA()(*idc8fd3f08e7ffc89e34663dc85de24b296c593e9aa28ecfd0c5c1fde908d1caa.IsNARequestBuilder) {
    return idc8fd3f08e7ffc89e34663dc85de24b296c593e9aa28ecfd0c5c1fde908d1caa.NewIsNARequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) IsNonText()(*i27090be239364b444a136b5826b6896a10d7ecf32ebda209d4c91af0f55455bf.IsNonTextRequestBuilder) {
    return i27090be239364b444a136b5826b6896a10d7ecf32ebda209d4c91af0f55455bf.NewIsNonTextRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) IsNumber()(*i1755304227922b0ac2038dc08b501d9ef710b978adb15d6211bfa3bd4baf16a2.IsNumberRequestBuilder) {
    return i1755304227922b0ac2038dc08b501d9ef710b978adb15d6211bfa3bd4baf16a2.NewIsNumberRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Iso_Ceiling()(*i59aff93824d19ad9146734be1350addc911717282bdfb6e5e50619b006989eaa.Iso_CeilingRequestBuilder) {
    return i59aff93824d19ad9146734be1350addc911717282bdfb6e5e50619b006989eaa.NewIso_CeilingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) IsOdd()(*ia4c89f7e9370bce5de71aa5a04bba1cb2dc0defa1923953b9aba644ff6115f29.IsOddRequestBuilder) {
    return ia4c89f7e9370bce5de71aa5a04bba1cb2dc0defa1923953b9aba644ff6115f29.NewIsOddRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) IsoWeekNum()(*i986d90630b8bd4704c82b2fdf90293b3ab0e120575e89da616037e92bf0059d1.IsoWeekNumRequestBuilder) {
    return i986d90630b8bd4704c82b2fdf90293b3ab0e120575e89da616037e92bf0059d1.NewIsoWeekNumRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Ispmt()(*i505aa17044c18b82f550f25288473c9ef4eb4704713abd4d25f343671d631204.IspmtRequestBuilder) {
    return i505aa17044c18b82f550f25288473c9ef4eb4704713abd4d25f343671d631204.NewIspmtRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Isref()(*i2479450dce1c1bcf5ee09162317898f4b82af362a855982508246d484edaaa8b.IsrefRequestBuilder) {
    return i2479450dce1c1bcf5ee09162317898f4b82af362a855982508246d484edaaa8b.NewIsrefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) IsText()(*i32e1f576fe43214f30cc8ffcd3f6853808eb7b877e9a328b9e8e2e647e6bf1c7.IsTextRequestBuilder) {
    return i32e1f576fe43214f30cc8ffcd3f6853808eb7b877e9a328b9e8e2e647e6bf1c7.NewIsTextRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Kurt()(*i22ae119a3184e61ad2aae889ae7224bab4ea903395898b3eb5756b8488676892.KurtRequestBuilder) {
    return i22ae119a3184e61ad2aae889ae7224bab4ea903395898b3eb5756b8488676892.NewKurtRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Large()(*icedf2711d1b11d1597a25ff1dceea21dc8397e613cff25b8ed7a4ff1cdfaf381.LargeRequestBuilder) {
    return icedf2711d1b11d1597a25ff1dceea21dc8397e613cff25b8ed7a4ff1cdfaf381.NewLargeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Lcm()(*i8a5aa2ca50c45a9aaa4a64a26c03d619e6ac6eb2ce39d9d7e0d0d75791d41d27.LcmRequestBuilder) {
    return i8a5aa2ca50c45a9aaa4a64a26c03d619e6ac6eb2ce39d9d7e0d0d75791d41d27.NewLcmRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Left()(*i4d961ff236da4359d12fbdfa7eee89fd03184cb0e49a2e2752677085a30472d5.LeftRequestBuilder) {
    return i4d961ff236da4359d12fbdfa7eee89fd03184cb0e49a2e2752677085a30472d5.NewLeftRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Leftb()(*i3767ab8e1ea6498037301289a27a8c90b420e3097e0ed91c5eef9ffc3dcfc2a8.LeftbRequestBuilder) {
    return i3767ab8e1ea6498037301289a27a8c90b420e3097e0ed91c5eef9ffc3dcfc2a8.NewLeftbRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Len()(*i4c5ae7091bd09004549bef0c9652ff1ae3fb0bf193fe9ded298576b924e200a1.LenRequestBuilder) {
    return i4c5ae7091bd09004549bef0c9652ff1ae3fb0bf193fe9ded298576b924e200a1.NewLenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Lenb()(*i35ba6be57981049c16486e05810aae9332e4f287c8a7fb56cba1de9585508ce4.LenbRequestBuilder) {
    return i35ba6be57981049c16486e05810aae9332e4f287c8a7fb56cba1de9585508ce4.NewLenbRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Ln()(*ib8859f36db353950f680ccddd5f4852712454947f41ab65661b9a34c8eeacb14.LnRequestBuilder) {
    return ib8859f36db353950f680ccddd5f4852712454947f41ab65661b9a34c8eeacb14.NewLnRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Log()(*i44c1d271b0ded00fc624c0e85f818ca370a3df578fdbd135e4c10a90d904bc6d.LogRequestBuilder) {
    return i44c1d271b0ded00fc624c0e85f818ca370a3df578fdbd135e4c10a90d904bc6d.NewLogRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Log10()(*i07ad831b4c905eef0f2b8147975bdc789d10f43b8261b423194792c65640c762.Log10RequestBuilder) {
    return i07ad831b4c905eef0f2b8147975bdc789d10f43b8261b423194792c65640c762.NewLog10RequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) LogNorm_Dist()(*ia8bfcef4032b0b7b3cfd0aa9712c6695c4c4baef57a47bff0df8c2d3e3cbe8de.LogNorm_DistRequestBuilder) {
    return ia8bfcef4032b0b7b3cfd0aa9712c6695c4c4baef57a47bff0df8c2d3e3cbe8de.NewLogNorm_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) LogNorm_Inv()(*i83ea1fbcbefdac4f93b0345b4d5d2a5e65ba70d393aa82950e87b9d5048a334e.LogNorm_InvRequestBuilder) {
    return i83ea1fbcbefdac4f93b0345b4d5d2a5e65ba70d393aa82950e87b9d5048a334e.NewLogNorm_InvRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Lookup()(*i64dcf16df745efb2a277f9da7811014fb9e46d69a69d3d2c0d98c40271d60c41.LookupRequestBuilder) {
    return i64dcf16df745efb2a277f9da7811014fb9e46d69a69d3d2c0d98c40271d60c41.NewLookupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Lower()(*i21b4bbaa6945c0f66867a02c17404afdb32643cba4e2e3a393bc38d19ded2c85.LowerRequestBuilder) {
    return i21b4bbaa6945c0f66867a02c17404afdb32643cba4e2e3a393bc38d19ded2c85.NewLowerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Match()(*ic252910cd61d1027a295d59661b3f0559e1293642f79081245cdb05f9599f7a4.MatchRequestBuilder) {
    return ic252910cd61d1027a295d59661b3f0559e1293642f79081245cdb05f9599f7a4.NewMatchRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Max()(*i99c3f166a423758c32dbc701aa5590ca9233e7d0fe1ac471c4373b792487f0c9.MaxRequestBuilder) {
    return i99c3f166a423758c32dbc701aa5590ca9233e7d0fe1ac471c4373b792487f0c9.NewMaxRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) MaxA()(*i22a50a05268af02ad93b3e099195bd0aa10043cf850f934baaa779347ff912fd.MaxARequestBuilder) {
    return i22a50a05268af02ad93b3e099195bd0aa10043cf850f934baaa779347ff912fd.NewMaxARequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Mduration()(*ie4c7b2ccc0fefd3b2166097de5afc1341fbe3df95ae3077cfb857b85eee03a24.MdurationRequestBuilder) {
    return ie4c7b2ccc0fefd3b2166097de5afc1341fbe3df95ae3077cfb857b85eee03a24.NewMdurationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Median()(*id6dcff57797b11fba1fcdb52de657107a8c156f9ad8d483dc42b2c441152fd77.MedianRequestBuilder) {
    return id6dcff57797b11fba1fcdb52de657107a8c156f9ad8d483dc42b2c441152fd77.NewMedianRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Mid()(*i82621983f40ab0734ef85bfbb618ffaa7cad9f67cfb178f33096ebf6673889f7.MidRequestBuilder) {
    return i82621983f40ab0734ef85bfbb618ffaa7cad9f67cfb178f33096ebf6673889f7.NewMidRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Midb()(*ide699df9204c2fb4c05840413b3fef64453c8456166397ced94412afbcd7572c.MidbRequestBuilder) {
    return ide699df9204c2fb4c05840413b3fef64453c8456166397ced94412afbcd7572c.NewMidbRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Min()(*i9a3f732bd7241183f682437f2f058d6f652e9ba5cc72ff3e03857f0788befb86.MinRequestBuilder) {
    return i9a3f732bd7241183f682437f2f058d6f652e9ba5cc72ff3e03857f0788befb86.NewMinRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) MinA()(*i5a73a60fca271f2a6e4ea5782835bbeb6b9e46a1ff77cbe422456e88a3a280f4.MinARequestBuilder) {
    return i5a73a60fca271f2a6e4ea5782835bbeb6b9e46a1ff77cbe422456e88a3a280f4.NewMinARequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Minute()(*i91014446a234d7aee31201c0b75a75f41c56ac2352caaafbe59f81ced3df23eb.MinuteRequestBuilder) {
    return i91014446a234d7aee31201c0b75a75f41c56ac2352caaafbe59f81ced3df23eb.NewMinuteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Mirr()(*i8384057cf305481cf65d4ea1b5ad309efd1eb6cbf2e43f5d6937c31a076dc3c2.MirrRequestBuilder) {
    return i8384057cf305481cf65d4ea1b5ad309efd1eb6cbf2e43f5d6937c31a076dc3c2.NewMirrRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Mod()(*iaf8b5876f5af8a9f7a5ef8e7f8b3e0fc749df12943de859ec1580bb6cb234ea9.ModRequestBuilder) {
    return iaf8b5876f5af8a9f7a5ef8e7f8b3e0fc749df12943de859ec1580bb6cb234ea9.NewModRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Month()(*i9e7061d8c4b334646d3b9442f20b3d0f0139b2c32cf905f7ec5cf4e479c36327.MonthRequestBuilder) {
    return i9e7061d8c4b334646d3b9442f20b3d0f0139b2c32cf905f7ec5cf4e479c36327.NewMonthRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Mround()(*ie8c1476b60a276425ece63c534f4e3c1ee0409af598fc7be59a08023d52ccacd.MroundRequestBuilder) {
    return ie8c1476b60a276425ece63c534f4e3c1ee0409af598fc7be59a08023d52ccacd.NewMroundRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) MultiNomial()(*id0142a852a9c647fb622b73201977e57aa86e47ce78f59bbc8799e24ad96c296.MultiNomialRequestBuilder) {
    return id0142a852a9c647fb622b73201977e57aa86e47ce78f59bbc8799e24ad96c296.NewMultiNomialRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) N()(*ie5716615358214f3b3b2e47c740f34ff277d58709f634223426fe9032bf24b6f.NRequestBuilder) {
    return ie5716615358214f3b3b2e47c740f34ff277d58709f634223426fe9032bf24b6f.NewNRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Na()(*i2f691a5ab87df22eea441e2516133fe7160951a84ac21da408c2c6fc21889c24.NaRequestBuilder) {
    return i2f691a5ab87df22eea441e2516133fe7160951a84ac21da408c2c6fc21889c24.NewNaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) NegBinom_Dist()(*if85e354c04136d477faa4e0f4230b6aa3f937d7654bcc233236b241bf7b1be3a.NegBinom_DistRequestBuilder) {
    return if85e354c04136d477faa4e0f4230b6aa3f937d7654bcc233236b241bf7b1be3a.NewNegBinom_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) NetworkDays()(*i13a285c284ba2120d7f47ea78b56d599c53ac40b7bf0f52dea113e716d2acb01.NetworkDaysRequestBuilder) {
    return i13a285c284ba2120d7f47ea78b56d599c53ac40b7bf0f52dea113e716d2acb01.NewNetworkDaysRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) NetworkDays_Intl()(*ifba3e217442e7b422f933a4732577709c619487368acff3f8aaa56e855079820.NetworkDays_IntlRequestBuilder) {
    return ifba3e217442e7b422f933a4732577709c619487368acff3f8aaa56e855079820.NewNetworkDays_IntlRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Nominal()(*i80c0df997d334f32be5b33abf7bb1214ab8f06300beaacc75d9acbd254bd7fed.NominalRequestBuilder) {
    return i80c0df997d334f32be5b33abf7bb1214ab8f06300beaacc75d9acbd254bd7fed.NewNominalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Norm_Dist()(*ie9f3949526768b26584d54e2031840cce1419e5425128f5260fc078d3eacfe3d.Norm_DistRequestBuilder) {
    return ie9f3949526768b26584d54e2031840cce1419e5425128f5260fc078d3eacfe3d.NewNorm_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Norm_Inv()(*if2cc99747d2a61f6ea9c0ec7aee2987e560c1e63342dd98f49246f082274726c.Norm_InvRequestBuilder) {
    return if2cc99747d2a61f6ea9c0ec7aee2987e560c1e63342dd98f49246f082274726c.NewNorm_InvRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Norm_S_Dist()(*i3631de4628d7bc943b5b2b5b9a5ae5b6f6b176596f019b2abe9c2cdba1aedea7.Norm_S_DistRequestBuilder) {
    return i3631de4628d7bc943b5b2b5b9a5ae5b6f6b176596f019b2abe9c2cdba1aedea7.NewNorm_S_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Norm_S_Inv()(*iefb8b5dc5714d55e2e4537e8bdb401e5c22894cefd038573046698e064963b32.Norm_S_InvRequestBuilder) {
    return iefb8b5dc5714d55e2e4537e8bdb401e5c22894cefd038573046698e064963b32.NewNorm_S_InvRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Not()(*i56c64d431ad2709306c509a9f9f0ec8e4d894d5c4bff5eeefdd4cbda7a7b7587.NotRequestBuilder) {
    return i56c64d431ad2709306c509a9f9f0ec8e4d894d5c4bff5eeefdd4cbda7a7b7587.NewNotRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Now()(*i1f4d3e0102ee9cebf20bdce62e14eec64e35b6cf5afbfae34c3eaad8d10a88e9.NowRequestBuilder) {
    return i1f4d3e0102ee9cebf20bdce62e14eec64e35b6cf5afbfae34c3eaad8d10a88e9.NewNowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Nper()(*i1a95ed2d2c1358a8e5a551f86342c22d27ead7fd9fea23c3e0f974d95289704c.NperRequestBuilder) {
    return i1a95ed2d2c1358a8e5a551f86342c22d27ead7fd9fea23c3e0f974d95289704c.NewNperRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Npv()(*i62f9d809bb32318e6fa706fa90f69b0be2c4f5ce7ab88f2190db66f8d8b74fe4.NpvRequestBuilder) {
    return i62f9d809bb32318e6fa706fa90f69b0be2c4f5ce7ab88f2190db66f8d8b74fe4.NewNpvRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) NumberValue()(*ie5f4e649a2430cb9ba32512c7cd00b5175d24789f7845a39081aa6e917d6ad49.NumberValueRequestBuilder) {
    return ie5f4e649a2430cb9ba32512c7cd00b5175d24789f7845a39081aa6e917d6ad49.NewNumberValueRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Oct2Bin()(*i851e9134195868314efcca697670811addf31221342640ba1a490891597a3664.Oct2BinRequestBuilder) {
    return i851e9134195868314efcca697670811addf31221342640ba1a490891597a3664.NewOct2BinRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Oct2Dec()(*if314cc65ab505afff71197b6b58adf355ac2583ec88039f4833d4245565cdb45.Oct2DecRequestBuilder) {
    return if314cc65ab505afff71197b6b58adf355ac2583ec88039f4833d4245565cdb45.NewOct2DecRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Oct2Hex()(*id0ab7dda529fcb6b25781c1d2cd57da3924ae8281797dc28cc48d2c8ba39a2b7.Oct2HexRequestBuilder) {
    return id0ab7dda529fcb6b25781c1d2cd57da3924ae8281797dc28cc48d2c8ba39a2b7.NewOct2HexRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Odd()(*iaf300ff360e1f70c118ab1245a8805547261654f42fcf5e927329cef9c94b60e.OddRequestBuilder) {
    return iaf300ff360e1f70c118ab1245a8805547261654f42fcf5e927329cef9c94b60e.NewOddRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) OddFPrice()(*i5b1782155c5c405bee8b8f5a33cba1abcf8e62c7784228ab4eb0aafed5551b0c.OddFPriceRequestBuilder) {
    return i5b1782155c5c405bee8b8f5a33cba1abcf8e62c7784228ab4eb0aafed5551b0c.NewOddFPriceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) OddFYield()(*i34c9271a0e30a434316e7cec8acb85de40a232501dd92511efc8c52ab3e0a708.OddFYieldRequestBuilder) {
    return i34c9271a0e30a434316e7cec8acb85de40a232501dd92511efc8c52ab3e0a708.NewOddFYieldRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) OddLPrice()(*i20c624b914d32116e4852ae299f9489ecbb99fd83655afa4d29dbd67b41372ff.OddLPriceRequestBuilder) {
    return i20c624b914d32116e4852ae299f9489ecbb99fd83655afa4d29dbd67b41372ff.NewOddLPriceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) OddLYield()(*ia4d363d4f1ca7b24dcccc1ef2cd633f884a8d61f1acf1699f6d8a72ebf14737b.OddLYieldRequestBuilder) {
    return ia4d363d4f1ca7b24dcccc1ef2cd633f884a8d61f1acf1699f6d8a72ebf14737b.NewOddLYieldRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Or()(*i9aa03510a433102dd03798716d94e40befa72b313c5eea5043a03f9c60bcf6fc.OrRequestBuilder) {
    return i9aa03510a433102dd03798716d94e40befa72b313c5eea5043a03f9c60bcf6fc.NewOrRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property functions in workbooks
func (m *FunctionsRequestBuilder) Patch(options *FunctionsRequestBuilderPatchOptions)(error) {
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
func (m *FunctionsRequestBuilder) Pduration()(*i83ff9d1937f074321dfcd778ed52ef8deac4669fd191114eb80bf60660c99517.PdurationRequestBuilder) {
    return i83ff9d1937f074321dfcd778ed52ef8deac4669fd191114eb80bf60660c99517.NewPdurationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Percentile_Exc()(*i4a4e247f672df0066887e62056a41a014663030a786167984fe09b18a220f24f.Percentile_ExcRequestBuilder) {
    return i4a4e247f672df0066887e62056a41a014663030a786167984fe09b18a220f24f.NewPercentile_ExcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Percentile_Inc()(*ibbb26bea14a287ff49e53453caa33dbd62106e41fcc9daafdd036da1c9ae6549.Percentile_IncRequestBuilder) {
    return ibbb26bea14a287ff49e53453caa33dbd62106e41fcc9daafdd036da1c9ae6549.NewPercentile_IncRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) PercentRank_Exc()(*i448ab36b74b378173a5767a9d3dc2209b2ec52bd89af54860a8c353a1249eabb.PercentRank_ExcRequestBuilder) {
    return i448ab36b74b378173a5767a9d3dc2209b2ec52bd89af54860a8c353a1249eabb.NewPercentRank_ExcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) PercentRank_Inc()(*icaef813c1a68be08058cf4c8982fec7f193c7f4ce4d210432c546739b2e837da.PercentRank_IncRequestBuilder) {
    return icaef813c1a68be08058cf4c8982fec7f193c7f4ce4d210432c546739b2e837da.NewPercentRank_IncRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Permut()(*idab688590eba2c52ae8ce39e7c029e05705eaf3846f9f3661da25fe15b15a2e0.PermutRequestBuilder) {
    return idab688590eba2c52ae8ce39e7c029e05705eaf3846f9f3661da25fe15b15a2e0.NewPermutRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Permutationa()(*idfce0e11b43129d348eab5e2d598a5e343fe5d432d4225dc1aec75c1017df6e4.PermutationaRequestBuilder) {
    return idfce0e11b43129d348eab5e2d598a5e343fe5d432d4225dc1aec75c1017df6e4.NewPermutationaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Phi()(*i670b8d255cccdd20e535f3c9f26bfa95cb9e18d53062194370910cf773a878f3.PhiRequestBuilder) {
    return i670b8d255cccdd20e535f3c9f26bfa95cb9e18d53062194370910cf773a878f3.NewPhiRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Pi()(*i2736fbe644ba9c2b350316e1ff478ada2a4c58fc4a3d8aeefe7904ef5cf7f8d5.PiRequestBuilder) {
    return i2736fbe644ba9c2b350316e1ff478ada2a4c58fc4a3d8aeefe7904ef5cf7f8d5.NewPiRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Pmt()(*i1c7ecd5d9304af3cf2d1f18063bb065457b82f3762eef0998c8aa1365f22f800.PmtRequestBuilder) {
    return i1c7ecd5d9304af3cf2d1f18063bb065457b82f3762eef0998c8aa1365f22f800.NewPmtRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Poisson_Dist()(*i15f4c5ffb740886c3c9fa4ff008942d605f56f15a5a9ac16ddce2cbce6a6a174.Poisson_DistRequestBuilder) {
    return i15f4c5ffb740886c3c9fa4ff008942d605f56f15a5a9ac16ddce2cbce6a6a174.NewPoisson_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Power()(*i02bf3208bc8afd502281b1c8c45241c968d956cfa84b08a4a8c61ed099189b43.PowerRequestBuilder) {
    return i02bf3208bc8afd502281b1c8c45241c968d956cfa84b08a4a8c61ed099189b43.NewPowerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Ppmt()(*i099b06710a4a149b79a3430055535b893ef0f2075c0f128716f0e326022844cc.PpmtRequestBuilder) {
    return i099b06710a4a149b79a3430055535b893ef0f2075c0f128716f0e326022844cc.NewPpmtRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Price()(*id5f48c79968fe4a26b921d06e11cdbc22451895ec74a40ff23bd2af6f9f6bfe4.PriceRequestBuilder) {
    return id5f48c79968fe4a26b921d06e11cdbc22451895ec74a40ff23bd2af6f9f6bfe4.NewPriceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) PriceDisc()(*i32abbd433bb4aa4005a887aab37f79f1409aa21e4ac36cb2eda8bc7d3fff6fe4.PriceDiscRequestBuilder) {
    return i32abbd433bb4aa4005a887aab37f79f1409aa21e4ac36cb2eda8bc7d3fff6fe4.NewPriceDiscRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) PriceMat()(*ied0c2d1a639b42db012bda18c0f9b2f7c4ebab2f9b1e3028b9209e5df07386ae.PriceMatRequestBuilder) {
    return ied0c2d1a639b42db012bda18c0f9b2f7c4ebab2f9b1e3028b9209e5df07386ae.NewPriceMatRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Product()(*iad8c455c383d91e2ddc21056c2cf1ad006674858b6e9d1b318fb40a4d1ead44a.ProductRequestBuilder) {
    return iad8c455c383d91e2ddc21056c2cf1ad006674858b6e9d1b318fb40a4d1ead44a.NewProductRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Proper()(*i75be6168523edf64b90022fdf9bec646c1bfc8ef1efa5fa3ac44b4afd2e3eb37.ProperRequestBuilder) {
    return i75be6168523edf64b90022fdf9bec646c1bfc8ef1efa5fa3ac44b4afd2e3eb37.NewProperRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Pv()(*i15bcb7ff9e078c7f041d0c54593e7b72a510b188ab533ff34d5acc37bbd04ed5.PvRequestBuilder) {
    return i15bcb7ff9e078c7f041d0c54593e7b72a510b188ab533ff34d5acc37bbd04ed5.NewPvRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Quartile_Exc()(*ia177a9f5c18691cc4406e66e13883d49fbdd9680bfb701c43311e417d1161835.Quartile_ExcRequestBuilder) {
    return ia177a9f5c18691cc4406e66e13883d49fbdd9680bfb701c43311e417d1161835.NewQuartile_ExcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Quartile_Inc()(*i559163cfcd0e3e255e4a9a51f12ccc72dfd4e1513c60909da25bbe744c7b7eed.Quartile_IncRequestBuilder) {
    return i559163cfcd0e3e255e4a9a51f12ccc72dfd4e1513c60909da25bbe744c7b7eed.NewQuartile_IncRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Quotient()(*i8e4cda1c8d4d879c336f3412bf3cf660300b4382767348a71dd7fb969b2a6192.QuotientRequestBuilder) {
    return i8e4cda1c8d4d879c336f3412bf3cf660300b4382767348a71dd7fb969b2a6192.NewQuotientRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Radians()(*if45416702cd1d18be7686d76f0e214785248198fdbc18dd4876ecf03112da042.RadiansRequestBuilder) {
    return if45416702cd1d18be7686d76f0e214785248198fdbc18dd4876ecf03112da042.NewRadiansRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Rand()(*i2766193728abd40f684dfaa8e25d37852148f793fdd189fd41e8836a8faf327b.RandRequestBuilder) {
    return i2766193728abd40f684dfaa8e25d37852148f793fdd189fd41e8836a8faf327b.NewRandRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) RandBetween()(*i0a886fcac7090403f01ae3a41ee46377ef9642eddf6fb911c9f8e55a4a6e35c7.RandBetweenRequestBuilder) {
    return i0a886fcac7090403f01ae3a41ee46377ef9642eddf6fb911c9f8e55a4a6e35c7.NewRandBetweenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Rank_Avg()(*if7e680f3499d68a27fe79595542c43bfb9cdcc74173f7ab77e839125313ae4a2.Rank_AvgRequestBuilder) {
    return if7e680f3499d68a27fe79595542c43bfb9cdcc74173f7ab77e839125313ae4a2.NewRank_AvgRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Rank_Eq()(*i9a47a986544d6a5f06b03d7cb7453c0870eb8874b2f226ff9499ae252d8e65e5.Rank_EqRequestBuilder) {
    return i9a47a986544d6a5f06b03d7cb7453c0870eb8874b2f226ff9499ae252d8e65e5.NewRank_EqRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Rate()(*if90a89d555d1426169b10dbb9ac2554cda2ce1f4f2f7839205cd2fcbfa03444a.RateRequestBuilder) {
    return if90a89d555d1426169b10dbb9ac2554cda2ce1f4f2f7839205cd2fcbfa03444a.NewRateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Received()(*iaa87c9aca478301cc253d3d7e5bfe18a8cd7a4b5b78b7a47fe04bd32cf9f8517.ReceivedRequestBuilder) {
    return iaa87c9aca478301cc253d3d7e5bfe18a8cd7a4b5b78b7a47fe04bd32cf9f8517.NewReceivedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Replace()(*i0ab188e1c48548c3af73eb596d61e636fc8b4d8a715e840444b9622e63740fbe.ReplaceRequestBuilder) {
    return i0ab188e1c48548c3af73eb596d61e636fc8b4d8a715e840444b9622e63740fbe.NewReplaceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ReplaceB()(*i55f82606d2062de63a09cfa317d19a38bc44cf9caca3f84be3b452b7888fc07f.ReplaceBRequestBuilder) {
    return i55f82606d2062de63a09cfa317d19a38bc44cf9caca3f84be3b452b7888fc07f.NewReplaceBRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Rept()(*i0650ae4116ca9eb8a4f165acdbbeb4eb9b2a87b6df15c716bcff863b121e82c5.ReptRequestBuilder) {
    return i0650ae4116ca9eb8a4f165acdbbeb4eb9b2a87b6df15c716bcff863b121e82c5.NewReptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Right()(*i049e9d1c53e9c7d3f4e9ec70418d7aec06590e4019424eb68b3e6c1dec18a845.RightRequestBuilder) {
    return i049e9d1c53e9c7d3f4e9ec70418d7aec06590e4019424eb68b3e6c1dec18a845.NewRightRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Rightb()(*iead39b340d7b27388c9912da5044632d1d33a0c34dc34ae5dcb1aa93456f1f70.RightbRequestBuilder) {
    return iead39b340d7b27388c9912da5044632d1d33a0c34dc34ae5dcb1aa93456f1f70.NewRightbRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Roman()(*if9ff5147ed19f76469c0737eaf1588b869564fa419d964e6acd6ddb95f1f65a3.RomanRequestBuilder) {
    return if9ff5147ed19f76469c0737eaf1588b869564fa419d964e6acd6ddb95f1f65a3.NewRomanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Round()(*ib316263ec23cf028a1c9c43f97a6c87006bd31b72eab6e609a69453c7b1e8249.RoundRequestBuilder) {
    return ib316263ec23cf028a1c9c43f97a6c87006bd31b72eab6e609a69453c7b1e8249.NewRoundRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) RoundDown()(*i6233f1d5ffabf3630193ff6919645c7ac85a1e1635e8136ae522a7bfd3ffb273.RoundDownRequestBuilder) {
    return i6233f1d5ffabf3630193ff6919645c7ac85a1e1635e8136ae522a7bfd3ffb273.NewRoundDownRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) RoundUp()(*ia08221458d3ce92ffa340c9aa72763f26ef30b9af826a69b3fb3297753f6d1e3.RoundUpRequestBuilder) {
    return ia08221458d3ce92ffa340c9aa72763f26ef30b9af826a69b3fb3297753f6d1e3.NewRoundUpRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Rows()(*id446936ca56486e000b19e23dbd525b966a48aee1300a88134721dd5cf9efd9e.RowsRequestBuilder) {
    return id446936ca56486e000b19e23dbd525b966a48aee1300a88134721dd5cf9efd9e.NewRowsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Rri()(*i5693d5ef4e4534c8934b5902523fea06b2534f831e4373120963f1d7186f3fc3.RriRequestBuilder) {
    return i5693d5ef4e4534c8934b5902523fea06b2534f831e4373120963f1d7186f3fc3.NewRriRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Sec()(*i5d2b89aed9f287e850042083336337a009be43b6571b38dd4a953ae60b0ecfd6.SecRequestBuilder) {
    return i5d2b89aed9f287e850042083336337a009be43b6571b38dd4a953ae60b0ecfd6.NewSecRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Sech()(*i0b9eecb7ccfe792400a9c77664400f2d2a8d08f97db4c334d61a360325a50d8d.SechRequestBuilder) {
    return i0b9eecb7ccfe792400a9c77664400f2d2a8d08f97db4c334d61a360325a50d8d.NewSechRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Second()(*id18fc1c9e53e811409aa175b1903d8a4231f5e1b21c1f6c08ae8bbd9b4fd2c3e.SecondRequestBuilder) {
    return id18fc1c9e53e811409aa175b1903d8a4231f5e1b21c1f6c08ae8bbd9b4fd2c3e.NewSecondRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) SeriesSum()(*ib8295a41b8411110d323251838d9fb6d72cec878915644fa9149be58e9f19c3e.SeriesSumRequestBuilder) {
    return ib8295a41b8411110d323251838d9fb6d72cec878915644fa9149be58e9f19c3e.NewSeriesSumRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Sheet()(*ifc0b96049e408483dcda28b9da9af8b3b9d0e3ab22420ca3ca0c69580846792f.SheetRequestBuilder) {
    return ifc0b96049e408483dcda28b9da9af8b3b9d0e3ab22420ca3ca0c69580846792f.NewSheetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Sheets()(*ie487f292c394041439201b1836fb377bbaf737623fb82bed4786600999e0a332.SheetsRequestBuilder) {
    return ie487f292c394041439201b1836fb377bbaf737623fb82bed4786600999e0a332.NewSheetsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Sign()(*i143e69103e0c49602fa219a4dd399818f412db7e7276d4263af8e6357a08e352.SignRequestBuilder) {
    return i143e69103e0c49602fa219a4dd399818f412db7e7276d4263af8e6357a08e352.NewSignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Sin()(*ibb028562b184f1114c572f1bfd4ecdd09c376808379f82fea18ab98530345d6a.SinRequestBuilder) {
    return ibb028562b184f1114c572f1bfd4ecdd09c376808379f82fea18ab98530345d6a.NewSinRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Sinh()(*ibd0b49ef1c39b06635846dd9dcb0e6f0a9011de6063d0aa2c4c2b1767d01c418.SinhRequestBuilder) {
    return ibd0b49ef1c39b06635846dd9dcb0e6f0a9011de6063d0aa2c4c2b1767d01c418.NewSinhRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Skew()(*i3a6c51172f28c572ffc43c7fd7e49866fd2b0c97b32fe94e0e28d58d09fc4276.SkewRequestBuilder) {
    return i3a6c51172f28c572ffc43c7fd7e49866fd2b0c97b32fe94e0e28d58d09fc4276.NewSkewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Skew_p()(*i3b46708bc24ffb2b82689186a5a6f750836ea124768d99b4a4fbc31efa01d9c5.Skew_pRequestBuilder) {
    return i3b46708bc24ffb2b82689186a5a6f750836ea124768d99b4a4fbc31efa01d9c5.NewSkew_pRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Sln()(*i566e7cb1af1639c06f9ebeba3ccb8a1fdc6c345eeefbfe6acac09d1582c12126.SlnRequestBuilder) {
    return i566e7cb1af1639c06f9ebeba3ccb8a1fdc6c345eeefbfe6acac09d1582c12126.NewSlnRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Small()(*i985e10a00be70c0d4c56e4933e1f13b0e6e96eafb096e52d812287d89fc04a26.SmallRequestBuilder) {
    return i985e10a00be70c0d4c56e4933e1f13b0e6e96eafb096e52d812287d89fc04a26.NewSmallRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Sqrt()(*ib8c96c5eea945469ccedea2b68135503234eeb7add02c3da84f01264bb80d249.SqrtRequestBuilder) {
    return ib8c96c5eea945469ccedea2b68135503234eeb7add02c3da84f01264bb80d249.NewSqrtRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) SqrtPi()(*i7b102aaeba7077cd75bd4b4829916459ef28397accd5982c90e6a3a60d059865.SqrtPiRequestBuilder) {
    return i7b102aaeba7077cd75bd4b4829916459ef28397accd5982c90e6a3a60d059865.NewSqrtPiRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Standardize()(*i33353f01662e530553a98ee8936d864fcf8de48750304038df77208337f06742.StandardizeRequestBuilder) {
    return i33353f01662e530553a98ee8936d864fcf8de48750304038df77208337f06742.NewStandardizeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) StDev_P()(*ia72c98a4751729a1e0614fc8ed05a45d4d39faa0ff7481aa4c31bae0ade72188.StDev_PRequestBuilder) {
    return ia72c98a4751729a1e0614fc8ed05a45d4d39faa0ff7481aa4c31bae0ade72188.NewStDev_PRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) StDev_S()(*i86ef0e159279f96f70079c0b4493ddddbd65f6f451676e153884ef85a549febf.StDev_SRequestBuilder) {
    return i86ef0e159279f96f70079c0b4493ddddbd65f6f451676e153884ef85a549febf.NewStDev_SRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) StDevA()(*i938a669433496fcdc36cb0d50f11ce775a2b7fbce462efe42da5f705832dcc26.StDevARequestBuilder) {
    return i938a669433496fcdc36cb0d50f11ce775a2b7fbce462efe42da5f705832dcc26.NewStDevARequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) StDevPA()(*i6635c9b9d1a681ece3749da33e90e966d20a0993e382e510bbd2789e2741ea27.StDevPARequestBuilder) {
    return i6635c9b9d1a681ece3749da33e90e966d20a0993e382e510bbd2789e2741ea27.NewStDevPARequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Substitute()(*ia4b182a5c215c546905ed2917110a5553ab0fd097232bd6b3a079abcd7d85723.SubstituteRequestBuilder) {
    return ia4b182a5c215c546905ed2917110a5553ab0fd097232bd6b3a079abcd7d85723.NewSubstituteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Subtotal()(*ie3a60eaa52f0e2c83224f957f6c5fd743ef6922115a77b7a30b636c2ddd0f617.SubtotalRequestBuilder) {
    return ie3a60eaa52f0e2c83224f957f6c5fd743ef6922115a77b7a30b636c2ddd0f617.NewSubtotalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Sum()(*if44efbdfc9b5c14c1fc3f5d0ddb7ed3bd98a90400a4e96b1d18314d4d7765659.SumRequestBuilder) {
    return if44efbdfc9b5c14c1fc3f5d0ddb7ed3bd98a90400a4e96b1d18314d4d7765659.NewSumRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) SumIf()(*i5dc621c21d96565fe810946c44544466dd6248d89e84a7626b434803d541c4b5.SumIfRequestBuilder) {
    return i5dc621c21d96565fe810946c44544466dd6248d89e84a7626b434803d541c4b5.NewSumIfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) SumIfs()(*i8ac88088cb1eeb8d34fde38032a24e268fd5ea5f3d5b1d1dd6920dd78c5f2f01.SumIfsRequestBuilder) {
    return i8ac88088cb1eeb8d34fde38032a24e268fd5ea5f3d5b1d1dd6920dd78c5f2f01.NewSumIfsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) SumSq()(*i3457ae9a71aebc3d34cffda9ff5c133b2776a91cefac7fc05350402850687859.SumSqRequestBuilder) {
    return i3457ae9a71aebc3d34cffda9ff5c133b2776a91cefac7fc05350402850687859.NewSumSqRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Syd()(*i43db37f22ff8e93c72e7ee1e05631ed0e75db3f46614da86fd2f4edc442ba8bb.SydRequestBuilder) {
    return i43db37f22ff8e93c72e7ee1e05631ed0e75db3f46614da86fd2f4edc442ba8bb.NewSydRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) T()(*i30a090b5650b04cc93000bec9f08421b6363c6d560327127a84916a3b09f0086.TRequestBuilder) {
    return i30a090b5650b04cc93000bec9f08421b6363c6d560327127a84916a3b09f0086.NewTRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) T_Dist()(*i7dfe45d82e3dda0f3316e137bedc1b70710e63e18987a0fb781c89837a0b5260.T_DistRequestBuilder) {
    return i7dfe45d82e3dda0f3316e137bedc1b70710e63e18987a0fb781c89837a0b5260.NewT_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) T_Dist_2T()(*i963416bc83cf905778043344d0bab1da3fca59fc6c8c051b4575a104cbdf28e3.T_Dist_2TRequestBuilder) {
    return i963416bc83cf905778043344d0bab1da3fca59fc6c8c051b4575a104cbdf28e3.NewT_Dist_2TRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) T_Dist_RT()(*ica5cc4cfc9127bfa086a9705314366a7ad715f4e0686b90fcaa0b3125318914e.T_Dist_RTRequestBuilder) {
    return ica5cc4cfc9127bfa086a9705314366a7ad715f4e0686b90fcaa0b3125318914e.NewT_Dist_RTRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) T_Inv()(*i27f60ca2023b9a4a5cc9d457ab7c1ebdefefd08ecd946d39b55f41e7c6bf329b.T_InvRequestBuilder) {
    return i27f60ca2023b9a4a5cc9d457ab7c1ebdefefd08ecd946d39b55f41e7c6bf329b.NewT_InvRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) T_Inv_2T()(*ic8b44e7573fd7c941987e253f26d7389a87bfca781b04b5bcbf2c27fb25f8616.T_Inv_2TRequestBuilder) {
    return ic8b44e7573fd7c941987e253f26d7389a87bfca781b04b5bcbf2c27fb25f8616.NewT_Inv_2TRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Tan()(*i4ace284db0c47bd0aede99134978f5c929609ce4e35b202bdc8f6500f9efbdbc.TanRequestBuilder) {
    return i4ace284db0c47bd0aede99134978f5c929609ce4e35b202bdc8f6500f9efbdbc.NewTanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Tanh()(*ib213aa115a5783efc6a93a4547d27e9aba61ae37bd2a46ec2535a3ed61aa0537.TanhRequestBuilder) {
    return ib213aa115a5783efc6a93a4547d27e9aba61ae37bd2a46ec2535a3ed61aa0537.NewTanhRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) TbillEq()(*i2cc67731e9c550a9cc653ccb9823cb407f72e05c8043783118760f0b7792187f.TbillEqRequestBuilder) {
    return i2cc67731e9c550a9cc653ccb9823cb407f72e05c8043783118760f0b7792187f.NewTbillEqRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) TbillPrice()(*i7afce34f3f2fd116a978ece18049d97b8250d2f8d24f54c2b7244b820740b61b.TbillPriceRequestBuilder) {
    return i7afce34f3f2fd116a978ece18049d97b8250d2f8d24f54c2b7244b820740b61b.NewTbillPriceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) TbillYield()(*i263f9c4268db82addce7e614d4c5702fa0452162ff16cd2a4b843da8d618b4e9.TbillYieldRequestBuilder) {
    return i263f9c4268db82addce7e614d4c5702fa0452162ff16cd2a4b843da8d618b4e9.NewTbillYieldRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Text()(*i38fc3cb72c39e10c123242af337410c10810bac285ded0e88d03eea4d7630c1c.TextRequestBuilder) {
    return i38fc3cb72c39e10c123242af337410c10810bac285ded0e88d03eea4d7630c1c.NewTextRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Time()(*iac60b2c91ba5e9073d79b2404aa84140abc5eb1fa5031e6fbf82ce642af5d310.TimeRequestBuilder) {
    return iac60b2c91ba5e9073d79b2404aa84140abc5eb1fa5031e6fbf82ce642af5d310.NewTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Timevalue()(*i869a657b2ae7ad505dee44529f6b7b72249331958fbba1421c6c7ce183f865db.TimevalueRequestBuilder) {
    return i869a657b2ae7ad505dee44529f6b7b72249331958fbba1421c6c7ce183f865db.NewTimevalueRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Today()(*i3f6bd1bd30b3d4eae2b33020bfff18913d339b66027e66bb84d38211e36bed13.TodayRequestBuilder) {
    return i3f6bd1bd30b3d4eae2b33020bfff18913d339b66027e66bb84d38211e36bed13.NewTodayRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Trim()(*ia2368438ab65c6ddbe0b26f6ba1296d19fcdb897869c0f0744ce722675daad33.TrimRequestBuilder) {
    return ia2368438ab65c6ddbe0b26f6ba1296d19fcdb897869c0f0744ce722675daad33.NewTrimRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) TrimMean()(*i68deb2e05608a96df2e0ca27f79d5923264ab2ba69b0b423931ee7b0a6582c6a.TrimMeanRequestBuilder) {
    return i68deb2e05608a96df2e0ca27f79d5923264ab2ba69b0b423931ee7b0a6582c6a.NewTrimMeanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) True()(*ifec202f6e2f8b4f71b4533ff606b32733c1244f51a6f1c2e6d45eabf4b17a8c5.TrueRequestBuilder) {
    return ifec202f6e2f8b4f71b4533ff606b32733c1244f51a6f1c2e6d45eabf4b17a8c5.NewTrueRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Trunc()(*i6bce0181254fb93d7930572c11d9dd1b75a83e5b3645600893fa1fff1d12d7b0.TruncRequestBuilder) {
    return i6bce0181254fb93d7930572c11d9dd1b75a83e5b3645600893fa1fff1d12d7b0.NewTruncRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Type_escaped()(*i12ce5b3d2a6128dd758ffd3fd51436cd30412f72f262fdbfc334abc06417e3f7.TypeRequestBuilder) {
    return i12ce5b3d2a6128dd758ffd3fd51436cd30412f72f262fdbfc334abc06417e3f7.NewTypeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Unichar()(*i88b361cc03b1e56e044d4445f4d89a0ecb754efe4ea824b27420646e43020419.UnicharRequestBuilder) {
    return i88b361cc03b1e56e044d4445f4d89a0ecb754efe4ea824b27420646e43020419.NewUnicharRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Unicode()(*i9d7b75a1b7e470f87e877b53d428a3510edf27826fd7887599769cfb74e9254d.UnicodeRequestBuilder) {
    return i9d7b75a1b7e470f87e877b53d428a3510edf27826fd7887599769cfb74e9254d.NewUnicodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Upper()(*i82fdfa05294be1ba769b77abf428e73287bc5a10a4c911c9625b2030af94e9e2.UpperRequestBuilder) {
    return i82fdfa05294be1ba769b77abf428e73287bc5a10a4c911c9625b2030af94e9e2.NewUpperRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Usdollar()(*ie3218b4af599d768c57d5aa1130d1ec1c6779febe7e16e6499a9456d46c026a4.UsdollarRequestBuilder) {
    return ie3218b4af599d768c57d5aa1130d1ec1c6779febe7e16e6499a9456d46c026a4.NewUsdollarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Value()(*i8451f7c3341a403be03d74fa7ff4bce0b6cff9a7985aa2bcc16a5eb72d4eb911.ValueRequestBuilder) {
    return i8451f7c3341a403be03d74fa7ff4bce0b6cff9a7985aa2bcc16a5eb72d4eb911.NewValueRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Var_P()(*i83006d85d2ec9e9c821e0df31373485e9bd4b920829e4a4980282767b0a72eb0.Var_PRequestBuilder) {
    return i83006d85d2ec9e9c821e0df31373485e9bd4b920829e4a4980282767b0a72eb0.NewVar_PRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Var_S()(*i9a7860e17c1a21c709fe83d415604d69e827b6a98f94305171b5ec3b01f769bd.Var_SRequestBuilder) {
    return i9a7860e17c1a21c709fe83d415604d69e827b6a98f94305171b5ec3b01f769bd.NewVar_SRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) VarA()(*ia08e5cb104669018ca6e491458a452542cef1b63353541697e68648e7b1031e8.VarARequestBuilder) {
    return ia08e5cb104669018ca6e491458a452542cef1b63353541697e68648e7b1031e8.NewVarARequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) VarPA()(*i6b66ab4ec1d212e8495dd66f10b6a82a906af5bdd344e286a87d49fdb988d600.VarPARequestBuilder) {
    return i6b66ab4ec1d212e8495dd66f10b6a82a906af5bdd344e286a87d49fdb988d600.NewVarPARequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Vdb()(*i20ffdeea2d03c7fc303f6b9e21d5f3c19cea0ae592ba935bfab22486e02a484d.VdbRequestBuilder) {
    return i20ffdeea2d03c7fc303f6b9e21d5f3c19cea0ae592ba935bfab22486e02a484d.NewVdbRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Vlookup()(*i297b0fb639f80c99fd017e82ff46d54acf62e7b366a48a1d74ec58773b7af6ff.VlookupRequestBuilder) {
    return i297b0fb639f80c99fd017e82ff46d54acf62e7b366a48a1d74ec58773b7af6ff.NewVlookupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Weekday()(*i1a087e75f3b9161d8ce9481ff6539f40aa80c4d8fb841cc20f5841f312630609.WeekdayRequestBuilder) {
    return i1a087e75f3b9161d8ce9481ff6539f40aa80c4d8fb841cc20f5841f312630609.NewWeekdayRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) WeekNum()(*i3bfcffb921445a5837441582c023df3ed7a99ae389fbed7fda482943e3fd5121.WeekNumRequestBuilder) {
    return i3bfcffb921445a5837441582c023df3ed7a99ae389fbed7fda482943e3fd5121.NewWeekNumRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Weibull_Dist()(*ieb5426682c782d2a15d99495fa6ea6438abe7e1e528e9fccdcc79bdf4e31105b.Weibull_DistRequestBuilder) {
    return ieb5426682c782d2a15d99495fa6ea6438abe7e1e528e9fccdcc79bdf4e31105b.NewWeibull_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) WorkDay()(*if8e737917131ca982223412916f7e20cc3b2fdc7e0c7fa5cac8c2dbcf066536a.WorkDayRequestBuilder) {
    return if8e737917131ca982223412916f7e20cc3b2fdc7e0c7fa5cac8c2dbcf066536a.NewWorkDayRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) WorkDay_Intl()(*ice2b692e5c82f169b0bfa2f37f0215a2eb1020ee060f49a315caa4420c43b6ce.WorkDay_IntlRequestBuilder) {
    return ice2b692e5c82f169b0bfa2f37f0215a2eb1020ee060f49a315caa4420c43b6ce.NewWorkDay_IntlRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Xirr()(*i9a8c3c1caa6ef8ce9239ac05447a657a5c1df8e3fc86936c6a80fe3fde5af182.XirrRequestBuilder) {
    return i9a8c3c1caa6ef8ce9239ac05447a657a5c1df8e3fc86936c6a80fe3fde5af182.NewXirrRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Xnpv()(*i4a6097157b72180db7ddd2c7f18b5bcf5fcb552868620e4eb3e0cabb9c634aa6.XnpvRequestBuilder) {
    return i4a6097157b72180db7ddd2c7f18b5bcf5fcb552868620e4eb3e0cabb9c634aa6.NewXnpvRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Xor()(*i1b3a3082854c5f5e5e7798d70e3ece3df8ecaaf43f8f08c2b5e4cd6cfc7516b0.XorRequestBuilder) {
    return i1b3a3082854c5f5e5e7798d70e3ece3df8ecaaf43f8f08c2b5e4cd6cfc7516b0.NewXorRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Year()(*i31051972fdf67777e1c1e16a89b980d25bb524eee6cd715e3eccfe9e99387cad.YearRequestBuilder) {
    return i31051972fdf67777e1c1e16a89b980d25bb524eee6cd715e3eccfe9e99387cad.NewYearRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) YearFrac()(*if6ca53a02d3f7eb4a0ef9b12ddc8dde0d6b2cc96ae1a41b91c028d7daf9ead56.YearFracRequestBuilder) {
    return if6ca53a02d3f7eb4a0ef9b12ddc8dde0d6b2cc96ae1a41b91c028d7daf9ead56.NewYearFracRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Yield()(*ic932478eb02695cb9c1766e2919fb21f692b10fd800981f7c8590cab8a333b56.YieldRequestBuilder) {
    return ic932478eb02695cb9c1766e2919fb21f692b10fd800981f7c8590cab8a333b56.NewYieldRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) YieldDisc()(*i18a2a6c3c741418f2ad549ddaa3c6adccf6ca9644ff9e7bf828710a10b7a6609.YieldDiscRequestBuilder) {
    return i18a2a6c3c741418f2ad549ddaa3c6adccf6ca9644ff9e7bf828710a10b7a6609.NewYieldDiscRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) YieldMat()(*ie650d4648a9f8150fff1faca1b39152f94ae4bef583ce99ac183fd9ff7f6f362.YieldMatRequestBuilder) {
    return ie650d4648a9f8150fff1faca1b39152f94ae4bef583ce99ac183fd9ff7f6f362.NewYieldMatRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Z_Test()(*i53dc72af898ff473ccd17dd44368585359f3b67d1ff8aa7e80aaf3102680c763.Z_TestRequestBuilder) {
    return i53dc72af898ff473ccd17dd44368585359f3b67d1ff8aa7e80aaf3102680c763.NewZ_TestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
