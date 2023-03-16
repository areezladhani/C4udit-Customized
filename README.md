<h1 align=center><code>c4udit</code></h1>

## Introduction

`c4udit` is a static analyzer for solidity contracts based on regular
expressions specifically crafted for [Code4Rena](https://code4rena.com) contests.

It is capable of finding low risk issues and gas optimization documented in the
[c4-common-issues](https://github.com/byterocket/c4-common-issues) repository.

Note that `c4udit` uses [c4-common-issues](https://github.com/byterocket/c4-common-issues)'s issue identifiers.

## Installation

First you need to have the Go toolchain installed. You can find instruction [here](https://go.dev/doc/install).

Then install `c4udit` with:

```
$ go install github.com/byterocket/c4udit@latest
```

To just build the binary:

```
$ git clone https://github.com/byterocket/c4udit
$ cd c4udit/
$ go build .
```

Now you should be able to run `c4udit` with:

```
$ ./c4udit
```

## Configuration

Before using the script, you need to update the `baseURL` and `baseDir` variables in the script file (analyzer.go). This allows the analyzer to generate correct links in the report.

1. Set `baseURL` to the base link of your GitHub repository, followed by `/blob/main/`. For example, if your repository is at `https://github.com/PartyDAO/party-contracts-c4`, set `baseURL` to `https://github.com/PartyDAO/party-contracts-c4/blob/main/`.

2. Set `baseDir` to the base directory of the repository on your local machine. For example, if the repository is located at `/path/to/party-contracts-c4`, set `baseDir` to `party-contracts-c4`.

## Usage

```
Usage:
	c4udit [flags] [files...]

Flags:
	-h    Print help text.
	-s    Save report as file.
```

## Example

Running `c4udit` against the `examples` directory:

```
$ ./c4udit examples/
Files analyzed:
- examples/Test.sol

Issues found:
 G001:
  examples/Test.sol::4 => uint256 a = 0;
  examples/Test.sol::12 => for (uint256 i = 0; i < array.length; i++) {

 G002:
  examples/Test.sol::12 => for (uint256 i = 0; i < array.length; i++) {

 G007:
  examples/Test.sol::6 => string b = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";

 G008:
  examples/Test.sol::13 => i = i / 2;

 L001:
  examples/Test.sol::16 => token.transferFrom(msg.sender, address(this), 100);

 L003:
  examples/Test.sol::1 => pragma solidity ^0.8.0;
```

Using the `-s` flag, `c4udit` will create a report in markdown format.
For an example check out the report in the `examples` directory [here](./examples/c4udit-report.md).

## License

Note that this tool is licensed as [free software](./LICENSE)!
