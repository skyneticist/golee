# golee

cli tool for git productivity

*written in go and packaged for consumption via npmjs.com*

[![Go Report Card](https://goreportcard.com/badge/github.com/skyneticist/golee)](https://goreportcard.com/report/github.com/skyneticist/golee)

<img src="https://storage.googleapis.com/gopherizeme.appspot.com/gophers/c31240b190ee8485d98aff995b22f8cc4edc8d10.png" width="250" height="250">

This cli tool is intended to make frequent git work in the command line even less time-consuming.

### Without golee
```
> git add .
> git commit -m "commit message goes here"
> git push --set-upstream origin branchName
```

### With golee 
```
> golee acpr "commit message goes here"
```


<!-- ![customgopher hh](https://storage.googleapis.com/gopherizeme.appspot.com/gophers/c31240b190ee8485d98aff995b22f8cc4edc8d10.png | ) -->

The main goals of this application are:

  ```
  - To ~significantly reduce the time typing when working in git cli, without relying on hooks 
  Note: (since I work on a team for a big company, it's preferred to have some versatility here)

  - Easy to use
  - To be as lightweight and efficient/performant as possible
 
  - Easy to distribute/release
  - Easy to update
  - Easy to work on
  ```
 
# Installing

There are multiple ways to download and install the git-golee cli tool. The preferred method is to use npm via the command-line--a lot of work was put into making
this tool as portable and convenient for others to use on ANY platform! 

Of course, you can clone the repository (or just download the necessary binary) and set the binary up in your environment PATH.

Or 

`go install` in this directory: `./cmd/net-helper/`

That's it. Now you can run golee using:

`> golee`

For the help prompt:

`> golee h`


## Simpler Install [Preferred Method!]

This section explains how to simply install the git-golee tool using node package registry (npm). This is the preferred method! 


`npm i -g @skyneticist/git-golee@latest`

That's it. Now you have access to git-golee via `golee`

note: *some *nix system users may experience permission denied! To remedy this, simply run `chmod +X` on the executable*

One thing that IS nice about this tool is how easy it is to install and to keep updated. Also, the size of the executable is small and the way its being added to PATH is fast, safe, convenient.

To achieve all of this, Goreleaser and a custom postinstall script is used in combination with a package.json file. This essentially allows for a platform-independent way to install the git-golee tool with a simple npm command.

*Please NOTE: there are way too many versions published here (I'm currently working on cleaning it up) - just know that v2.0.0 and later are working!*

[Packgage on NPM](https://www.npmjs.com/package/@skyneticist/git-golee)

*Currently working on a way to set the command to `gg` instead of `golee` without setting an alias*

# Usage 

![image](https://user-images.githubusercontent.com/81132371/140797865-f1f1d778-84d4-4d23-8218-ac327840ba59.png)

*invoke*

`golee`

or

`gg`

*commands*

One of the most often set of commands used when working on a user story is:
```  
  > git add .
  
  > git commit -m "message here"
  
  > git push
```


With golee, we can do the same with:
``` 
 > golee acp "commit message here"
```

or
 
``` 
 > gg acp "commit message here"
```


Similarly, when pushing a fresh remote, we usually would do:
```
  > git add .
 
  > git commit -m "message here"
  
  > git push --set-upstream origin <branch>` or `git push -u origin HEAD
```

With golee:

  `> golee acpr "commit message here"`
  
or

  `> gg acpr "commit message here"`
  
Renaming a branch can be a bit burdensome,
but not with git-golee: 

  `> golee rn <updatedBranchName>`
  
or
  
  `> gg rn <updatedBranchName>`
  

# Other Useful Commands

There are a few other useful commands available for taking care of things like configuration and logging. 

### Set multiple authors globally in git config

`> golee auth <author1> <author2>`

`> gg auth <author1> <author2>`

  
