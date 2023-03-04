pipeline {
    agent any

    tools {
        go 'go1.19.5'
    }

    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
        dockerImage = ''
        tagName = "127.0.0.1:8085/actions:latest"
        registry = "127.0.0.1:8085"
        registryCredentials = "nexus_creds"
    }

    stages {
        // Clon app from git
        stage('Git Repo Cloning') {
            steps {
                checkout([
                    $class: 'GitSCM',
                    branches: [[name: '*/master']],
                    extensions: [],
                    userRemoteConfigs: [[credentialsId: 'git_creds', url: 'https://github.com/ruauka/for_actions.git']]])
                println "Cloning done"
            }
        }
        // Make unit-tests
        stage('Unit Tests') {
            steps {
                sh 'go test ./... -cover'
                println "Unit Tests done"
            }
        }
        // Building Docker images
        stage('Building image') {
            steps{
                script {
                    dockerImage = docker.build tagName
                }
            }
        }
        // Uploading Docker images into Nexus Registry
        stage('Uploading to Nexus') {
            steps{
                script {
                    docker.withRegistry('http://'+registry, registryCredentials) {
                        dockerImage.push()
                    }
                }
            }
        }
        // Stopping Docker containers for cleaner Docker run
        stage('stop previous containers') {
            steps {
                sh 'docker ps -f name=actions -q | xargs --no-run-if-empty docker container stop'
                // sh 'docker container ls -a -fname=container_name -q | xargs -r docker container rm'
            }
      }
        stage('Docker Run') {
           steps{
                script {
                    sh 'docker run -d --rm --name actions -p 8080:8000 ' + tagName
                }
            }
        }
    }
}
