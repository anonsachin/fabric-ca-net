const express = require('express');
const app = express();
const cors = require('cors');
const port = 3000;

//requirements
const addUsage = require("./addUsage.js");
const getUsage = require("./getUsage.js");

// server
app.use(cors());
app.use(express.json());
app.use(express.urlencoded({extended: true}));
app.set('title','Counterfeit Drug Detection');

app.get('/',(req,res) => res.send('Counterfeit Drug Detection Server UP!!'));

app.get('/getUsage',(req,res) => {
    getUsage.run(req.body.org,req.body.channelName,req.body.contractName,req.body.cpuName)
        .then(cpu => {
            console.log("Got the cpu!")
            const result = {
                Status: "Success",
                Payload: cpu,
            };
            res.status(200).json(result);
        })
        .catch(e => {
            const result = {
                Status: "Faliure",
                Error: e,
            };
            res.status(500).send(result)
        })
});

app.post('/addUsage',(req,res) => {
    addUsage.run(req.body.org,req.body.channelName,req.body.contractName,req.body.cpuName,req.body.value1,req.body.value2)
        .then(cpu => {
            console.log("Got the cpu!")
            const result = {
                Status: "Success",
                Payload: cpu,
            };
            res.status(200).json(result);
        })
        .catch(e => {
            const result = {
                Status: "Faliure",
                Error: e,
            };
            res.status(500).send(result)
        })
});


// startup
app.listen(port, () => console.log("Counterfeit Drug Detection App listening ..."));