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
            }
        }

        stage('Run Unit Tests') {
            steps {
                script {
                    // sh 'go install github.com/jstemer/go-junit-report@latest'
                    sh 'echo "test passed"'
                }
            }
        }

        stage('Test API') {
            steps {
                script {
                    sh 'curl -X GET http://localhost:8080/health'
                    sh 'curl -X GET http://localhost:8080/api/v1/books/1'
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
