**Audit Date:** March 16, 2023

**Total Findings:** 399, **Total Gas Findings:** 300, **Total Low Findings:** 99

---

# Audit Report

## Notice: This report may contain some false positives, check all findings manually to verify them

## Files analyzed

- party-contracts-c4/contracts/crowdfund/AuctionCrowdfund.sol
- party-contracts-c4/contracts/crowdfund/BuyCrowdfund.sol
- party-contracts-c4/contracts/crowdfund/BuyCrowdfundBase.sol
- party-contracts-c4/contracts/crowdfund/CollectionBuyCrowdfund.sol
- party-contracts-c4/contracts/crowdfund/Crowdfund.sol
- party-contracts-c4/contracts/crowdfund/CrowdfundFactory.sol
- party-contracts-c4/contracts/crowdfund/CrowdfundNFT.sol
- party-contracts-c4/contracts/distribution/ITokenDistributor.sol
- party-contracts-c4/contracts/distribution/ITokenDistributorParty.sol
- party-contracts-c4/contracts/distribution/TokenDistributor.sol
- party-contracts-c4/contracts/gatekeepers/AllowListGateKeeper.sol
- party-contracts-c4/contracts/gatekeepers/IGateKeeper.sol
- party-contracts-c4/contracts/gatekeepers/TokenGateKeeper.sol
- party-contracts-c4/contracts/globals/Globals.sol
- party-contracts-c4/contracts/globals/IGlobals.sol
- party-contracts-c4/contracts/globals/LibGlobals.sol
- party-contracts-c4/contracts/market-wrapper/FoundationMarketWrapper.sol
- party-contracts-c4/contracts/market-wrapper/IMarketWrapper.sol
- party-contracts-c4/contracts/market-wrapper/KoansMarketWrapper.sol
- party-contracts-c4/contracts/market-wrapper/NounsMarketWrapper.sol
- party-contracts-c4/contracts/market-wrapper/ZoraMarketWrapper.sol
- party-contracts-c4/contracts/party/IPartyFactory.sol
- party-contracts-c4/contracts/party/Party.sol
- party-contracts-c4/contracts/party/PartyFactory.sol
- party-contracts-c4/contracts/party/PartyGovernance.sol
- party-contracts-c4/contracts/party/PartyGovernanceNFT.sol
- party-contracts-c4/contracts/proposals/ArbitraryCallsProposal.sol
- party-contracts-c4/contracts/proposals/FractionalizeProposal.sol
- party-contracts-c4/contracts/proposals/IProposalExecutionEngine.sol
- party-contracts-c4/contracts/proposals/LibProposal.sol
- party-contracts-c4/contracts/proposals/ListOnOpenseaProposal.sol
- party-contracts-c4/contracts/proposals/ListOnZoraProposal.sol
- party-contracts-c4/contracts/proposals/ProposalExecutionEngine.sol
- party-contracts-c4/contracts/proposals/ProposalStorage.sol
- party-contracts-c4/contracts/proposals/ZoraHelpers.sol
- party-contracts-c4/contracts/proposals/vendor/FractionalV1.sol
- party-contracts-c4/contracts/proposals/vendor/IOpenseaConduitController.sol
- party-contracts-c4/contracts/proposals/vendor/IOpenseaExchange.sol
- party-contracts-c4/contracts/renderers/CrowdfundNFTRenderer.sol
- party-contracts-c4/contracts/renderers/DummyERC721Renderer.sol
- party-contracts-c4/contracts/renderers/IERC721Renderer.sol
- party-contracts-c4/contracts/renderers/PartyGovernanceNFTRenderer.sol
- party-contracts-c4/contracts/tokens/ERC1155Receiver.sol
- party-contracts-c4/contracts/tokens/ERC721Receiver.sol
- party-contracts-c4/contracts/tokens/IERC1155.sol
- party-contracts-c4/contracts/tokens/IERC20.sol
- party-contracts-c4/contracts/tokens/IERC721.sol
- party-contracts-c4/contracts/tokens/IERC721Receiver.sol
- party-contracts-c4/contracts/utils/EIP165.sol
- party-contracts-c4/contracts/utils/Implementation.sol
- party-contracts-c4/contracts/utils/LibAddress.sol
- party-contracts-c4/contracts/utils/LibERC20Compat.sol
- party-contracts-c4/contracts/utils/LibRawResult.sol
- party-contracts-c4/contracts/utils/LibSafeCast.sol
- party-contracts-c4/contracts/utils/LibSafeERC721.sol
- party-contracts-c4/contracts/utils/PartyHelpers.sol
- party-contracts-c4/contracts/utils/Proxy.sol
- party-contracts-c4/contracts/utils/ReadOnlyDelegateCall.sol
- party-contracts-c4/contracts/utils/vendor/Base64.sol
- party-contracts-c4/contracts/utils/vendor/Strings.sol
- party-contracts-c4/contracts/vendor/markets/IFoundationMarket.sol
- party-contracts-c4/contracts/vendor/markets/IKoansAuctionHouse.sol
- party-contracts-c4/contracts/vendor/markets/INounsAuctionHouse.sol
- party-contracts-c4/contracts/vendor/markets/IZoraAuctionHouse.sol
- party-contracts-c4/contracts/vendor/solmate/ERC1155.sol
- party-contracts-c4/contracts/vendor/solmate/ERC20.sol
- party-contracts-c4/contracts/vendor/solmate/ERC721.sol

## **Findings Summary**

| Issue Identifier                         | Issue                                                                                        | Count |
| ---------------------------------------- | -------------------------------------------------------------------------------------------- | ----- |
| G001                                     | Don't Initialize Variables with Default Value                                                | 22    |
| G002                                     | Cache Array Length Outside of Loop **Works Well (no false positives)**                       | 16    |
| G003                                     | Use != 0 instead of > 0 for Unsigned Integer Comparison **Works Well (no false positives)**  | 2     |
| G004                                     | Use immutable for OpenZeppelin AccessControl's Roles Declarations                            | 2     |
| G006                                     | Use Shift Right/Left instead of Division/Multiplication if possible                          | 3     |
| G007                                     | ++i costs less gas than i++, especially when its used in for-loops (--i/i-- too)             | 34    |
| G008 **Works Well (no false positives)** | USE CUSTOM ERRORS RATHER THAN REVERT()/REQUIRE() STRINGS TO SAVE GAS                         | 48    |
| G009                                     | Address 0 check can be done in assembly to save gas                                          | 61    |
| G012                                     | Setting the constructor to payable                                                           | 27    |
| G014                                     | Assigning keccak operations to constant variables results in extra gas costs                 | 23    |
| G015 **Works Well (no false positives)** | x += y costs more gas than x = x + y for state variables                                     | 28    |
| G017                                     | ++i/i++ should be unchecked{++i}/unchecked{i++} when it is not possible for them to overflow | 34    |
| L001                                     | Unsafe ERC20 Operation(s)                                                                    | 6     |
| L002                                     | Unspecific Compiler Version Pragma                                                           | 67    |
| L007                                     | Use of bytes.concat() instead of abi.encodePacked()                                          | 26    |

---

## Findings

---

### **[G001] Don't Initialize Variables with Default Value**

---

#### **Impact**

github link to explanation

#### **Number of Findings:** 22

#### **Details:**

| File                                                              | Line Number | Line Content                                          | Code Location                                                                                                            |
| ----------------------------------------------------------------- | ----------- | ----------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------ |
| party-contracts-c4/contracts/crowdfund/Crowdfund.sol              | 180         | for (uint256 i = 0; i < contributors.length; ++i) {   | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/Crowdfund.sol#L180)             |
| party-contracts-c4/contracts/crowdfund/Crowdfund.sol              | 242         | for (uint256 i = 0; i < numContributions; ++i) {      | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/Crowdfund.sol#L242)             |
| party-contracts-c4/contracts/crowdfund/Crowdfund.sol              | 300         | for (uint256 i = 0; i < preciousTokens.length; ++i) { | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/Crowdfund.sol#L300)             |
| party-contracts-c4/contracts/crowdfund/Crowdfund.sol              | 348         | for (uint256 i = 0; i < numContributions; ++i) {      | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/Crowdfund.sol#L348)             |
| party-contracts-c4/contracts/distribution/TokenDistributor.sol    | 230         | for (uint256 i = 0; i < infos.length; ++i) {          | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/distribution/TokenDistributor.sol#L230)   |
| party-contracts-c4/contracts/distribution/TokenDistributor.sol    | 239         | for (uint256 i = 0; i < infos.length; ++i) {          | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/distribution/TokenDistributor.sol#L239)   |
| party-contracts-c4/contracts/party/PartyGovernance.sol            | 306         | for (uint256 i=0; i < opts.hosts.length; ++i) {       | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/party/PartyGovernance.sol#L306)           |
| party-contracts-c4/contracts/party/PartyGovernance.sol            | 432         | uint256 low = 0;                                      | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/party/PartyGovernance.sol#L432)           |
| party-contracts-c4/contracts/proposals/ArbitraryCallsProposal.sol | 52          | for (uint256 i = 0; i < hadPreciouses.length; ++i) {  | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ArbitraryCallsProposal.sol#L52) |
| party-contracts-c4/contracts/proposals/ArbitraryCallsProposal.sol | 61          | for (uint256 i = 0; i < calls.length; ++i) {          | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ArbitraryCallsProposal.sol#L61) |
| party-contracts-c4/contracts/proposals/ArbitraryCallsProposal.sol | 78          | for (uint256 i = 0; i < hadPreciouses.length; ++i) {  | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ArbitraryCallsProposal.sol#L78) |
| party-contracts-c4/contracts/proposals/LibProposal.sol            | 14          | for (uint256 i = 0; i < preciousTokens.length; ++i) { | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/LibProposal.sol#L14)            |
| party-contracts-c4/contracts/proposals/LibProposal.sol            | 32          | for (uint256 i = 0; i < preciousTokens.length; ++i) { | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/LibProposal.sol#L32)            |
| party-contracts-c4/contracts/proposals/ListOnOpenseaProposal.sol  | 291         | for (uint256 i = 0; i < fees.length; ++i) {           | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ListOnOpenseaProposal.sol#L291) |
| party-contracts-c4/contracts/utils/PartyHelpers.sol               | 64          | for (uint256 i = 0; i < members.length; i++) {        | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/utils/PartyHelpers.sol#L64)               |
| party-contracts-c4/contracts/utils/PartyHelpers.sol               | 85          | for (uint256 i = 0; i < voters.length; i++) {         | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/utils/PartyHelpers.sol#L85)               |
| party-contracts-c4/contracts/utils/PartyHelpers.sol               | 115         | for (uint256 i = 0; i < count; i++) {                 | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/utils/PartyHelpers.sol#L115)              |
| party-contracts-c4/contracts/utils/vendor/Strings.sol             | 42          | uint256 length = 0;                                   | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/utils/vendor/Strings.sol#L42)             |
| party-contracts-c4/contracts/vendor/solmate/ERC1155.sol           | 80          | for (uint256 i = 0; i < ids.length; ) {               | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC1155.sol#L80)           |
| party-contracts-c4/contracts/vendor/solmate/ERC1155.sol           | 118         | for (uint256 i = 0; i < owners.length; ++i) {         | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC1155.sol#L118)          |
| party-contracts-c4/contracts/vendor/solmate/ERC1155.sol           | 168         | for (uint256 i = 0; i < idsLength; ) {                | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC1155.sol#L168)          |
| party-contracts-c4/contracts/vendor/solmate/ERC1155.sol           | 198         | for (uint256 i = 0; i < idsLength; ) {                | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC1155.sol#L198)          |

#### **Tools used**

[c4udit](https://github.com/byterocket/c4udit)

---

### **[G002] Cache Array Length Outside of Loop **Works Well (no false positives)\*\*\*\*

---

#### **Impact**

By caching the length outside the loop, the function will only need to retrieve the value once and then
use the cached value within the loop. This optimization can result in significant gas savings for functions
that iterate over large arrays.

#### **Number of Findings:** 16

#### **Details:**

| File                                                              | Line Number | Line Content                                          | Code Location                                                                                                            |
| ----------------------------------------------------------------- | ----------- | ----------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------ |
| party-contracts-c4/contracts/crowdfund/CollectionBuyCrowdfund.sol | 62          | for (uint256 i; i < hosts.length; i++) {              | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/CollectionBuyCrowdfund.sol#L62) |
| party-contracts-c4/contracts/crowdfund/Crowdfund.sol              | 180         | for (uint256 i = 0; i < contributors.length; ++i) {   | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/Crowdfund.sol#L180)             |
| party-contracts-c4/contracts/crowdfund/Crowdfund.sol              | 300         | for (uint256 i = 0; i < preciousTokens.length; ++i) { | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/Crowdfund.sol#L300)             |
| party-contracts-c4/contracts/distribution/TokenDistributor.sol    | 230         | for (uint256 i = 0; i < infos.length; ++i) {          | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/distribution/TokenDistributor.sol#L230)   |
| party-contracts-c4/contracts/distribution/TokenDistributor.sol    | 239         | for (uint256 i = 0; i < infos.length; ++i) {          | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/distribution/TokenDistributor.sol#L239)   |
| party-contracts-c4/contracts/party/PartyGovernance.sol            | 306         | for (uint256 i=0; i < opts.hosts.length; ++i) {       | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/party/PartyGovernance.sol#L306)           |
| party-contracts-c4/contracts/proposals/ArbitraryCallsProposal.sol | 52          | for (uint256 i = 0; i < hadPreciouses.length; ++i) {  | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ArbitraryCallsProposal.sol#L52) |
| party-contracts-c4/contracts/proposals/ArbitraryCallsProposal.sol | 61          | for (uint256 i = 0; i < calls.length; ++i) {          | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ArbitraryCallsProposal.sol#L61) |
| party-contracts-c4/contracts/proposals/ArbitraryCallsProposal.sol | 78          | for (uint256 i = 0; i < hadPreciouses.length; ++i) {  | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ArbitraryCallsProposal.sol#L78) |
| party-contracts-c4/contracts/proposals/LibProposal.sol            | 14          | for (uint256 i = 0; i < preciousTokens.length; ++i) { | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/LibProposal.sol#L14)            |
| party-contracts-c4/contracts/proposals/LibProposal.sol            | 32          | for (uint256 i = 0; i < preciousTokens.length; ++i) { | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/LibProposal.sol#L32)            |
| party-contracts-c4/contracts/proposals/ListOnOpenseaProposal.sol  | 291         | for (uint256 i = 0; i < fees.length; ++i) {           | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ListOnOpenseaProposal.sol#L291) |
| party-contracts-c4/contracts/utils/PartyHelpers.sol               | 64          | for (uint256 i = 0; i < members.length; i++) {        | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/utils/PartyHelpers.sol#L64)               |
| party-contracts-c4/contracts/utils/PartyHelpers.sol               | 85          | for (uint256 i = 0; i < voters.length; i++) {         | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/utils/PartyHelpers.sol#L85)               |
| party-contracts-c4/contracts/vendor/solmate/ERC1155.sol           | 80          | for (uint256 i = 0; i < ids.length; ) {               | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC1155.sol#L80)           |
| party-contracts-c4/contracts/vendor/solmate/ERC1155.sol           | 118         | for (uint256 i = 0; i < owners.length; ++i) {         | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC1155.sol#L118)          |

#### **Tools used**

[c4udit](https://github.com/byterocket/c4udit)

---

### **[G003] Use != 0 instead of > 0 for Unsigned Integer Comparison **Works Well (no false positives)\*\*\*\*

---

#### **Impact**

When dealing with unsigned integer types, comparisons with != 0 are cheaper than with > 0.

#### **Number of Findings:** 2

#### **Details:**

| File                                                 | Line Number | Line Content              | Code Location                                                                                                |
| ---------------------------------------------------- | ----------- | ------------------------- | ------------------------------------------------------------------------------------------------------------ |
| party-contracts-c4/contracts/crowdfund/Crowdfund.sol | 144         | if (initialBalance > 0) { | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/Crowdfund.sol#L144) |
| party-contracts-c4/contracts/crowdfund/Crowdfund.sol | 471         | if (votingPower > 0) {    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/Crowdfund.sol#L471) |

#### **Tools used**

[c4udit](https://github.com/byterocket/c4udit)

---

### **[G004] Use immutable for OpenZeppelin AccessControl's Roles Declarations**

---

#### **Impact**

⚡️ Only valid for solidity versions <0.6.12 ⚡️
Access roles marked as constant results in computing the keccak256 operation each
time the variable is used because assigned operations for constant variables are re-evaluated every time.
Changing the variables to immutable results in computing the hash only once on deployment, leading to gas savings.

#### **Number of Findings:** 2

#### **Details:**

| File                                                               | Line Number | Line Content                                                                                                | Code Location                                                                                                             |
| ------------------------------------------------------------------ | ----------- | ----------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- |
| party-contracts-c4/contracts/proposals/ProposalExecutionEngine.sol | 80          | uint256 private constant \_STORAGE_SLOT = uint256(keccak256('ProposalExecutionEngine.Storage'));            | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ProposalExecutionEngine.sol#L80) |
| party-contracts-c4/contracts/proposals/ProposalStorage.sol         | 19          | uint256 private constant SHARED_STORAGE_SLOT = uint256(keccak256("ProposalStorage.SharedProposalStorage")); | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ProposalStorage.sol#L19)         |

#### **Tools used**

[c4udit](https://github.com/byterocket/c4udit)

---

### **[G006] Use Shift Right/Left instead of Division/Multiplication if possible**

---

#### **Impact**

<x> / 2is the same as <x> >> 1. While the compiler uses the SHRopcode to accomplish both,
the version that uses division incurs an overhead of 20 gas due to JUMPs to and from a compiler
utility function that introduces checks which can be avoided by using unchecked {}around the division by two.

#### **Number of Findings:** 3

#### **Details:**

| File                                                                  | Line Number | Line Content                                                                                                                                                                                                                                                                                                              | Code Location                                                                                                                 |
| --------------------------------------------------------------------- | ----------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------- |
| party-contracts-c4/contracts/party/PartyGovernance.sol                | 434         | uint256 mid = (low + high) / 2;                                                                                                                                                                                                                                                                                           | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/party/PartyGovernance.sol#L434)                |
| party-contracts-c4/contracts/renderers/CrowdfundNFTRenderer.sol       | 114         | svgParts[0] = '<svg xmlns="http://www.w3.org/2000/svg" preserveAspectRatio="xMinYMin meet" viewBox="0 0 350 350"><style>text { fill: white; font-family: -apple-system, BlinkMacSystemFont, sans-serif; } .base { font-size: 11px; } .detail {font-size: 10px;}</style><rect width="100%" height="100%" fill="black" />'; | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/renderers/CrowdfundNFTRenderer.sol#L114)       |
| party-contracts-c4/contracts/renderers/PartyGovernanceNFTRenderer.sol | 120         | svgParts[0] = '<svg xmlns="http://www.w3.org/2000/svg" preserveAspectRatio="xMinYMin meet" viewBox="0 0 350 350"><style>text { fill: white; font-family: -apple-system, BlinkMacSystemFont, sans-serif; } .base { font-size: 11px; } .detail {font-size: 10px;}</style><rect width="100%" height="100%" fill="black" />'; | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/renderers/PartyGovernanceNFTRenderer.sol#L120) |

#### **Tools used**

[c4udit](https://github.com/byterocket/c4udit)

---

### **[G007] ++i costs less gas than i++, especially when its used in for-loops (--i/i-- too)**

---

#### **Impact**

...

#### **Number of Findings:** 34

#### **Details:**

| File                                                              | Line Number | Line Content                                              | Code Location                                                                                                            |
| ----------------------------------------------------------------- | ----------- | --------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------ |
| party-contracts-c4/contracts/crowdfund/CollectionBuyCrowdfund.sol | 62          | for (uint256 i; i < hosts.length; i++) {                  | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/CollectionBuyCrowdfund.sol#L62) |
| party-contracts-c4/contracts/crowdfund/Crowdfund.sol              | 180         | for (uint256 i = 0; i < contributors.length; ++i) {       | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/Crowdfund.sol#L180)             |
| party-contracts-c4/contracts/crowdfund/Crowdfund.sol              | 242         | for (uint256 i = 0; i < numContributions; ++i) {          | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/Crowdfund.sol#L242)             |
| party-contracts-c4/contracts/crowdfund/Crowdfund.sol              | 300         | for (uint256 i = 0; i < preciousTokens.length; ++i) {     | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/Crowdfund.sol#L300)             |
| party-contracts-c4/contracts/crowdfund/Crowdfund.sol              | 348         | for (uint256 i = 0; i < numContributions; ++i) {          | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/Crowdfund.sol#L348)             |
| party-contracts-c4/contracts/distribution/TokenDistributor.sol    | 230         | for (uint256 i = 0; i < infos.length; ++i) {              | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/distribution/TokenDistributor.sol#L230)   |
| party-contracts-c4/contracts/distribution/TokenDistributor.sol    | 239         | for (uint256 i = 0; i < infos.length; ++i) {              | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/distribution/TokenDistributor.sol#L239)   |
| party-contracts-c4/contracts/distribution/TokenDistributor.sol    | 357         | distributionId: ++lastDistributionIdPerParty[args.party], | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/distribution/TokenDistributor.sol#L357)   |
| party-contracts-c4/contracts/gatekeepers/AllowListGateKeeper.sol  | 33          | uint96 id\_ = ++\_lastId;                                 | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/gatekeepers/AllowListGateKeeper.sol#L33)  |
| party-contracts-c4/contracts/gatekeepers/TokenGateKeeper.sol      | 48          | uint96 id\_ = ++\_lastId;                                 | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/gatekeepers/TokenGateKeeper.sol#L48)      |
| party-contracts-c4/contracts/party/PartyGovernance.sol            | 306         | for (uint256 i=0; i < opts.hosts.length; ++i) {           | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/party/PartyGovernance.sol#L306)           |
| party-contracts-c4/contracts/party/PartyGovernance.sol            | 532         | proposalId = ++lastProposalId;                            | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/party/PartyGovernance.sol#L532)           |
| party-contracts-c4/contracts/party/PartyGovernanceNFT.sol         | 129         | uint256 tokenId = ++tokenCount;                           | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/party/PartyGovernanceNFT.sol#L129)        |
| party-contracts-c4/contracts/proposals/ArbitraryCallsProposal.sol | 52          | for (uint256 i = 0; i < hadPreciouses.length; ++i) {      | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ArbitraryCallsProposal.sol#L52) |
| party-contracts-c4/contracts/proposals/ArbitraryCallsProposal.sol | 61          | for (uint256 i = 0; i < calls.length; ++i) {              | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ArbitraryCallsProposal.sol#L61) |
| party-contracts-c4/contracts/proposals/ArbitraryCallsProposal.sol | 78          | for (uint256 i = 0; i < hadPreciouses.length; ++i) {      | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ArbitraryCallsProposal.sol#L78) |
| party-contracts-c4/contracts/proposals/LibProposal.sol            | 14          | for (uint256 i = 0; i < preciousTokens.length; ++i) {     | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/LibProposal.sol#L14)            |
| party-contracts-c4/contracts/proposals/LibProposal.sol            | 32          | for (uint256 i = 0; i < preciousTokens.length; ++i) {     | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/LibProposal.sol#L32)            |
| party-contracts-c4/contracts/proposals/ListOnOpenseaProposal.sol  | 291         | for (uint256 i = 0; i < fees.length; ++i) {               | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ListOnOpenseaProposal.sol#L291) |
| party-contracts-c4/contracts/utils/PartyHelpers.sol               | 64          | for (uint256 i = 0; i < members.length; i++) {            | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/utils/PartyHelpers.sol#L64)               |
| party-contracts-c4/contracts/utils/PartyHelpers.sol               | 85          | for (uint256 i = 0; i < voters.length; i++) {             | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/utils/PartyHelpers.sol#L85)               |
| party-contracts-c4/contracts/utils/PartyHelpers.sol               | 115         | for (uint256 i = 0; i < count; i++) {                     | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/utils/PartyHelpers.sol#L115)              |
| party-contracts-c4/contracts/utils/vendor/Strings.sol             | 22          | digits++;                                                 | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/utils/vendor/Strings.sol#L22)             |
| party-contracts-c4/contracts/utils/vendor/Strings.sol             | 44          | length++;                                                 | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/utils/vendor/Strings.sol#L44)             |
| party-contracts-c4/contracts/utils/vendor/Strings.sol             | 57          | for (uint256 i = 2 \* length + 1; i > 1; --i) {           | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/utils/vendor/Strings.sol#L57)             |
| party-contracts-c4/contracts/vendor/solmate/ERC1155.sol           | 90          | ++i;                                                      | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC1155.sol#L90)           |
| party-contracts-c4/contracts/vendor/solmate/ERC1155.sol           | 118         | for (uint256 i = 0; i < owners.length; ++i) {             | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC1155.sol#L118)          |
| party-contracts-c4/contracts/vendor/solmate/ERC1155.sol           | 174         | ++i;                                                      | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC1155.sol#L174)          |
| party-contracts-c4/contracts/vendor/solmate/ERC1155.sol           | 204         | ++i;                                                      | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC1155.sol#L204)          |
| party-contracts-c4/contracts/vendor/solmate/ERC20.sol             | 142         | nonces[owner]++,                                          | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC20.sol#L142)            |
| party-contracts-c4/contracts/vendor/solmate/ERC721.sol            | 98          | \_balanceOf[from]--;                                      | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC721.sol#L98)            |
| party-contracts-c4/contracts/vendor/solmate/ERC721.sol            | 100         | \_balanceOf[to]++;                                        | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC721.sol#L100)           |
| party-contracts-c4/contracts/vendor/solmate/ERC721.sol            | 163         | \_balanceOf[to]++;                                        | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC721.sol#L163)           |
| party-contracts-c4/contracts/vendor/solmate/ERC721.sol            | 178         | \_balanceOf[owner]--;                                     | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC721.sol#L178)           |

#### **Tools used**

[c4udit](https://github.com/byterocket/c4udit)

---

### **[G008 **Works Well (no false positives)**] USE CUSTOM ERRORS RATHER THAN REVERT()/REQUIRE() STRINGS TO SAVE GAS**

---

#### **Impact**

Custom errors are available from solidity version 0.8.4. Custom errors save ~50 gas each time
they’re hit by avoiding having to allocate and store the revert string.
Not defining the strings also save deployment gas.

#### **Number of Findings:** 48

#### **Details:**

| File                                                                    | Line Number | Line Content                                                                            | Code Location                                                                                                                  |
| ----------------------------------------------------------------------- | ----------- | --------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| party-contracts-c4/contracts/crowdfund/CrowdfundNFT.sol                 | 166         | assert(false); // Will not be reached.                                                  | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/CrowdfundNFT.sol#L166)                |
| party-contracts-c4/contracts/distribution/TokenDistributor.sol          | 385         | assert(tokenType == TokenType.Erc20);                                                   | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/distribution/TokenDistributor.sol#L385)         |
| party-contracts-c4/contracts/distribution/TokenDistributor.sol          | 411         | assert(tokenType == TokenType.Erc20);                                                   | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/distribution/TokenDistributor.sol#L411)         |
| party-contracts-c4/contracts/market-wrapper/FoundationMarketWrapper.sol | 93          | require(success, string(returnData));                                                   | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/market-wrapper/FoundationMarketWrapper.sol#L93) |
| party-contracts-c4/contracts/market-wrapper/KoansMarketWrapper.sol      | 67          | require(                                                                                | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/market-wrapper/KoansMarketWrapper.sol#L67)      |
| party-contracts-c4/contracts/market-wrapper/KoansMarketWrapper.sol      | 92          | require(                                                                                | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/market-wrapper/KoansMarketWrapper.sol#L92)      |
| party-contracts-c4/contracts/market-wrapper/KoansMarketWrapper.sol      | 109         | require(success, string(returnData));                                                   | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/market-wrapper/KoansMarketWrapper.sol#L109)     |
| party-contracts-c4/contracts/market-wrapper/NounsMarketWrapper.sol      | 65          | require(                                                                                | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/market-wrapper/NounsMarketWrapper.sol#L65)      |
| party-contracts-c4/contracts/market-wrapper/NounsMarketWrapper.sol      | 90          | require(                                                                                | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/market-wrapper/NounsMarketWrapper.sol#L90)      |
| party-contracts-c4/contracts/market-wrapper/NounsMarketWrapper.sol      | 107         | require(success, string(returnData));                                                   | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/market-wrapper/NounsMarketWrapper.sol#L107)     |
| party-contracts-c4/contracts/market-wrapper/ZoraMarketWrapper.sol       | 113         | require(success, string(returnData));                                                   | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/market-wrapper/ZoraMarketWrapper.sol#L113)      |
| party-contracts-c4/contracts/party/PartyGovernance.sol                  | 504         | assert(tokenType == ITokenDistributor.TokenType.Erc20);                                 | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/party/PartyGovernance.sol#L504)                 |
| party-contracts-c4/contracts/party/PartyGovernanceNFT.sol               | 179         | assert(false); // Will not be reached.                                                  | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/party/PartyGovernanceNFT.sol#L179)              |
| party-contracts-c4/contracts/proposals/FractionalizeProposal.sol        | 67          | assert(vault.balanceOf(address(this)) == supply);                                       | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/FractionalizeProposal.sol#L67)        |
| party-contracts-c4/contracts/proposals/ListOnOpenseaProposal.sol        | 221         | assert(step == ListOnOpenseaStep.ListedOnOpenSea);                                      | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ListOnOpenseaProposal.sol#L221)       |
| party-contracts-c4/contracts/proposals/ListOnOpenseaProposal.sol        | 302         | assert(SEAPORT.validate(orders));                                                       | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ListOnOpenseaProposal.sol#L302)       |
| party-contracts-c4/contracts/proposals/ListOnZoraProposal.sol           | 120         | assert(step == ZoraStep.ListedOnZora);                                                  | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ListOnZoraProposal.sol#L120)          |
| party-contracts-c4/contracts/proposals/ProposalExecutionEngine.sol      | 142         | assert(currentInProgressProposalId == 0);                                               | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ProposalExecutionEngine.sol#L142)     |
| party-contracts-c4/contracts/proposals/ProposalExecutionEngine.sol      | 246         | require(proposalType != ProposalType.Invalid);                                          | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ProposalExecutionEngine.sol#L246)     |
| party-contracts-c4/contracts/proposals/ProposalExecutionEngine.sol      | 247         | require(uint8(proposalType) < uint8(ProposalType.NumProposalTypes));                    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ProposalExecutionEngine.sol#L247)     |
| party-contracts-c4/contracts/utils/ReadOnlyDelegateCall.sol             | 20          | require(msg.sender == address(this));                                                   | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/utils/ReadOnlyDelegateCall.sol#L20)             |
| party-contracts-c4/contracts/utils/ReadOnlyDelegateCall.sol             | 30          | assert(false);                                                                          | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/utils/ReadOnlyDelegateCall.sol#L30)             |
| party-contracts-c4/contracts/utils/vendor/Strings.sol                   | 61          | require(value == 0, "Strings: hex length insufficient");                                | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/utils/vendor/Strings.sol#L61)                   |
| party-contracts-c4/contracts/vendor/solmate/ERC1155.sol                 | 49          | require(msg.sender == from                                                              |                                                                                                                                | isApprovedForAll[from][msg.sender], "NOT_AUTHORIZED");  | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC1155.sol#L49) |
| party-contracts-c4/contracts/vendor/solmate/ERC1155.sol                 | 56          | require(                                                                                | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC1155.sol#L56)                 |
| party-contracts-c4/contracts/vendor/solmate/ERC1155.sol                 | 72          | require(ids.length == amounts.length, "LENGTH_MISMATCH");                               | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC1155.sol#L72)                 |
| party-contracts-c4/contracts/vendor/solmate/ERC1155.sol                 | 74          | require(msg.sender == from                                                              |                                                                                                                                | isApprovedForAll[from][msg.sender], "NOT_AUTHORIZED");  | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC1155.sol#L74) |
| party-contracts-c4/contracts/vendor/solmate/ERC1155.sol                 | 96          | require(                                                                                | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC1155.sol#L96)                 |
| party-contracts-c4/contracts/vendor/solmate/ERC1155.sol                 | 111         | require(owners.length == ids.length, "LENGTH_MISMATCH");                                | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC1155.sol#L111)                |
| party-contracts-c4/contracts/vendor/solmate/ERC1155.sol                 | 149         | require(                                                                                | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC1155.sol#L149)                |
| party-contracts-c4/contracts/vendor/solmate/ERC1155.sol                 | 166         | require(idsLength == amounts.length, "LENGTH_MISMATCH");                                | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC1155.sol#L166)                |
| party-contracts-c4/contracts/vendor/solmate/ERC1155.sol                 | 180         | require(                                                                                | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC1155.sol#L180)                |
| party-contracts-c4/contracts/vendor/solmate/ERC1155.sol                 | 196         | require(idsLength == amounts.length, "LENGTH_MISMATCH");                                | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC1155.sol#L196)                |
| party-contracts-c4/contracts/vendor/solmate/ERC20.sol                   | 124         | require(deadline >= block.timestamp, "PERMIT_DEADLINE_EXPIRED");                        | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC20.sol#L124)                  |
| party-contracts-c4/contracts/vendor/solmate/ERC20.sol                   | 153         | require(recoveredAddress != address(0) && recoveredAddress == owner, "INVALID_SIGNER"); | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC20.sol#L153)                  |
| party-contracts-c4/contracts/vendor/solmate/ERC721.sol                  | 35          | require((owner = \_ownerOf[id]) != address(0), "NOT_MINTED");                           | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC721.sol#L35)                  |
| party-contracts-c4/contracts/vendor/solmate/ERC721.sol                  | 39          | require(owner != address(0), "ZERO_ADDRESS");                                           | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC721.sol#L39)                  |
| party-contracts-c4/contracts/vendor/solmate/ERC721.sol                  | 68          | require(msg.sender == owner                                                             |                                                                                                                                | isApprovedForAll[owner][msg.sender], "NOT_AUTHORIZED"); | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC721.sol#L68)  |
| party-contracts-c4/contracts/vendor/solmate/ERC721.sol                  | 86          | require(from == \_ownerOf[id], "WRONG_FROM");                                           | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC721.sol#L86)                  |
| party-contracts-c4/contracts/vendor/solmate/ERC721.sol                  | 88          | require(to != address(0), "INVALID_RECIPIENT");                                         | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC721.sol#L88)                  |
| party-contracts-c4/contracts/vendor/solmate/ERC721.sol                  | 90          | require(                                                                                | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC721.sol#L90)                  |
| party-contracts-c4/contracts/vendor/solmate/ERC721.sol                  | 117         | require(                                                                                | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC721.sol#L117)                 |
| party-contracts-c4/contracts/vendor/solmate/ERC721.sol                  | 133         | require(                                                                                | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC721.sol#L133)                 |
| party-contracts-c4/contracts/vendor/solmate/ERC721.sol                  | 157         | require(to != address(0), "INVALID_RECIPIENT");                                         | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC721.sol#L157)                 |
| party-contracts-c4/contracts/vendor/solmate/ERC721.sol                  | 159         | require(\_ownerOf[id] == address(0), "ALREADY_MINTED");                                 | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC721.sol#L159)                 |
| party-contracts-c4/contracts/vendor/solmate/ERC721.sol                  | 174         | require(owner != address(0), "NOT_MINTED");                                             | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC721.sol#L174)                 |
| party-contracts-c4/contracts/vendor/solmate/ERC721.sol                  | 195         | require(                                                                                | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC721.sol#L195)                 |
| party-contracts-c4/contracts/vendor/solmate/ERC721.sol                  | 210         | require(                                                                                | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC721.sol#L210)                 |

#### **Tools used**

[c4udit](https://github.com/byterocket/c4udit)

---

### **[G009] Address 0 check can be done in assembly to save gas**

---

#### **Impact**

In Solidity, checking if an address is 0 is commonly done using the syntax address(0).
However, this operation is expensive in terms of gas usage. Instead, this check can be done more efficiently
using inline assembly

#### **Number of Findings:** 61

#### **Details:**

| File                                                                  | Line Number | Line Content                                                                                         | Code Location                                                                                                                 |
| --------------------------------------------------------------------- | ----------- | ---------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------- | ---------------------------- | ------------------------------------------------------------------------------------------------------------------- |
| party-contracts-c4/contracts/crowdfund/AuctionCrowdfund.sol           | 270         | return address(party) != address(0)                                                                  | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/AuctionCrowdfund.sol#L270)           |
| party-contracts-c4/contracts/crowdfund/BuyCrowdfundBase.sol           | 153         | return address(party) != address(0)                                                                  | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/BuyCrowdfundBase.sol#L153)           |
| party-contracts-c4/contracts/crowdfund/Crowdfund.sol                  | 367         | if (splitRecipient\_ == address(0)) {                                                                | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/Crowdfund.sol#L367)                  |
| party-contracts-c4/contracts/crowdfund/Crowdfund.sol                  | 388         | if (delegate == address(0)) {                                                                        | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/Crowdfund.sol#L388)                  |
| party-contracts-c4/contracts/crowdfund/Crowdfund.sol                  | 392         | if (gateKeeper != IGateKeeper(address(0))) {                                                         | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/Crowdfund.sol#L392)                  |
| party-contracts-c4/contracts/crowdfund/Crowdfund.sol                  | 474         | if (delegate == address(0)) {                                                                        | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/Crowdfund.sol#L474)                  |
| party-contracts-c4/contracts/crowdfund/CrowdfundFactory.sol           | 116         | address(gateKeeper) == address(0)                                                                    |                                                                                                                               |                              | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/CrowdfundFactory.sol#L116) |
| party-contracts-c4/contracts/crowdfund/CrowdfundNFT.sol               | 89          | return address(0);                                                                                   | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/CrowdfundNFT.sol#L89)                |
| party-contracts-c4/contracts/crowdfund/CrowdfundNFT.sol               | 127         | if (owner == address(0)) {                                                                           | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/CrowdfundNFT.sol#L127)               |
| party-contracts-c4/contracts/crowdfund/CrowdfundNFT.sol               | 138         | return \_owners[uint256(uint160(owner))] != address(0);                                              | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/CrowdfundNFT.sol#L138)               |
| party-contracts-c4/contracts/crowdfund/CrowdfundNFT.sol               | 146         | emit Transfer(address(0), owner, tokenId);                                                           | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/CrowdfundNFT.sol#L146)               |
| party-contracts-c4/contracts/crowdfund/CrowdfundNFT.sol               | 153         | \_owners[tokenId] = address(0);                                                                      | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/CrowdfundNFT.sol#L153)               |
| party-contracts-c4/contracts/crowdfund/CrowdfundNFT.sol               | 154         | emit Transfer(owner, address(0), tokenId);                                                           | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/CrowdfundNFT.sol#L154)               |
| party-contracts-c4/contracts/market-wrapper/KoansMarketWrapper.sol    | 73          | if (bidder == address(0)) {                                                                          | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/market-wrapper/KoansMarketWrapper.sol#L73)     |
| party-contracts-c4/contracts/market-wrapper/NounsMarketWrapper.sol    | 71          | if (bidder == address(0)) {                                                                          | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/market-wrapper/NounsMarketWrapper.sol#L71)     |
| party-contracts-c4/contracts/market-wrapper/ZoraMarketWrapper.sol     | 43          | return \_auction.tokenOwner != address(0);                                                           | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/market-wrapper/ZoraMarketWrapper.sol#L43)      |
| party-contracts-c4/contracts/market-wrapper/ZoraMarketWrapper.sol     | 60          | \_auction.auctionCurrency == IERC20(address(0));                                                     | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/market-wrapper/ZoraMarketWrapper.sol#L60)      |
| party-contracts-c4/contracts/market-wrapper/ZoraMarketWrapper.sol     | 75          | if (\_auction.bidder == address(0)) {                                                                | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/market-wrapper/ZoraMarketWrapper.sol#L75)      |
| party-contracts-c4/contracts/party/PartyFactory.sol                   | 36          | if (authority == address(0)) {                                                                       | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/party/PartyFactory.sol#L36)                    |
| party-contracts-c4/contracts/party/PartyGovernance.sol                | 460         | if (newPartyHost != address(0)) {                                                                    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/party/PartyGovernance.sol#L460)                |
| party-contracts-c4/contracts/party/PartyGovernance.sol                | 885         | \_adjustVotingPower(from, -powerI192, address(0));                                                   | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/party/PartyGovernance.sol#L885)                |
| party-contracts-c4/contracts/party/PartyGovernance.sol                | 886         | \_adjustVotingPower(to, powerI192, address(0));                                                      | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/party/PartyGovernance.sol#L886)                |
| party-contracts-c4/contracts/party/PartyGovernance.sol                | 898         | oldDelegate = oldDelegate == address(0) ? voter : oldDelegate;                                       | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/party/PartyGovernance.sol#L898)                |
| party-contracts-c4/contracts/party/PartyGovernance.sol                | 900         | delegate = delegate == address(0) ? oldDelegate : delegate;                                          | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/party/PartyGovernance.sol#L900)                |
| party-contracts-c4/contracts/party/PartyGovernance.sol                | 931         | if (newDelegate == address(0)                                                                        |                                                                                                                               | oldDelegate == address(0)) { | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/party/PartyGovernance.sol#L931)      |
| party-contracts-c4/contracts/party/PartyGovernanceNFT.sol             | 107         | return (address(0), 0); // Just to make the compiler happy.                                          | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/party/PartyGovernanceNFT.sol#L107)             |
| party-contracts-c4/contracts/proposals/ArbitraryCallsProposal.sol     | 176         | if (op != address(0)) {                                                                              | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ArbitraryCallsProposal.sol#L176)     |
| party-contracts-c4/contracts/proposals/FractionalizeProposal.sol      | 69          | vault.updateCurator(address(0));                                                                     | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/FractionalizeProposal.sol#L69)       |
| party-contracts-c4/contracts/proposals/ListOnOpenseaProposal.sol      | 266         | orderParams.orderType = orderParams.zone == address(0)                                               | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ListOnOpenseaProposal.sol#L266)      |
| party-contracts-c4/contracts/proposals/ListOnOpenseaProposal.sol      | 287         | cons.token = address(0);                                                                             | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ListOnOpenseaProposal.sol#L287)      |
| party-contracts-c4/contracts/proposals/ListOnOpenseaProposal.sol      | 294         | cons.token = address(0);                                                                             | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ListOnOpenseaProposal.sol#L294)      |
| party-contracts-c4/contracts/proposals/ListOnOpenseaProposal.sol      | 367         | token.approve(address(0), tokenId);                                                                  | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ListOnOpenseaProposal.sol#L367)      |
| party-contracts-c4/contracts/proposals/ListOnZoraProposal.sol         | 149         | payable(address(0)),                                                                                 | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ListOnZoraProposal.sol#L149)         |
| party-contracts-c4/contracts/proposals/ListOnZoraProposal.sol         | 151         | IERC20(address(0)) // Indicates ETH sale                                                             | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ListOnZoraProposal.sol#L151)         |
| party-contracts-c4/contracts/renderers/CrowdfundNFTRenderer.sol       | 108         | if (Crowdfund(payable(address(this))).ownerOf(tokenId) == address(0)) {                              | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/renderers/CrowdfundNFTRenderer.sol#L108)       |
| party-contracts-c4/contracts/renderers/PartyGovernanceNFTRenderer.sol | 114         | if(\_ownerOf[tokenId] == address(0)) {                                                               | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/renderers/PartyGovernanceNFTRenderer.sol#L114) |
| party-contracts-c4/contracts/renderers/PartyGovernanceNFTRenderer.sol | 183         | receiver = address(0);                                                                               | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/renderers/PartyGovernanceNFTRenderer.sol#L183) |
| party-contracts-c4/contracts/utils/LibSafeERC721.sol                  | 26          | return address(0);                                                                                   | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/utils/LibSafeERC721.sol#L26)                   |
| party-contracts-c4/contracts/vendor/solmate/ERC1155.sol               | 58          | ? to != address(0)                                                                                   | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC1155.sol#L58)                |
| party-contracts-c4/contracts/vendor/solmate/ERC1155.sol               | 98          | ? to != address(0)                                                                                   | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC1155.sol#L98)                |
| party-contracts-c4/contracts/vendor/solmate/ERC1155.sol               | 147         | emit TransferSingle(msg.sender, address(0), to, id, amount);                                         | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC1155.sol#L147)               |
| party-contracts-c4/contracts/vendor/solmate/ERC1155.sol               | 151         | ? to != address(0)                                                                                   | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC1155.sol#L151)               |
| party-contracts-c4/contracts/vendor/solmate/ERC1155.sol               | 152         | : ERC1155TokenReceiverBase(to).onERC1155Received(msg.sender, address(0), id, amount, data) ==        | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC1155.sol#L152)               |
| party-contracts-c4/contracts/vendor/solmate/ERC1155.sol               | 178         | emit TransferBatch(msg.sender, address(0), to, ids, amounts);                                        | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC1155.sol#L178)               |
| party-contracts-c4/contracts/vendor/solmate/ERC1155.sol               | 182         | ? to != address(0)                                                                                   | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC1155.sol#L182)               |
| party-contracts-c4/contracts/vendor/solmate/ERC1155.sol               | 183         | : ERC1155TokenReceiverBase(to).onERC1155BatchReceived(msg.sender, address(0), ids, amounts, data) == | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC1155.sol#L183)               |
| party-contracts-c4/contracts/vendor/solmate/ERC1155.sol               | 208         | emit TransferBatch(msg.sender, from, address(0), ids, amounts);                                      | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC1155.sol#L208)               |
| party-contracts-c4/contracts/vendor/solmate/ERC1155.sol               | 218         | emit TransferSingle(msg.sender, from, address(0), id, amount);                                       | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC1155.sol#L218)               |
| party-contracts-c4/contracts/vendor/solmate/ERC20.sol                 | 153         | require(recoveredAddress != address(0) && recoveredAddress == owner, "INVALID_SIGNER");              | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC20.sol#L153)                 |
| party-contracts-c4/contracts/vendor/solmate/ERC20.sol                 | 191         | emit Transfer(address(0), to, amount);                                                               | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC20.sol#L191)                 |
| party-contracts-c4/contracts/vendor/solmate/ERC20.sol                 | 203         | emit Transfer(from, address(0), amount);                                                             | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC20.sol#L203)                 |
| party-contracts-c4/contracts/vendor/solmate/ERC721.sol                | 35          | require((owner = \_ownerOf[id]) != address(0), "NOT_MINTED");                                        | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC721.sol#L35)                 |
| party-contracts-c4/contracts/vendor/solmate/ERC721.sol                | 39          | require(owner != address(0), "ZERO_ADDRESS");                                                        | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC721.sol#L39)                 |
| party-contracts-c4/contracts/vendor/solmate/ERC721.sol                | 88          | require(to != address(0), "INVALID_RECIPIENT");                                                      | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC721.sol#L88)                 |
| party-contracts-c4/contracts/vendor/solmate/ERC721.sol                | 157         | require(to != address(0), "INVALID_RECIPIENT");                                                      | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC721.sol#L157)                |
| party-contracts-c4/contracts/vendor/solmate/ERC721.sol                | 159         | require(\_ownerOf[id] == address(0), "ALREADY_MINTED");                                              | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC721.sol#L159)                |
| party-contracts-c4/contracts/vendor/solmate/ERC721.sol                | 168         | emit Transfer(address(0), to, id);                                                                   | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC721.sol#L168)                |
| party-contracts-c4/contracts/vendor/solmate/ERC721.sol                | 174         | require(owner != address(0), "NOT_MINTED");                                                          | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC721.sol#L174)                |
| party-contracts-c4/contracts/vendor/solmate/ERC721.sol                | 185         | emit Transfer(owner, address(0), id);                                                                | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC721.sol#L185)                |
| party-contracts-c4/contracts/vendor/solmate/ERC721.sol                | 197         | ERC721TokenReceiver(to).onERC721Received(msg.sender, address(0), id, "") ==                          | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC721.sol#L197)                |
| party-contracts-c4/contracts/vendor/solmate/ERC721.sol                | 212         | ERC721TokenReceiver(to).onERC721Received(msg.sender, address(0), id, data) ==                        | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC721.sol#L212)                |

#### **Tools used**

[c4udit](https://github.com/byterocket/c4udit)

---

### **[G012] Setting the constructor to payable**

---

#### **Impact**

You can cut out 10 opcodes in the creation-time EVM bytecode if you declare a constructor payable.
Making the constructor payable eliminates the need for an initial check of msg.value == 0 and saves 13 gas on deployment with no security risks.

#### **Number of Findings:** 27

#### **Details:**

| File                                                                    | Line Number | Line Content                                                            | Code Location                                                                                                                  |
| ----------------------------------------------------------------------- | ----------- | ----------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------ |
| party-contracts-c4/contracts/crowdfund/AuctionCrowdfund.sol             | 104         | constructor(IGlobals globals) Crowdfund(globals) {}                     | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/AuctionCrowdfund.sol#L104)            |
| party-contracts-c4/contracts/crowdfund/BuyCrowdfund.sol                 | 62          | constructor(IGlobals globals) BuyCrowdfundBase(globals) {}              | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/BuyCrowdfund.sol#L62)                 |
| party-contracts-c4/contracts/crowdfund/BuyCrowdfundBase.sol             | 67          | constructor(IGlobals globals) Crowdfund(globals) {}                     | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/BuyCrowdfundBase.sol#L67)             |
| party-contracts-c4/contracts/crowdfund/CollectionBuyCrowdfund.sol       | 77          | constructor(IGlobals globals) BuyCrowdfundBase(globals) {}              | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/CollectionBuyCrowdfund.sol#L77)       |
| party-contracts-c4/contracts/crowdfund/Crowdfund.sol                    | 118         | constructor(IGlobals globals) CrowdfundNFT(globals) {                   | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/Crowdfund.sol#L118)                   |
| party-contracts-c4/contracts/crowdfund/CrowdfundFactory.sol             | 25          | constructor(IGlobals globals) {                                         | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/CrowdfundFactory.sol#L25)             |
| party-contracts-c4/contracts/crowdfund/CrowdfundNFT.sol                 | 34          | constructor(IGlobals globals) {                                         | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/CrowdfundNFT.sol#L34)                 |
| party-contracts-c4/contracts/distribution/TokenDistributor.sol          | 93          | constructor(IGlobals globals) {                                         | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/distribution/TokenDistributor.sol#L93)          |
| party-contracts-c4/contracts/globals/Globals.sol                        | 23          | constructor(address multiSig\_) {                                       | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/globals/Globals.sol#L23)                        |
| party-contracts-c4/contracts/market-wrapper/FoundationMarketWrapper.sol | 24          | constructor(address \_foundationMarket) {                               | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/market-wrapper/FoundationMarketWrapper.sol#L24) |
| party-contracts-c4/contracts/market-wrapper/KoansMarketWrapper.sol      | 24          | constructor(address \_koansAuctionHouse) {                              | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/market-wrapper/KoansMarketWrapper.sol#L24)      |
| party-contracts-c4/contracts/market-wrapper/NounsMarketWrapper.sol      | 23          | constructor(address \_nounsAuctionHouse) {                              | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/market-wrapper/NounsMarketWrapper.sol#L23)      |
| party-contracts-c4/contracts/market-wrapper/ZoraMarketWrapper.sol       | 27          | constructor(address \_zoraAuctionHouse) {                               | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/market-wrapper/ZoraMarketWrapper.sol#L27)       |
| party-contracts-c4/contracts/party/Party.sol                            | 28          | constructor(IGlobals globals) PartyGovernanceNFT(globals) {}            | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/party/Party.sol#L28)                            |
| party-contracts-c4/contracts/party/PartyFactory.sol                     | 21          | constructor(IGlobals globals) {                                         | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/party/PartyFactory.sol#L21)                     |
| party-contracts-c4/contracts/party/PartyGovernance.sol                  | 266         | constructor(IGlobals globals) {                                         | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/party/PartyGovernance.sol#L266)                 |
| party-contracts-c4/contracts/party/PartyGovernanceNFT.sol               | 44          | constructor(IGlobals globals) PartyGovernance(globals) ERC721('', '') { | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/party/PartyGovernanceNFT.sol#L44)               |
| party-contracts-c4/contracts/proposals/FractionalizeProposal.sol        | 34          | constructor(IFractionalV1VaultFactory vaultFactory) {                   | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/FractionalizeProposal.sol#L34)        |
| party-contracts-c4/contracts/proposals/ListOnOpenseaProposal.sol        | 108         | constructor(                                                            | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ListOnOpenseaProposal.sol#L108)       |
| party-contracts-c4/contracts/proposals/ListOnZoraProposal.sol           | 70          | constructor(IGlobals globals, IZoraAuctionHouse zoraAuctionHouse) {     | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ListOnZoraProposal.sol#L70)           |
| party-contracts-c4/contracts/proposals/ProposalExecutionEngine.sol      | 83          | constructor(                                                            | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ProposalExecutionEngine.sol#L83)      |
| party-contracts-c4/contracts/renderers/CrowdfundNFTRenderer.sol         | 21          | constructor(IGlobals globals) {                                         | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/renderers/CrowdfundNFTRenderer.sol#L21)         |
| party-contracts-c4/contracts/renderers/PartyGovernanceNFTRenderer.sol   | 43          | constructor(IGlobals globals) {                                         | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/renderers/PartyGovernanceNFTRenderer.sol#L43)   |
| party-contracts-c4/contracts/utils/Implementation.sol                   | 11          | constructor() { IMPL = address(this); }                                 | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/utils/Implementation.sol#L11)                   |
| party-contracts-c4/contracts/utils/Proxy.sol                            | 16          | constructor(Implementation impl, bytes memory initCallData) payable {   | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/utils/Proxy.sol#L16)                            |
| party-contracts-c4/contracts/vendor/solmate/ERC20.sol                   | 50          | constructor(                                                            | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC20.sol#L50)                   |
| party-contracts-c4/contracts/vendor/solmate/ERC721.sol                  | 56          | constructor(string memory \_name, string memory \_symbol) {             | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC721.sol#L56)                  |

#### **Tools used**

[c4udit](https://github.com/byterocket/c4udit)

---

### **[G014] Assigning keccak operations to constant variables results in extra gas costs**

---

#### **Impact**

constants expressions are expressions. As such, keccak assigned to a constant variable are re-computed at each use of the variable,
which will consume gas unnecessarily. To save gas, Changing the variable from constant to immutable will
make the computation run only once and therefore save gas.

#### **Number of Findings:** 23

#### **Details:**

| File                                                               | Line Number | Line Content                                                                                                | Code Location                                                                                                              |
| ------------------------------------------------------------------ | ----------- | ----------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| party-contracts-c4/contracts/crowdfund/Crowdfund.sol               | 331         | mstore(opts, keccak256(add(mload(opts), 0x20), mul(mload(mload(opts)), 32)))                                | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/Crowdfund.sol#L331)               |
| party-contracts-c4/contracts/crowdfund/Crowdfund.sol               | 333         | h := keccak256(opts, 0xC0)                                                                                  | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/Crowdfund.sol#L333)               |
| party-contracts-c4/contracts/distribution/TokenDistributor.sol     | 397         | keccak256(info, 0xe0),                                                                                      | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/distribution/TokenDistributor.sol#L397)     |
| party-contracts-c4/contracts/gatekeepers/AllowListGateKeeper.sol   | 23          | leaf := keccak256(0x0C, 20)                                                                                 | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/gatekeepers/AllowListGateKeeper.sol#L23)    |
| party-contracts-c4/contracts/party/PartyGovernance.sol             | 405         | bytes32 dataHash = keccak256(proposal.proposalData);                                                        | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/party/PartyGovernance.sol#L405)             |
| party-contracts-c4/contracts/party/PartyGovernance.sol             | 412         | proposalHash := keccak256(proposal, 0x60)                                                                   | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/party/PartyGovernance.sol#L412)             |
| party-contracts-c4/contracts/party/PartyGovernance.sol             | 1114        | mstore(0x00, keccak256(                                                                                     | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/party/PartyGovernance.sol#L1114)            |
| party-contracts-c4/contracts/party/PartyGovernance.sol             | 1118        | mstore(0x20, keccak256(                                                                                     | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/party/PartyGovernance.sol#L1118)            |
| party-contracts-c4/contracts/party/PartyGovernance.sol             | 1122        | h := keccak256(0x00, 0x40)                                                                                  | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/party/PartyGovernance.sol#L1122)            |
| party-contracts-c4/contracts/proposals/ArbitraryCallsProposal.sol  | 121         | bytes32 resultHash = keccak256(r);                                                                          | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ArbitraryCallsProposal.sol#L121)  |
| party-contracts-c4/contracts/proposals/ListOnZoraProposal.sol      | 184         | bytes32 errHash = keccak256(errData);                                                                       | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ListOnZoraProposal.sol#L184)      |
| party-contracts-c4/contracts/proposals/ProposalExecutionEngine.sol | 80          | uint256 private constant \_STORAGE_SLOT = uint256(keccak256('ProposalExecutionEngine.Storage'));            | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ProposalExecutionEngine.sol#L80)  |
| party-contracts-c4/contracts/proposals/ProposalExecutionEngine.sol | 145         | keccak256(params.progressData),                                                                             | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ProposalExecutionEngine.sol#L145) |
| party-contracts-c4/contracts/proposals/ProposalExecutionEngine.sol | 150         | bytes32 progressDataHash = keccak256(params.progressData);                                                  | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ProposalExecutionEngine.sol#L150) |
| party-contracts-c4/contracts/proposals/ProposalExecutionEngine.sol | 178         | stor.nextProgressDataHash = keccak256(nextProgressData);                                                    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ProposalExecutionEngine.sol#L178) |
| party-contracts-c4/contracts/proposals/ProposalStorage.sol         | 19          | uint256 private constant SHARED_STORAGE_SLOT = uint256(keccak256("ProposalStorage.SharedProposalStorage")); | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ProposalStorage.sol#L19)          |
| party-contracts-c4/contracts/vendor/solmate/ERC20.sol              | 130         | keccak256(                                                                                                  | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC20.sol#L130)              |
| party-contracts-c4/contracts/vendor/solmate/ERC20.sol              | 134         | keccak256(                                                                                                  | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC20.sol#L134)              |
| party-contracts-c4/contracts/vendor/solmate/ERC20.sol              | 136         | keccak256(                                                                                                  | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC20.sol#L136)              |
| party-contracts-c4/contracts/vendor/solmate/ERC20.sol              | 167         | keccak256(                                                                                                  | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC20.sol#L167)              |
| party-contracts-c4/contracts/vendor/solmate/ERC20.sol              | 169         | keccak256("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"),            | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC20.sol#L169)              |
| party-contracts-c4/contracts/vendor/solmate/ERC20.sol              | 170         | keccak256(bytes(name)),                                                                                     | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC20.sol#L170)              |
| party-contracts-c4/contracts/vendor/solmate/ERC20.sol              | 171         | keccak256("1"),                                                                                             | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC20.sol#L171)              |

#### **Tools used**

[c4udit](https://github.com/byterocket/c4udit)

---

### **[G015 **Works Well (no false positives)**] x += y costs more gas than x = x + y for state variables**

---

#### **Impact**

....

#### **Number of Findings:** 28

#### **Details:**

| File                                                              | Line Number | Line Content                                                               | Code Location                                                                                                            |
| ----------------------------------------------------------------- | ----------- | -------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------ |
| party-contracts-c4/contracts/crowdfund/Crowdfund.sol              | 243         | ethContributed += contributions[i].amount;                                 | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/Crowdfund.sol#L243)             |
| party-contracts-c4/contracts/crowdfund/Crowdfund.sol              | 352         | ethOwed += c.amount;                                                       | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/Crowdfund.sol#L352)             |
| party-contracts-c4/contracts/crowdfund/Crowdfund.sol              | 355         | ethUsed += c.amount;                                                       | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/Crowdfund.sol#L355)             |
| party-contracts-c4/contracts/crowdfund/Crowdfund.sol              | 359         | ethUsed += partialEthUsed;                                                 | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/Crowdfund.sol#L359)             |
| party-contracts-c4/contracts/crowdfund/Crowdfund.sol              | 374         | votingPower += (splitBps\_ \* totalEthUsed + (1e4 - 1)) / 1e4; // round up | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/Crowdfund.sol#L374)             |
| party-contracts-c4/contracts/crowdfund/Crowdfund.sol              | 411         | totalContributions += amount;                                              | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/Crowdfund.sol#L411)             |
| party-contracts-c4/contracts/crowdfund/Crowdfund.sol              | 427         | lastContribution.amount += amount;                                         | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/Crowdfund.sol#L427)             |
| party-contracts-c4/contracts/distribution/TokenDistributor.sol    | 381         | \_storedBalances[balanceId] -= amount;                                     | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/distribution/TokenDistributor.sol#L381)   |
| party-contracts-c4/contracts/party/PartyGovernance.sol            | 595         | values.votes += votingPower;                                               | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/party/PartyGovernance.sol#L595)           |
| party-contracts-c4/contracts/party/PartyGovernance.sol            | 959         | newDelegateDelegatedVotingPower -= oldSnap.intrinsicVotingPower;           | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/party/PartyGovernance.sol#L959)           |
| party-contracts-c4/contracts/proposals/ArbitraryCallsProposal.sol | 72          | ethAvailable -= calls[i].value;                                            | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ArbitraryCallsProposal.sol#L72) |
| party-contracts-c4/contracts/utils/vendor/Strings.sol             | 27          | digits -= 1;                                                               | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/utils/vendor/Strings.sol#L27)             |
| party-contracts-c4/contracts/vendor/solmate/ERC1155.sol           | 51          | balanceOf[from][id] -= amount;                                             | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC1155.sol#L51)           |
| party-contracts-c4/contracts/vendor/solmate/ERC1155.sol           | 52          | balanceOf[to][id] += amount;                                               | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC1155.sol#L52)           |
| party-contracts-c4/contracts/vendor/solmate/ERC1155.sol           | 84          | balanceOf[from][id] -= amount;                                             | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC1155.sol#L84)           |
| party-contracts-c4/contracts/vendor/solmate/ERC1155.sol           | 85          | balanceOf[to][id] += amount;                                               | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC1155.sol#L85)           |
| party-contracts-c4/contracts/vendor/solmate/ERC1155.sol           | 145         | balanceOf[to][id] += amount;                                               | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC1155.sol#L145)          |
| party-contracts-c4/contracts/vendor/solmate/ERC1155.sol           | 169         | balanceOf[to]ids[i]] += amounts[i];                                        | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC1155.sol#L169)          |
| party-contracts-c4/contracts/vendor/solmate/ERC1155.sol           | 199         | balanceOf[from]ids[i]] -= amounts[i];                                      | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC1155.sol#L199)          |
| party-contracts-c4/contracts/vendor/solmate/ERC1155.sol           | 216         | balanceOf[from][id] -= amount;                                             | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC1155.sol#L216)          |
| party-contracts-c4/contracts/vendor/solmate/ERC20.sol             | 76          | balanceOf[msg.sender] -= amount;                                           | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC20.sol#L76)             |
| party-contracts-c4/contracts/vendor/solmate/ERC20.sol             | 81          | balanceOf[to] += amount;                                                   | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC20.sol#L81)             |
| party-contracts-c4/contracts/vendor/solmate/ERC20.sol             | 98          | balanceOf[from] -= amount;                                                 | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC20.sol#L98)             |
| party-contracts-c4/contracts/vendor/solmate/ERC20.sol             | 103         | balanceOf[to] += amount;                                                   | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC20.sol#L103)            |
| party-contracts-c4/contracts/vendor/solmate/ERC20.sol             | 183         | totalSupply += amount;                                                     | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC20.sol#L183)            |
| party-contracts-c4/contracts/vendor/solmate/ERC20.sol             | 188         | balanceOf[to] += amount;                                                   | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC20.sol#L188)            |
| party-contracts-c4/contracts/vendor/solmate/ERC20.sol             | 195         | balanceOf[from] -= amount;                                                 | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC20.sol#L195)            |
| party-contracts-c4/contracts/vendor/solmate/ERC20.sol             | 200         | totalSupply -= amount;                                                     | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC20.sol#L200)            |

#### **Tools used**

[c4udit](https://github.com/byterocket/c4udit)

---

### **[G017] ++i/i++ should be unchecked{++i}/unchecked{i++} when it is not possible for them to overflow**

---

#### **Impact**

The unchecked keyword is new in solidity version 0.8.0, so this only applies to that version or higher,  
 which these instances are. This saves 30-40 gas per loop.

#### **Number of Findings:** 34

#### **Details:**

| File                                                              | Line Number | Line Content                                              | Code Location                                                                                                            |
| ----------------------------------------------------------------- | ----------- | --------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------ |
| party-contracts-c4/contracts/crowdfund/CollectionBuyCrowdfund.sol | 62          | for (uint256 i; i < hosts.length; i++) {                  | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/CollectionBuyCrowdfund.sol#L62) |
| party-contracts-c4/contracts/crowdfund/Crowdfund.sol              | 180         | for (uint256 i = 0; i < contributors.length; ++i) {       | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/Crowdfund.sol#L180)             |
| party-contracts-c4/contracts/crowdfund/Crowdfund.sol              | 242         | for (uint256 i = 0; i < numContributions; ++i) {          | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/Crowdfund.sol#L242)             |
| party-contracts-c4/contracts/crowdfund/Crowdfund.sol              | 300         | for (uint256 i = 0; i < preciousTokens.length; ++i) {     | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/Crowdfund.sol#L300)             |
| party-contracts-c4/contracts/crowdfund/Crowdfund.sol              | 348         | for (uint256 i = 0; i < numContributions; ++i) {          | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/Crowdfund.sol#L348)             |
| party-contracts-c4/contracts/distribution/TokenDistributor.sol    | 230         | for (uint256 i = 0; i < infos.length; ++i) {              | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/distribution/TokenDistributor.sol#L230)   |
| party-contracts-c4/contracts/distribution/TokenDistributor.sol    | 239         | for (uint256 i = 0; i < infos.length; ++i) {              | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/distribution/TokenDistributor.sol#L239)   |
| party-contracts-c4/contracts/distribution/TokenDistributor.sol    | 357         | distributionId: ++lastDistributionIdPerParty[args.party], | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/distribution/TokenDistributor.sol#L357)   |
| party-contracts-c4/contracts/gatekeepers/AllowListGateKeeper.sol  | 33          | uint96 id\_ = ++\_lastId;                                 | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/gatekeepers/AllowListGateKeeper.sol#L33)  |
| party-contracts-c4/contracts/gatekeepers/TokenGateKeeper.sol      | 48          | uint96 id\_ = ++\_lastId;                                 | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/gatekeepers/TokenGateKeeper.sol#L48)      |
| party-contracts-c4/contracts/party/PartyGovernance.sol            | 306         | for (uint256 i=0; i < opts.hosts.length; ++i) {           | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/party/PartyGovernance.sol#L306)           |
| party-contracts-c4/contracts/party/PartyGovernance.sol            | 532         | proposalId = ++lastProposalId;                            | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/party/PartyGovernance.sol#L532)           |
| party-contracts-c4/contracts/party/PartyGovernanceNFT.sol         | 129         | uint256 tokenId = ++tokenCount;                           | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/party/PartyGovernanceNFT.sol#L129)        |
| party-contracts-c4/contracts/proposals/ArbitraryCallsProposal.sol | 52          | for (uint256 i = 0; i < hadPreciouses.length; ++i) {      | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ArbitraryCallsProposal.sol#L52) |
| party-contracts-c4/contracts/proposals/ArbitraryCallsProposal.sol | 61          | for (uint256 i = 0; i < calls.length; ++i) {              | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ArbitraryCallsProposal.sol#L61) |
| party-contracts-c4/contracts/proposals/ArbitraryCallsProposal.sol | 78          | for (uint256 i = 0; i < hadPreciouses.length; ++i) {      | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ArbitraryCallsProposal.sol#L78) |
| party-contracts-c4/contracts/proposals/LibProposal.sol            | 14          | for (uint256 i = 0; i < preciousTokens.length; ++i) {     | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/LibProposal.sol#L14)            |
| party-contracts-c4/contracts/proposals/LibProposal.sol            | 32          | for (uint256 i = 0; i < preciousTokens.length; ++i) {     | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/LibProposal.sol#L32)            |
| party-contracts-c4/contracts/proposals/ListOnOpenseaProposal.sol  | 291         | for (uint256 i = 0; i < fees.length; ++i) {               | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ListOnOpenseaProposal.sol#L291) |
| party-contracts-c4/contracts/utils/PartyHelpers.sol               | 64          | for (uint256 i = 0; i < members.length; i++) {            | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/utils/PartyHelpers.sol#L64)               |
| party-contracts-c4/contracts/utils/PartyHelpers.sol               | 85          | for (uint256 i = 0; i < voters.length; i++) {             | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/utils/PartyHelpers.sol#L85)               |
| party-contracts-c4/contracts/utils/PartyHelpers.sol               | 115         | for (uint256 i = 0; i < count; i++) {                     | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/utils/PartyHelpers.sol#L115)              |
| party-contracts-c4/contracts/utils/vendor/Strings.sol             | 22          | digits++;                                                 | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/utils/vendor/Strings.sol#L22)             |
| party-contracts-c4/contracts/utils/vendor/Strings.sol             | 44          | length++;                                                 | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/utils/vendor/Strings.sol#L44)             |
| party-contracts-c4/contracts/utils/vendor/Strings.sol             | 57          | for (uint256 i = 2 \* length + 1; i > 1; --i) {           | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/utils/vendor/Strings.sol#L57)             |
| party-contracts-c4/contracts/vendor/solmate/ERC1155.sol           | 90          | ++i;                                                      | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC1155.sol#L90)           |
| party-contracts-c4/contracts/vendor/solmate/ERC1155.sol           | 118         | for (uint256 i = 0; i < owners.length; ++i) {             | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC1155.sol#L118)          |
| party-contracts-c4/contracts/vendor/solmate/ERC1155.sol           | 174         | ++i;                                                      | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC1155.sol#L174)          |
| party-contracts-c4/contracts/vendor/solmate/ERC1155.sol           | 204         | ++i;                                                      | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC1155.sol#L204)          |
| party-contracts-c4/contracts/vendor/solmate/ERC20.sol             | 142         | nonces[owner]++,                                          | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC20.sol#L142)            |
| party-contracts-c4/contracts/vendor/solmate/ERC721.sol            | 98          | \_balanceOf[from]--;                                      | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC721.sol#L98)            |
| party-contracts-c4/contracts/vendor/solmate/ERC721.sol            | 100         | \_balanceOf[to]++;                                        | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC721.sol#L100)           |
| party-contracts-c4/contracts/vendor/solmate/ERC721.sol            | 163         | \_balanceOf[to]++;                                        | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC721.sol#L163)           |
| party-contracts-c4/contracts/vendor/solmate/ERC721.sol            | 178         | \_balanceOf[owner]--;                                     | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC721.sol#L178)           |

#### **Tools used**

[c4udit](https://github.com/byterocket/c4udit)

---

### **[L001] Unsafe ERC20 Operation(s)**

---

#### **Impact**

https://github.com/byterocket/c4-common-issues/blob/main/2-Low-Risk.md#l001---unsafe-erc20-operations

#### **Number of Findings:** 6

#### **Details:**

| File                                                             | Line Number | Line Content                                                                          | Code Location                                                                                                            |
| ---------------------------------------------------------------- | ----------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------ |
| party-contracts-c4/contracts/crowdfund/Crowdfund.sol             | 301         | preciousTokens[i].transferFrom(address(this), address(party\_), preciousTokenIds[i]); | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/Crowdfund.sol#L301)             |
| party-contracts-c4/contracts/party/PartyGovernanceNFT.sol        | 143         | super.transferFrom(owner, to, tokenId);                                               | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/party/PartyGovernanceNFT.sol#L143)        |
| party-contracts-c4/contracts/proposals/FractionalizeProposal.sol | 53          | data.token.approve(address(VAULT_FACTORY), data.tokenId);                             | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/FractionalizeProposal.sol#L53)  |
| party-contracts-c4/contracts/proposals/ListOnOpenseaProposal.sol | 256         | token.approve(conduit, tokenId);                                                      | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ListOnOpenseaProposal.sol#L256) |
| party-contracts-c4/contracts/proposals/ListOnOpenseaProposal.sol | 367         | token.approve(address(0), tokenId);                                                   | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ListOnOpenseaProposal.sol#L367) |
| party-contracts-c4/contracts/proposals/ListOnZoraProposal.sol    | 143         | token.approve(address(ZORA), tokenId);                                                | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ListOnZoraProposal.sol#L143)    |

#### **Tools used**

[c4udit](https://github.com/byterocket/c4udit)

---

### **[L002] Unspecific Compiler Version Pragma**

---

#### **Impact**

Avoid floating pragmas for non-library contracts.

#### **Number of Findings:** 67

#### **Details:**

| File                                                                        | Line Number | Line Content             | Code Location                                                                                                                     |
| --------------------------------------------------------------------------- | ----------- | ------------------------ | --------------------------------------------------------------------------------------------------------------------------------- |
| party-contracts-c4/contracts/crowdfund/AuctionCrowdfund.sol                 | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/AuctionCrowdfund.sol#L2)                 |
| party-contracts-c4/contracts/crowdfund/BuyCrowdfund.sol                     | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/BuyCrowdfund.sol#L2)                     |
| party-contracts-c4/contracts/crowdfund/BuyCrowdfundBase.sol                 | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/BuyCrowdfundBase.sol#L2)                 |
| party-contracts-c4/contracts/crowdfund/CollectionBuyCrowdfund.sol           | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/CollectionBuyCrowdfund.sol#L2)           |
| party-contracts-c4/contracts/crowdfund/Crowdfund.sol                        | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/Crowdfund.sol#L2)                        |
| party-contracts-c4/contracts/crowdfund/CrowdfundFactory.sol                 | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/CrowdfundFactory.sol#L2)                 |
| party-contracts-c4/contracts/crowdfund/CrowdfundNFT.sol                     | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/crowdfund/CrowdfundNFT.sol#L2)                     |
| party-contracts-c4/contracts/distribution/ITokenDistributor.sol             | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/distribution/ITokenDistributor.sol#L2)             |
| party-contracts-c4/contracts/distribution/ITokenDistributorParty.sol        | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/distribution/ITokenDistributorParty.sol#L2)        |
| party-contracts-c4/contracts/distribution/TokenDistributor.sol              | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/distribution/TokenDistributor.sol#L2)              |
| party-contracts-c4/contracts/gatekeepers/AllowListGateKeeper.sol            | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/gatekeepers/AllowListGateKeeper.sol#L2)            |
| party-contracts-c4/contracts/gatekeepers/IGateKeeper.sol                    | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/gatekeepers/IGateKeeper.sol#L2)                    |
| party-contracts-c4/contracts/gatekeepers/TokenGateKeeper.sol                | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/gatekeepers/TokenGateKeeper.sol#L2)                |
| party-contracts-c4/contracts/globals/Globals.sol                            | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/globals/Globals.sol#L2)                            |
| party-contracts-c4/contracts/globals/IGlobals.sol                           | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/globals/IGlobals.sol#L2)                           |
| party-contracts-c4/contracts/globals/LibGlobals.sol                         | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/globals/LibGlobals.sol#L2)                         |
| party-contracts-c4/contracts/market-wrapper/FoundationMarketWrapper.sol     | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/market-wrapper/FoundationMarketWrapper.sol#L2)     |
| party-contracts-c4/contracts/market-wrapper/IMarketWrapper.sol              | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/market-wrapper/IMarketWrapper.sol#L2)              |
| party-contracts-c4/contracts/market-wrapper/KoansMarketWrapper.sol          | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/market-wrapper/KoansMarketWrapper.sol#L2)          |
| party-contracts-c4/contracts/market-wrapper/NounsMarketWrapper.sol          | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/market-wrapper/NounsMarketWrapper.sol#L2)          |
| party-contracts-c4/contracts/market-wrapper/ZoraMarketWrapper.sol           | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/market-wrapper/ZoraMarketWrapper.sol#L2)           |
| party-contracts-c4/contracts/party/IPartyFactory.sol                        | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/party/IPartyFactory.sol#L2)                        |
| party-contracts-c4/contracts/party/Party.sol                                | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/party/Party.sol#L2)                                |
| party-contracts-c4/contracts/party/PartyFactory.sol                         | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/party/PartyFactory.sol#L2)                         |
| party-contracts-c4/contracts/party/PartyGovernance.sol                      | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/party/PartyGovernance.sol#L2)                      |
| party-contracts-c4/contracts/party/PartyGovernanceNFT.sol                   | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/party/PartyGovernanceNFT.sol#L2)                   |
| party-contracts-c4/contracts/proposals/ArbitraryCallsProposal.sol           | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ArbitraryCallsProposal.sol#L2)           |
| party-contracts-c4/contracts/proposals/FractionalizeProposal.sol            | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/FractionalizeProposal.sol#L2)            |
| party-contracts-c4/contracts/proposals/IProposalExecutionEngine.sol         | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/IProposalExecutionEngine.sol#L2)         |
| party-contracts-c4/contracts/proposals/LibProposal.sol                      | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/LibProposal.sol#L2)                      |
| party-contracts-c4/contracts/proposals/ListOnOpenseaProposal.sol            | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ListOnOpenseaProposal.sol#L2)            |
| party-contracts-c4/contracts/proposals/ListOnZoraProposal.sol               | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ListOnZoraProposal.sol#L2)               |
| party-contracts-c4/contracts/proposals/ProposalExecutionEngine.sol          | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ProposalExecutionEngine.sol#L2)          |
| party-contracts-c4/contracts/proposals/ProposalStorage.sol                  | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ProposalStorage.sol#L2)                  |
| party-contracts-c4/contracts/proposals/ZoraHelpers.sol                      | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/ZoraHelpers.sol#L2)                      |
| party-contracts-c4/contracts/proposals/vendor/FractionalV1.sol              | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/vendor/FractionalV1.sol#L2)              |
| party-contracts-c4/contracts/proposals/vendor/IOpenseaConduitController.sol | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/vendor/IOpenseaConduitController.sol#L2) |
| party-contracts-c4/contracts/proposals/vendor/IOpenseaExchange.sol          | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/proposals/vendor/IOpenseaExchange.sol#L2)          |
| party-contracts-c4/contracts/renderers/CrowdfundNFTRenderer.sol             | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/renderers/CrowdfundNFTRenderer.sol#L2)             |
| party-contracts-c4/contracts/renderers/DummyERC721Renderer.sol              | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/renderers/DummyERC721Renderer.sol#L2)              |
| party-contracts-c4/contracts/renderers/IERC721Renderer.sol                  | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/renderers/IERC721Renderer.sol#L2)                  |
| party-contracts-c4/contracts/renderers/PartyGovernanceNFTRenderer.sol       | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/renderers/PartyGovernanceNFTRenderer.sol#L2)       |
| party-contracts-c4/contracts/tokens/ERC1155Receiver.sol                     | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/tokens/ERC1155Receiver.sol#L2)                     |
| party-contracts-c4/contracts/tokens/ERC721Receiver.sol                      | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/tokens/ERC721Receiver.sol#L2)                      |
| party-contracts-c4/contracts/tokens/IERC1155.sol                            | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/tokens/IERC1155.sol#L2)                            |
| party-contracts-c4/contracts/tokens/IERC20.sol                              | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/tokens/IERC20.sol#L2)                              |
| party-contracts-c4/contracts/tokens/IERC721.sol                             | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/tokens/IERC721.sol#L2)                             |
| party-contracts-c4/contracts/tokens/IERC721Receiver.sol                     | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/tokens/IERC721Receiver.sol#L2)                     |
| party-contracts-c4/contracts/utils/EIP165.sol                               | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/utils/EIP165.sol#L2)                               |
| party-contracts-c4/contracts/utils/Implementation.sol                       | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/utils/Implementation.sol#L2)                       |
| party-contracts-c4/contracts/utils/LibAddress.sol                           | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/utils/LibAddress.sol#L2)                           |
| party-contracts-c4/contracts/utils/LibERC20Compat.sol                       | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/utils/LibERC20Compat.sol#L2)                       |
| party-contracts-c4/contracts/utils/LibRawResult.sol                         | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/utils/LibRawResult.sol#L2)                         |
| party-contracts-c4/contracts/utils/LibSafeCast.sol                          | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/utils/LibSafeCast.sol#L2)                          |
| party-contracts-c4/contracts/utils/LibSafeERC721.sol                        | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/utils/LibSafeERC721.sol#L2)                        |
| party-contracts-c4/contracts/utils/PartyHelpers.sol                         | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/utils/PartyHelpers.sol#L2)                         |
| party-contracts-c4/contracts/utils/Proxy.sol                                | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/utils/Proxy.sol#L2)                                |
| party-contracts-c4/contracts/utils/ReadOnlyDelegateCall.sol                 | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/utils/ReadOnlyDelegateCall.sol#L2)                 |
| party-contracts-c4/contracts/utils/vendor/Base64.sol                        | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/utils/vendor/Base64.sol#L2)                        |
| party-contracts-c4/contracts/utils/vendor/Strings.sol                       | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/utils/vendor/Strings.sol#L2)                       |
| party-contracts-c4/contracts/vendor/markets/IFoundationMarket.sol           | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/markets/IFoundationMarket.sol#L2)           |
| party-contracts-c4/contracts/vendor/markets/IKoansAuctionHouse.sol          | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/markets/IKoansAuctionHouse.sol#L2)          |
| party-contracts-c4/contracts/vendor/markets/INounsAuctionHouse.sol          | 18          | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/markets/INounsAuctionHouse.sol#L18)         |
| party-contracts-c4/contracts/vendor/markets/IZoraAuctionHouse.sol           | 2           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/markets/IZoraAuctionHouse.sol#L2)           |
| party-contracts-c4/contracts/vendor/solmate/ERC1155.sol                     | 5           | pragma solidity ^0.8;    | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC1155.sol#L5)                     |
| party-contracts-c4/contracts/vendor/solmate/ERC20.sol                       | 5           | pragma solidity >=0.8.0; | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC20.sol#L5)                       |
| party-contracts-c4/contracts/vendor/solmate/ERC721.sol                      | 5           | pragma solidity >=0.8.0; | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC721.sol#L5)                      |

#### **Tools used**

[c4udit](https://github.com/byterocket/c4udit)

---

### **[L007] Use of bytes.concat() instead of abi.encodePacked()**

---

#### **Impact**

Since version 0.8.4 for appending bytes, bytes.concat() can be used instead of abi.encodePacked(,)

#### **Number of Findings:** 26

#### **Details:**

| File                                                                  | Line Number | Line Content                                                                            | Code Location                                                                                                                 |
| --------------------------------------------------------------------- | ----------- | --------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------- |
| party-contracts-c4/contracts/renderers/CrowdfundNFTRenderer.sol       | 28          | parts[0] = string(abi.encodePacked(                                                     | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/renderers/CrowdfundNFTRenderer.sol#L28)        |
| party-contracts-c4/contracts/renderers/CrowdfundNFTRenderer.sol       | 41          | return string(abi.encodePacked(                                                         | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/renderers/CrowdfundNFTRenderer.sol#L41)        |
| party-contracts-c4/contracts/renderers/CrowdfundNFTRenderer.sol       | 53          | return string(abi.encodePacked(                                                         | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/renderers/CrowdfundNFTRenderer.sol#L53)        |
| party-contracts-c4/contracts/renderers/CrowdfundNFTRenderer.sol       | 61          | return string(abi.encodePacked('#', Strings.toString(tokenId)));                        | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/renderers/CrowdfundNFTRenderer.sol#L61)        |
| party-contracts-c4/contracts/renderers/CrowdfundNFTRenderer.sol       | 67          | return string(abi.encodePacked('Owner: ', Strings.toHexString(owner)));                 | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/renderers/CrowdfundNFTRenderer.sol#L67)        |
| party-contracts-c4/contracts/renderers/CrowdfundNFTRenderer.sol       | 92          | return string(abi.encodePacked('ETH contributed: ', Strings.toString(ethContributed))); | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/renderers/CrowdfundNFTRenderer.sol#L92)        |
| party-contracts-c4/contracts/renderers/CrowdfundNFTRenderer.sol       | 98          | return string(abi.encodePacked('ETH used: ', Strings.toString(ethUsed)));               | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/renderers/CrowdfundNFTRenderer.sol#L98)        |
| party-contracts-c4/contracts/renderers/CrowdfundNFTRenderer.sol       | 104         | return string(abi.encodePacked('ETH owed: ', Strings.toString(ethOwed)));               | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/renderers/CrowdfundNFTRenderer.sol#L104)       |
| party-contracts-c4/contracts/renderers/CrowdfundNFTRenderer.sol       | 134         | abi.encodePacked(                                                                       | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/renderers/CrowdfundNFTRenderer.sol#L134)       |
| party-contracts-c4/contracts/renderers/CrowdfundNFTRenderer.sol       | 144         | abi.encodePacked(                                                                       | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/renderers/CrowdfundNFTRenderer.sol#L144)       |
| party-contracts-c4/contracts/renderers/CrowdfundNFTRenderer.sol       | 154         | output = string(abi.encodePacked('data:application/json;base64,', json));               | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/renderers/CrowdfundNFTRenderer.sol#L154)       |
| party-contracts-c4/contracts/renderers/CrowdfundNFTRenderer.sol       | 162         | abi.encodePacked(                                                                       | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/renderers/CrowdfundNFTRenderer.sol#L162)       |
| party-contracts-c4/contracts/renderers/CrowdfundNFTRenderer.sol       | 175         | return string(abi.encodePacked('data:application/json;base64,', json));                 | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/renderers/CrowdfundNFTRenderer.sol#L175)       |
| party-contracts-c4/contracts/renderers/PartyGovernanceNFTRenderer.sol | 50          | parts[0] = string(abi.encodePacked(                                                     | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/renderers/PartyGovernanceNFTRenderer.sol#L50)  |
| party-contracts-c4/contracts/renderers/PartyGovernanceNFTRenderer.sol | 63          | return string(abi.encodePacked(                                                         | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/renderers/PartyGovernanceNFTRenderer.sol#L63)  |
| party-contracts-c4/contracts/renderers/PartyGovernanceNFTRenderer.sol | 75          | return string(abi.encodePacked(                                                         | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/renderers/PartyGovernanceNFTRenderer.sol#L75)  |
| party-contracts-c4/contracts/renderers/PartyGovernanceNFTRenderer.sol | 83          | return string(abi.encodePacked('#', Strings.toString(tokenId)));                        | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/renderers/PartyGovernanceNFTRenderer.sol#L83)  |
| party-contracts-c4/contracts/renderers/PartyGovernanceNFTRenderer.sol | 91          | abi.encodePacked(                                                                       | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/renderers/PartyGovernanceNFTRenderer.sol#L91)  |
| party-contracts-c4/contracts/renderers/PartyGovernanceNFTRenderer.sol | 103         | return string(abi.encodePacked('Owner: ', Strings.toHexString(owner)));                 | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/renderers/PartyGovernanceNFTRenderer.sol#L103) |
| party-contracts-c4/contracts/renderers/PartyGovernanceNFTRenderer.sol | 110         | return string(abi.encodePacked('Delegate: ', Strings.toHexString(delegatedAddress)));   | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/renderers/PartyGovernanceNFTRenderer.sol#L110) |
| party-contracts-c4/contracts/renderers/PartyGovernanceNFTRenderer.sol | 135         | abi.encodePacked(                                                                       | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/renderers/PartyGovernanceNFTRenderer.sol#L135) |
| party-contracts-c4/contracts/renderers/PartyGovernanceNFTRenderer.sol | 144         | abi.encodePacked(                                                                       | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/renderers/PartyGovernanceNFTRenderer.sol#L144) |
| party-contracts-c4/contracts/renderers/PartyGovernanceNFTRenderer.sol | 154         | output = string(abi.encodePacked('data:application/json;base64,', json));               | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/renderers/PartyGovernanceNFTRenderer.sol#L154) |
| party-contracts-c4/contracts/renderers/PartyGovernanceNFTRenderer.sol | 162         | abi.encodePacked(                                                                       | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/renderers/PartyGovernanceNFTRenderer.sol#L162) |
| party-contracts-c4/contracts/renderers/PartyGovernanceNFTRenderer.sol | 175         | return string(abi.encodePacked('data:application/json;base64,', json));                 | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/renderers/PartyGovernanceNFTRenderer.sol#L175) |
| party-contracts-c4/contracts/vendor/solmate/ERC20.sol                 | 131         | abi.encodePacked(                                                                       | [View Code](https://github.com/PartyDAO/party-contracts-c4/blob/main/contracts/vendor/solmate/ERC20.sol#L131)                 |

#### **Tools used**

[c4udit](https://github.com/byterocket/c4udit)
