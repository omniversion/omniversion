## Why use `omniversion`?

### The problem

Maintaining a modern software stack across deployments is hard work.

Given the manual effort needed just to keep dependencies up-to-date, avoid version conflicts and patch vulnerabilities, it is no surprise that these tasks are often neglected.

Version control systems, package managers and vulnerability scanners offer partial solutions, but they leave a lot to be desired:

* There is no obvious way of ensuring version consistency across package managers, e.g. when an `apt` package and an `npm` dependency need to be in sync.
* Multiple package managers need to be called in turn to answer simple questions like "Is there anything to patch?".
* Repeating this for each version currently deployed on a server is time-consuming and error-prone.
* Many software versions are not actually controlled by package managers, including - more often than not - the package managers themselves.
* Versions kept in configuration files are also frequently unmanaged, leading to hidden inconsistencies.
* Package managers differ greatly in their syntax, features, terminology and underlying model.

### The solution

For our daily maintenance work, we wanted a single dashboard containing all relevant information - across package managers, across servers and even across projects.

Given that we perform maintenance fairly frequently, we found that the added complexity of automation was worthwhile, saving time and preventing human error.

This is why we built the `omniversion` toolbox.

We make it available as Free Open Source Software in the hope that it might benefit other people as well.
