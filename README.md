Parakeet
--------

Attempting to reduce the suckery of releases accross Github projects. Parakeet foolishly pecks at `git log --oneline` output betweeen git commits to add complete github link to originating Github issues for documentation.

![](https://media.giphy.com/media/ceHKRKMR6Ojao/giphy.gif)
[source](https://giphy.com/gifs/bird-explosion-parakeet-ceHKRKMR6Ojao)

Expects the Github "Squash and Merge" notation as Git log lines: "{hash} {PR title message} {Issue Number}"


## Example

Process the git commit log between two tags, process logs, and dump to clipboard(#yearofthelinuxdesktop).
`git log --oneline tagOld..tagNew | parakeet github.com/ropes/parakeet | xclip`

Output from dumping the example file included in project:
```cat eg.gitlog | ./parakeet https://github.com/ropes/parakeet
[666666e11](https://github.com/ropes/parakeet/commit/666666e11) Guardrails to cache usage refactored [1287](https://github.com/ropes/parakeet/issues/1287)
[131313131](https://github.com/ropes/parakeet/commit/131313131) major refactoring to clean things up. [1302](https://github.com/ropes/parakeet/issues/1302)
```