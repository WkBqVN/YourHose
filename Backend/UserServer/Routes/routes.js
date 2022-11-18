const express = require('express')
const app = express()
const port = 4050

app.get('/', (req, res) => {
    console.log(req)
    res.send('Hello World!')
})

app.listen(port, () => {
    console.log(`Example app listening on port ${port}`)
})
