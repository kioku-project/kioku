# Frontend development

This documents outlines all the important information to understand how to run and develop the frontend of Kioku.

## Table of Contents

- [Introduction](#introduction)
- [Local development](#local-development)
  - [VSCode Extensions](#visual-studio-code-extensions)
  - [Storybook](#storybook)
  - [Chromatic](#chromatic)
- [Create a component](#create-a-component)
- [Internationalization](#internationalization)

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

### Storybook

Storybook is an open-source tool for developing and testing UI components in isolation. It enables developers to create components independently and showcase components interactively in an isolated development environment. The main benefit of Storybook is that it aids in building a component library, offers a playground for components, and serves as fantastic documentation for teams to understand how to use the components. It makes UI development faster and easier, while improving component reuse and consistency.
![chrome_SXJWCN6Yuo](https://github.com/kioku-project/kioku/assets/60541979/9cb21aa9-8c1b-4582-83d6-a95d770fbebf)

#### Setup Storybook

To use [Storybook](https://storybook.js.org/) locally, run `npm run storybook` inside of the `frontend` folder and open [localhost:6006](http://localhost:6006) if it does not open automatically.

New to Storybook? Learn how to write stories [here](https://storybook.js.org/docs/react/writing-stories/introduction).

### Chromatic

Chromatic is a development tool created by developers of Storybook. It's used for testing, and documenting UI components and maintaining UI libraries. Chromatic simplifies the process of collaborating and maintaining UI components by providing features for visual testing, component documentation, and publishing. It also integrates with GitHub and offers automated Continuous Integration and Deployment processes, making it easier for teams to review and merge code changes.

#### Setup Chromatic

Chromatic should be run on every pull request that introduces changes to the frontend to ensure that components don't change unexpectedly.
To use [Chromatic](https://www.chromatic.com/) you have to add the [chromatic-project-token](https://www.chromatic.com/manage?appId=63e354941aa15501d3467f88&view=configure) to your .env file inside of the `frontend folder`.

```
CHROMATIC_PROJECT_TOKEN=
```

To publish storybook to Chromatic run `npm run chromatic`.

## Create a component

If you want to create a new component, add it to a suitable category folder inside of the `components` folder.
A component should always have an implementation in Typescript and a story with the same name ending with `.stories.ts`

- Components should not handle business logic themselves, they should solely display given data.
- Business logic should be inside of a page or generalized in functions outside of the component itself
- A component should have stories for all different versions. For example: if a button can be toggled between a primary and secondary style, both cases should be covered by stories.

## Internationalization

To internationalize Kioku in different languages, we use Next.js' native capability for [internationalized routing](https://nextjs.org/docs/pages/building-your-application/routing/internationalization) as well as the [Lingui](https://lingui.dev/) library. Our current policy is that English is the default language and a German translation for all texts has to be provided. However, if you would like to add another language, contributions are always welcome!

### Next.js' native internationalized routing.

Next.js natively supports internationalized routing, which means that it interprets `/de/home` as a request for the `home` page with German text. Additionally, all navigation using Next.js’s `Link` component and the `useRouter()` hook will automatically modify the paths to maintain the user’s language preference.

### Lingui

Lingui is a translation library that provides various macros to simplify the translation process, including plural handling and conditionals. The following section will briefly explain the most important concepts of Lingui:

#### The `<Trans>` macro

The `<Trans>` macro will probably be used most often when you want to create an internationalized component
```jsx
import { Trans } from "@lingui/macro";

function render() {
  return (
    <>
      <h1>
        <Trans>LinguiJS example</Trans>
      </h1>
      <p>
        <Trans>
          Hello <a href="/profile">{name}</a>.
        </Trans>
      </p>
    </>
  );
}
```
By simply wrapping text inside of components with the `<Trans>` macro, the text will be automatically recognized as translatable text.

#### The `useLingui()` macro

It is not always possible to wrap text within the `<Trans>` component. Take, for example, this situation, where you want to pass a string as a property to a component:
```jsx
import { msg } from "@lingui/macro";
import { useLingui } from "@lingui/react";

export default function ImageWithCaption() {
  const { _ } = useLingui();

  return <img src="..." alt={_(msg`Image caption`)} />;
}
```
Here, we want to be able to translate the image caption but we can’t use the `<Trans>` component in the `alt` property so we have to do it like this:

#### Translate outside of React components
If you need to translate text that resides in functions not part of React components, you cannot use the `<Trans>` component or the `useLingui()` hook. Instead, you will need to use the `t` macro.
```jsx
import { t } from "@lingui/macro";

export function showAlert() {
  alert(t`Warning! Something went wrong!`);
}
```

#### Plural
With Lingui, handling pluralization in frontend text is straightforward. You can define plural text as follows, and the appropriate form will be automatically selected based on the provided number.
```jsx
plural(numBooks, {
  one: "# book",
  other: "# books",
});
```

#### Extraction and Compilation
After developing new components, you will have to generate the translation files using `npm run extract`.
After that, the new source texts will be automatically added to `locale/de/messages.po`. You can search for the new texts and translate them inside of this file.
Finally, run `npm run compile` to compile the translations into Typescript files that can be used by the frontend.