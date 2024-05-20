# go-crud
- This is a simple CRUD application written in Golang.
- Package `gorilla/mux` implements a request router and dispatcher for matching incoming requests to their respective handler.
- Automated image building and docker push happens at master commit using `Github Action workflow`
- Database used : `MongoDB`

Steps to start the application using docker
--------------------------------------------
1. Update the .env file with your own `MONGO_CONNECTION_STRING` or you can choose to pass with option `-e MONGO_CONNECTION_STRING=<conn_string>` in docker run command.
2. Build docker image from root directory of the application using `docker build --tag go-crud .`
3. Start the container using `docker run -d -p 8000:8000 --name go-crud go-crud`.
4. Run `curl localhost:8000/api/movies` to test if the application is up.
5. Once done, stop the running container in the background using `docker stop go-crud`

Github action workflow (CI/CD)
------------------------------
A github actions workflow file has been setup (.github/workflows/main.yml). When we push the changes to master branch, action gets triggered. It checkouts the codebase, builds and creates the image, and then finally pushes it over dockerhub.
Prerequisites :
1. Open the repository Settings, and go to Secrets and variables > Actions.
2. Create a new Repository secrets named DOCKER_USERNAME and your Docker ID as value.
3. Create a new Personal Access Token (PAT) for Docker Hub. You can name this token docker-tutorial. Make sure access permissions include Read and Write.
4. Add the PAT as a second Repository secrets in your GitHub repository, with the name DOCKERHUB_TOKEN.

Testing over kubernetes using docker desktop
---------------------------------------------
Prerequisite:
1. Setup docker desktop on local system. Need to install kubectl as well if running on linux.
Docker Desktop includes a standalone Kubernetes server and client, as well as Docker CLI integration that runs on your machine.
The Kubernetes server runs locally within your Docker instance, is not configurable, and is a single-node cluster. It runs within a Docker container on your local system, and is only for local testing.
2. Image is already published over dockerhub using github workflow.

Steps to depoy and test:
1. Update the `go-crud-kubernetes.yaml` file to change the `MONGO_CONNECTION_STRING` env variable value.
2. Navigate to the project directory and deploy your application to Kubernetes using `kubectl apply -f go-crud-kubernetes.yaml`
3. Check if deployed successfully using `kubectl get deployments` and `kubectl get services`
4. Run `curl localhost:30001/api/movies`


References : 
1. https://pkg.go.dev/net/http
2. https://dev.to/hackmamba/build-a-rest-api-with-golang-and-mongodb-gorillamux-version-57fh
3. https://docs.docker.com/language/golang/build-images/
4. https://docs.docker.com/language/golang/run-containers/
5. https://docs.docker.com/desktop/install/ubuntu/
6. https://docs.docker.com/desktop/kubernetes/#install-and-turn-on-kubernetes
7. https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/
8. https://docs.docker.com/language/golang/configure-ci-cd/

