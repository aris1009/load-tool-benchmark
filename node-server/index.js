const app = require('express')();

let reqCount = 0;

app.get('/', (req, res) => {
    console.log(`req incoming ${Math.random().toString()}`)
    reqCount++;
    res.end('OK')
});

app.listen(3005, () => console.log(`Server listens on port: 3005`));

process.on('SIGINT', () => {
    console.log('\n\nTotal request count: ', reqCount)
    console.log('Exiting...')
    process.exit()
})