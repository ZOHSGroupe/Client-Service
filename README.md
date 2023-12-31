## Description

Client-Service is an Application Programming Interface (API) to handle add,update,delete and get client of ZOHS company.
## Installation :
```bash
# install requirements
$ go get -u github.com/go-sql-driver/mysql github.com/jinzhu/inflection github.com/gorilla/mux github.com/joho/godotenv github.com/jinzhu/inflection github.com/dgrijalva/jwt-go
```
## Running the app : 
```bash
# Run application
$ go run main.go
```
## Build Docker image : 
```bash
# build a docker image
$ docker build -t client-service .
```
## Running the app in the Docker : 
```bash
# Run docker image
$ docker run -p 5000:5000 client-service
```
## Running with Docker compose :
```bash
# Run docker compose
$ docker compose up
```
## Stay in touch :
- Author - [Ouail Laamiri](https://www.linkedin.com/in/ouaillaamiri/) , [Hajar Sadik](https://www.linkedin.com/in/hajar-sadik-b27594268?miniProfileUrn=urn%3Ali%3Afs_miniProfile%3AACoAAEGU9k8BLhHzARArf7SblDplgU6ufFqc-0w&lipi=urn%3Ali%3Apage%3Ad_flagship3_search_srp_all%3BwRckppvzRNiFd9xB2PswKw%3D%3D)
- Test - [Postman](https://www.postman.com/avionics-meteorologist-32935362/workspace/postman-api-fundamentals-student-expert/collection/29141176-8abc7d24-d6b7-4b89-a56e-4e4a6e5c7a5a?action=share&creator=29141176)
- Documentation - [Postman](https://documenter.getpostman.com/view/27440422/2s9YkuXxNb)

## License

Client-Service is [MIT licensed](LICENSE).


