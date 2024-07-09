# CMS

![Tests](https://github.com/YanSystems/cms/actions/workflows/tests.yml/badge.svg) [![codecov](https://codecov.io/gh/YanSystems/cms/graph/badge.svg?token=okvGgYV5UR)](https://codecov.io/gh/YanSystems/cms) [![Go Report](https://goreportcard.com/badge/YanSystems/cms)](https://goreportcard.com/report/YanSystems/cms) [![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/YanSystems/cms/blob/main/LICENSE)


A content management (system) microservice that exposes a REST API and persists content in MongoDB.

## Running the service locally

You will need to store a MongoDB connection string as an environment variable named `YAN_CMS_DB_URI`. This microservice doesn't load from `.env` due to relative pathing issues, so you will need to store it in your shell session.
```
export YAN_CMS_DB_URI="my-connection-string"
```
Should you need it to persist across shell session, be sure to store it in `~/.bashrc`

You can now run the server (make sure you have go version `1.22.4`),
```
make run
```

you can run the service in a container. First, build the docker image,
```
make image
```
To start the container, run `make up`. To stop it, run `make down`

## License

This CMS microservice is [MIT licensed.](https://github.com/YanSystems/cms/blob/main/LICENSE)
