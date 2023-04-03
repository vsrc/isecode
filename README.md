# isecode

Tool to populate your code with traceable and secure error codes

## Problem

Essential part of any project, especially customer facing is proper and secure error handling. When error happens and customer reports it, it would be nice to know the context of the error and where it exactly occured. For example, if user tried to login and we tried to fetch a record about this user from the db, there is a number of things that could go wrong. Maybe db server is down, maybe number of connections is exhausted, maybe user deleted their account previously, maybe somebody accessed the db and altered the data. For each of those situations your code can do a check but for a security reasons you will never disclose this information to a user. For the same reason, you wouldn't return stack trace to the user as it could give confidential information aobut your system to the end user which can be used to breach the infrastructure.

Instead it's very often to return a cryptic error like:

`"Something went wrong, please contact us."`

When trying to debug the problem and help the customer, it's very hard to troubleshoot given the above (on purpose) non-verbose error. Customer might explain what they tried to do and what was the issue but that information is not always complete and reliable. As explained in the example above it wouldn't exactly pinpoint the underlying reason for why customer is facing the issue.

It would be nice to be able to pinpoint where problem occurs to the specific line of the code.

## Solution

I made this simple tool to help with this problem. Whenver you find a line of code where you need to return an error but can't explain the user why error happend (for example, Internal Server Error), you can add predefined string like `"ISE_CODE"` and run this cli tool. Tool will search all the matching files in given path containing your code and replace the string `"ISE_CODE"` with a unique number. So if you wrote a code that returns an error with following text:

`"ERROR ISE_CODE: Something went wrong, please contact us."`

this tool will find all the occurences of the `"ISE_CODE"` string and replace it with the unique numbers to look like:

`"ERROR 7248152: Something went wrong, please contact us."`

Now if your user reports having an issue, they can give you this error number and you can search your code base for this number and instantly pinpoint the code line where issue occured. All of that without exposing too much information to the user.

## Requirements

This tool is written in golang. It doesn't require golang to run, but it does require golang to compile from source. It also depends on config file (read bellow more about it).

## How to download?

At the moment I haven't compiled the binary release so you will need to do it yourself. Luckily the process is very simple and it requires just one line command:

```sh
go install github.com/vsrc/isecode@[version]
```

For example:

```sh
go install github.com/vsrc/isecode@v1.0.1
```

If you have any problems compiling your binary, you can file an issue and I will try to compile it and release it for your architecture platform.

## How to use?

Simpliest way to use it by copying file `isecode.json` into your project working directory and then running command:

```sh
isecode
```

Optionally you can run the command with following options:

```sh

 # specify path to your project working directory or any directory where this tool should look for files to inject isecode numbers (default to the current directory where you run the command)
-p [VALUE] or --path [VALUE]

# path to the config file (default is isecode.json in the current directory where you run the command)
-c [VALUE] or --config [VALUE]

# show help
-h or --help

```

## How to configure?

This tool relies on being supplied with configuration file which has to be in the json format. For example please take a look at the `isecode.json` provided in this repository. Config file has to contain following settings:

- `"LAST_NUMBER"`: number from which this tool will incrementally add isecode numbers in your code. For each placeholder you put in your code, this tool will replace it with `"LAST_NUMBER"` + 1 and after successfully injecting the number in your code it will overwrite `"LAST_NUMBER"` in your config file with that number.

- `"MATCH_STRING"`: a placeholder string this tool will look for in your code and replace it with a number, for example: `"ISE_CODE"`

- `"MATCH_FILES"`: a regex expression if you want to filter which type of files this tool should look for, for example: `"*.go"` will set to look only for files that end with `.go` extension

- `"CODE_PREFIX"`: an optional parameter that can be appended to every isecode number, useful for defining a scope of the errors.


## What isecode stands for?

An acronym for Internal Server Error Code. Not imaginative name, I agree but if you have suggestions feel free to create an issue and I will take a look.

## TODO Whishlist

These are the things I would like to improve this project with. If you want to suggest any other feel free to file an issue or contact me. If you want to improve this code feel free to submit a PR, I will gladly take a look.

- [ ] Set default options for config parameters making it optional to have configuration file
- [ ] If there is enough interest, add a watcher so that this tool can optionally be a long running process, injecting isecode numbers on file save
- [x] ~~If there is enough interest, add a format option for the isecode injection so if you want to inject specific error message beside the code number everyhere it occurs, you don't have to do it manually~~ resolved by adding the option for prefix in the v1.0.1, it accepts any string
- [] tests

## Updates

- v1.0.1 added option in the configuration file to add a code prefix which will prepend error codes with whatever string you put there. Useful if you use this tool on multiple libraries and you import them into the same project. Thanks to [@yangjuncode](https://github.com/yangjuncode) and [@pycckuu](https://github.com/pycckuu) for suggesting this in the issue [#1](https://github.com/vsrc/isecode/issues/1)
