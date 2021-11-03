node('node') {
    stage('Checkout') {
      lastStage = env.STAGE_NAME
      checkout scm
    }

    stage('Test') {
      sh './build -v test'
    }
}
