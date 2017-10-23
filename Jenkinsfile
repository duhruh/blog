pipeline {
    agent {
        dockerfile true
    }

    stages {
        stage('Build') {
            steps {
                sh 'pwd'
                sh "go version"
                sh 'cd /go/src/github.com/duhruh/blog'
                sh 'pwd'

                sh 'cd /go/src/github.com/duhruh/blog && ls'
                sh 'cd /go/src/github.com/duhruh/blog && glide install'
                sh "cd /go/src/github.com/duhruh/blog && go run cmd/task/main.go build --version=v.0.0.0-alpha.1 --build=${env.BUILD_NUMBER}"
            }
        }
        stage('Test') {
            steps {
                sh 'go test $(glide nv)'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
            }
        }
    }
}