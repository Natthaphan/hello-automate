pipeline {
    agent any

    stages {
        stage('unit test and report') {
            steps {
                sh label: '', script: '''cd helloworld
                go mod download
                go test -v 2>&1 | go-junit-report > report.xml'''
            }
        }
    }
}