const express = require('express')
const app = express()
const port = 3000

app.use('/', express.static('./public'))

app.get('/*', (req, res) => {
    res.json({message: 'Imagine all sorts of cool, custom views here. 🌈✨🌈✨'});
});

app.listen(port, () => console.log(`Example app listening on port ${port}!`))