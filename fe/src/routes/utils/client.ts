import { error } from '@sveltejs/kit';
import type { iRangeData, iResponse, iCorrectionData, iCredentials, iSaveCorrection, iCorrectionResponse } from './types';
import axios, { type AxiosRequestConfig } from 'axios';
export const get = async <t>(info: string): Promise<iResponse<t>> => {
	try {
		let config = {
			method: 'get',
			maxBodyLength: Infinity,
			withCredentials: true,
			url: `http://localhost/wizard/${info}`,
			headers: {
				'Content-Type': 'application/json'
			},
		};

		const res = await axios
			.request(config)
			.then((response) => {
				return response.data;
			})
			.catch((error) => {
				throw error;
			});

		return res
	} catch (error) {
		console.log(error);
		throw error;
	}
};
export const post = async <t,k>(info: string, data: k): Promise<t> => {
	try {
		//https://api.boluswizard.io/
		let config: AxiosRequestConfig = {
			method: 'post',
			maxBodyLength: Infinity,
			withCredentials: true,
			url: `http://localhost/wizard/${info}`,
			headers: {
				'Content-Type': 'application/json'
			},
			data: data as k
		};

		const res = await axios
			.request(config)
			.then((response) => {
				return response.data as t;
			})
			.catch((error) => {
				throw error;
			});

		return res as t
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
export const saveCorrection = async <t>(body: iSaveCorrection): Promise<iCorrectionResponse> => {
	try {

		body.Key = "";
		body.TimeStamp = null;

		const res = await post<iSaveCorrection>('Corrections', body);
		return res.Data;
	} catch (error) {
		console.error(error);
	}
};

export const calculateCorrection = async (bg: number, carbs: number): Promise<iCorrectionData> => {
	try {
		const res = await get<iCorrectionData>(`NewCorrection?Bg=${bg}&Carbs=${carbs}`)
			.then(data => data.Data);
		return res;
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
			return true;
		}
	} catch (error) {
		console.error(error);
		return false;
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
			return true;
		}
	} catch (error) {
		console.error(error);
		return false;
		// update message
		// this.updateMessage('Oops! Something went wrong.', error);
	}
};
export const test = async <t>(info: string): Promise<t> => {
	try {
		const res = await fetch(`http://127.0.0.1/wizard/Duration`, {
			method: 'GET',
			credentials: 'include',
			headers: {
				'content-type': 'application/json'
			}
		})
			.then((data) => data.json())
			.catch((err) => err);

		console.log(res.Data.duration);
		return res.Data;
	} catch (error) {
		console.log(error);
		throw error;
	}
};
export const axTest = async (): Promise<any> => {
	try {

		let config = {
			method: 'get',
			maxBodyLength: Infinity,
			withCredentials: true,
			url: 'http://localhost/wizard/Duration',
			headers: {
				'Content-Type': 'application/json'
			},
		};

		const res = await axios
			.request(config)
			.then((response) => {
				response.data.Data.duration;
			})
			.catch((error) => {
				throw error;
			});

		return res
	} catch (error) {
		console.log(error);
		return error;
	}
};
