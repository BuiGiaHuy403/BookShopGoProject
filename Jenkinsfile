pipeline {
    agent any
    tools {
        go '1.23.4'
    }

    stages {
        stage('Checkout Code') {
            steps {
                checkout scm
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    sh 'docker compose build'
                }
            }
        }
         stage('Run Containers') {
            steps {
                script {
                    sh 'docker compose up -d'
                }
                script {
                    sh './wait-for-it.sh postgres-db:5432 --timeout=30 --strict'
                }
            }
        }

        stage('Run Unit Tests') {
            steps {
                script {
                    // sh 'go install github.com/jstemer/go-junit-report@latest'
                    sh 'go test ./... -v'
                }
            }
        }

        stage('Test API') {
            steps {
                script {
                    sh 'curl -X GET http://localhost:7000/health'
                }
            }
        }

        stage('Clean Up') {
            steps {
                script {
                    sh 'docker compose down'
                }
            }
        }
    }

    post {
        always {
            junit 'test-report.xml'
        }
    }
}
