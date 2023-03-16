package analyzer

// AllIssues returns the list of all issues.
func AllIssues() []Issue {
	return append(GasOpIssues(), LowRiskIssues()...)
}

// GasOpIssues returns the list of all gas optimization issues.
func GasOpIssues() []Issue {
	return []Issue{
		// G001 - Don't Initialize Variables with Default Value
		{
			"G001 **Works Well (no false positives)**",
			GASOP,
			"Don't Initialize Variables with Default Value",
			"When initializing a variable with a default value, it can lead to unnecessary gas consumption",
			`\b(uint\d{0,3}|uint256|bool)\s+(\w+)\s*=\s*(0|false|address\(0\)|""|'')\s*(;.*)?$`,
		},
		// G002 - Cache Array Length Outside of Loop
		{
			"G002",
			GASOP,
			"Cache Array Length Outside of Loop **Works Well (no false positives)**",
			`By caching the length outside the loop, the function will only need to retrieve the value once and then 
			use the cached value within the loop. This optimization can result in significant gas savings for functions 
			that iterate over large arrays.`,
			`for\s*\((.*;)?\s*.*(<|>|<=|>=)\s*[^;]+\.length\s*;`,
		},
		// G003 - Use != 0 instead of > 0 for Unsigned Integer Comparison
		{
			"G003",
			GASOP,
			"Use != 0 instead of > 0 for Unsigned Integer Comparison **Works Well (no false positives)**",
			"When dealing with unsigned integer types, comparisons with != 0 are cheaper than with > 0.",
			`\s*> ?0`,
		},
		// G006 - Use immutable for OpenZeppelin AccessControl's Roles Declarations
		{
			"G006",
			GASOP,
			"Use immutable for OpenZeppelin AccessControl's Roles Declarations",
			`⚡️ Only valid for solidity versions <0.6.12 ⚡️
			Access roles marked as constant results in computing the keccak256 operation each 
			time the variable is used because assigned operations for constant variables are re-evaluated every time.
			Changing the variables to immutable results in computing the hash only once on deployment, leading to gas savings.`,
			`constant.*keccak256`,
		},
		// G007 - Long Revert Strings
		{
			"G007",
			GASOP,
			"Long Revert Strings **Works Well (no false positives)**",
			"For the output on a require, Each extra memory word of bytes past the original 32 incurs an MSTORE which costs 3 gas",
			`require\(.+\".{33,}\"`, // Anything in a require and between "'s with at least 33 characters
		},
		// G008 - Use Shift Right/Left instead of Division/Multiplication if possible
		{
			"G008",
			GASOP,
			"Use Shift Right/Left instead of Division/Multiplication if possible",
			`<x> / 2is the same as <x> >> 1. While the compiler uses the SHRopcode to accomplish both, 
			the version that uses division incurs an overhead of 20 gas due to JUMPs to and from a compiler 
			utility function that introduces checks which can be avoided by using unchecked {}around the division by two.`,
			`(/[2,4,8]|/ [2,4,8]|\*[2,4,8]|\* [2,4,8])`, // Matches division or multiplication by 2, 4, or 8
		},
		{
			"G009",
			GASOP,
			"++i costs less gas than i++, especially when its used in for-loops (--i/i-- too)",
			"...",
			`\+\+|--`,
		},
		{
			"G010 **Works Well (no false positives)**",
			GASOP,
			"USE CUSTOM ERRORS RATHER THAN REVERT()/REQUIRE() STRINGS TO SAVE GAS",
			`Custom errors are available from solidity version 0.8.4. Custom errors save ~50 gas each time 
			 they’re hit by avoiding having to allocate and store the revert string.
			 Not defining the strings also save deployment gas.`,
			`\b(require|assert)\b`,
		},
		{
			"G011",
			GASOP,
			"Address 0 check can be done in assembly to save gas",
			`In Solidity, checking if an address is 0 is commonly done using the syntax address(0). 
			However, this operation is expensive in terms of gas usage. Instead, this check can be done more efficiently 
			using inline assembly`,
			`address\s*\(\s*0[x0]*\s*\)`,
		},
		{
			"G012",
			GASOP,
			"Using private rather than public for constants, saves gas **Works Well (no false positives)**",
			"public constants generate a getter function, which increases the bytecode size and gas usage.",
			`public\s+constant\s+\S+\s*=.*;`,
		},
		{
			"G013",
			GASOP,
			"Functions guaranteed to revert when called by normal users can be marked payable",
			`Marking the function as payable will lower the gas cost for legitimate callers because the 
			compiler will not include checks for whether a payment was provided. `,
			`\b(?:onlyOwner|onlyAdmin)\b`,
		},
		{
			"G014",
			GASOP,
			"Setting the constructor to payable",
			`You can cut out 10 opcodes in the creation-time EVM bytecode if you declare a constructor payable. 
			Making the constructor payable eliminates the need for an initial check of msg.value == 0 and saves 13 gas on deployment with no security risks.`,
			`\bconstructor\b`,
		},
		{
			"G015",
			GASOP,
			"10 ** 18 can be changed to 1e18 and save some gas",
			"10 ** 18 can be changed to 1e18 to avoid unnecessary arithmetic operation and save gas.",
			`10\s*\*\*\s*18\b`,
		},
		{
			"G016",
			GASOP,
			"Assigning keccak operations to constant variables results in extra gas costs",
			`constants expressions are expressions. As such, keccak assigned to a constant variable are re-computed at each use of the variable, 
			which will consume gas unnecessarily. To save gas, Changing the variable from constant to immutable will 
			make the computation run only once and therefore save gas.`,

			`\b(keccak|keccak256)\b`,
		},
		{
			"G017 **Works Well (no false positives)**",
			GASOP,
			"x += y costs more gas than x = x + y for state variables",
			"....",
			`[\+\-]=`,
		},
		{
			"G018",
			GASOP,
			"Splitting require() statements that use && saves gas - (saves 8 gas per &&)",
			`Instead of using the && operator in a single require statement to check multiple conditions, 
			using multiple require statements with 1 condition per require statement will save 8 GAS per &&. 
			The gas difference would only be realized if the revert condition is realized(met).`,

			`&&(?!\s*=)`,
		},
		{
			"G019",
			GASOP,
			`++i/i++ should be unchecked{++i}/unchecked{i++} when it is not possible for them to overflow`,
			`The unchecked keyword is new in solidity version 0.8.0, so this only applies to that version or higher,  
			which these instances are. This saves 30-40 gas per loop.`,
			`\+\+|--`,
		},
		{
			"G020 **Works Well (no false positives)**",
			GASOP,
			`Don’t compare boolean expressions to boolean literals`, 
			"if (<x> == true) => if (<x>), if (<x> == false) => if (!<x>)",
			`== (true|false)\b`,
		},
	}
}

// LowRiskIssues returns the list of all low risk issues.
func LowRiskIssues() []Issue {
	return []Issue{
		// L001 - Unsafe ERC20 Operation(s)
		{
			"L001",
			LOW,
			"Unsafe ERC20 Operation(s)",
			"https://github.com/byterocket/c4-common-issues/blob/main/2-Low-Risk.md#l001---unsafe-erc20-operations",
			`\.transfer\(|\.transferFrom\(|\.approve\(`, // ".tranfer(", ".transferFrom(" or ".approve("
		},
		// L003 - Unspecific Compiler Version Pragma
		{
			"L003",
			LOW,
			"Unspecific Compiler Version Pragma",
			"Avoid floating pragmas for non-library contracts.",
			"pragma solidity (\\^|>)", // "pragma solidity ^" or "pragma solidity >"
		},
		// L005 - Do not use Deprecated Library Functions
		{
			"L005",
			LOW,
			"Do not use Deprecated Library Functions",
			`The usage of deprecated library functions should be discouraged.
			This issue is mostly related to OpenZeppelin libraries.`,
			`_setupRole\(|safeApprove\(`, // _setupRole and safeApprove are common deprecated lib functions
		},
				// L005 - Do not use Deprecated Library Functions
		{
			"L006",
			LOW,
			"Use Ownable2StepUpgradeable instead of OwnableUpgradeable contract",
			`Contracts implementing access control's, e.g. owner, should consider implementing a Two-Step Transfer pattern.
			Otherwise it's possible that the role mistakenly transfers ownership to the wrong address, resulting in a loss of the role.`,
			`\bOwnableUpgradeable\b`, // _setupRole and safeApprove are common deprecated lib functions
		},
		{
			"L007",
			LOW,
			"Use safeTransferOwnership instead of transferOwnership function",
			"Use a 2 structure transferOwnership which is safer. safeTransferOwnership, use is more secure due to 2-stage ownership transfer.",
			`\bOwnable\b`, // _setupRole and safeApprove are common deprecated lib functions
		},
		{
			"L008",
			LOW,
			"The nonReentrant modifier should occur before all other modifiers",
			"This is a best-practice to protect against reentrancy in other modifiers.",
			`\bnonReentrant\b`, // _setupRole and safeApprove are common deprecated lib functions
		},
		{
			"L009",
			LOW,
			"Use of bytes.concat() instead of abi.encodePacked()",
			"Since version 0.8.4 for appending bytes, bytes.concat() can be used instead of abi.encodePacked(,)",
			`\babi.encodePacked\b`, // _setupRole and safeApprove are common deprecated lib functions
		},
		{
			"L010",
			LOW,
			"For modern and more readable code; update import usages",
			"only import what you need Specific imports with curly braces allow us to apply this rule better.",
			`\import\b`, // _setupRole and safeApprove are common deprecated lib functions
		},
		{
			"L011",
			LOW,
			"Use of abi.encodePacked() instead of abi.encode()",
			"...",
			`\encode\b`, // _setupRole and safeApprove are common deprecated lib functions
		},
		{
			"L012",
			LOW,
			"Avoid using tx.origin",
			`tx.origin is a global variable in Solidity that returns the address of the account that sent the transaction +
			Using the variable could make a contract vulnerable if an authorized account calls a malicious contract. +
			You can impersonate a user using a third party contract.`,
			`tx\.origin\b`, // _setupRole and safeApprove are common deprecated lib functions
		},
	}
}
