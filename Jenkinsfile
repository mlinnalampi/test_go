pipeline {
	agent any
	stages {
		stage('Checkout') {
			steps {
				checkout scm
			}
		}
		stage('Build') {
			steps {
				sh 'go build'
			}
		}
		stage('Unit test') {
			steps {
				sh 'go test'
			}
		}
	}
}
