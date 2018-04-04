#!/usr/bin/env groovy

pipeline {
  agent { label 'executor-v2' }

  options {
    timestamps()
    buildDiscarder(logRotator(daysToKeepStr: '30'))
  }

  stages {
    stage('Build Go binaries') {
      steps {
        sh './build.sh'
        archiveArtifacts artifacts: 'output/*', fingerprint: true
      }
    }
    stage('Run unit tests') {
      steps {
        sh './test.sh'
        junit 'output/junit.xml'
        sh 'sudo chown -R jenkins:jenkins .'  // bad docker mount creates unreadable files TODO fix this
      }
    }

    stage('Package distribution tarballs') {
      steps {
        sh './package.sh'
        archiveArtifacts artifacts: 'output/dist/*', fingerprint: true
      }
    }
  }

  post {
    always {
      cleanupAndNotify(currentBuild.currentResult)
    }
  }
}
