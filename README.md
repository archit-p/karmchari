# Karmchari : REST Job Manager
Karmchari is a job manager written in Go. It allows for adding a job and updating and reading job states through a REST API.

## Running
Perhaps the simplest way to run Karmchari is through the included `Dockerfile`.
```bash
docker build -t karmchari .
docker run -p 51463:51463 -it karmchari
```
on Linux based PCs, simply run
```bash
./docker-build.sh
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
## License
[The MIT License](LICENSE.md)
