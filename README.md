# GoLogging :mag:

> A simple implementation applying some logging concerns

## Intent

This project is a simple API with just one route, that responds if a value is integer or not. <br>
The intent is develop a simple logging system into this for log some debugs, infos or errors.

## Some tools used in this project

1. [Echo](https://echo.labstack.com/) (golang HTTP framework)
2. [Logrus](https://github.com/sirupsen/logrus) (for Debug/Info/Error logs)

## How to run this locally

> It's simple, all application is encapsuled into a docker container, so with few steps you can run this project. <br>

1. Build the container:

```sh
make up
```

2. Get into application:

```sh
make app
```

3. Install dependencies (you just need run this once):

```sh
make install
```

4. Now, simply run the application:

```sh
make run
```

> Application is running on port 2001

## How to use

To test you only need execute requests to `/isInt` passing the query param `value`. <br>
To make things easly, I developed a simple script to execute HTTP requests to this API, so... <br>

1. To run requests from your CLI, simply run:

```sh
sh ./request.sh <your-value-here>
```
