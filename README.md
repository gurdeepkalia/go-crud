# go-crud

Steps to start the application
1. Update the .env file with your own `MONGO_CONNECTION_STRING`
2. Build docker image from root directory of the application using `docker build --tag go-crud .`
3. Start the container using `docker run -d -p 8000:8000 --name go-crud go-crud`.
4. Run `curl localhost:8000/api/movies` to test if the application is up.
5. Once done, stop the running container in the background using `docker stop go-crud`

References : 
1. https://pkg.go.dev/net/http
2. https://dev.to/hackmamba/build-a-rest-api-with-golang-and-mongodb-gorillamux-version-57fh
3. https://docs.docker.com/language/golang/build-images/
4. https://docs.docker.com/language/golang/run-containers/
