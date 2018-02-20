---
title: Go Start
layout: default
excerpt: The Go language infers the type of static variable created  ...
version: Page Template md Dtd 02-16-18
navigation_weight: 8
categories: template
---
# {{ page.title }}

{{ page.excerpt }}

{% include toc.md %}

## Visual Studio Code

- Create a new **My_GO** directory from the Finder in Mac Os High Sierra at the location of your choice.

- Now, fire up a copy of your favorite IDE or text editor.

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

This is a test sentence for the internal page anchor #2 in red => For: Go Succinctly [[2](#GOSUCCINCTLY){:.red}].

The reference points to the second internal citation [[Syncfusion dot com](#GOSUCCINCTLY){:.red}] below.

By clicking on either the red [2] above, or the red [Syncfusion dot com] ...
{:.red}

The reader will be hyperlinked to the bottom of the page directly to the referenced citation.

### Go Commands

Other commands available in **Go** are listed below using the following one-word code statement at the local Terminal prompt `$`, as follows:

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

**Note**. The columns in the above table are set to be *centered* via Kramdown.

Therefore, both the title row and the data rows should all display as *centered* within their respective columns.

## Main dot go

Your *main dot go* file should reside in the subdirectory of your instance under the `src` subdirectory of your **My_GO** directory.

**Note**. To move down a level `..` using the bash *change directory* `cd` command, run the following command from your local Terminal `$` prompt, as follows:

```liquid
{% raw %}
cd ..
{% endraw %}
```

### The Go file format

The **Go** file format starts with a `package main` statement and is followed by your `import` statement, if any, and your *main function* `func main`, as follows:

```liquid
{% raw %}
package main

import (
    "fmt"
)

// Original main function given
func main() {
    fmt.Printf("Hello there, World!\n")
}

// Note. You may declare and assign static variables regular way in Go ( var, name, type ), as follow:

func main() {
    // Declare and assign variables
    var message string = (`Hello there, World!`)
    var newline string = ("\n")

    // Declare methods
    fmt.Printf(message + newline)
}
{% endraw %}
```

### Execution

Create the following three subdirectories inside your **My_GO** directory ...

1. `bin`

1. `pkg`

1. `src`

Under `src` place your *main dot go** program inside a `hello` subdirectory

### Change Directories

Invoke the bash *change directory* `cd` command

- From the local Terminal prompt `$` of your **My_GO** directory

- `cd` from your **My_GO** directory to the `src` subdirectory, as follows:

```liquid
{% raw %}
cd src
{% endraw %}
```

Next, `cd` once again from the `src` subdirectory to the `hello` subdirectory, as follows:

```liquid
{% raw %}
cd hello
{% endraw %}
```

From the `hello` subdirectory of your `src` subdirectory ( within your **My_GO** directory ) ...

- Run the following command from the local Terminal prompt `$` to execute your program

```liquid
{% raw %}
go run main.go
{% endraw %}
```

Returns,

```liquid
{% raw %}
Hello there, World!
{% endraw %}
```

## Go-static operator

The *Go-static* operator `:=` condenses the declaration and assignment steps when creating variables, constants, and other features of the **Go** programming language.

Recall from above the *regular way* declaration and assignment of the `message` and `newline` static variables, both of the `string` type, as follows:

```liquid
{% raw %}
// Declare and assign variables
    var message string = (`Hello there, World!`)
    var newline string = ("\n")
{% endraw %}
```

**Note.** The declaration and assignment steps are separated by a single `=` sign when declaring and assigning a static variable *regular way* in **Go**.

### Consolidation

The *Go-static* operator `:=` allows for the consolidation of the declaration and assignment steps of a static variable, as follows:

```liquid
{% raw %}
// Create variables using the Go-static operator
    message := `Hello there, World!`
    newline := "\n"
{% endraw %}
```

**Note**. You can even kiss the parenthesis `(...)` goodbye, as well, when creating a variable using the *Go-static* operator `:=`

### Demonstration

Let's now plug that new-fangled set of static variables created with the *Go-static* operator `:=` into our main `func`, run the program, and check the results, as follows:

```liquid
{% raw %}
package main

import (
    "fmt"
)

func main() {
    // Create variables using the Go-static operator
    message := `Hello there, World!`
    newline := "\n"

    // Declare methods
    fmt.Printf(message + newline)
}
{% endraw %}
```

Returns,

```liquid
{% raw %}
Hello there, World!
{% endraw %}
```

**Note**. The **Go** language infers the type of static variable created from the input given by the programmer on the right-side of the *Go-static* operator `:=`

## Import Code

More to come ...

## Last Subtitle

**Note**. The above synopsis was derived from an article written by Blank [[2](#BLANK){:.red}].

### Raw Code Block

```liquid
{% raw %}
Enjoy the successful output!
{% endraw %}
```

{% include sources-and-uses.md %}

1. {:#GOSUCCINCTLY}Code lines courtesy of [Go Succinctly by Mark Lewin](https://www.syncfusion.com/resources/techportal/details/ebooks/Go_Succinctly){:title="Click to Visit the Go Succinctly page at Syncfusion dot com"}{:target="_blank"}. Published by © 2017 [Syncfusion.com](https://www.syncfusion.com/){:title="Click to Visit Syncfusion dot com"}{:target="_blank"}.

### External Sources

- {:#SOURCELINKS}The [Project Source Links](https://mminail.github.io/Shell/Source-Shell-Links.htm){:title='Click to Visit the Source Links page of the Shell Lessons Project at Concepts Library'}{:target='_blank'} page of the Shell Lessons Project at Concepts Library. Published by © 2017 - 2018 [Mminail.github.io](https://mminail.github.io/){:title='Click to Visit the Home Page of the Concepts Library of the Medical Marijuana Initiative of North America - International Limited, an Arizona Benefit Corporation'}{:target='_blank'}.

**Note**. This page crafted with {{ page.version }}.
