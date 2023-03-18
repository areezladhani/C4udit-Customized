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
			"G001",
			GASOP,
			"Don't Initialize Variables with Default Value",
			"https://github.com/areezladhani/C4udit-Customized/blob/main/Findings-explained/Gas-findings.md#g001---dont-initialize-variables-with-default-value",
			`\b(uint\d{0,3}|uint256|bool)\s+(\w+)\s*=\s*(0|false|address\(0\)|""|'')\s*(;.*)?$`,
		},
		// G002 - Cache Array Length Outside of Loop
		{
			"G002",
			GASOP,
			"Cache Array Length Outside of Loop",
			`https://github.com/areezladhani/C4udit-Customized/blob/main/Findings-explained/Gas-findings.md#g002---cache-array-length-outside-of-loop`,
			`for\s*\((.*;)?\s*.*(<|>|<=|>=)\s*[^;]+\.length\s*;`,
		},
		// G003 - Use != 0 instead of > 0 for Unsigned Integer Comparison
		{
			"G003",
			GASOP,
			"Use != 0 instead of > 0 for Unsigned Integer Comparison",
			"https://github.com/areezladhani/C4udit-Customized/blob/main/Findings-explained/Gas-findings.md#g003---use--0-instead-of--0-for-unsigned-integer-comparison",
			`\s*> ?0`,
		},
		// G006 - Use immutable for OpenZeppelin AccessControl's Roles Declarations
		{
			"G004",
			GASOP,
			"Use immutable for OpenZeppelin AccessControl's Roles Declarations",
			`https://github.com/areezladhani/C4udit-Customized/blob/main/Findings-explained/Gas-findings.md#g004---use-immutable-for-openzeppelin-accesscontrols-roles-declarations`,
			`constant.*keccak256`,
		},
		// G007 - Long Revert Strings
		{
			"G005",
			GASOP,
			"Long Revert Strings",
			"https://github.com/areezladhani/C4udit-Customized/blob/main/Findings-explained/Gas-findings.md#g005---long-revert-strings",
			`require\(.+\".{33,}\"`, // Anything in a require and between "'s with at least 33 characters
		},
		// G008 - Use Shift Right/Left instead of Division/Multiplication if possible
		{
			"G006",
			GASOP,
			"Use Shift Right/Left instead of Division/Multiplication if possible",
			`https://github.com/areezladhani/C4udit-Customized/blob/main/Findings-explained/Gas-findings.md#g006---use-shift-rightleft-instead-of-divisionmultiplication-if-possible`,
			`(/[2,4,8]|/ [2,4,8]|\*[2,4,8]|\* [2,4,8])`, // Matches division or multiplication by 2, 4, or 8
		},
		{
			"G007",
			GASOP,
			"++i costs less gas than i++, especially when its used in for-loops (--i/i-- too)",
			"https://github.com/areezladhani/C4udit-Customized/blob/main/Findings-explained/Gas-findings.md#g007---i-costs-less-gas-than-i-especially-when-used-in-for-loops---ii---too",
			`\+\+|--`,
		},
		{
			"G008",
			GASOP,
			"USE CUSTOM ERRORS RATHER THAN REVERT()/REQUIRE() STRINGS TO SAVE GAS",
			`https://github.com/areezladhani/C4udit-Customized/blob/main/Findings-explained/Gas-findings.md#g008---use-custom-errors-rather-than-revertrequire-strings-to-save-gas`,
			`\b(require|assert)\b`,
		},
		{
			"G009",
			GASOP,
			"Address 0 check can be done in assembly to save gas",
			`https://github.com/areezladhani/C4udit-Customized/blob/main/Findings-explained/Gas-findings.md#g009---address-0-check-can-be-done-in-assembly-to-save-gas`,
			`address\s*\(\s*0[x0]*\s*\)`,
		},
		{
			"G010",
			GASOP,
			"Using private rather than public for constants, saves gas",
			"https://github.com/areezladhani/C4udit-Customized/blob/main/Findings-explained/Gas-findings.md#g010---using-private-rather-than-public-for-constants-saves-gas",
			`public\s+constant\s+\S+\s*=.*;`,
		},
		{
			"G011",
			GASOP,
			"Functions guaranteed to revert when called by normal users can be marked payable",
			`https://github.com/areezladhani/C4udit-Customized/blob/main/Findings-explained/Gas-findings.md#g011---functions-guaranteed-to-revert-when-called-by-normal-users-can-be-marked-payable`,
			`\b(?:onlyOwner|onlyAdmin)\b`,
		},
		{
			"G012",
			GASOP,
			"Setting the constructor to payable",
			`https://github.com/areezladhani/C4udit-Customized/blob/main/Findings-explained/Gas-findings.md#g012---setting-the-constructor-to-payable`,
			`\bconstructor\b`,
		},
		{
			"G013",
			GASOP,
			"10 ** 18 can be changed to 1e18 and save some gas",
			"https://github.com/areezladhani/C4udit-Customized/blob/main/Findings-explained/Gas-findings.md#g013---use-1e18-instead-of-10--18-to-save-gas",
			`10\s*\*\*\s*18\b`,
		},
		{
			"G014",
			GASOP,
			"Assigning keccak operations to constant variables results in extra gas costs",
			`https://github.com/areezladhani/C4udit-Customized/blob/main/Findings-explained/Gas-findings.md#g014---assign-keccak-operations-to-immutable-variables-instead-of-constant-variables-to-save-gas`,
			`\b(keccak|keccak256)\b`,
		},
		{
			"G015",
			GASOP,
			"x += y costs more gas than x = x + y for state variables",
			"https://github.com/areezladhani/C4udit-Customized/blob/main/Findings-explained/Gas-findings.md#g015---use-x--x--y-instead-of-x--y-for-state-variables-to-save-gas",
			`[\+\-]=`,
		},
		{
			"G016",
			GASOP,
			"Splitting require() statements that use && saves gas - (saves 8 gas per &&)",
			`https://github.com/areezladhani/C4udit-Customized/blob/main/Findings-explained/Gas-findings.md#g016---splitting-require-statements-that-use--saves-gas`,
			`&&(?!\s*=)`,
		},
		{
			"G017",
			GASOP,
			`++i/i++ should be unchecked{++i}/unchecked{i++} when it is not possible for them to overflow`,
			`https://github.com/areezladhani/C4udit-Customized/blob/main/Findings-explained/Gas-findings.md#g017---use-unchecked-keyword-for-ii-when-overflow-is-not-possible`,
			`\+\+|--`,
		},
		{
			"G018",
			GASOP,
			`Don’t compare boolean expressions to boolean literals`, 
			"https://github.com/areezladhani/C4udit-Customized/blob/main/Findings-explained/Gas-findings.md#g018---dont-compare-boolean-expressions-to-boolean-literals",
			`== (true|false)\b`,
		},
		{
			"G019",
			GASOP,
			"Avoid using uints/ints smaller than 32 bytes (256 bits)",
			"https://github.com/areezladhani/C4udit-Customized/blob/main/Findings-explained/Gas-findings.md#g020---avoid-using-uintsints-smaller-than-32-bytes-256-bits",
			`\b(uint|int)(8|16|24|32|40|48|56|64)?(\[\])?\b(?<!uint256)(?<!uint\b)`,
		},
		{
			"G020",
			GASOP,
			"Use caution with initializers to avoid front-running attacks",
			"https://github.com/areezladhani/C4audit-Customized/blob/main/Findings-explained/Gas-findings.md#g021---use-caution-with-initializers-to-avoid-front-running-attacks",
			`\binitialize\b`,
		},
		{
			"G021",
			GASOP,
			"Event is missing indexed fields",
			"https://github.com/areezladhani/C4audit-Customized/blob/main/Findings-explained/Gas-findings.md#g022---event-is-missing-indexed-fields",
			`\bevent\b`,
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
			"https://github.com/areezladhani/C4udit-Customized/blob/main/Findings-explained/Low-Findings.md#l001---unsafe-erc20-operations",
			`\.transfer\(|\.transferFrom\(|\.approve\(`, // ".tranfer(", ".transferFrom(" or ".approve("
		},
		// L003 - Unspecific Compiler Version Pragma
		{
			"L002",
			LOW,
			"Unspecific Compiler Version Pragma",
			"https://github.com/areezladhani/C4udit-Customized/blob/main/Findings-explained/Low-Findings.md#l002---unspecific-compiler-version-pragma",
			"pragma solidity (\\^|>)", // "pragma solidity ^" or "pragma solidity >"
		},
		// L005 - Do not use Deprecated Library Functions
		{
			"L003",
			LOW,
			"Do not use Deprecated Library Functions",
			`https://github.com/areezladhani/C4udit-Customized/blob/main/Findings-explained/Low-Findings.md#l003---do-not-use-deprecated-library-functions`,
			`_setupRole\(|safeApprove\(`, // _setupRole and safeApprove are common deprecated lib functions
		},
				// L005 - Do not use Deprecated Library Functions
		{
			"L004",
			LOW,
			"Use Ownable2StepUpgradeable instead of OwnableUpgradeable contract",
			`https://github.com/areezladhani/C4udit-Customized/blob/main/Findings-explained/Low-Findings.md#l004---use-ownable2stepupgradeable-instead-of-ownableupgradeable-contract`,
			`\bOwnableUpgradeable\b`, // _setupRole and safeApprove are common deprecated lib functions
		},
		{
			"L005",
			LOW,
			"Use safeTransferOwnership instead of transferOwnership function",
			"https://github.com/areezladhani/C4udit-Customized/blob/main/Findings-explained/Low-Findings.md#l005---use-safetransferownership-instead-of-transferownership-function",
			`\bOwnable\b`, // _setupRole and safeApprove are common deprecated lib functions
		},
		{
			"L006",
			LOW,
			"The nonReentrant modifier should occur before all other modifiers",
			"https://github.com/areezladhani/C4udit-Customized/blob/main/Findings-explained/Low-Findings.md#l006---the-nonreentrant-modifier-should-occur-before-all-other-modifiers",
			`\bnonReentrant\b`, // _setupRole and safeApprove are common deprecated lib functions
		},
		{
			"L007",
			LOW,
			"Use of bytes.concat() instead of abi.encodePacked()",
			"https://github.com/areezladhani/C4udit-Customized/blob/main/Findings-explained/Low-Findings.md#l007---use-of-bytesconcat-instead-of-abiencodepacked",
			`\babi.encodePacked\b`, // _setupRole and safeApprove are common deprecated lib functions
		},
		{
			"L008",
			LOW,
			"For modern and more readable code; update import usages",
			"https://github.com/areezladhani/C4udit-Customized/blob/main/Findings-explained/Low-Findings.md#l008---update-import-usages-for-modern-and-more-readable-code",
			`\import\b`, // _setupRole and safeApprove are common deprecated lib functions
		},
		{
			"L009",
			LOW,
			"Use of abi.encodePacked() instead of abi.encode()",
			"https://github.com/areezladhani/C4udit-Customized/blob/main/Findings-explained/Low-Findings.md#l009---use-abiencodepacked-instead-of-abiencode",
			`\encode\b`, // _setupRole and safeApprove are common deprecated lib functions
		},
		{
			"L010",
			LOW,
			"Avoid using tx.origin",
			`https://github.com/areezladhani/C4udit-Customized/blob/main/Findings-explained/Low-Findings.md#l010---avoid-using-txorigin`,
			`tx\.origin\b`, // _setupRole and safeApprove are common deprecated lib functions
		},
	}
}
