#!/usr/bin/env groovy

pipeline {
  agent { label 'executor-v2' }

  options {
    timestamps()
    buildDiscarder(logRotator(numToKeepStr: '30'))
  }

  triggers {
    cron(getDailyCronString())
  }

  stages {
    stage('Build Docker image') {
      steps {
        sh './bin/build'
      }
    }

    stage('Test') {
      parallel {
        stage('Test Postgres') {
          steps {
            sh './test/test postgres'
          }
        }

        stage('Test MySQL') {
          steps {
            sh './test/test mysql'
          }
        }

        stage('Test MSSQL') {
          steps {
            sh './test/test mssql'
          }
        }
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
