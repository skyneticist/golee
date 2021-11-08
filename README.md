# git-golee

cli tool for git productivity

[![Go Report Card](https://goreportcard.com/badge/github.com/skyneticist/golee)](https://goreportcard.com/report/github.com/skyneticist/golee)

This cli tool is intended to make frequent git work in the command line even less time-consuming.

### Without git-golee
```
> git add .
> git commit -m "commit message goes here"
> git push --set-upstream origin branchName
```

### With git-golee 
```
> golee acpr "commit message goes here"
```


<img src="https://storage.googleapis.com/gopherizeme.appspot.com/gophers/c31240b190ee8485d98aff995b22f8cc4edc8d10.png" width="250" height="250">
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

Of course, you can clone the repository and set up the path for the included exe or do:

`go install` in this directory: `./cmd/net-helper/`

That's it. Now you can run golee using:

`> golee`


## Simpler Install [WIP!]

This section explains how to simply install the git-golee tool using node package registry (npm). This is the preferred method! 


`npm i -g @skyneticist/git-golee@latest`

That's it. Now you have access to git-golee via `golee`

One thing that IS nice about this tool is how easy it is to install and to keep updated. Also, the size of the executable is small and the way its being added to PATH is fast, safe, convenient.

To achieve all of this, Goreleaser and a custom postinstall script is used in combination with a package.json file. This essentially allows for a platform-independent way to install the git-golee tool with a simple npm command.

[Packgage on NPM](https://www.npmjs.com/package/@skyneticist/git-golee)

*Currently working on a way to set the command to `gg` instead of `golee` without setting an alias*

# Usage 

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


With git-golee, we can do the same with:
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

With git-golee:

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

  
