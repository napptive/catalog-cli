# Contributing

- [Getting Involved](#getting-involved)
- [How can you contribute](#how-can-you-contribute)
- [Code contributions](#code-contributions)
- [Best practices](#to-make-a-good-contribution-keep-in-mind-the-following-points)

## Getting Involved

Welcome to Napptive! This guide will help you know the necessary steps to contribute to our community.
Please, first of all, read our [code of conduct](code-of-conduct.md).

## How can you contribute

- **Help us to improve documentation.** We are working on improving our documentation. If you find any errors or you think that something is missing, please let us know by opening an issue or a pull request.

- **Have you found a bug?**
We will be very grateful if you inform us by oppening an issue.
Be clear and concise in the definition of the bug. Indicate the steps to reproduce it, which version you are using, on which platform. Include any additional information that may help to understand the bug.
  
- **Open an issue.**
Do you see something that needs work? Do you have any new ideas? Would you like to discuss any functionality? Feel free to open an issue. Fill the issue template being as clear and concise as you can.

- **Have you found a security problem?**
We would appreciate if you could send us a message to community@napptive.com informing us of the problem

If you have any questions, you don’t know how to contribute, or you need help, feel free to use [Napptive’s Gitter](https://gitter.im/napptive/community), and we will be happy to help you.

## Code contributions

- **Fork the repository.**
- **Clone the repository.**
  ```bash
  git clone https://github.com/<your-username>/catalog-cli.git
  ```
- **Create a branch**: Relevent branch name is preferable.
  ```bash
  git checkout -b <branch-name>
  ```
- **Make your changes:** You can use the [development guide](DEVELOPMENT.md) to help you.
- **Commit your changes:** Write understandable and clear git commit messages.
  ```bash
  git commit -m "commit message"
  ```
- **Push your changes.**
  ```bash
    git push origin <branch-name>
    ```
- **Create a pull request:** Be sure to include a clear and detailed PR description that explains the purpose of the PR. Don't forget that the reviewer must know what you are trying to solve, why, and how. Select at least one person to review your PR.  
Fill all the required information in PR template:

```md
### Checklist
- [ ] Code compiles correctly
- [ ] Created tests which fail without the change (if possible)
- [ ] All tests passing
- [ ] Extended the README/documentation, if necessary

#### What does this PR do?

#### Where should the reviewer start?

#### What is missing?

#### How should this be manually tested?

#### Any background context you want to provide?

#### What are the associated tickets?

#### Screenshots (if appropriate)

#### Questions
```
## Best practices
**To make a good contribution keep in mind the following points**:


- Write understandable and clear git commit messages is really helpfull.
- Take care of the quality of the code following the [official guide](https://go.dev/doc/effective_go).
- Use [gofmt](https://pkg.go.dev/cmd/gofmt) tool.
- Add comments in the code in order to make it more understandable to the community.
- Be sure to include the corresponding tests (unit test and integration test if proceed)
---
> If it is four first contribution, you should sign our CLA document.
