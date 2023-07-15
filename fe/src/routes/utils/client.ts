import type { iRangeData, iResponse, iCorrectionData, iCredentials, iSaveCorrection, iCorrectionResponse, duration, Signup, Target, Isf, Ratio } from './types';
import axios, { type AxiosRequestConfig } from 'axios';
import Cookies from 'js-cookie';
import { goto } from '$app/navigation';

export const get = async <t>(info: string): Promise<iResponse<t>> => {
	try {

		const res = await fetch(`http://127.0.0.1/${info}`,{
				method:'GET',
				credentials: 'include',
				headers: {
					'Content-Type':'application/json',
				}
			})
			.then((response) => {
				return response.json();
			})
			.catch((error) => {
				throw error;
			});

		return res as iResponse<t>
	} catch (error) {
		console.log(error);
		throw error;
	}
};
export const post = async <t>(info: string, data: t): Promise<iResponse<t>> => {
	try {
		const res = await fetch(`http://127.0.0.1/${info}`,{
			method:'POST',
			credentials: 'include',
			headers: {
				'Content-Type':'application/json',
			},
			body: JSON.stringify(data)
		})
		.then((response) => {
			return response.json();
		})
		.catch((error) => {
			throw error;
		});
		console.log(res)

		return res as iResponse<t>
	} catch (error) {
		console.log(error);
		throw error;
	}
};

export const saveTargets = async (rows: Target[]) => {
	try {
		await post<Target[]>('wizard/Targets', rows)
			.catch((err) => {
				throw err;
			});
	} catch (error) {
		throw error
	}
};

export const getTargets = async () =>{
	try {
		const res = await get<Target[]>('wizard/Targets')
			.then(data => data);

		return res as iResponse<Target[]>
	} catch (error) {
		throw error
	}
}
export const getCurrentTarget = async () =>{
	try {
		const res = await get<Target>('wizard/CurrentTarget')
			.then(data => data);

		return res as iResponse<Target>
	} catch (error) {
		throw error
	}
}

export const saveRatios = async (rows: Ratio[]) =>{
	try {
		const res = await post<Ratio[]>("wizard/Ratios",rows)
			.then(data => data)

		return res as iResponse<Ratio[]>
	} catch (error) {
		throw error
	}
}
export const getRatios = async () => {
	try {
		const res = await get<Ratio[]>("wizard/Ratios")
			.then(data => data)

		return res as iResponse<Ratio[]>
	} catch (error) {
		throw error
	}
}
export const getCurrentRatio = async () => {
	try {
		const res = await get<Ratio>("wizard/CurrentRatio")
			.then(data => data)

		return res as iResponse<Ratio>
	} catch (error) {
		throw error
	}
}
export const saveISFs = async (rows: Isf[]) =>{
	try {
		const res = await post<Isf[]>("wizard/ISFs",rows)
			.then(data => data)

		return res as iResponse<Isf[]>
	} catch (error) {
		throw error
	}
}
export const getISFs = async () => {
	try {
		const res = await get<Isf[]>("wizard/ISFs")
			.then(data => data)

		return res as iResponse<Isf[]>
	} catch (error) {
		throw error
	}
}
export const getCurrentISF = async () => {
	try {
		const res = await get<Isf>("wizard/CurrentISF")
			.then(data => data)

		return res as iResponse<Isf>
	} catch (error) {
		throw error
	}
}
export const saveDuration = async (rows: duration) =>{
	try {
		const res = await post<duration>("wizard/Duration",rows)
			.then(data => data)

		return res as iResponse<duration>
	} catch (error) {
		throw error
	}
}
export const getDuration = async () => {
	try {
		const res = await get<duration>("wizard/Duration")
			.then(data => data)

		return res as iResponse<duration>
	} catch (error) {
		throw error
	}
}

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
export const userSignup = async (creds: iCredentials): Promise<boolean> => {
	try {
		const res = await fetch(`http://127.0.0.1/SignUp`, {
			method: 'POST',
			body: JSON.stringify(creds),
			headers: { 'content-type': 'application/json' }
		});

		if (res.ok) {
			// assign token to local storage
			const response = await res.json();	
			
			Cookies.set('authToken', response.Data.Token, {
				//httpOnly: true,
				secure: false,
				//sameSite: 'strict',
				path: `/`,
				expires: 60 * 60 * 24
			});		
			setTimeout(() =>{
				goto(`/user/${response.Data.Username}`);
			},500);
			
			return true;
		}
		throw Error("SignUp Error")
	} catch (error) {
		console.error(error);
		return false;
	}
};

// Signin
export const userSignin = async (creds: iCredentials): Promise<boolean> => {
	try {
		// setup post request for login
		const res = await fetch(`http://127.0.0.1/SignIn`, {
			method: 'POST',
			body: JSON.stringify(creds),
			headers: { 'content-type': 'application/json' }
		});

		if (res.ok) {
			// assign token to local storage
			const response = await res.json();	
			
			Cookies.set('authToken', response.Data.Token, {
				//httpOnly: true,
				secure: false,
				//sameSite: 'strict',
				path: `/`,
				expires: 60 * 60 * 24
			});

			setTimeout(() =>{
				goto(`/user/${response.Data.Username}`);
			},500);
			
			return true;
		}
		throw Error("SignIn Error")
	} catch (error) {
		console.error(error);
		return error;
	}
};