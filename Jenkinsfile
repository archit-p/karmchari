pipeline {
    agent any
    tools {
        go 'go-1.14.7'
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
