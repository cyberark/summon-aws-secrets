#!/usr/bin/env groovy
@Library("product-pipelines-shared-library") _

pipeline {
  agent { label 'conjur-enterprise-common-agent' }

  options {
    timestamps()
    buildDiscarder(logRotator(daysToKeepStr: '30'))
  }

  triggers {
    cron(getDailyCronString())
  }

  stages {

    stage('Get InfraPool AzureExecutorV2 Agent') {
      steps {
        script {
          INFRAPOOL_EXECUTORV2_AGENT_0 = getInfraPoolAgent.connected(type: "ExecutorV2", quantity: 1, duration: 1)[0]
        }
      }
    }

    stage('Validate') {
      parallel {
        stage('Changelog') {
          steps { parseChangelog(INFRAPOOL_EXECUTORV2_AGENT_0) }
        }
      }
    }

    stage('Build Go binaries') {
      stages {
        stage('Release artifacts') {
          when {
            buildingTag()
          }
          steps {
            script {
              INFRAPOOL_EXECUTORV2_AGENT_0.agentSh './bin/build'
              INFRAPOOL_EXECUTORV2_AGENT_0.agentStash name: 'dist', includes: 'dist/*'
              unstash 'dist'
              archiveArtifacts artifacts: 'dist/*', fingerprint: true
            }
          }
        }
        stage('Snapshot artifacts') {
          when {
            not {
              buildingTag()
            }
          }
          steps {
            script {
              INFRAPOOL_EXECUTORV2_AGENT_0.agentSh './bin/build --snapshot'
              INFRAPOOL_EXECUTORV2_AGENT_0.agentStash name: 'dist2', includes: 'dist/*'
              unstash 'dist2'
              archiveArtifacts artifacts: 'dist/*', fingerprint: true
            }
          }
        }
      }
    }

    stage('Run unit tests') {
      steps {
        script {
          INFRAPOOL_EXECUTORV2_AGENT_0.agentSh './bin/test.sh'
          INFRAPOOL_EXECUTORV2_AGENT_0.agentStash name: 'output', includes: 'output/*'
          unstash 'output'
          junit 'output/junit.xml'
          INFRAPOOL_EXECUTORV2_AGENT_0.agentSh 'sudo chown -R jenkins:jenkins .'  // bad docker mount creates unreadable files TODO fix this
          cobertura autoUpdateHealth: true, autoUpdateStability: true, coberturaReportFile: 'output/coverage.xml', conditionalCoverageTargets: '30, 0, 0', failUnhealthy: true, failUnstable: false, lineCoverageTargets: '30, 0, 0', maxNumberOfBuilds: 0, methodCoverageTargets: '30, 0, 0', onlyStable: false, sourceEncoding: 'ASCII', zoomCoverageChart: false
          sh 'mv output/c.out .'
          codacy action: 'reportCoverage', filePath: "output/coverage.xml"
        }
      }
    }
  }

  post {
    always {
      releaseInfraPoolAgent(".infrapool/release_agents")
    }
  }
}
