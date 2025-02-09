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
                    sh '''
                    apt update && apt install -y make
                    make all
                    '''
                }
            }
        }
    }
    //      stage('Run Containers') {
    //         steps {
    //             script {
    //                 sh '''
    //                 docker compose up -d
    //                 '''
    //             }

    //         }
    //     }

    //     stage('Run Unit Tests') {
    //         steps {
    //             script {
    //                 // sh 'go install github.com/jstemer/go-junit-report@latest'
    //               sh '''
    //                     docker compose run --rm test-runner sh -c "go test ./... -v 2>&1 | go-junit-report > test-report.xml"
    //                 '''
    //             }
    //         }
    //     }

      

    //     stage('Clean Up') {
    //         steps {
    //             script {
    //                 sh 'docker compose down'
    //             }
    //         }
    //     }
    // }

    post {
        always {
            junit 'test-report.xml'
        }
    }
}

