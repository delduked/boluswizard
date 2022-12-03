//import { IUser } from "http://localhost:80/schemas/i-user";
class LoginPage {
    constructor() {
        this.username = document.getElementById('username');
        this.password = document.getElementById('password');
        this.message = document.getElementById('message');
        this.login = document.getElementById('login');
        this.signup = document.getElementById('signup');
        // api service
        this.apiService = document.querySelector("body").getAttribute('data-api-service');
        console.log(this.apiService);
        this.listen();
    }
    // login
    async Login() {
        try {
            // get user credentials
            const creds = await {
                Username: this.username.value,
                Password: this.password.value
            };
            // setup post request for login
            const res = await fetch(`${this.apiService}SignIn`, {
                method: "POST",
                body: JSON.stringify(creds),
                credentials:'include',
                headers: { "content-type": "application/json" }
            });
            if (res.ok) {
                // assign token to local storage
                const responseJson = await res.json();
                console.log(responseJson.Data);
                //setTimeout(() => { document.cookie = document.cookie + `authToken=${responseJson.Data.Token};path=/wizard`; }, 500);
                //setTimeout(() => { document.cookie = document.cookie + `authToken=${responseJson.Data.Token};path=/u/${responseJson.Data.Username}`; }, 1000);
                this.updateMessage("Login successful!");
                setTimeout(() => { document.location.href = `/u/${responseJson.Data.Username}/home`; }, 1500);
            }
        }
        catch (error) {
            // update message
            this.updateMessage("Oops! Something went wrong.", error);
        }
    }
    // update message
    updateMessage(message, error) {
        console.log(error);
        console.log(message);
        this.message.textContent = message;
    }
    // Validate user input
    validateInput() {
        // validate input
        console.log("validating");
        if (!this.username.value || this.username.value == '' || !this.password.value || this.password.value == '') {
            return false;
        }
        return true;
    }
    goToSignUpPage() {
        document.location.href = "/SignUp";
    }
    // Listen for button clicks
    listen() {
        console.log("Listening for user input...");
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
