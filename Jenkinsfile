pipeline {
    agent {
        docker {
            image 'golang:1.23'
        }
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

        stage('Run Unit Tests') {
            steps {
                script {
                    sh 'go install github.com/jstemer/go-junit-report@latest'
                    sh 'go test ./... -v | go-junit-report > test-report.xml'
                }
            }
        }

        stage('Run Containers') {
            steps {
                script {
                    sh 'docker compose up -d'
                }
            }
        }

        stage('Test API') {
            steps {
                script {
                    sh 'curl -X GET http://localhost:8080/health'
                }
            }
        }

        stage('Clean Up') {
            steps {
                script {
                    sh 'docker-compose down'
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
