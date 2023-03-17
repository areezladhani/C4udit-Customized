# Low Findings

## L001 - Unsafe ERC20 Operations

**Category**: LOW

### Description

Unsafe ERC20 operations, such as `transfer()`, `transferFrom()`, and `approve()`, may return false on failure, which could lead to unexpected behavior if not handled correctly. Use safer alternatives such as OpenZeppelin's SafeERC20 library to mitigate this risk.

### Example

Unsafe ERC20 operation:

```solidity
contract Example {
    IERC20 public token;

    function transferToken(address recipient, uint256 amount) public {
        require(token.transfer(recipient, amount), "Token transfer failed");
    }
}
```

Safer alternative using OpenZeppelin's SafeERC20 library:

```solidity
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";


contract Example {
    using SafeERC20 for IERC20;

    IERC20 public token;

    function transferToken(address recipient, uint256 amount) public {
        token.safeTransfer(recipient, amount);
    }
}
```

### Steps to mitigate/change the issue

Identify occurrences of unsafe ERC20 operations in your contract.
Replace them with safer alternatives, such as the SafeERC20 library from OpenZeppelin.
Import and use the SafeERC20 library correctly in your contract.
Test the contract to ensure that the behavior remains as expected after making the change.

## L002 - Unspecific Compiler Version Pragma

**Category**: LOW

### Description

Unspecific compiler version pragma can lead to potential security risks and unexpected behavior when new compiler versions are released. It is recommended to use a fixed version pragma for non-library contracts to ensure consistency and maintainability.

### Example

Unspecific compiler version pragma:

```solidity
pragma solidity ^0.8.0;

contract Example {
    // ...
}
```

Recommended fixed version pragma:

```solidity
pragma solidity 0.8.0;

contract Example {
    // ...
}
```

### Steps to mitigate/change the issue

Identify the occurrences of unspecific compiler version pragma in your contract.
Replace them with a fixed version pragma, corresponding to the compiler version you are using.
Test the contract to ensure that the behavior remains as expected after making the change.

## L003 - Do not use Deprecated Library Functions

**Category**: LOW

### Description

Using deprecated library functions can lead to potential security risks and code maintainability issues. This issue is mostly related to OpenZeppelin libraries. It is recommended to replace deprecated functions with their updated counterparts to ensure the contract remains secure and up-to-date.

### Example

Usage of deprecated library function:

```solidity
import "@openzeppelin/contracts/access/AccessControl.sol";

contract Example is AccessControl {
    function assignRole(address account) public {
        _setupRole(DEFAULT_ADMIN_ROLE, account);
    }
}
```

Recommended usage of updated library function:

```solidity
import "@openzeppelin/contracts/access/AccessControl.sol";

contract Example is AccessControl {
    function assignRole(address account) public {
        grantRole(DEFAULT_ADMIN_ROLE, account);
    }
}
```

### Steps to mitigate/change the issue

Identify the occurrences of deprecated library functions in your contract.
Replace them with their updated counterparts from the latest version of the library.
Test the contract to ensure that the behavior remains as expected after making the change.

## L004 - Use Ownable2StepUpgradeable instead of OwnableUpgradeable contract

**Category**: LOW

### Description

Contracts implementing access control, such as owner roles, should consider implementing a Two-Step Transfer pattern. This pattern helps mitigate the risk of mistakenly transferring ownership to the wrong address, which could result in a loss of control over the contract. Instead of using `OwnableUpgradeable`, it is recommended to use `Ownable2StepUpgradeable` to enforce the Two-Step Transfer pattern.

### Example

Usage of `OwnableUpgradeable` contract:

```solidity
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

contract Example is OwnableUpgradeable {
    // ...
}
```

Recommended usage of Ownable2StepUpgradeable contract:

```solidity
import "path/to/Ownable2StepUpgradeable.sol";

contract Example is Ownable2StepUpgradeable {
    // ...
}
```

### Steps to mitigate/change the issue

Replace the OwnableUpgradeable contract import with Ownable2StepUpgradeable in your contract.
Update the contract inheritance to use Ownable2StepUpgradeable instead of OwnableUpgradeable.
Ensure your contract implementation is compatible with the Two-Step Transfer pattern.
Test the contract to ensure that the behavior remains as expected after making the change.

## L005 - Use safeTransferOwnership instead of transferOwnership function

**Category**: LOW

### Description

It is recommended to use a two-step structure for transferring ownership, which is safer and more secure. The `safeTransferOwnership` function, as part of a two-step transfer process, provides an additional layer of security compared to using the `transferOwnership` function directly. This method helps prevent accidental transfers of ownership to unintended addresses.

### Example

Usage of `transferOwnership` function:

```solidity
import "@openzeppelin/contracts/access/Ownable.sol";

contract Example is Ownable {
    function changeOwner(address newOwner) external onlyOwner {
        transferOwnership(newOwner);
    }
}
```

Recommended usage of safeTransferOwnership function:

```solidity
import "path/to/OwnableWithSafeTransfer.sol";

contract Example is OwnableWithSafeTransfer {
    function changeOwner(address newOwner) external onlyOwner {
        safeTransferOwnership(newOwner);
    }
}
```

### Steps to mitigate/change the issue

Replace the Ownable contract import with a custom OwnableWithSafeTransfer contract that implements the safeTransferOwnership function.
Update the contract inheritance to use OwnableWithSafeTransfer instead of Ownable.
Replace the transferOwnership function calls with safeTransferOwnership function calls.
Implement the necessary logic for the two-step ownership transfer process.
Test the contract to ensure that the behavior remains as expected after making the change.

## L006 - The nonReentrant modifier should occur before all other modifiers

**Category**: LOW

### Description

As a best practice, the `nonReentrant` modifier should be placed before all other modifiers in a function definition to protect against potential reentrancy attacks that could occur in other modifiers.

### Example

Incorrect modifier order:

```solidity
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";

contract Example is ReentrancyGuard {
    function withdraw(uint256 amount) external onlyOwner nonReentrant {
        // Withdraw logic
    }
}
```

Correct modifier order:

```solidity
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";

contract Example is ReentrancyGuard {
    function withdraw(uint256 amount) external nonReentrant onlyOwner {
        // Withdraw logic
    }
}
```

### Steps to mitigate/change the issue

Identify functions where the nonReentrant modifier is not placed before all other modifiers.
Update the function definitions to place the nonReentrant modifier before all other modifiers.
Test the contract to ensure that the behavior remains as expected after making the change.

## L007 - Use of bytes.concat() instead of abi.encodePacked()

**Category**: LOW

### Description

Starting from Solidity version 0.8.4, it is recommended to use `bytes.concat()` for appending bytes instead of `abi.encodePacked()` for better readability and simplicity.

### Example

Using `abi.encodePacked()`:

```solidity
pragma solidity ^0.8.0;

contract Example {
    function concatenateBytes(bytes memory a, bytes memory b) public pure returns (bytes memory) {
        return abi.encodePacked(a, b);
    }
}
```

New approach using bytes.concat():

```solidity
pragma solidity ^0.8.4;

contract Example {
    function concatenateBytes(bytes memory a, bytes memory b) public pure returns (bytes memory) {
        return bytes.concat(a, b);
    }
}
```

### Steps to mitigate/change the issue

Identify instances where abi.encodePacked() is used for concatenating bytes.
Replace abi.encodePacked() with bytes.concat() for improved readability and simplicity.
Test the contract to ensure that the behavior remains as expected after making the change.

## L008 - Update import usages for modern and more readable code

**Category**: LOW

### Description

To make the code more modern and readable, it is recommended to use specific imports with curly braces. This approach ensures that only the required parts of the imported file are used, reducing clutter and improving code readability.

### Example

Older approach with general imports:

```solidity
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract MyToken is ERC20 {
    constructor() ERC20("MyToken", "MTK") {
    }
}
```

New approach with specific imports:

```solidity
pragma solidity ^0.8.0;

import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract MyToken is ERC20 {
    constructor() ERC20("MyToken", "MTK") {
    }
}
```

### Steps to mitigate/change the issue

Identify instances where general imports are used.
Replace general imports with specific imports using curly braces.
Ensure that only the required components are imported, reducing clutter and improving readability.
Test the contract to ensure that the behavior remains as expected after making the change.

## L009 - Use abi.encodePacked() instead of abi.encode()

**Category**: LOW

### Description

In certain cases, using `abi.encodePacked()` instead of `abi.encode()` can provide more efficient and compact encoding of data. It is important to carefully consider the specific use case to ensure that the selected encoding method is appropriate.

### Example

Using `abi.encode()`:

```solidity
pragma solidity ^0.8.0;

contract MyContract {
    function encodeData(uint256 a, address b) public pure returns (bytes memory) {
        return abi.encode(a, b);
    }
}
```

Using abi.encodePacked():

```solidity
pragma solidity ^0.8.0;

contract MyContract {
    function encodeData(uint256 a, address b) public pure returns (bytes memory) {
        return abi.encodePacked(a, b);
    }
}
```

### Steps to mitigate/change the issue

Identify instances where abi.encode() is used.
Analyze the specific use case to determine whether abi.encodePacked() is more appropriate.
Replace abi.encode() with abi.encodePacked() if it provides a more efficient and compact encoding method for the given use case.
Test the contract to ensure that the behavior remains as expected after making the change.

## L010 - Avoid using tx.origin

**Category**: LOW

### Description

`tx.origin` is a global variable in Solidity that returns the address of the account that sent the transaction. Using the variable could make a contract vulnerable if an authorized account calls a malicious contract. It is possible to impersonate a user using a third-party contract. To avoid such security risks, use `msg.sender` instead of `tx.origin`.

### Example

Using `tx.origin`:

```solidity
pragma solidity ^0.8.0;

contract MyContract {
    address public owner;

    constructor() {
        owner = tx.origin;
    }

    function doSomething() public {
        require(tx.origin == owner, "Not authorized");
        // ...
    }
}
```

Using msg.sender:

```solidity
pragma solidity ^0.8.0;

contract MyContract {
    address public owner;

    constructor() {
        owner = msg.sender;
    }

    function doSomething() public {
        require(msg.sender == owner, "Not authorized");
        // ...
    }
}
```

### Steps to mitigate/change the issue

Identify instances where tx.origin is used.
Replace tx.origin with msg.sender to avoid potential security risks.
Test the contract to ensure that the behavior remains as expected after making the change.
