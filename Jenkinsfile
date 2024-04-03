pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                script {
                    // Build the Go application
                    sh 'go build -o CI-CD-Pipeline .'
                }
            }
        }
        
        stage('Docker Build') {
            steps {
                script {
                    // Build Docker image
                    docker.build('sagnik3788/ci-cd-pipeline:latest')
                }
            }
        }

        stage('Docker Push') {
            steps {
                script {
                    // Push Docker image to registry

                    docker.withRegistry('https://index.docker.io/v1/', [username: 'sagnik3788', password: 'hicutie']) {

                        docker.image('sagnik3788/ci-cd-pipeline:latest').push('latest')
                    }
                }
            }
        }
    }
}
