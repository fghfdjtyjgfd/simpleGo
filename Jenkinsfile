pipeline {
    agent { 
        node {label 'slave-2'} 
    }

    environment {
        BRANCH = "${env.BRANCH_NAME}"
    }

    stages {
        stage('Test') {
            // it will be run test if it is a dev or master.
            when {
                expression {
                    BRANCH == 'dev' || BRANCH == 'main'
                }
            }
            steps {
                echo 'testing the application ...'
            }
        }
        stage('Prepare job definition') {
            steps {
                script {
                    echo 'Pulling...' + env.BRANCH_NAME
                    withDockerRegistry(credentialsId: 'd2330c30-d21b-4670-84ed-37dd988b506a') {
                        sh '''
                            docker build --no-cache -t unknopu/test${BRANCH} .
                            docker push unknopu/test${BRANCH}
                        '''
                    }
                }
            }
            post {
                always { 
                    cleanWs()
                }
                success {
                    echo 'SUCCESS ...!!!'
                }
                failure {
                    echo 'ITS FAILED, NOOB !!!'
                }
            }
        }
    }
}


