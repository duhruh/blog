pipeline {
    agent any

    stages {
        stage('Build') {
            node("go") {
                // Install the desired Go version
                def root = tool name: 'Go 1.8', type: 'go'

                // Export environment variables pointing to the directory where Go was installed
                withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
                    sh 'go version'
                }

                steps {
                                sh 'curl https://glide.sh/get | sh'
                                sh 'glide install'
                                sh 'go run cmd/task/main.go build --version=v.0.0.0-alpha.1 --build=${env.BUILD_NUMBER}'
                            }
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