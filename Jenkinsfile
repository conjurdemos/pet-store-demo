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

    stage('Scan Docker image') {
      parallel {
        stage('Scan Docker image for fixable issues') {
          steps{
            script {
              TAG = sh(returnStdout: true, script: 'cat VERSION')
            }
            scanAndReport("demo-app:${TAG}", "HIGH", false)
          }
        }
        stage('Scan Docker image for all issues') {
          steps{
            script {
              TAG = sh(returnStdout: true, script: 'cat VERSION')
            }
            scanAndReport("demo-app:${TAG}", "NONE", true)
          }
        }
      }
    }

    stage('Publish Docker image to registry') {
      when {
        branch 'main'
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
