import type { iRangeData, iResponse, iCorrectionData, iCredentials } from './types';

export const get = async <t>(info: string): Promise<t> => {
	try {
		const res = await fetch(`${process.env.ApiServiceUrl}wizard/${info}`, {
			method: 'GET',
			credentials: 'include',
			headers: { 
				'content-type': 'application/json', 
			}
		})
			.then((data) => data.json())
			.catch((err) => err);

		console.log(res.Data);
		return res.Data;
	} catch (error) {
		console.log(error);
		throw error;
	}
};
export const post = async <t>(info: string, data: string): Promise<t> => {
	try {
		//https://api.boluswizard.io/
		const res = await fetch(`${process.env.ApiServiceUrl}wizard/${info}`, {
			method: 'POST',
			credentials: 'include',
			body: data,
			headers: { 
				'content-type': 'application/json', 
			}
		})
			.then((data) => data.json())
			.catch((err) => err);

		console.log(res.Data);
		return res.Data;
	} catch (error) {
		console.log(error);
		throw error;
	}
};
export const getRange = async (): Promise<iRangeData[]> => {
	try {
		const currentNanoseconds = Date.now() * 1000000;
		// two days 172800000000000
		const res = await get<iRangeData[]>(
			`CorrectionRange?Start=${currentNanoseconds - 172800000000000}&End=${currentNanoseconds}`
		);

		return res;
	} catch (error) {
		throw error;
	}
};
export const saveCorrection = async (): Promise<iResponse> => {
	try {
		var raw = await JSON.stringify([
			{
				//   Bg: +this.bg.value,
				//   Bolus: +this.Bolus.textContent!,
				//   Carbs: +this.carbs.value,
			}
		]);

		const res = await post<iResponse>('Corrections', raw);
        return res
	} catch (error) {
		console.error(error);
	}
};

export const calculateCorrection = async (bg: number, carbs: number): Promise<iCorrectionData> => {
	try {
		const response = await get<iCorrectionData>(`NewCorrection?Bg=${bg}&Carbs=${carbs}`);

		//   this.CarbCorrection.innerHTML = response.CarbCorrection.toString();
		//   this.BgCorrection.innerHTML = response.BgCorrection.toString();
		//   this.ActiveInsulinReduction.innerHTML = response.ActiveInsulinReduction.toString();
		//   this.Bolus.innerHTML = response.Bolus.toString();
        return response
	} catch (error) {
		console.error(error);
	}
};

// Signup
export const Signup = async (creds: iCredentials): Promise<boolean> => {
	try {
		// get user credentials
		//   const creds: iCredentials = await {
		//     Username: this.username.value,
		//     Password: this.password.value
		//   }

		// setup post request for login
		const res = await fetch(`${process.env.ApiServiceUrl}SignUp`, {
			method: 'POST',
			body: JSON.stringify(creds),
			headers: { 'content-type': 'application/json' }
		});

		if (res.ok) {
			// assign token to local storage
			const responseJson = await res.json();
			console.log(responseJson.Data);
			setTimeout(() => {
				document.cookie = document.cookie + `authToken=${responseJson.Data.Token};path=/wizard`;
			}, 1000);
			setTimeout(() => {
				document.cookie =
					document.cookie +
					`authToken=${responseJson.Data.Token};path=/u/${responseJson.Data.Username}`;
			}, 1000);
			// this.updateMessage("Login successful!");

			setTimeout(() => {
				document.location.href = `/u/${responseJson.Data.Username}/home`;
			}, 1000);
            return true
		}
	} catch (error) {
        console.error(error)
        return false
		// update message
		// this.updateMessage("Oops! Something went wrong.", error);
	}
};

// Signin
export const Signin = async (creds: iCredentials): Promise<boolean> => {
	try {
		// setup post request for login
		const res = await fetch(`${process.env.ApiServiceUrl}SignIn`, {
			method: 'POST',
			body: JSON.stringify(creds),
			headers: { 'content-type': 'application/json' }
		});

		if (res.ok) {
			// assign token to local storage
			const responseJson = await res.json();
			console.log(responseJson.Data);
			setTimeout(() => {
				document.cookie = document.cookie + `authToken=${responseJson.Data.Token};path=/wizard`;
			}, 1000);
			setTimeout(() => {
				document.cookie =
					document.cookie +
					`authToken=${responseJson.Data.Token};path=/u/${responseJson.Data.Username}`;
			}, 1000);
			// this.updateMessage('Login successful!');

			setTimeout(() => {
				document.location.href = `/user/${responseJson.Data.Username}`;
			}, 1000);
            return true
		}
	} catch (error) {
        console.error(error)
        return false
		// update message
		// this.updateMessage('Oops! Something went wrong.', error);
	}
};
