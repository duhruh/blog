pipeline {
    agent {
        docker { image 'devtransition/golang-glide' }
    }

    stages {
        stage('Build') {
            steps {
                sh "mkdir -p /go/src/github.com/duhruh"
                dir("/go/src/github.com/duhruh") {
                    checkout scm
                    dir("/go/src/github.com/duhruh/blog"){
                        sh 'echo $GOPATH'
                        sh 'GO15VENDOREXPERIMENT=1 glide install'
                        sh "go run cmd/task/main.go build --version=v.0.0.0-alpha.1 --build=${env.BUILD_NUMBER}"
                    }
                }
            }
        }
        stage('Test') {
            steps {
                dir("/go/src/github.com/duhruh/blog"){
                    sh 'go test $(glide nv)'
                }
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
            }
        }
    }
}