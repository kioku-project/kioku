# Frontend development

This documents outlines all the important information to understand how to run and develop the frontend of Kioku.

## Table of Contents

- [Introduction](#introduction)
- [Local development](#local-development)
  - [VSCode Extensions](#visual-studio-code-extensions)
  - [Storybook](#storybook)
  - [Chromatic](#chromatic)
- [Create a component](#create-a-component)

## Introduction

The frontend of Kioku is written in [React](https://react.dev/), using the [NextJS](https://nextjs.org/) framework. All the styling is done directly inside of the components, using [TailwindCSS](https://tailwindcss.com/). To document all of our frontend components, we use [Storybook](https://storybook.js.org/) in conjunction with [Chromatic](https://www.chromatic.com/) for visual regression testing.

## Local development

NextJS offers a development mode with live reload that can speed up the development of the frontend significantly. In order to use this mode, the `frontend` container should be commented out of the `docker-compose.yml` file.
Additionally, the `frontend_proxy` has to be exposed on port `80`.

```yaml
frontend_proxy:
  build:
    context: backend
    dockerfile: services/frontend/Dockerfile
  container_name: kioku-frontend_proxy
  restart: always
  env_file:
    - ./backend/.env
  depends_on:
    - db
  ports:
    - 80:80
```

Afterwards, you will have to adjust the frontend rewrite in the `next.config.js` file:

```javascript
async rewrites() {
    return [
        {
            source: "/api/:path*",
            destination: "http://localhost:80/api/:path*",
        },
    ];
},
```

> [!IMPORTANT]  
> Please remember to exclude these changes from your pull request as they should only be used locally!

Furthermore, configure the frontend `.env` file

```
cp .env.example .env
```

Finally, you can start all of the backend services:

```
docker compose up -d
```

and start the frontend separately after navigating into the `frontend` folder

```
npm install
```

```
npm run dev
```

### Visual Studio Code Extensions

We recommend to install the [Prettier](https://marketplace.visualstudio.com/items?itemName=esbenp.prettier-vscode), the [SonarLint](https://marketplace.visualstudio.com/items?itemName=SonarSource.sonarlint-vscode) and the [TailwindCSS](https://marketplace.visualstudio.com/items?itemName=bradlc.vscode-tailwindcss) extensions inside of Visual Studio Code for an improved development experience.
The Prettier extension formats the code according to our guidelines.

> [!IMPORTANT]
> Please remember to format your code before committing it!

The SonarLint extension adds linting for SonarCloud errors.
The TailwindCSS extension adds documentation windows to tailwind classes inside of HTML that explains which styles get applied.

## Storybook

Storybook is an open-source tool for developing and testing UI components in isolation. It enables developers to create components independently and showcase components interactively in an isolated development environment. The main benefit of Storybook is that it aids in building a component library, offers a playground for components, and serves as fantastic documentation for teams to understand how to use the components. It makes UI development faster and easier, while improving component reuse and consistency.
![chrome_SXJWCN6Yuo](https://github.com/kioku-project/kioku/assets/60541979/9cb21aa9-8c1b-4582-83d6-a95d770fbebf)

### Setup Storybook

To use [Storybook](https://storybook.js.org/) locally, run `npm run storybook` inside of the `frontend` folder and open [localhost:6006](http://localhost:6006) if it does not open automatically.

New to Storybook? Learn how to write stories [here](https://storybook.js.org/docs/react/writing-stories/introduction).

## Chromatic

Chromatic is a development tool created by developers of Storybook. It's used for testing, and documenting UI components and maintaining UI libraries. Chromatic simplifies the process of collaborating and maintaining UI components by providing features for visual testing, component documentation, and publishing. It also integrates with GitHub and offers automated Continuous Integration and Deployment processes, making it easier for teams to review and merge code changes.

## Setup Chromatic

Chromatic should be run on every pull request that introduces changes to the frontend to ensure that components don't change unexpectedly.
To use [Chromatic](https://www.chromatic.com/) you have to add the [chromatic-project-token](https://www.chromatic.com/manage?appId=63e354941aa15501d3467f88&view=configure) to your .env file inside of the `frontend folder`.

```
CHROMATIC_PROJECT_TOKEN=
```

To publish storybook to Chromatic run `npm run chromatic`.

# Create a component

If you want to create a new component, add it to a suitable category folder inside of the `components` folder.
A component should always have an implementation in Typescript and a story with the same name ending with `.stories.ts`

- Components should not handle business logic themselves, they should solely display given data.
- Business logic should be inside of a page or generalized in functions outside of the component itself
- A component should have stories for all different versions. For example: if a button can be toggled between a primary and secondary style, both cases should be covered by stories.
