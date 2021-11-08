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
> gg acpr "commit message goes here"
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
### npm install is work in progress
Of course, you can clone the repository and set up the path for the included exe or do:

`go install` in this directory: `./cmd/net-helper/`

That's it. Now you can run golee using:

`> net-helper`

NOTE: The name `net-helper` is temporary and will be `gg` in the future.


## Simpler Install [WIP!]

This section explains how to simply install the git-golee tool using node package registry (npm). Please note, this is currently being developed and does not work as intended.
Will update readme when able.

`npm i -g @skyneticist/git-golee@latest`


That's it. Now you have access to git-golee via `net-helper`

One thing that IS nice about this tool is how easy it is to install and to keep updated. Also, the size of the executable is small and the way its being added to PATH is fast, safe, convenient.

To achieve all of this, Goreleaser and a custom postinstall script is used in combination with a package.json file. This essentially allows for a platform-independent way to install the git-golee tool with a simple npm command.

[Packgage on NPM](https://www.npmjs.com/package/@skyneticist/git-golee)

*Currently working on a way to set the command to `gg` instead of `net-helper` without setting an alias*

# Usage 

`net-helper`

`gg`

One of the most often set of commands used when working on a user story is:
  
  `git add .`
  
  `git commit -m "message here"`
  
  `git push`
  
With git-golee, we can do the same with:
 
 `net-helper acp "commit message here"`
 
 or
 
 `gg acp "commit message here"`
  
Similarly, when pushing a fresh remote, we usually would do:

  `git add .`
 
  `git commit -m "message here"`
  
  `git push --set-upstream origin <branch>` or `git push -u origin HEAD`
  
With git-golee:

  `net-helper acpr "commit message here"`
  
or

  `gg acpr "commit message here"`
  
Renaming a branch can be a bit burdensome:
  
  ``

Alternatively, we could do:

  `net-helper rn <updatedBranchName>`
  
or
  
  `gg rn <updatedBranchName>`
  
