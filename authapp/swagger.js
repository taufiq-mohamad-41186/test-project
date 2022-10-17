const swaggerAutogen = require('swagger-autogen')()

const outputFile = './docs/swagger_output.json'
const endpointsFiles = ['./handler/rest.js']

swaggerAutogen(outputFile, endpointsFiles).then(() => {
    require('./index.js')
})