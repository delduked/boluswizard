<script lang="ts">
	import { fly ,fade} from 'svelte/transition';
	import { userSignup } from '../utils/client';

	let Username: string;
	let Password: string;
</script>

<div

	class="hero min-h-screen bg-base-200"
	style="background-image: url(https://boluswizard.io/assets/loginpage/bg.gif);"
>
	<div 
	class="hero-content flex-col lg:flex-row-reverse">
		<div class="text-center lg:text-left lg:ml-4 text-white sm:max-w-lg">
			<h1 class="text-5xl font-bold">Sign Up Now!</h1>
			<p class="py-6">
				Provident cupiditate voluptatem et in. Quaerat fugiat ut assumenda excepturi exercitationem
				quasi. In deleniti eaque aut repudiandae et a id nisi.
			</p>
		</div>
		<div class="card flex-shrink-0 w-full max-w-sm shadow-2xl bg-base-100 ml-4">
			<div class="card-body">
				<div class="form-control">
					<label for="username" class="label">
						<span class="label-text">Username</span>
					</label>
					<input
						id="username"
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
					<label for="signin" class="label">
						<a id="signin" href="/signin" class="label-text-alt link link-hover">Sign In?</a>
					</label>
				</div>
				<div class="form-control mt-6">
					<button
						on:click={() => {
							userSignup({ Username, Password });
						}}
						class="btn btn-primary">Sign up</button
					>
				</div>
			</div>
		</div>
	</div>
</div>
