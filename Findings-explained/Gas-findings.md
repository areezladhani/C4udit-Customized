# Gas Findings

## G001 - Don't Initialize Variables with Default Value

**Category**: GASOP

### Description

Initializing variables with their default values is redundant and can lead to unnecessary gas consumption. In Solidity, variables have default values when they are declared but not initialized. For instance, uint and uint256 variables have a default value of 0, while bool variables have a default value of false. Explicitly setting these default values when declaring a variable increases the contract's bytecode size, leading to higher deployment and execution costs.

### Example

Instead of initializing a variable with a default value like this:

```solidity
uint256 public counter = 0;
bool public isActive = false;
```

You should simply declare the variable without initializing it, as the default values are automatically assigned:

```solidity
uint256 public counter;
bool public isActive;
```

#### Steps to mitigate/change the issue

Identify the variables that are explicitly initialized with default values.
Remove the explicit initialization and let the variables take their default values implicitly.
Test the contract to ensure that the behavior remains as expected after removing the explicit initializations.

## G002 - Cache Array Length Outside of Loop

**Category**: GASOP

### Description

Caching the length of an array outside the loop can optimize gas usage, as the function only needs to retrieve the value once and then use the cached value within the loop. This optimization can result in significant gas savings for functions that iterate over large arrays.

### Example

Instead of using the array length directly in the loop condition like this:

```solidity
for (uint i = 0; i < myArray.length; i++) {
    // Perform some operation
}
```

You should cache the array length outside the loop and use the cached value:

```solidity
uint256 arrayLength = myArray.length;
for (uint i = 0; i < arrayLength; i++) {
// Perform some operation
}
```

#### Steps to mitigate/change the issue

Identify loops where the array length is used directly in the loop condition.
Cache the array length in a separate variable outside the loop.
Replace the direct use of the array length in the loop condition with the cached value.
Test the contract to ensure that the behavior remains as expected after making the change.

## G003 - Use != 0 instead of > 0 for Unsigned Integer Comparison

**Category**: GASOP

### Description

When dealing with unsigned integer types, comparisons using `!= 0` are cheaper in terms of gas than using `> 0`.

### Example

Instead of using `> 0` for unsigned integer comparison like this:

```solidity
uint256 value = 42;
if (value > 0) {
    // Perform some operation
}
```

You should use != 0:

```solidity
uint256 value = 42;
if (value != 0) {
    // Perform some operation
}
```

#### Steps to mitigate/change the issue

Identify instances where unsigned integers are compared using > 0.
Replace the > 0 comparison with != 0.
Test the contract to ensure that the behavior remains as expected after making the change.

## G004 - Use immutable for OpenZeppelin AccessControl's Roles Declarations

**Category**: GASOP

### Description

⚡️ Only valid for Solidity versions <0.6.12 ⚡️

Access roles marked as `constant` result in computing the `keccak256` operation each time the variable is used. This is because assigned operations for `constant` variables are re-evaluated every time. Changing the variables to `immutable` results in computing the hash only once on deployment, leading to gas savings.

### Example

Instead of using `constant` for OpenZeppelin AccessControl's roles declarations like this:

```solidity
bytes32 public constant ROLE = keccak256("MY_ROLE");
```

You should use immutable:

```solidity
bytes32 public immutable ROLE = keccak256("MY_ROLE");
```

### Steps to mitigate/change the issue

Identify instances where OpenZeppelin AccessControl's roles are declared using constant.
Replace the constant keyword with immutable.
Test the contract to ensure that the behavior remains as expected after making the change.

## G005 - Long Revert Strings

**Category**: GASOP

### Description

For the output on a `require`, each extra memory word of bytes past the original 32 incurs an `MSTORE` operation which costs 3 gas.

### Example

Instead of using a long revert string like this:

```solidity
require(someCondition, "This is a very long revert string that consumes more gas than necessary.");
```

You should use a shorter revert string:

```solidity
require(someCondition, "Short revert string.");
```

### Steps to mitigate/change the issue

Identify instances where long revert strings (33 characters or more) are used in require statements.
Replace the long revert strings with shorter, more concise descriptions.
Test the contract to ensure that the behavior remains as expected after making the change.

## G006 - Use Shift Right/Left instead of Division/Multiplication if possible

**Category**: GASOP

### Description

`<x> / 2` is the same as `<x> >> 1`. While the compiler uses the `SHR` opcode to accomplish both, the version that uses division incurs an overhead of 20 gas due to `JUMPs` to and from a compiler utility function that introduces checks which can be avoided by using `unchecked {}` around the division by two.

### Example

Instead of using division or multiplication:

```solidity
uint256 result = value / 2;
uint256 result2 = value * 4;
```

You should use shift right/left operations:

```solidity
uint256 result = value >> 1;
uint256 result2 = value << 2;
```

### Steps to mitigate/change the issue

Identify instances where division or multiplication by 2, 4, or 8 are used in the code.
Replace these instances with the corresponding shift right/left operations.
Test the contract to ensure that the behavior remains as expected after making the change.

## G007 - ++i costs less gas than i++, especially when used in for-loops (--i/i-- too)

**Category**: GASOP

### Description

Using `++i` or `--i` (pre-increment/decrement) is generally more gas-efficient than using `i++` or `i--` (post-increment/decrement), especially when used in for-loops.

### Example

Instead of using post-increment/decrement:

```solidity
for (uint256 i = 0; i < array.length; i++) {
    // ...
}

for (uint256 i = array.length; i > 0; i--) {
    // ...
}
```

You should use pre-increment/decrement:

```solidity
for (uint256 i = 0; i < array.length; ++i) {
    // ...
}

for (uint256 i = array.length; i > 0; --i) {
    // ...
}
```

### Steps to mitigate/change the issue

Identify instances where i++, i--, ++i, or --i are used in the code.
Replace i++ and i-- with ++i and --i, respectively.
Test the contract to ensure that the behavior remains as expected after making the change.

## G008 - Use Custom Errors rather than revert()/require() strings to save gas

**Category**: GASOP

### Description

From Solidity version 0.8.4 onwards, custom errors can be used instead of revert() and require() strings to save gas. Custom errors save approximately 50 gas each time they are hit, by avoiding the need to allocate and store the revert string. Additionally, not defining the strings saves deployment gas.

### Example

Instead of using revert()/require() strings:

```solidity
require(isValid, "Invalid input");
```

You should use custom errors:

```solidity
error InvalidInput();

function example() public {
    if (!isValid) {
        revert InvalidInput();
    }
}
```

### Steps to mitigate/change the issue

Identify instances where revert() or require() with strings are used in the code.
Replace them with custom error declarations and revert statements with the custom error names.
Test the contract to ensure that the behavior remains as expected after making the change.

## G009 - Address 0 check can be done in assembly to save gas

**Category**: GASOP

### Description

In Solidity, checking if an address is 0 is commonly done using the syntax `address(0)`. However, this operation is expensive in terms of gas usage. Instead, this check can be done more efficiently using inline assembly.

### Example

Instead of using the `address(0)` syntax:

```solidity
require(receiver != address(0), "Receiver address cannot be 0");
```

You should use inline assembly:

```solidity
function isZeroAddress(address addr) private pure returns (bool) {
    assembly {
        return iszero(eq(addr, 0))
    }
}

require(!isZeroAddress(receiver), "Receiver address cannot be 0");
```

### Steps to mitigate/change the issue

Identify instances where address(0) is used in the code.
Replace them with an inline assembly function to perform the 0 address check more efficiently.
Test the contract to ensure that the behavior remains as expected after making the change.

## G010 - Using private rather than public for constants, saves gas

**Category**: GASOP

### Description

Public constants generate a getter function, which increases the bytecode size and gas usage. Using private rather than public for constants can save gas.

### Example

Instead of using a public constant:

```solidity
contract Example {
    uint256 public constant RATE = 100;
}
```

You should use a private constant:

```solidity
contract Example {
    uint256 private constant RATE = 100;
}
```

### Steps to mitigate/change the issue

Identify instances where public constants are used in the code.
Replace them with private constants.
Test the contract to ensure that the behavior remains as expected after making the change.

## G011 - Functions guaranteed to revert when called by normal users can be marked payable

**Category**: GASOP

### Description

Marking the function as payable will lower the gas cost for legitimate callers because the compiler will not include checks for whether a payment was provided.

### Example

Instead of having a non-payable function with access control:

```solidity
contract Example {
    modifier onlyOwner {
        require(msg.sender == owner, "Not the owner");
        _;
    }

    function restrictedFunction() public onlyOwner {
        // Do something
    }
}
```

You can mark the function as payable to reduce gas costs:

```solidity
contract Example {
    modifier onlyOwner {
        require(msg.sender == owner, "Not the owner");
        _;
    }

    function restrictedFunction() public payable onlyOwner {
        // Do something
    }
}
```

### Steps to mitigate/change the issue

Identify functions that are guaranteed to revert when called by normal users due to access control modifiers (e.g., onlyOwner, onlyAdmin).
Mark these functions as payable.
Test the contract to ensure that the behavior remains as expected after making the change.

## G012 - Setting the constructor to payable

**Category**: GASOP

### Description

You can cut out 10 opcodes in the creation-time EVM bytecode if you declare a constructor payable. Making the constructor payable eliminates the need for an initial check of msg.value == 0 and saves 13 gas on deployment with no security risks.

### Example

Instead of having a non-payable constructor:

```solidity
contract Example {
    uint256 public value;

    constructor(uint256 _value) public {
        value = _value;
    }
}
```

You can mark the constructor as payable to reduce gas costs:

```solidity
contract Example {
    uint256 public value;

    constructor(uint256 _value) public payable {
        value = _value;
    }
}
```

### Steps to mitigate/change the issue

Identify the constructor in your contract.
Mark the constructor as payable.
Test the contract to ensure that the behavior remains as expected after making the change.

## G013 - Use 1e18 instead of 10 \*\* 18 to save gas

**Category**: GASOP

### Description

10 \*\* 18 can be changed to 1e18 to avoid unnecessary arithmetic operation and save gas.

### Example

Instead of using the expression `10 ** 18`:

```solidity
contract Example {
    uint256 public constant weiPerEther = 10 ** 18;

    function convertToWei(uint256 _ethers) public pure returns (uint256) {
        return _ethers * weiPerEther;
    }
}
```

You can use 1e18 to reduce gas costs:

```solidity
contract Example {
    uint256 public constant weiPerEther = 1e18;

    function convertToWei(uint256 _ethers) public pure returns (uint256) {
        return _ethers * weiPerEther;
    }
}
```

### Steps to mitigate/change the issue

Identify occurrences of 10 \*\* 18 in your contract.
Replace them with 1e18.
Test the contract to ensure that the behavior remains as expected after making the change.

## G014 - Assign keccak operations to immutable variables instead of constant variables to save gas

**Category**: GASOP

### Description

Assigning keccak operations to constant variables results in extra gas costs. Constants expressions are expressions, and as such, keccak assigned to a constant variable are re-computed at each use of the variable, which will consume gas unnecessarily. To save gas, change the variable from constant to immutable, making the computation run only once and therefore save gas.

### Example

Instead of using a constant variable for keccak operation:

```solidity
contract Example {
    bytes32 constant ROLE_ADMIN = keccak256("ROLE_ADMIN");

    function isAdmin(address _account) public view returns (bool) {
        return hasRole(ROLE_ADMIN, _account);
    }
}
```

You can use an immutable variable to reduce gas costs:

```solidity
contract Example {
    bytes32 immutable ROLE_ADMIN = keccak256("ROLE_ADMIN");

    function isAdmin(address _account) public view returns (bool) {
        return hasRole(ROLE_ADMIN, _account);
    }
}
```

### Steps to mitigate/change the issue

Identify occurrences of keccak operations assigned to constant variables in your contract.
Replace the constant keyword with immutable.
Test the contract to ensure that the behavior remains as expected after making the change.

## G015 - Use x = x + y instead of x += y for state variables to save gas

**Category**: GASOP

### Description

Using `x += y` for state variables costs more gas than using `x = x + y`. This is because the `+=` and `-=` operators require extra gas for updating the state variable compared to the standard assignment operator.

### Example

Instead of using `+=` for state variables:

```solidity
contract Example {
    uint256 public counter;

    function increment(uint256 value) public {
        counter += value;
    }
}
```

You can use x = x + y to reduce gas costs:

```solidity
contract Example {
    uint256 public counter;

    function increment(uint256 value) public {
        counter = counter + value;
    }
}
```

### Steps to mitigate/change the issue

Identify occurrences of += or -= operators used for state variables in your contract.
Replace the += or -= operators with the standard assignment operator, like x = x + y or x = x - y.
Test the contract to ensure that the behavior remains as expected after making the change.

## G016 - Splitting require() statements that use && saves gas

**Category**: GASOP

### Description

Splitting `require()` statements that use the `&&` operator to check multiple conditions can save gas. By using multiple `require()` statements with one condition per `require()` statement, you can save 8 gas per `&&`. The gas difference is only realized if the revert condition is met.

### Example

Instead of using a single `require()` statement with multiple conditions:

```solidity
contract Example {
    function withdraw(uint256 amount, address recipient) public {
        require(amount > 0 && recipient != address(0), "Invalid amount or recipient");
        // ...
    }
}
```

You can split the require() statements to save gas:

```solidity
contract Example {
    function withdraw(uint256 amount, address recipient) public {
        require(amount > 0, "Invalid amount");
        require(recipient != address(0), "Invalid recipient");
        // ...
    }
}
```

### Steps to mitigate/change the issue

Identify occurrences of require() statements using the && operator in your contract.
Split each require() statement containing multiple conditions into multiple require() statements with one condition per statement.
Test the contract to ensure that the behavior remains as expected after making the change.

## G017 - Use unchecked keyword for ++i/i++ when overflow is not possible

**Category**: GASOP

### Description

When using Solidity version 0.8.0 or higher, use the `unchecked` keyword with `++i` or `i++` when it is not possible for them to overflow. This can save 30-40 gas per loop.

### Example

Instead of using a loop without the `unchecked` keyword:

```solidity
contract Example {
    uint256 public counter = 0;

    function increment() public {
        counter++;
    }
}
```

You can use the unchecked keyword when it is not possible for the counter to overflow:

```solidity
contract Example {
    uint256 public counter = 0;

    function increment() public {
        unchecked {
            counter++;
        }
    }
}
```

### Steps to mitigate/change the issue

Identify occurrences of ++i or i++ in your contract where an overflow is not possible.
Use the unchecked keyword for these occurrences.
Test the contract to ensure that the behavior remains as expected after making the change.

## G018 - Don't compare boolean expressions to boolean literals

**Category**: GASOP

### Description

Don't compare boolean expressions to boolean literals (`true` or `false`). Instead, use the expression directly or negate it if necessary. This can save gas by simplifying the condition.

### Example

Instead of comparing a boolean expression to a boolean literal:

```solidity
contract Example {
    bool public isActive = true;

    function checkStatus() public view returns (bool) {
        if (isActive == true) {
            return true;
        }
        return false;
    }
}
```

Use the expression directly or negate it if necessary:

```solidity
contract Example {
    bool public isActive = true;

    function checkStatus() public view returns (bool) {
        if (isActive) {
            return true;
        }
        return false;
    }
}
```

### Steps to mitigate/change the issue

Identify occurrences of boolean expressions being compared to boolean literals in your contract.
Remove the comparison to the boolean literal and use the expression directly, or negate it if necessary.
Test the contract to ensure that the behavior remains as expected after making the change.
