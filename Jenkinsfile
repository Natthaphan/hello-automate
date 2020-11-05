pipeline {
    agent any

    stages {
        stage('unit test and report') {
            steps {
                sh label: '', script: '''cd helloworld
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
            sh label: '', script: 'docker-compose up -d --build --force-recreate'

            post {
                always {
                    sh label: '', script: 'docker-compose down'
                }
            }
        }
    }
}