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
<<<<<<< HEAD:jenkinsfile
                    docker.withRegistry('https://index.docker.io/v1/', [credentialsId: 'docker-hub-credentials']) {
=======
                    docker.withRegistry('https://index.docker.io/v1/', [username: 'sagnik3788', password: 'coder@981']) {
>>>>>>> e85ec8e5590c99c45f71d112b0586494661554e4:Jenkinsfile
                        docker.image('sagnik3788/ci-cd-pipeline:latest').push('latest')
                    }
                }
            }
        }
    }
}
