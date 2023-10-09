
![kioku](https://github.com/kioku-project/kioku/assets/60541979/1f827df3-5882-4285-913f-47f04b26196b)

![GitHub contributors](https://img.shields.io/github/contributors/kioku-project/kioku)
[![Backend Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=kioku-project_kioku_services&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=kioku-project_kioku_services)
[![Frontend Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=kioku-project_kioku_frontend&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=kioku-project_kioku_frontend)
[![Backend Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=kioku-project_kioku_services&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=kioku-project_kioku_services)
[![Frontend Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=kioku-project_kioku_frontend&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=kioku-project_kioku_frontend)
[![Storybook link](https://github.com/storybookjs/brand/blob/master/badge/badge-storybook.svg)](https://main--63e354941aa15501d3467f88.chromatic.com)


# Welcome to Kioku!
The cloud native flashcard application that focuses on collaborative content creation.

## Features

<table>
<tbody>
<tr>
<td><picture>
  <source media="(prefers-color-scheme: dark)" srcset="https://github.com/kioku-project/kioku/assets/60541979/0cc0f108-f0e7-49c5-bafc-02123afdf514">
  <source media="(prefers-color-scheme: light)" srcset="https://github.com/kioku-project/kioku/assets/60541979/2e21c7e1-304d-4f2c-9328-3f0486c12d0c">
  <img alt="Collaborative icon" src="">
</picture>
  <b>Collaborative</b>

Collaborate with your friends and fellow students in groups and work on shared decks. Learn together and motivate each other!

</td>
<td><picture>
  <source media="(prefers-color-scheme: dark)" srcset="https://github.com/kioku-project/kioku/assets/60541979/a3dbfa9e-0b38-477d-b1eb-705f30b45eda">
  <source media="(prefers-color-scheme: light)" srcset="https://github.com/kioku-project/kioku/assets/60541979/91a42bb7-8985-435f-8c61-c8a1f455dd7c">
  <img alt="Individual icon" src="">
</picture>
  <b>Individual</b>

Create and customize your own flashcards tailored to your needs and preferences. Set your own pace with our spaced repetition system to maximize your potential!


</td>
</tr>
<tr>
<td><picture>
  <source media="(prefers-color-scheme: dark)" srcset="https://github.com/kioku-project/kioku/assets/60541979/d2753ef0-6e62-48f5-9b1c-fc4aea450a1b">
  <source media="(prefers-color-scheme: light)" srcset="https://github.com/kioku-project/kioku/assets/60541979/f18f855c-6f26-49a7-84eb-18ef94838c69">
  <img alt="Compatible icon" src="">
</picture>
  <b>Compatible</b>

Kioku is compatible with Anki, allowing you to import and export your existing decks into our application while taking advantage of Kioku's collaborative features!
</td>
<td><picture>
  <source media="(prefers-color-scheme: dark)" srcset="https://github.com/kioku-project/kioku/assets/60541979/5c31d81d-dedd-4abf-8012-6b5b5ca430f4">
  <source media="(prefers-color-scheme: light)" srcset="https://github.com/kioku-project/kioku/assets/60541979/fc2b34df-ee6e-48cd-b1d5-3cbdd0eb0f2c">
  <img alt="Informative icon" src="">
</picture>
  <b>Informative</b>

We provide you with detailed statistics and insights into your study progress. Identify areas of improvement to optimize your strategy for maximum effectiveness!
</td>
</tr>
<tr>
<td><picture>
  <source media="(prefers-color-scheme: dark)" srcset="https://github.com/kioku-project/kioku/assets/60541979/ee59ba9e-e3ee-4dfd-a170-3f8438523309">
  <source media="(prefers-color-scheme: light)" srcset="https://github.com/kioku-project/kioku/assets/60541979/d12880ec-2d74-4146-a8f8-b2ade59688c9">
  <img alt="Available icon" src="">
</picture>
  <b>Available</b>

Access your flashcards everywhere and at any time. Switch seamlessly between multiple platforms and never miss a learning opportunity again!
</td>
<td>
    <picture>
  <source media="(prefers-color-scheme: dark)" srcset="https://github.com/kioku-project/kioku/assets/60541979/76142857-6749-419a-becc-86b234f16d42">
  <source media="(prefers-color-scheme: light)" srcset="https://github.com/kioku-project/kioku/assets/60541979/87bfe7ed-e292-4cd1-a6a3-ac26c0ea27ca">
  <img alt="Entertaining icon" src="">
</picture>
  <b vertical-align="center">Entertaining</b>

Achievements and leaderboards make learning more engaging and motivating. Kioku helps you to achieve better results and stay on track with your personal learning goals!
</td>
</tr>
</tbody>
</table>

# Index
1. [Getting started](#getting-started)
2. [Frontend development](./doc/frontend_development.md)
3. [Backend development](./doc/backend_development.md)
4. [Deployment](./doc/deployment.md)

# Getting started
In order to run Kioku locally, first clone the repository
```
git clone https://github.com/kioku-project/kioku
```
and configure the `.env` file in the `backend` folder
```
cp .env.example .env
```

> [!WARNING]
> The example environment file is populated with default values, be sure to change all values before using the application in production!

Finally, start frontend and backend using docker compose
```
docker compose up
```
after that, the frontend is reachable at `http://localhost:3000`  
to reach pgAdmin, open `http://localhost:3002`  
the micro dashboard is reachable at `http://localhost:3001`
