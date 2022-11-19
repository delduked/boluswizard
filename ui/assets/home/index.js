// const Chartist = require('chartist');
// const moment = require('moment');
class ProjectInput {
    constructor() {
        // save buttons
        this.SaveIsf = document.getElementById('SaveIsf');
        this.SaveRatio = document.getElementById('SaveRatio');
        this.SaveTarget = document.getElementById('SaveTarget');
        this.SaveDuration = document.getElementById('SaveDuration');
        this.SaveCorrection = document.getElementById('SaveCorrection');
        // edit
        this.bg = document.getElementById('bg');
        this.carbs = document.getElementById('carbs');
        this.EditIsf = document.getElementById('EditIsf');
        this.EditRatio = document.getElementById('EditRatio');
        this.EditTarget = document.getElementById('EditTarget');
        this.EditDuration = document.getElementById('EditDuration');
        this.EditCorrection = document.getElementById('EditCorrection');
        this.GetCorrection = document.getElementById('GetCorrection');
        // output
        this.CarbCorrection = document.getElementById('CarbCorrection');
        this.BgCorrection = document.getElementById('BGCorrection');
        this.ActiveInsulinReduction = document.getElementById('ActiveInsulinReduction');
        this.Bolus = document.getElementById('Bolus');
        this.EditorTitle = document.querySelector(".window-header h3");
        // toast
        this.ToastError = document.getElementById('Error');
        this.ToastSuccess = document.getElementById('Success');
        this.ToastErrorMessage = document.getElementById('Errormessage');
        // signout
        this.SignOut = document.querySelector('[href="/login"]');
        // nav
        this.navBtns = document.querySelectorAll('.nav-btn');
        this.initSettings();
        this.drawGraph();
        this.configure();
        console.log("class initatilized");
    }
    configure() {
        this.GetCorrection.addEventListener('click', event => {
            this.validateInput() ? this.calculateCorrection() : console.log('invalid input');
            event.preventDefault();
        });
        this.SaveTarget.addEventListener('click', event => {
            console.log(event);
            const data = this.post('Targets', this.EditTarget.value);
            if (data instanceof Error) {
                this.EditTarget.value = JSON.stringify(data, null, 4);
            }
        });
        this.SaveRatio.addEventListener('click', event => {
            console.log(event);
            const data = this.post('Ratios', this.EditRatio.value);
            if (data instanceof Error) {
                this.EditRatio.value = JSON.stringify(data, null, 4);
            }
        });
        this.SaveIsf.addEventListener('click', event => {
            const data = this.post('ISFs', this.EditIsf.value);
            if (data instanceof Error) {
                console.log(event);
                this.EditIsf.value = JSON.stringify(data, null, 4);
            }
        });
        this.SaveDuration.addEventListener('click', event => {
            const data = this.post('Duration', this.EditDuration.value);
            if (data instanceof Error) {
                console.log(event);
                this.EditDuration.value = JSON.stringify(data, null, 4);
            }
        });
        this.SaveCorrection.addEventListener('click', event => {
            console.log(event);
            this.validateInput() ? this.saveCorrection() : console.log('invalid input');
            location.reload();
        });
        this.navBtns.forEach(ele => {
            ele.addEventListener('click', function () {
                console.log("nav btn clicked");
                // Get the dropdown-menu associated with this nav button (insert the id of your menu)
                let dropDownMenu = document.getElementById('header-menu');
                // Toggle the nav-btn and the dropdown menu
                ele.classList.toggle('active');
                dropDownMenu.classList.toggle('active');
            });
            this.SignOut.addEventListener('click', event => {
                document.cookie = document.cookie + "max-age=0;";
            });
        });
    }
    async initSettings() {
        try {
            const isf = await this.get('ISFs');
            const ratio = await this.get('Ratios');
            const target = await this.get('Targets');
            const duration = await this.get('Duration');
            this.EditIsf.textContent = await JSON.stringify(isf, null, 4);
            this.EditRatio.textContent = await JSON.stringify(ratio, null, 4);
            this.EditTarget.textContent = await JSON.stringify(target, null, 4);
            this.EditDuration.textContent = await JSON.stringify(duration, null, 4);
        }
        catch (error) {
            console.log(error);
        }
    }
    validateInput() {
        // validate input
        if (!this.bg.value || this.bg.value == '' || +this.bg.value == 0 || !this.carbs.value || this.carbs.value == '' || +this.carbs.value == 0) {
            return false;
        }
        return true;
    }
    async calculateCorrection() {
        try {
            const response = await this.get(`NewCorrection?Bg=${this.bg.value}&Carbs=${this.carbs.value}`);
            this.CarbCorrection.innerHTML = response.CarbCorrection.toString();
            this.BgCorrection.innerHTML = response.BgCorrection.toString();
            this.ActiveInsulinReduction.innerHTML = response.ActiveInsulinReduction.toString();
            this.Bolus.innerHTML = response.Bolus.toString();
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
            if (range == null) {
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
            const res = await this.post('Corrections', raw);
        }
        catch (error) {
            console.log(error);
        }
    }
    async saveDuration() {
        try {
        }
        catch (error) {
        }
    }
    async getRange() {
        try {
            const currentNanoseconds = Date.now() * 1000000;
            // two days 172800000000000
            const res = await this.get(`CorrectionRange?Start=${currentNanoseconds - 172800000000000}&End=${currentNanoseconds}`);
            return res;
        }
        catch (error) {
            throw error;
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
            console.log(res.Data);
            return res.Data;
        }
        catch (error) {
            console.log(error);
            throw error;
        }
    }
    showErrorPopup(error) {
        this.ToastError.style.bottom = '0px';
        this.ToastErrorMessage.innerHTML = error;
        setTimeout(() => {
            this.ToastError.style.bottom = '-200px';
        }, 5000);
    }
    showSuccessPopup() {
        this.ToastSuccess.style.bottom = '0px';
        setTimeout(() => {
            this.ToastSuccess.style.bottom = '-200px';
        }, 5000);
    }
}
const asdf = new ProjectInput();
