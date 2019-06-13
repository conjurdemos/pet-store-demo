#!/usr/bin/env groovy

pipeline {
  agent { label 'executor-v2' }

  options {
    timestamps()
    buildDiscarder(logRotator(numToKeepStr: '30'))
  }

  environment {
      OPENSHIFT_URL         = 'https://openshift-311.itci.conjur.net:8443'
      OPENSHIFT_USER        = credentials('oc-311-user')
      OPENSHIFT_PASSWORD    = credentials('oc-311-password')

      DOCKER_REGISTRY_PATH  = 'docker-registry-default.openshift-311.itci.conjur.net'
      PROJECT_NAME          = 'p2p-pet-store'
  }

  stages {
    stage('Build Docker image') {
      steps {
        sh './bin/build'
      }
    }

    stage('Test') {
      steps {
        sh './test/test postgres'
        sh './test/test mysql'
      }
    }

    stage('Deploy to OpenShift') {
      steps {
        sh './bin/deploy'
      }
    }

    // stage('Publish Docker image to registry') {
    //   when {
    //     branch 'master'
    //   }

    //   steps {
    //     sh './bin/publish'
    //   }
    // }
  }

  post {
    always {
      cleanupAndNotify(currentBuild.currentResult)
    }
  }
}