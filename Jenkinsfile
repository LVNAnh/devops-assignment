pipeline {
    agent any

    environment {
        DOCKER_IMAGE = 'lvnanh/server_assignment'
        DOCKER_TAG = 'latest'
        PROD_SERVER = 'ec2-47-129-236-253.ap-southeast-1.compute.amazonaws.com'
        PROD_USER = 'ubuntu'
        EMAIL_RECIPIENT = 'levunhatanh1997@gmail.com'
    }

    stages {
        stage('Clone Repository') {
            steps {
                git branch: 'master', url: 'https://github.com/LVNAnh/devops-assignment.git'
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    echo 'Building Docker image for linux/amd64 platform...'
                    docker.build("${DOCKER_IMAGE}:${DOCKER_TAG}", "--platform linux/amd64 .")
                }
            }
        }

        stage('Run Tests') {
            steps {
                echo 'Running tests...'
            }
        }

        stage('Push to Docker Hub') {
            steps {
                script {
                    docker.withRegistry('https://index.docker.io/v1/', 'docker-hub-credentials') {
                        docker.image("${DOCKER_IMAGE}:${DOCKER_TAG}").push()
                    }
                }
            }
        }

        stage('Deploy Golang to DEV') {
            steps {
                script {
                    echo 'Clearing server_golang-related images and containers...'
                    sh '''
                        docker container stop server-golang || echo "No container named server-golang to stop"
                        docker container rm server-golang || echo "No container named server-golang to remove"
                        docker image rm ${DOCKER_IMAGE}:${DOCKER_TAG} || echo "No image ${DOCKER_IMAGE}:${DOCKER_TAG} to remove"
                    '''

                    echo 'Deploying to DEV environment...'
                    sh '''
                        docker image pull ${DOCKER_IMAGE}:${DOCKER_TAG}

                        # Kiểm tra và tạo network nếu cần
                        docker network inspect dev >/dev/null 2>&1 || docker network create dev

                        docker container run -d --rm --name server-golang -p 3006:3005 --network dev ${DOCKER_IMAGE}:${DOCKER_TAG}
                    '''
                }
            }
        }

        stage('Deploy to Production on AWS') {
            steps {
                script {
                    echo 'Deploying to Production...'
                    sshagent(['aws-ssh-key']) {
                        sh '''
                            ssh -o StrictHostKeyChecking=no ${PROD_USER}@${PROD_SERVER} << EOF
                            docker container stop server-golang || echo "No container to stop"
                            docker container rm server-golang || echo "No container to remove"
                            docker image rm ${DOCKER_IMAGE}:${DOCKER_TAG} || echo "No image to remove"
                            docker image pull ${DOCKER_IMAGE}:${DOCKER_TAG} || { echo "Failed to pull image"; exit 1; }
                            docker container run -d --rm --name server-golang -p 3006:3005 ${DOCKER_IMAGE}:${DOCKER_TAG} || { echo "Failed to run container"; exit 1; }
EOF
                        '''
                    }
                }
            }
        }
    }

    post {
        success {
            emailext(
                subject: 'Jenkins Build Successful: ${JOB_NAME} #${BUILD_NUMBER}',
                body: '''<p>Good news!</p>
                         <p>The build ${JOB_NAME} #${BUILD_NUMBER} was successful.</p>
                         <p>Check details at: <a href="${BUILD_URL}">${BUILD_URL}</a></p>''',
                to: "${EMAIL_RECIPIENT}"
            )
        }
        failure {
            emailext(
                subject: 'Jenkins Build Failed: ${JOB_NAME} #${BUILD_NUMBER}',
                body: '''<p>Unfortunately, the build ${JOB_NAME} #${BUILD_NUMBER} failed.</p>
                         <p>Check details at: <a href="${BUILD_URL}">${BUILD_URL}</a></p>''',
                to: "${EMAIL_RECIPIENT}"
            )
        }
        always {
            cleanWs()
        }
    }
}
