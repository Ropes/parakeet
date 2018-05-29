Parakeet
--------

Attempting to reduce the suckery of releases accross Github projects. Parakeet foolishly pecks at `git log --oneline` output betweeen git commits to add complete github link to originating Github issues for documentation.


## Example

`git log --oneline tagOld..tagNew | parakeet github.com/ropes/parakeet | xclip`

```cat eg.gitlog | ./parakeet https://github.com/ropes/parakeet
[666666e11](https://github.com/ropes/parakeet/commit/666666e11) Guardrails to cache usage refactored [1287](https://github.com/ropes/parakeet/issues/1287)
[131313131](https://github.com/ropes/parakeet/commit/131313131) major refactoring to clean things up. [1302](https://github.com/ropes/parakeet/issues/1302)
```