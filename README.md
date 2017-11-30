# Data Together API

<!-- Repo Badges for: Github Project, Slack, License-->

[![GitHub](https://img.shields.io/badge/project-Data_Together-487b57.svg?style=flat-square)](http://github.com/datatogether)
[![Slack](https://img.shields.io/badge/slack-Archivers-b44e88.svg?style=flat-square)](https://archivers-slack.herokuapp.com/)
[![License](https://img.shields.io/github/license/datatogether/api.svg)](./LICENSE)

Serves the Data Together JSON API.

This component in isolation doesn't do much, but when connected to the project as a whole it provides a variety of endpoints over https, which can be used to extract information from the Data Together database.

**API Documentation:** [`https://api.archivers.co/docs/`](https://api.archivers.co/docs/)

## License & Copyright

Copyright (C) 2017 Data Together
This program is free software: you can redistribute it and/or modify it under
the terms of the GNU Affero General Public License as published by the Free Software
Foundation, version 3.0.

This program is distributed in the hope that it will be useful, but WITHOUT ANY
WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A
PARTICULAR PURPOSE.

See the [`LICENSE`](./LICENSE) file for details.

## Getting Involved

We would love involvement from more people! If you notice any errors or would like to submit changes, please see our [Contributing Guidelines](./.github/CONTRIBUTING.md). 

We use GitHub issues for [tracking bugs and feature requests](https://github.com/datatogether/api/issues) and Pull Requests (PRs) for [submitting changes](https://github.com/datatogether/api/pulls)

## Installation and Local Deployment

- install [docker](https://www.docker.com/) and [docker-compose](https://docs.docker.com/compose/)
  - if you are new to docker, you may want to test your install with the [default test case](https://docs.docker.com/compose/gettingstarted/)
- in the project directory, run ``docker-compose build && docker-compose up`
- the API should now be running at `http://localhost:3200/`, and the following endpoints should work:
  - http://localhost:3200/urls for a list of URL's in the database
  - http://localhost:3200/urls/{id} for an indivudual URL
  - http://localhost:3200/primers for a list of primers
  - http://localhost:3200/users for a list of users
  - http://localhost:3200/users/{id} for an individual user
  - http://localhost:3200/primers for a list of primers
  - http://localhost:3200/primers/{id} for an individual primer
  - http://localhost:3200/sources for a list of data sources
  - http://localhost:3200/sources/{id} for an individual source
  - http://localhost:3200/repositories for a list of repositories
  - http://localhost:3200/repository/{id} for an in ndividual repository
  - http://localhost:3200/coverage for a coverage tree associated with root URLs
  - http://localhost:3200/collections for a list of collections
  - http://localhost:3200/collections/{id}  for an individual collectoin
see below for more information

### Generating Documentation

THe API documentation is OpenAPI/Swagger compliant and is generated by the `spectacle` node module. You will likely want to explore these docs, so you should generate them!

1. Install [spectacle](https://github.com/sourcey/spectacle) with `npm -g install specatacle-docs`
2. To dynamically generate API docs as you work, run `spectacle -d open_api.yaml` and edit `open_api.yaml` to see changes
3. Generate Static docs with `spectacle open_api.yaml`
4. Commit. Rinse. Repeat.

## Development

Right now modifying & updating code is a huge pain, but this is at least a start.

