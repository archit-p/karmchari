# Karmchari : REST Job Manager
[![Golang](https://www.vectorlogo.zone/logos/golang/golang-ar21.svg)](https://golang.org)
[![Redis](https://www.vectorlogo.zone/logos/redis/redis-ar21.svg)](https://redis.io)
[![Build Status](https://img.shields.io/travis/archit-p/karmchari?style=flat-square)](https://travis-ci.org/archit-p/karmchari)
[![License](https://img.shields.io/github/license/archit-p/karmchari?style=flat-square)](LICENSE.md)  
Karmchari is a general purpose job manager written in Go. It allows for adding a job and updating and reading job states through a REST API.

## Usecases
Having Karmchari maintain a global state of running processes can help give better control on the jobs running on various sytems and workers.
1. Asynchronously signal workers to pause or kill jobs.
2. Keep track of long running jobs.

## Running
Karmchari uses Redis for data storage. A quick way to run Redis is using its Docker image.
```bash
docker run -p 6379:6379 --name redis-karmchari -d redis
```
Perhaps the simplest way to run Karmchari is through the included `Dockerfile`, and linking to the redis container.
```bash
docker build -t karmchari .
docker run -p 51463:51463 --name karmchari-prod --link redis-karmchari:redis -d karmchari
```
on Linux based PCs, simply run
```bash
./docker-build.sh
```
## Options
```
port  : port to start the app on ex. 51463
shost : host for the redis instance ex. localhost:6379
```

## Endpoints
Currently two endpoints are supported - `registerJob` and `jobState`.
### 1. registerJob (Method : POST)
#### Description
Add a new job.
#### Parameters
```
type : type of job - can be upload, export or teams
```
#### Returns
```
On success the job ID - MD5 hash string.
On failure appropriate error codes.
```
### 2. jobState (Method : POST)
#### Description
Change the state of a job.
#### Parameters
```
id : job id
command : what to do with the job - can be start, pause or kill
```
#### Returns
```
On success the job ID and new state.
On failure appropriate error codes.
```
### 3. jobState (Method: GET)
#### Descrtiption
Read the state of a job
#### Parameters
```
id : job id
```
#### Returns
```
On success returns the job state (start, pause or kill)
On failure appropriate error codes.
```
## Examples
Karmchari's APIs can be accessed through an HTTP client such as `curl` or `wget`.
#### Creating a job
```bash
curl -d "type=teams" http://localhost:51463/registerJob
Output: Created! id = 45c02fc01519814156e94adbd1902279
```
#### Put a job into hibernation
```bash
curl -d "id=45c02fc01519814156e94adbd1902279&command=pause" http://localhost:51463/jobState
Output: 45c02fc01519814156e94adbd1902279 -- pause
```
#### Read a jobs state
```bash
curl http://localhost:51463/jobState?id=45c02fc01519814156e94adbd1902279
Output: 45c02fc01519814156e94adbd1902279 -- pause
```
