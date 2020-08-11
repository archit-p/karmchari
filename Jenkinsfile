pipeline {
    agent any
    tools {
        go 'alpine'
    }
    environment {
        GO111MODULE = 'on'
    }
	stages {
		stage('Compile') {
			steps {
				sh 'go build'
			}
		}
	}
}
