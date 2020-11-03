# Github to Redmine issues binder

Webhook for converting Redmine issues IDs (eg #53344, #33773) in the Pull Requests body to corresponding issue link.

## Deployment

Docker image is available [here](https://hub.docker.com/r/alphatroya/github-redmine-binder)

The container uses 8933 port and handles `/github` endpoint

Following ENV variables should be set:
- `GITHUB_ACCESS_TOKEN`: token for PR editor (scope should be `repo` or greater)
- `REDMINE_HOST`: Redmine host address

