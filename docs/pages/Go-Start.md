---
title: Go Start
layout: default
navigation_weight: 9
---
# Go Start

Create a new **My_GO** directory from the Finder in Mac Os High Sierra at the location of your choice.

{% include toc-flammarion.md %}

## Visual Studio Code

Now, fire up a copy of your favorite IDE or text editor.

### Go Version

*File*, *Open* your new **My_GO** directory in **Visual Studio Code**.

From the *View*, *Integrated Terminal* in **Visual Studio Code** ...

- Type the *go version* command at the local command prompt `$`, as follows:

```liquid
{% raw %}
go version
{% endraw %}
```

Returns,

```liquid
{% raw %}
go version go1.9.2 darwin/amd64
{% endraw %}
```

### Bash Profile

In the dot bash underscore profile file `.bash_profile` of your home directory ...

- Add a statement under your **Symlinks** subheading designating the path to your new **My_GO** directory

- By exporting your new settings for the **GOPATH** variable, as follows:

```liquid
{% raw %}
# Symlinks
#
# Add symlink to GOPATH variable
export GOPATH=$HOME/documents/My_GO
# Finished adapting your PATH environment variable
#
{% endraw %}
```

**Note.** The **Go** program will search for your files to execute following the path established via the **GOPATH** variable

### Check It

1. Restart your Terminal window by first clicking the *Kill Terminal* icon located in the upper right-hand corner of the **View Terminal** window in Visual Studio Code

1. Next, select *File*, *View Terminal* once again to reestablish the **View Terminal** window

In order to verify the new path to your **My_GO** directory ...

- Use the **Go** `env` command at the local **My_GO** prompt `$` to check for path accuracy by piping the result through `grep`

```liquid
{% raw %}
env | grep GOPATH
{% endraw %}
```

Returns,

```liquid
{% raw %}
GOPATH=/Users/yourHomeDirectory/documents/My_GO
{% endraw %}
```

**Note**. The previous $HOME entry in bash corresponds to the home directory of your development machine.

### Anchor Test

This is a test sentence for the internal page anchor #2 in red => For: Go Succinctly [[2](#GoSuccinctly){:.red}].

The reference points to the second internal citation [[Syncfusion dot com](#GoSuccinctly){:.red}] below.

By clicking on either the red [2] above, or the red [Syncfusion dot co] ...
{:.red}

The reader will be hyperlinked to the bottom of the page directly to the referenced citation.

### Go Commands

Other commands available in **Go** are listed using the following one-word code statement at the local Terminal prompt `$`, as follows:

```liquid
{% raw %}
go
{% endraw %}
```

Returns,

|-----------------------------+-----------------------|
| Go commands | Description |
|:---------------------------:|:---------------------:|
| build | compile packages and dependencies |
| clean | remove object files |
| doc | show documentation for package or symbol |
| env | print Go environment information |
| bug | start a bug report |
| fix | run go tool fix on packages |
| fmt | run gofmt on package sources |
| generate | generate Go files by processing source |
| get | download and install packages and dependencies |
| install | compile and install packages and dependencies |
| list | list packages |
| run | compile and run Go program |
| test | test packages |
| tool | run specified go tool |
| version | print Go version |
| vet | run go tool vet on packages |

{% include sources-and-uses.md %}

1. {:#GoSuccinctly}Code lines courtesy of [Go Succinctly by Mark Lewin](https://www.syncfusion.com/resources/techportal/details/ebooks/Go_Succinctly){:title="Click to Visit the Go Succinctly page at Syncfusion dot com"}{:target="_blank"}. Published by © 2017 [Syncfusion.com](https://www.syncfusion.com/){:title="Click to Visit Syncfusion dot com"}{:target="_blank"}.

### External Sources

- The [Project Source Links](https://mminail.github.io/Shell/Source-Shell-Links.htm){:title="Click to Visit the Source Links page of the Shell Lessons Project at GitHub pages"}{:target="_blank"} page of the Shell Lessons Project. Published by © 2017 [Mminail.github.io](https://mminail.github.io/){:title="Click to Visit the Concept Library of the Medical Marijuana Initiative of North America - International Limited, an Arizona Benefit Corporation"}{:target="_blank"}.
