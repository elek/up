pipeline {
    agent {
        docker {
            label 'main'
            image 'storjlabs/ci:latest'
            alwaysPull true
            args '-u root:root --cap-add SYS_PTRACE -v "/tmp/gomod":/go/pkg/mod'
        }
    }

    options {
          timeout(time: 26, unit: 'MINUTES')
    }

    environment {
        GOTRACEBACK = 'all'
    }

    stages {
        stage('Checkout') {
            checkout scm
        }

        stage('Test') {
            sh './build -v test'
       }
    }

    post {
        always {
            sh "chmod -R 777 ." // ensure Jenkins agent can delete the working directory
            deleteDir()
        }
    }
}
