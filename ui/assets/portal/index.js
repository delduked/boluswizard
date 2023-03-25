const Chartist = require("chartist");
const moment = require("moment");
class ProjectInput {
    constructor() {
        //this.signin()
        // input
        this.bg = document.getElementById('bg');
        this.carbs = document.getElementById('carbs');
        this.ratios = document.getElementById('ratios');
        this.isfs = document.getElementById('isfs');
        this.targets = document.getElementById('targets');
        this.Editor = document.getElementById('editor');
        // output
        this.CarbCorrection = document.getElementById('CarbCorrection');
        this.BgCorrection = document.getElementById('BGCorrection');
        this.ActiveInsulinReduction = document.getElementById('InsulinReduction');
        this.Bolus = document.getElementById('Result');
        this.EditorTitle = document.querySelector(".window-header h3");
        // buttons
        this.close = document.getElementById('close');
        this.modal = document.getElementById('modal');
        this.update = document.getElementById('update');
        this.success = document.getElementById('success');
        this.save = document.getElementById('save');
        this.drawGraph();
        this.configure();
        console.log("class initatilized");
    }
    configure() {
        this.bg.addEventListener('keypress', (event) => {
            if (event.key === "Enter") {
                this.validateInput() ? this.calculateCorrection() : console.log('invalid input');
                event.preventDefault();
            }
        });
        this.carbs.addEventListener('keypress', (event) => {
            if (event.key === "Enter") {
                this.validateInput() ? this.calculateCorrection() : console.log('invalid input');
                event.preventDefault();
            }
        });
        this.close.addEventListener('click', event => {
            this.modal.classList.remove("show");
            this.modal.classList.add("hide");
        });
        this.targets.addEventListener('click', event => {
            console.log(event);
            this.initEditor("Targets");
            this.EditorTitle.textContent = "BG Targets";
            this.modal.classList.remove("hide");
            this.modal.classList.add("show");
            this.clearEditor();
        });
        this.ratios.addEventListener('click', event => {
            console.log(event);
            this.initEditor("Ratios");
            this.EditorTitle.textContent = "Carb Ratios";
            this.modal.classList.remove("hide");
            this.modal.classList.add("show");
            this.clearEditor();
        });
        this.isfs.addEventListener('click', event => {
            console.log(event);
            this.initEditor("ISFs");
            this.EditorTitle.textContent = "Insulin Sensitivity Factors";
            this.modal.classList.remove("hide");
            this.modal.classList.add("show");
            this.clearEditor();
        });
        this.update.addEventListener('click', event => {
            console.log(event);
            this.updateEditor();
        });
        this.save.addEventListener('click', event => {
            console.log(event);
            this.validateInput() ? this.saveCorrection() : console.log('invalid input');
            this.graph.clear();
            this.drawGraph();
        });
    }
    validateInput() {
        // validate input
        console.log("validating");
        if (!this.bg.value || this.bg.value == '' || +this.bg.value == 0 || !this.carbs.value || this.carbs.value == '' || +this.carbs.value == 0) {
            return false;
        }
        return true;
    }
    async calculateCorrection() {
        try {
            const response = await fetch(`http://localhost:81/wizard/NewCorrection?Bg=${this.bg.value}&Carbs=${this.carbs.value}`, {
                method: "GET",
                credentials: "include"
            })
                .then(data => data.json())
                .catch(err => err);
            console.log(response);
            const CorrectionResponse = await response;
            if (CorrectionResponse.Status == 200) {
                this.CarbCorrection.innerHTML = CorrectionResponse.Data.CarbCorrection.toString();
                this.BgCorrection.innerHTML = CorrectionResponse.Data.BgCorrection.toString();
                this.ActiveInsulinReduction.innerHTML = CorrectionResponse.Data.ActiveInsulinReduction.toString();
                this.Bolus.innerHTML = CorrectionResponse.Data.Bolus.toString();
            }
            else {
                console.log(CorrectionResponse);
                throw Error("Ooops...");
            }
        }
        catch (error) {
            console.log(error);
        }
    }
    async drawGraph() {
        try {
            const range = await this.getRange();
            console.log(range);
            let coordinates;
            if (range.data == null) {
                coordinates = [{
                        x: new Date(),
                        y: 0
                    },
                    {
                        x: new Date(),
                        y: 0
                    }];
            }
            else {
                coordinates = range.map((data) => {
                    return {
                        x: new Date(data.TimeStamp / 1000000),
                        y: data.Bg
                    };
                });
            }
            let options = await {
                low: 0,
                showArea: true
            };
            let chartType = await {
                axisX: {
                    type: Chartist.FixedScaleAxis,
                    divisor: 5,
                    labelInterpolationFnc: function (value) {
                        return moment(value).format('HH:mm:ss');
                    }
                }
            };
            let data = await {
                series: [{
                        name: 'series-2',
                        data: coordinates,
                    }
                ]
            };
            this.graph = await new Chartist.Line('#my-chart', data, chartType, options);
        }
        catch (error) {
            console.log(error);
        }
    }
    async saveCorrection() {
        try {
            var raw = await JSON.stringify([
                {
                    Bg: +this.bg.value,
                    Bolus: +this.Bolus.textContent,
                    Carbs: +this.carbs.value,
                }
            ]);
            const res = await fetch('http://localhost:81/wizard/Corrections', {
                method: "POST",
                credentials: "include",
                body: raw,
                headers: { "content-type": "application/json" },
            });
            const data = await res.json();
            console.log(`correction: ${raw}`);
            console.log(data);
        }
        catch (error) {
            console.log(error);
        }
    }
    async getRange() {
        try {
            const currentNanoseconds = Date.now() * 1000000;
            // two days 172800000000000
            // const RangeResponse = await fetch(`http://localhost:81/wizard/CorrectionRange?Start=1653705254489676000&End=1653705256507473800`,{
            const res = await fetch(`http://localhost:81/wizard/CorrectionRange?Start=${currentNanoseconds - 172800000000000}&End=${currentNanoseconds}`, {
                method: "GET",
                credentials: "include"
            })
                .then(data => data.json())
                .catch(err => err);
            if (res.Status != 200) {
                throw Error(res.Error);
            }
            console.log(res);
            return res.Data;
        }
        catch (error) {
            console.log(error);
            throw error;
        }
    }
    async signin() {
        try {
            const token = await fetch("http://localhost:81/Signin", {
                method: "POST",
                body: JSON.stringify({ "Username": "nate@nated.ca", "Password": "n4th4n43l" }),
                headers: { "content-type": "application/json" },
                credentials: "include"
            })
                .then(data => {
                console.log(data);
                return data.json();
            });
        }
        catch (error) {
            console.log(error);
        }
    }
    async get(info) {
        try {
            const res = await fetch(`http://localhost:81/wizard/${info}`, {
                method: "GET",
                credentials: "include"
            })
                .then(data => data.json())
                .catch(err => err);
            if (res.Status != 200) {
                throw Error(res.Error);
            }
            res.Data.forEach(object => {
                delete object['Key'];
            });
            console.log(res.Data);
            return res.Data;
        }
        catch (error) {
            console.log(error);
            throw error;
        }
    }
    async post(info, data) {
        try {
            const res = await fetch(`http://localhost:81/wizard/${info}`, {
                method: "POST",
                credentials: "include",
                body: data,
                headers: { "content-type": "application/json" },
            })
                .then(data => data.json())
                .catch(err => err);
            if (res.Status != 200) {
                throw Error(res.Error);
            }
            console.log(res.Data);
            return res.Data;
        }
        catch (error) {
            console.log(error);
            throw error;
        }
    }
    async initEditor(info) {
        try {
            const res = await this.get(info);
            console.log(`initEditor: ${info}`);
            this.success.textContent = "";
            console.log(res);
            this.update.dataset.update = info;
            this.Editor.value = JSON.stringify(res, null, 4);
        }
        catch (error) {
            this.success.textContent = "Oops! Something went wrong...";
            console.log(error);
        }
    }
    async updateEditor() {
        try {
            const endpoint = this.update.dataset.update;
            await this.post(endpoint, this.Editor.value);
            this.success.textContent = "Success!";
        }
        catch (error) {
            console.log(error);
            this.success.textContent = "Oops! Something went wrong...";
        }
    }
    clearEditor() {
        this.Editor.textContent = "";
    }
}
const asdf = new ProjectInput();
