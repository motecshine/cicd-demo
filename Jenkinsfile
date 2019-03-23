
node('go-jnlp') {
    stage('Git CLone') {
        checkout scm
        script {
            build_tag = sh(returnStdout: true, script: 'git rev-parse --short HEAD').trim()
        }
    }
    
    stage('Build Binary') {
        sh "cd ${WORKSPACE} && /usr/local/go/bin/go build -o cicd-demo ."
    }

    stage('Build Docker Image') {
        sh "docker build -t motecshine/cicd-demo-${env.BRANCH_NAME}:${build_tag} ."
    }

    stage('Commit Docker Image') {
        withCredentials([usernamePassword(credentialsId: 'DockerHub', passwordVariable: 'dockerHubPassword', usernameVariable: 'dockerHubUser')]) {
            sh "docker login -u ${dockerHubUser} -p ${dockerHubPassword}"
            sh "docker push motecshine/default/cicd-demo-${env.BRANCH_NAME}:${build_tag}"
        }
    }
    // 在这里将会 批量sed k8s config yaml 来适配当前版本, 要么replace 要么 create
    stage('Notify K8S Deployagent') {     
    }
}