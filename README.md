# git-golee

cli tool for git productivity

[![Go Report Card](https://goreportcard.com/badge/github.com/skyneticist/golee)](https://goreportcard.com/report/github.com/skyneticist/golee)


This cli tool is intended to make frequent git work in the command line more convenient and less time-consuming.

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
-----------

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
 
  `gg acp "message here"`
  
Similarly, when pushing a fresh remote, we usually would do:

  `git add .`
 
  `git commit -m "message here"`
  
  `git push --set-upstream origin <branch>` or `git push -u origin HEAD`
  
With git-golee:

  `gg acpr "message here"`
  
Renaming a branch can be a bit burdensome:
  
  ``

Alternatively, we could do:

  `gg rn <updatedBranchName>`
