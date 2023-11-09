# Welcome to Kioku's contribution guidelines

We are delighted that you want to help us to make Kioku even better! To make this as easy and transparent as possible for you, we have summarized the guidelines we follow in this document. Don't be afraid to make a mistake, we appreciate every contribution. If anything is still unclear to you, don't hesitate to contact us and ask for help. We look forward to your contribution to the project and are always happy to help.

# Index

1. [Git](#git)
   - [Branches](#branches)
   - [Commits](#commits)
2. [Backend](#backend)
3. [Frontend](#frontend)

# Git

### Branches

When creating branches, we follow the following pattern:

`[type]/[area]/[individual name]`

- `type`: What is the purpose of your changes? Currently we are using the following tags
  - feature: for implementing new features
  - fix: for bugfixes
  - hotfix: for urgent fixes such as security vulnerabilities and more important bugs that severely affect the usability of our application
  - refactor: for changes that only improve the source code and have little or no impact on the functionality of Kioku
  - test: for everything concerning testing
  - deployment: for all deployment topics
  - docs: for updating the docs
- `area`: Which area of the application do the changes affect? If none or both apply to your changes, simply leave it out.
  - frontend
  - backend
- `individual name`: This is where you should try to summarize your changes as short and precise as possible. If you need several words, separate them with a dash and write everything in lower case.

### Commits

When committing changes, please write your messages according to the following convention:

This commit will `Your commit message`

Write the first letter of your message in upper case and the rest in lower case.

### Pull requests

When you open a pull request, remember to label it and link the corresponding issue. Pull requests should always concern only one specific subject. Please use the template and fill in the required information.

# Backend

# Frontend

If you make changes to the frontend, please ensure the following points.

- Please run Prettier before commiting changes.
- Stories should be available for each component and its different variants.
- A corresponding translation in English and German should be provided for each text. If you do not speak German, we will support you with the translation.

If you are not familiar with any of these points, take a look at our [frontend documentation](./docs/frontend_development.md) or ask us for help.
