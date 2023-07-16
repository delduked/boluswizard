import 'dotenv/config';
import jwt from 'jsonwebtoken';
import Cookies from 'js-cookie';
export const load = ({ request }) => {
	function authentication(): {success: boolean,token: string} {
		let futureDate = new Date();
		futureDate.setDate(futureDate.getDate() + 1);
		let success;
		let token;
		request.headers
			.get('Cookie')
			.split(';')
			.forEach((data) => {
				if (data.includes('authToken')) {
					const authToken = data.split('=')[1];
					console.log('authToken');
					console.log(authToken);
					jwt.verify(
						authToken,
						'n4th4n43ln4th4n43ln4th4n43l',
						{ algorithm: 'HS256' },
						function (err, decoded) {
							if (decoded) {
								Cookies.set("authToken", authToken, {
									httpOnly: false,
									secure: false,
									path: `/`,
									sameSite: 'lax',
									expires: futureDate,
									//domain: '*'
								})
								token = authToken
								success = true;
							} else {
								success = false;
							}
						}
					);
				}
			});
			
		return {success,token};
	};
	const res = authentication()

	return {
		auth: res.success,
		token: res.token
	};
};
