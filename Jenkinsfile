pipeline {
    agent any

    stages {
        stage('unit test and report') {
            steps {
                sh label: 'cd', script: '''cd helloworld
                go mod download
                go test -v 2>&1 | go-junit-report > report.xml'''
            }
            post {
                always {
                    junit 'helloworld/report.xml'
                }
            }
        }

        stage('build and start application') {

            steps {
                sh label: 'docker-compose', script: 'docker-compose up -d --build --force-recreate'
            }
        }

        stage('run api test') {

            steps {
                sh label: 'robot', script: '''cd test/api
                robot greeting.robot'''
            }

            post {
                always {
                    robot outputPath: '', unstableThreshold: 100.0
                    sh label: '', script: 'docker-compose down'
                }
            }
        }
    }
}