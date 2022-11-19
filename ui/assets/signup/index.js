//import { IUser } from "http://localhost:80/schemas/i-user";
class SignUp {
    constructor() {
        this.username = document.getElementById('username');
        this.password = document.getElementById('password');
        this.message = document.getElementById('message');
        this.signup = document.getElementById('signup');
        this.login = document.getElementById('login');
        this.listen();
    }
    // login
    async Signup() {
        try {
            // get user credentials
            const creds = await {
                Username: this.username.value,
                Password: this.password.value
            };
            // setup post request for login
            const res = await fetch("http://localhost:81/SignUp", {
                method: "POST",
                body: JSON.stringify(creds),
                headers: { "content-type": "application/json" }
            });
            if (res.ok) {
                // assign token to local storage
                const responseJson = await res.json();
                console.log(responseJson.Data);
                setTimeout(() => { document.cookie = document.cookie + `authToken=${responseJson.Data.Token};path=/wizard`; }, 1000);
                setTimeout(() => { document.cookie = document.cookie + `authToken=${responseJson.Data.Token};path=/u/${responseJson.Data.Username}`; }, 1000);
                this.updateMessage("Login successful!");
                setTimeout(() => { document.location.href = `/u/${responseJson.Data.Username}/home`; }, 1000);
            }
        }
        catch (error) {
            // update message
            this.updateMessage("Oops! Something went wrong.", error);
        }
    }
    goToLoginPage() {
        document.location.href = "/login";
    }
    // update message
    updateMessage(message, error) {
        error ? () => { console.log(error); console.log(message); } : console.log(message);
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
    // Listen for button clicks
    listen() {
        console.log("Listening for user input...");
        this.signup.addEventListener('click', event => {
            // validate user input
            this.validateInput() ? this.Signup() : this.updateMessage("That's some invalid input mate!");
            event.preventDefault();
        });
        this.login.addEventListener('click', event => {
            this.goToLoginPage();
        });
    }
}
const asdf = new SignUp();
