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
                   sh '''
                        echo "Waiting for database to be ready..."
                        until nc -z postgres-db 5432; do
                            sleep 5
                        done
                        echo "Database is ready. Running tests..."
                        go test ./... -v 
                    '''
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
