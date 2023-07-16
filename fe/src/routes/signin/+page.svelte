<script>
	import Alert from './../components/Alert.svelte';
	import { fly, fade } from 'svelte/transition';
	import { redirect } from '@sveltejs/kit';
	import { userSignin } from '../utils/client';
	import { goto } from '$app/navigation';
	import Cookies from 'js-cookie';
	let Username;
	let Password;

	let promise = null;
	async function signin() {
		try {
			if (!Username || !Password) {
				throw Error("Can not sign in null account")
			} else {
				const res = await userSignin({ Username, Password })
					.then(data => data)
					.catch(err => {throw err});
				return res
			}
		} catch (error) {
			throw error
		}
	}

	function handleSignin() {
		promise = signin()

		setTimeout(() => {
			promise = null;
		}, 2000);
	}

</script>

<div
	class="hero min-h-screen bg-base-200"
	style="background-image: url(https://boluswizard.io/assets/loginpage/bg.gif);"
>
	<div
		in:fly={{ y: -5, duration: 500, delay: 500 }}
		out:fly={{ y: 5, duration: 500 }}
		class="hero-content flex-col lg:flex-row-reverse"
	>
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
					<label for="email" class="label">
						<span class="label-text">Email</span>
					</label>
					<input
						id="email"
						bind:value={Username}
						type="text"
						placeholder="email"
						class="input input-bordered"
					/>
				</div>
				<div class="form-control">
					<label for="password" class="label">
						<span class="label-text">Password</span>
					</label>
					<input
						id="password"
						bind:value={Password}
						type="text"
						placeholder="password"
						class="input input-bordered"
					/>
					<label for="signup" class="label">
						<a id="signup" href="signup" class="label-text-alt link link-hover">Signup?</a>
					</label>
				</div>
				<div class="form-control mt-6">
					<button
						on:click|preventDefault={handleSignin}
						class="btn btn-primary">Signin</button
					>
				</div>
				<Alert message={"Signin"} {promise} />
			</div>
		</div>
	</div>
</div>