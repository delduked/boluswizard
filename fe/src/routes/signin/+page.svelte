<script>
	import { Navigate } from 'svelte-routing';
	import jwt_decode from 'jwt-decode';
	import Cookies from 'js-cookie';
	import { onMount } from 'svelte';
	import { Signin } from '../utils/client';

	let Username;
	let Password;

	const signin = Signin({ Username, Password }).then((data) => {
		if (data == Username) {
			Navigate(`/user/${Username}`);
		}
	});
</script>

<div
	class="hero min-h-screen bg-base-200"
	style="background-image: url(https://boluswizard.io/assets/loginpage/bg.gif);"
>
	<div class="hero-content flex-col lg:flex-row-reverse">
		<div class="text-center lg:text-left text-white sm:max-w-lg">
			<h1 class="text-5xl font-bold">Sign In!</h1>
			<p class="py-6">
				Provident cupiditate voluptatem et in. Quaerat fugiat ut assumenda excepturi exercitationem
				quasi. In deleniti eaque aut repudiandae et a id nisi.
			</p>
		</div>
		<div class="card flex-shrink-0 w-full max-w-sm shadow-2xl bg-base-100 ml-4">
			<div class="card-body">
				<div class="form-control">
					<label class="label">
						<span class="label-text">Email</span>
					</label>
					<input
						bind:value={Username}
						type="text"
						placeholder="email"
						class="input input-bordered"
					/>
				</div>
				<div class="form-control">
					<label class="label">
						<span class="label-text">Password</span>
					</label>
					<input
						bind:value={Password}
						type="text"
						placeholder="password"
						class="input input-bordered"
					/>
					<label class="label">
						<a href="signup" class="label-text-alt link link-hover">Signup?</a>
					</label>
				</div>
				<div class="form-control mt-6">
					<button
						on:click={() => {
							Signin({ Username, Password }).then((data) => {
								if (data == Username) {
									Navigate(`/user/${Username}`);
								}
							});
						}}
						class="btn btn-primary">Signin</button
					>
				</div>
			</div>
		</div>
	</div>
</div>
