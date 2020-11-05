## Steps (Self-test)

- Run unit test & report

  ```sh
  # install go-junit-report
  go get -u github.com/jstemmer/go-junit-report

  cd helloworld
  go mod download
  go test -v 2>&1 | go-junit-report > report.xml
  ```

- Build package(create image)
  ```sh
  docker image build -t helloworld .
  ```
- Start application
  ```sh
  docker-compose up -d --build --force-recreate
  ```
- Run API test

  ```sh
  # install lib
  pip install robotframework
  pip install robotframework-requests

  # cd test/api
  robot greeting.robot
  ```

- Stop application
  ```sh
  # cd ../..
  docker-compose down
  ```
