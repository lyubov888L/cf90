# CloudFoundry CLI - cf90 [![Build Status](https://travis-ci.org/igm/cf90.png?branch=master)](https://travis-ci.org/igm/cf90)

`cf90` is a command line tool to interact with cloud foundry instance. It is entirely written in GO language.
The tool is still under development.

## Installation
`$ go get -u github.com/igm/cf90`

This tool uses CF for GO library (http://github.com/igm/cf)

## Supported commands

```
Commands:
  help            - Shows this help message, use [COMMAND] for command parameters
  license         - show license information
Application
  app.create      - Create new application
  app.delete      - Delete application
  app.detail      - Show application info
  app.list        - Show a list of apps
  app.map         - Map application to given host and domain (route must already exist)
  app.push        - Push application
  app.start       - Start application
  app.stop        - Stop application
  app.unmap       - Unmap application from given host and domain (route must already exist)
Domain
  domain.list     - Show a list of domains
Organization
  org.detail      - Show organization detail
  org.list        - Show all organizations
Route
  route.create    - Create a route in current space
  route.delete    - Delete route
  route.list      - Show all routes
Service
  service.list    - Show a list of services
Space
  space.list      - Show all spaces in organization
  space.use       - Set default space (and also organization)
Target
  target.add      - Add new target
  target.info     - Show target information
  target.list     - Show known targets
  target.login    - Login to target
  target.logout   - Logout from target
  target.rm       - Remove target from the list of known targets
  target.use      - Set current target
```
