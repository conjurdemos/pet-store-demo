#!/usr/bin/env groovy

pipeline {
  agent { label 'executor-v2' }

  options {
    timestamps()
    buildDiscarder(logRotator(numToKeepStr: '30'))
  }

  stages {
    stage('Build Docker image') {
      steps {
        sh './bin/build'
      }
    }

    stage('Test') {
      steps {
        sh './test/test.sh postgres'
        sh './test/test.sh mysql'
      }
    }

    stage('Publish Docker image to registry') {
      when {
        branch 'master'
      }

      steps {
        sh './bin/publish'
      }
    }
  }

  post {
    always {
      cleanupAndNotify(currentBuild.currentResult)
    }
  }
}