// const Chartist = require('chartist');
// const moment = require('moment');

interface iCorrectionResponse {
  Status: number;
  Error: any;
  Data: iCorrectionData;
}
interface iResponse {
  Status: number;
  Error: any;
  Data: any;
}
interface iCorrectionData {
  BgCorrection: number;
  CarbCorrection: number;
  ActiveInsulinReduction: number;
  Bolus: number;
}
interface iRangeResponse {
  Status: number;
  Error: any;
  Data: iRangeData[];
}
interface iRangeData {
  Key: string;
  Bg: number;
  Carbs: number;
  Bolus: number;
  TimeStamp: number;
}

interface iRatio {
  Key?: string;
  End: string
  Start: string
  Ratio: number;
}
interface iIsf {
  Key?: string;
  End: string
  Start: string
  Sensitivity: number;
}
interface duration {
  Duration: string;

}
interface iTarget {
  Key?: string;
  End: string
  Start: string
  High: number;
  Low: number;
  Sensitivity: number;
}


interface iCoordinates {
  x: any;
  y: number;
}

class ProjectInput {

  // input
  bg: HTMLInputElement;
  carbs: HTMLInputElement;

  // output
  CarbCorrection: HTMLHtmlElement;
  BgCorrection: HTMLHtmlElement;
  ActiveInsulinReduction: HTMLHtmlElement;
  Bolus: HTMLHtmlElement;
  
  EditorTitle: HTMLHtmlElement;
  
  // save buttons
  SaveIsf: HTMLInputElement;
  SaveRatio: HTMLInputElement;
  SaveTarget: HTMLInputElement;
  SaveDuration: HTMLInputElement;
  SaveCorrection: HTMLInputElement;
  GetCorrection: HTMLInputElement;
  
  // edit
  EditIsf: HTMLInputElement;
  EditRatio: HTMLInputElement;
  EditTarget: HTMLInputElement;
  EditDuration: HTMLInputElement;
  EditCorrection: HTMLInputElement;

  // toast
  ToastError: HTMLHtmlElement;
  ToastErrorMessage: HTMLHtmlElement;
  ToastSuccess: HTMLHtmlElement;

  // signout
  SignOut: HTMLAnchorElement;

  // nav
  navBtns: any;

  // api service
  apiService: String;

  graph: any;

  constructor(){
    
    // save buttons
    this.SaveIsf = document.getElementById('SaveIsf')! as HTMLInputElement;
    this.SaveRatio = document.getElementById('SaveRatio')! as HTMLInputElement;
    this.SaveTarget = document.getElementById('SaveTarget')! as HTMLInputElement;
    this.SaveDuration = document.getElementById('SaveDuration')! as HTMLInputElement;
    this.SaveCorrection = document.getElementById('SaveCorrection')! as HTMLInputElement;
    
    // edit
    this.bg = document.getElementById('bg')! as HTMLInputElement;
    this.carbs = document.getElementById('carbs')! as HTMLInputElement;
    this.EditIsf = document.getElementById('EditIsf')! as HTMLInputElement;
    this.EditRatio = document.getElementById('EditRatio')! as HTMLInputElement;
    this.EditTarget = document.getElementById('EditTarget')! as HTMLInputElement;
    this.EditDuration = document.getElementById('EditDuration')! as HTMLInputElement;
    this.EditCorrection = document.getElementById('EditCorrection')! as HTMLInputElement;
    
    this.GetCorrection = document.getElementById('GetCorrection')! as HTMLInputElement;

    // output
    this.CarbCorrection = document.getElementById('CarbCorrection')! as HTMLHtmlElement;
    this.BgCorrection = document.getElementById('BGCorrection')! as HTMLHtmlElement;
    this.ActiveInsulinReduction = document.getElementById('ActiveInsulinReduction')! as HTMLHtmlElement;
    this.Bolus = document.getElementById('Bolus')! as HTMLHtmlElement;
    this.EditorTitle = document.querySelector(".window-header h3")! as HTMLHtmlElement;

    // toast
    this.ToastError = document.getElementById('Error')! as HTMLHtmlElement;
    this.ToastSuccess = document.getElementById('Success')! as HTMLHtmlElement;
    this.ToastErrorMessage = document.getElementById('Errormessage')! as HTMLHtmlElement;

    // signout
    this.SignOut = document.querySelector('[href="/login"]') as HTMLAnchorElement;

    // nav
    this.navBtns = document.querySelectorAll('.nav-btn')! as NodeListOf<Element>;

    // api service
    this.apiService = document.querySelector("body").getAttribute('data-api-service')! as string;

    this.initSettings()
    this.drawGraph()
    this.configure();
    console.log("class initatilized")
  }

  public configure() {
    this.GetCorrection.addEventListener('click', event => {

        this.validateInput() ? this.calculateCorrection() : console.log('invalid input');
        event.preventDefault();
    });

    this.SaveTarget.addEventListener('click', event => {
        console.log(event)

        const data = this.post<iTarget[]>('Targets',this.EditTarget.value)
        if (data !instanceof Error){
              this.EditTarget.value = JSON.stringify(data, null, 4);
        }
    });
    this.SaveRatio.addEventListener('click', event => {
        console.log(event)

        const data = this.post<iRatio[]>('Ratios',this.EditRatio.value)
        if (data !instanceof Error){
              this.EditRatio.value = JSON.stringify(data, null, 4);
        }
    });
    this.SaveIsf.addEventListener('click', event => {
            const data = this.post<iIsf[]>('ISFs',this.EditIsf.value)
            if (data !instanceof Error){
                  console.log(event)
                  this.EditIsf.value = JSON.stringify(data, null, 4);
            }
    });
    this.SaveDuration.addEventListener('click', event => {
            const data = this.post<duration>('Duration',this.EditDuration.value)
            if (data !instanceof Error){
                  console.log(event)
                  this.EditDuration.value = JSON.stringify(data, null, 4);
            }
    });
    this.SaveCorrection.addEventListener('click', event => {
        console.log(event);

        this.validateInput() ? this.saveCorrection() : console.log('invalid input');
        location.reload();
        
    });

    this.navBtns.forEach( ele => {
      ele.addEventListener('click', function() {
            console.log("nav btn clicked")
          // Get the dropdown-menu associated with this nav button (insert the id of your menu)
          let dropDownMenu = document.getElementById('header-menu')! as HTMLHtmlElement;
  
          // Toggle the nav-btn and the dropdown menu
          ele.classList.toggle('active');
          dropDownMenu.classList.toggle('active');
      });
    
    this.SignOut.addEventListener('click', event =>{
      document.cookie = document.cookie + "max-age=0;"
    })
  });

  }

  public async initSettings() {
      try {
            const isf = await this.get<iIsf[]>('ISFs')
            const ratio = await this.get<iRatio[]>('Ratios')
            const target = await this.get<iTarget[]>('Targets')
            const duration = await this.get<duration[]>('Duration')
      
            this.EditIsf.textContent = await JSON.stringify(isf, null, 4);
            this.EditRatio.textContent = await JSON.stringify(ratio, null, 4);
            this.EditTarget.textContent = await JSON.stringify(target, null, 4);
            this.EditDuration.textContent = await JSON.stringify(duration, null, 4);
      } catch (error) {
            console.log(error)
      }

  }

  public validateInput(): boolean {
    // validate input
    if ( !this.bg.value || this.bg.value == '' || +this.bg.value == 0 || !this.carbs.value || this.carbs.value == '' || +this.carbs.value == 0) {
      return false;
    } 
    
    return true
  }

  public async calculateCorrection(){
    try {
      const response = await this.get<iCorrectionData>(`NewCorrection?Bg=${this.bg.value}&Carbs=${this.carbs.value}`)

      this.CarbCorrection.innerHTML = response.CarbCorrection.toString();
      this.BgCorrection.innerHTML = response.BgCorrection.toString();
      this.ActiveInsulinReduction.innerHTML = response.ActiveInsulinReduction.toString();
      this.Bolus.innerHTML = response.Bolus.toString();

    } catch (error) {
      console.log(error)
    }
  }

  public async drawGraph() {
    try {

      const range: iRangeData[] = await this.getRange();
      console.log(range);

      let coordinates: iCoordinates[];
      if (range == null) {

        coordinates = [{
            x: new Date(),
            y: 0
        },
        {
          x: new Date(),
          y: 0
        }]

      } else {

        coordinates = range.map((data: iRangeData) => {
          return {
            x: new Date(data.TimeStamp / 1000000),
            y: data.Bg
          }
        })

      }

      let options = await {
        low: 0,
        showArea: true
      }
      let chartType = await {
        axisX: {
          type: Chartist.FixedScaleAxis,
          divisor: 5,
          labelInterpolationFnc: function (value: any) {
            return moment(value).format('HH:mm:ss');
          }
        }
      }
      let data = await {
        series: [{
          name: 'series-2',
          data: coordinates,
        }
        ]
      }

     this.graph = await new Chartist.Line('#my-chart', data, chartType, options);
    } catch (error) {
      console.log(error)
    }
  }

  public async saveCorrection() {
    try {
      
      var raw = await JSON.stringify([
        {
          Bg: +this.bg.value,
          Bolus: +this.Bolus.textContent!,
          Carbs: +this.carbs.value,
        }
      ]);
      
      const res = await this.post<iResponse>('Corrections',raw)

      
    } catch (error) {
      console.log(error)
    }
  }

  public async saveDuration(){
    try {
      
    } catch (error) {
      
    }
  }

  public async getRange(): Promise<iRangeData[]> {
    try {
      const currentNanoseconds = Date.now() * 1000000;
      // two days 172800000000000
      const res = await this.get<iRangeData[]>(`CorrectionRange?Start=${currentNanoseconds - 172800000000000}&End=${currentNanoseconds}`);

      return res;
    } catch (error) {
      throw error
    }
  }

  public async get<t>(info: string): Promise<t> {
    try {
      const res = await fetch(`${this.apiService}wizard/${info}`,{
        method: "GET",
        credentials: "include"
      })
      .then(data => data.json())
      .catch(err => err);

      console.log(res.Data);
      return res.Data;
    } catch (error) {
      console.log(error);
      throw error;
    }
  }
  public async post<t>(info: string, data: string): Promise<t> {
    try {
      const res = await fetch(`${this.apiService}wizard/${info}`,{
        method: "POST",
        credentials: "include",
        body: data,
        headers: {"content-type":"application/json"},
      })
      .then(data => data.json())
      .catch(err => err);

      console.log(res.Data);
      return res.Data;
    } catch (error) {
      console.log(error);
      throw error;
    }
  }

  public showErrorPopup(error: any) {
    this.ToastError.style.bottom = '0px';
    this.ToastErrorMessage.innerHTML = error;
    setTimeout(() => {
      this.ToastError.style.bottom = '-200px';
    }, 5000);
  }
  public showSuccessPopup() {
    this.ToastSuccess.style.bottom = '0px';
    setTimeout(() => {
      this.ToastSuccess.style.bottom = '-200px';
    }, 5000);
  }

}

const asdf = new ProjectInput();