//import { IUser } from "http://localhost:80/schemas/i-user";

interface iCredentials {
    Username: string;
    Password: string;
}

class LoginPage {
  
  // input
  username: HTMLInputElement;
  password: HTMLInputElement;
  login: HTMLInputElement;
  message: HTMLInputElement;
  signup : HTMLElement;

  // api service
  apiService: string;

  constructor(){

    this.username = document.getElementById('username')! as HTMLInputElement;
    this.password = document.getElementById('password')! as HTMLInputElement;
    this.message = document.getElementById('message')! as HTMLInputElement;
    this.login = document.getElementById('login')! as HTMLInputElement;
    this.signup = document.getElementById('signup')! as HTMLElement;

    // api service
    this.apiService = document.querySelector("body").getAttribute('data-api-service')! as string;
    console.log(this.apiService)
    this.listen();
  }

  // login
  private async Login(){
    try {
      // get user credentials
      const creds: iCredentials = await {
        Username: this.username.value,
        Password: this.password.value
      }

      // setup post request for login
      const res = await fetch(`${this.apiService}SignIn`,{
        method: "POST",
        body: JSON.stringify(creds),
        headers: {"content-type":"application/json"}
      });
      
      if (res.ok) {
        
        // assign token to local storage
        const responseJson = await res.json()
        console.log(responseJson.Data)
        setTimeout(() => {document.cookie = document.cookie +`authToken=${responseJson.Data.Token};path=/wizard`},1000);
        setTimeout(() => {document.cookie = document.cookie +`authToken=${responseJson.Data.Token};path=/u/${responseJson.Data.Username}`},1000);
        this.updateMessage("Login successful!");

        setTimeout(() => {document.location.href = `/u/${responseJson.Data.Username}/home`},1000)
      }

      
    } catch (error) {

      // update message
      this.updateMessage("Oops! Something went wrong.", error);
      
    }
  }

  // update message
  private updateMessage(message: string, error?: any) {
    console.log(error)
    console.log(message);
    this.message.textContent = message;
  }

  // Validate user input
  private validateInput(): boolean {
    // validate input
    console.log("validating");
    if ( !this.username.value || this.username.value == '' || !this.password.value || this.password.value == '') {
      return false;
    } 
    
    return true
  }

  private goToSignUpPage(){
    document.location.href = "/SignUp";
  }

  // Listen for button clicks
  private listen() {
    console.log("Listening for user input...")
    this.login.addEventListener('click', event => {
        
        // validate user input
        this.validateInput() ? this.Login() : this.updateMessage("That's some invalid input mate!");
        event.preventDefault();
    });
    this.signup.addEventListener('click', event => {
        
      this.goToSignUpPage();
    });
  }

}

const asdf = new LoginPage();